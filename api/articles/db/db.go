package db

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"vitess.io/vitess/go/vt/vitessdriver"
)

type DatabaseManager struct {
	db    *sql.DB
	dbBEI *sql.DB // (as an example) shard in Beijing
	dbHKG *sql.DB // (as an example) shard in Hong Kong
}

func NewDatabaseManager() (dbm *DatabaseManager, err error) {
	dbm = &DatabaseManager{}
	dbm.db, err = vitessdriver.Open("traefik:9112", "articles@master")
	if err != nil {
		return nil, err
	}
	dbm.dbBEI, err = vitessdriver.Open("traefik:9112", "articles:-80@master")
	if err != nil {
		return nil, err
	}
	dbm.dbHKG, err = vitessdriver.Open("traefik:9112", "articles:80-@master")
	if err != nil {
		return nil, err
	}
	dbm.db.SetConnMaxLifetime(time.Second * 5)
	dbm.dbBEI.SetConnMaxLifetime(time.Second * 5)
	dbm.dbHKG.SetConnMaxLifetime(time.Second * 5)
	return dbm, err
}

func (dbm *DatabaseManager) GetArticleByID(ID uint64) (article Article, err error) {
	qc := `WHERE _id LIKE ?`
	article, err = dbm.fetchArticle(qc, dbm.db, strconv.Itoa(int(ID)))
	if err != nil {
		log.Println(err.Error())
		return article, err
	}
	return article, nil
}

func (dbm *DatabaseManager) GetArticlesOfCategory(category string) (articles []Article, err error) {
	qc := `WHERE category=?`
	articles, err = dbm.fetchArticles(qc, dbm.db, category)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return articles, nil
}

func (dbm *DatabaseManager) GetArticlesFromRegion(region int) (articles []Article, err error) {
	db := &sql.DB{}
	if region == 1 {
		db = dbm.dbBEI
	} else if region == 2 {
		db = dbm.dbHKG
	} else {
		log.Println("no db region selected")
		db = dbm.db
	}
	qc := ``
	articles, err = dbm.fetchArticles(qc, db, region)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return articles, nil
}

func (dbm *DatabaseManager) fetchArticle(qc string, db *sql.DB, args ...interface{}) (article Article, err error) {
	err = db.QueryRow(`SELECT * FROM article `+qc, args...).Scan(
		&article.ID, &article.Timestamp, &article.ID2, &article.AID, &article.Title, &article.Category,
		&article.Abstract, &article.ArticleTags, &article.Authors, &article.Language, &article.Text,
		&article.Image, &article.Video)
	return article, err
}

func (dbm *DatabaseManager) fetchArticles(qc string, db *sql.DB, args ...interface{}) (articles []Article, err error) {
	rows, err := db.Query(`SELECT * FROM article `+qc, args...)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		var article Article
		err = rows.Scan(&article.ID, &article.Timestamp, &article.ID2, &article.AID, &article.Title, &article.Category,
			&article.Abstract, &article.ArticleTags, &article.Authors, &article.Language, &article.Text,
			&article.Image, &article.Video)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (dbm *DatabaseManager) InsertArticle(article Article) (newID int, err error) {
	sql := `INSERT INTO article (timestamp,id,aid,title,category,abstract,articleTags,authors,language,text,image,video) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`

	id, err := dbm.insertArticle(sql, article)
	if err != nil {
		return 0, err
	}
	newID = int(id)
	return newID, nil
}

func (dbm *DatabaseManager) insertArticle(sql string, article Article) (newID int64, err error) {
	query, err := dbm.db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer query.Close()

	res, err := query.Exec(article.Timestamp, article.ID2, article.AID, article.Title, article.Category,
		article.Abstract, article.ArticleTags, article.Authors, article.Language, article.Text,
		article.Image, article.Video)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

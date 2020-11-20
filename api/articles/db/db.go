package db

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseManager struct {
	db *sql.DB
}

func NewDatabaseManager() (dbm *DatabaseManager, err error) {
	dbm = &DatabaseManager{}
	connStr := "user@tcp(traefik:3000)/articles"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	dbm.db = db
	dbm.db.SetConnMaxLifetime(time.Second * 1)
	return dbm, err
}

func (dbm *DatabaseManager) GetArticleByID(ID uint64) (article Article, err error) {
	qc := `WHERE _id=?`
	article, err = dbm.fetchArticle(qc, strconv.Itoa(int(ID)))
	if err != nil {
		log.Println(err.Error())
		return article, err
	}
	return article, nil
}

func (dbm *DatabaseManager) GetArticlesOfCategory(region string) (articles []Article, err error) {
	qc := `WHERE category=?`
	articles, err = dbm.fetchArticles(qc, region)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (dbm *DatabaseManager) fetchArticle(qc string, args ...interface{}) (article Article, err error) {
	err = dbm.db.QueryRow(`SELECT * FROM article `+qc, args...).Scan(
		&article.ID, &article.Timestamp, &article.ID2, &article.AID, &article.Title, &article.Category,
		&article.Abstract, &article.ArticleTags, &article.Authors, &article.Language, &article.Text,
		&article.Image, &article.Video)
	return article, err
}

func (dbm *DatabaseManager) fetchArticles(qc string, args ...interface{}) (articles []Article, err error) {
	rows, err := dbm.db.Query(`SELECT * FROM article `+qc, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var article Article
		err = rows.Scan(&article.ID, &article.Timestamp, &article.ID2, &article.AID, &article.Title, &article.Category,
			&article.Abstract, &article.ArticleTags, &article.Authors, &article.Language, &article.Text,
			&article.Image, &article.Video)
		if err != nil {
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

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

func (dbm *DatabaseManager) GetArticlesOfRegion(region string) (articles []Article, err error) {
	qc := `WHERE region=?`
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

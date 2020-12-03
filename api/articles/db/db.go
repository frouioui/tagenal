package db

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"vitess.io/vitess/go/vt/vitessdriver"
)

// DatabaseManager contains the connection pools of
// the Vitess MySQL cluster.
// The main connection pool that should be used is
// the `db *sql.DB` one, the others are just for
// proof-of-concept-purpose.
type DatabaseManager struct {
	db *sql.DB

	// connection pool used for the purpose of example
	// connects only to the shard -80 of Vitess Cluster
	dbBEI *sql.DB

	// connection pool used for the purpose of example
	// connects only to the shard 80- of Vitess Cluster
	dbHKG *sql.DB
}

func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var parentCtx opentracing.SpanContext
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan != nil {
		parentCtx = parentSpan.Context()
	}

	tracer := opentracing.GlobalTracer()

	span := tracer.StartSpan(
		method,
		opentracing.ChildOf(parentCtx),
		opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
		ext.SpanKindRPCClient,
	)
	defer span.Finish()

	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	} else {
		md = md.Copy()
	}

	mdw := MetaDataWriter{md}
	err := tracer.Inject(span.Context(), opentracing.TextMap, mdw)
	if err != nil {
		span.LogFields(otlog.String("inject-error", err.Error()))
	}

	newCtx := metadata.NewOutgoingContext(ctx, md)
	err = invoker(newCtx, method, req, reply, cc, opts...)
	if err != nil {
		span.LogFields(otlog.String("call-error", err.Error()))
	}
	return err
}

func createVitessDriverConfig(target string) vitessdriver.Configuration {
	return vitessdriver.Configuration{
		Address:         "vitess-zone1-vtgate-srv.vitess:15999",
		Target:          target,
		GRPCDialOptions: []grpc.DialOption{grpc.WithUnaryInterceptor(UnaryClientInterceptor)},
	}
}

// NewDatabaseManager will return a newly created DatabaseManager,
// the DatabaseManager will contain the initialized connection pools
// to Vitess MySQL cluster.
func NewDatabaseManager() (dbm *DatabaseManager, err error) {
	dbm = &DatabaseManager{}
	dbm.db, err = vitessdriver.OpenWithConfiguration(createVitessDriverConfig("articles@master"))
	if err != nil {
		return nil, err
	}
	dbm.dbBEI, err = vitessdriver.OpenWithConfiguration(createVitessDriverConfig("articles:-80@master"))
	if err != nil {
		return nil, err
	}
	dbm.dbHKG, err = vitessdriver.OpenWithConfiguration(createVitessDriverConfig("articles:80-@master"))
	if err != nil {
		return nil, err
	}
	dbm.db.SetConnMaxLifetime(time.Second * 5)
	dbm.dbBEI.SetConnMaxLifetime(time.Second * 5)
	dbm.dbHKG.SetConnMaxLifetime(time.Second * 5)
	return dbm, err
}

// GetArticleByID will fetch an article from Vitess corresponding
// to the given unique ID.
func (dbm *DatabaseManager) GetArticleByID(ctx context.Context, vtspanctx string, ID uint64) (article Article, err error) {
	qc := `WHERE _id LIKE ?`
	article, err = dbm.fetchArticle(ctx, vtspanctx, qc, dbm.db, strconv.Itoa(int(ID)))
	if err != nil {
		log.Println(err.Error())
		return article, err
	}
	return article, nil
}

// GetArticlesOfCategory will return all the articles belonging to
// the given category.
func (dbm *DatabaseManager) GetArticlesOfCategory(ctx context.Context, vtspanctx, category string) (articles []Article, err error) {
	qc := `WHERE category=?`
	articles, err = dbm.fetchArticles(ctx, vtspanctx, qc, dbm.db, category)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return articles, nil
}

// GetArticlesFromRegion is used as an experiment. It queries all the
// articles that are stored ONLY in the given region ID.
//
// To do such query, the function will not use the default connection
// pool, instead it will use the shard specific connection pool.
//
// RegionID:
// 1 = Beijing
// 2 = Hong Kong
func (dbm *DatabaseManager) GetArticlesFromRegion(ctx context.Context, vtspanctx string, region int) (articles []Article, err error) {
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
	articles, err = dbm.fetchArticles(ctx, vtspanctx, qc, db, region)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return articles, nil
}

func (dbm *DatabaseManager) fetchArticle(ctx context.Context, vtspanctx, qc string, db *sql.DB, args ...interface{}) (article Article, err error) {
	psql := vtspanctx + `SELECT * FROM article ` + qc
	err = db.QueryRowContext(ctx, psql, args...).Scan(
		&article.ID, &article.Timestamp, &article.ID2, &article.AID, &article.Title, &article.Category,
		&article.Abstract, &article.ArticleTags, &article.Authors, &article.Language, &article.Text,
		&article.Image, &article.Video)
	return article, err
}

func (dbm *DatabaseManager) fetchArticles(ctx context.Context, vtspanctx, qc string, db *sql.DB, args ...interface{}) (articles []Article, err error) {
	psql := vtspanctx + ` SELECT * FROM article ` + qc
	rows, err := db.QueryContext(ctx, psql, args...)
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

// InsertArticle will insert the given Article in Vitess MySQL cluster
// the new ID will be returned, in addition to an error if there is any.
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

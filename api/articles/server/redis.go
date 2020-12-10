package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"

	"github.com/frouioui/tagenal/api/articles/db"

	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
)

var rdc *redis.ClusterClient

const (
	defMasterHostname = "redis-cluster.redis"

	defMasterPort = "6379"
)

type redisServiceConfig struct {
	hostname string
	port     string
}

type redisClusterConfig struct {
	master redisServiceConfig
}

func newRedisServiceConfig(hostname, port, defHostname, defPort string) redisServiceConfig {
	if hostname == "" {
		hostname = defHostname
	}
	if port == "" {
		port = defPort
	}
	return redisServiceConfig{
		hostname: hostname,
		port:     port,
	}
}

func (rdcc *redisClusterConfig) getAddrsArray() []string {
	return []string{
		fmt.Sprintf("%s:%s", rdcc.master.hostname, rdcc.master.port),
	}
}

func initRedisClusterClient() error {
	rdcc := redisClusterConfig{
		master: newRedisServiceConfig(os.Getenv("REDIS_MASTER_HOSTNAME"), os.Getenv("REDIS_MASTER_PORT"), defMasterHostname, defMasterPort),
	}

	rdc = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: rdcc.getAddrsArray(),
		NewClient: func(opt *redis.Options) *redis.Client {
			node := redis.NewClient(opt)
			node.AddHook(redisotel.TracingHook{})
			return node
		},
	})
	rdc.AddHook(redisotel.TracingHook{})
	return nil
}

func setCacheArticle(ctx context.Context, query string, data db.Article) error {
	parentSpan := opentracing.SpanFromContext(ctx)
	span := opentracing.StartSpan("redis set cache article", opentracing.ChildOf(parentSpan.Context()))
	defer span.Finish()

	err := rdc.Set(opentracing.ContextWithSpan(context.Background(), span), query, &data, time.Minute).Err()
	if err != nil {
		if err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogFields(otlog.String("error", err.Error()))
		}
		return err
	}
	return nil
}

func getCacheArticle(ctx context.Context, query string, data db.Article) (db.Article, error) {
	parentSpan := opentracing.SpanFromContext(ctx)
	span := opentracing.StartSpan("redis get cache article", opentracing.ChildOf(parentSpan.Context()))
	defer span.Finish()

	str := rdc.Get(ctx, query)
	if err := str.Err(); err != nil {
		if err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogFields(otlog.String("error", err.Error()))
			log.Println(err)
		}
		return db.Article{}, err
	}
	if err := str.Scan(&data); err != nil {
		ext.Error.Set(span, true)
		span.LogFields(otlog.String("error", err.Error()))
		log.Println(err)
		return db.Article{}, err
	}
	return data, nil
}

func setCacheArticles(ctx context.Context, query string, data db.ArticleArray) error {
	err := rdc.Set(ctx, query, &data, time.Minute).Err()
	return err
}

func getCacheArticles(ctx context.Context, query string, data db.ArticleArray) (db.ArticleArray, error) {
	str := rdc.Get(ctx, query)
	err := str.Err()
	if err != nil {
		log.Println(err)
		return db.ArticleArray{}, err
	}
	err = str.Scan(&data)
	if err != nil {
		log.Println(err)
		return db.ArticleArray{}, err
	}
	return data, str.Err()
}

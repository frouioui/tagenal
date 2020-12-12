package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/frouioui/tagenal/api/users/db"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"

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

func wrapperTracing(ctx context.Context, action string) (opentracing.Span, context.Context) {
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan == nil {
		return nil, ctx
	}
	span := opentracing.StartSpan("redis "+action, opentracing.ChildOf(parentSpan.Context()))
	return span, opentracing.ContextWithSpan(context.Background(), span)
}

func setCacheUser(ctx context.Context, query string, data db.User) error {
	span, sctx := wrapperTracing(ctx, "set cache")
	defer span.Finish()
	var err error

	err = rdc.Set(sctx, query, &data, time.Minute).Err()
	if err != nil {
		if err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogFields(otlog.String("error", err.Error()))
		}
		return err
	}
	return nil
}

func getCacheUser(ctx context.Context, query string, data db.User) (db.User, error) {
	span, sctx := wrapperTracing(ctx, "get cache")
	defer span.Finish()
	var err error

	str := rdc.Get(sctx, query)
	if err = str.Err(); err != nil {
		if err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogFields(otlog.String("error", err.Error()))
			log.Println(err)
		}
		return db.User{}, err
	}

	err = str.Scan(&data)
	if err != nil {
		ext.Error.Set(span, true)
		span.LogFields(otlog.String("error", err.Error()))
		log.Println(err)
		return db.User{}, err
	}
	return data, nil
}

func setCacheUsers(ctx context.Context, query string, data db.UserArray) error {
	span, sctx := wrapperTracing(ctx, "set cache")
	defer span.Finish()
	var err error

	err = rdc.Set(sctx, query, &data, time.Minute).Err()
	if err != nil {
		if err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogFields(otlog.String("error", err.Error()))
		}
		return err
	}
	return nil
}

func getCacheUsers(ctx context.Context, query string, data db.UserArray) (db.UserArray, error) {
	span, sctx := wrapperTracing(ctx, "get cache")
	defer span.Finish()
	var err error

	str := rdc.Get(sctx, query)
	if err = str.Err(); err != nil {
		if err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogFields(otlog.String("error", err.Error()))
			log.Println(err)
		}
		return db.UserArray{}, err
	}

	err = str.Scan(&data)
	if err != nil {
		ext.Error.Set(span, true)
		span.LogFields(otlog.String("error", err.Error()))
		log.Println(err)
		return db.UserArray{}, err
	}
	return data, nil
}

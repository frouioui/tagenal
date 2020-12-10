package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/frouioui/tagenal/api/users/db"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
)

var rdc *redis.ClusterClient
var crdc *cache.Cache

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
	crdc = cache.New(&cache.Options{
		Redis:      rdc,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	return nil
}

func setCacheUser(ctx context.Context, query string, data db.User) error {
	err := rdc.Set(ctx, query, &data, time.Minute).Err()
	return err
}

func getCacheUser(ctx context.Context, query string, data db.User) (db.User, error) {
	str := rdc.Get(ctx, query)
	err := str.Err()
	if err != nil {
		log.Println(err)
		return db.User{}, err
	}
	err = str.Scan(&data)
	if err != nil {
		log.Println(err)
		return db.User{}, err
	}
	return data, str.Err()
}

func setCacheUsers(ctx context.Context, query string, data db.UserArray) error {
	err := rdc.Set(ctx, query, &data, time.Minute).Err()
	return err
}

func getCacheUsers(ctx context.Context, query string, data db.UserArray) (db.UserArray, error) {
	str := rdc.Get(ctx, query)
	err := str.Err()
	if err != nil {
		log.Println(err)
		return []db.User{}, err
	}
	err = str.Scan(&data)
	if err != nil {
		log.Println(err)
		return []db.User{}, err
	}
	return data, str.Err()
}

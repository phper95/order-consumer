package main

import (
	"gitee.com/phper95/pkg/cache"
	"gitee.com/phper95/pkg/logger"
	"gitee.com/phper95/pkg/trace"
	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
	"order-consumer/conf"
	"order-consumer/global"
)

func init() {
	global.LoadConfig()
	initRedisClient()
	initMongoClient()
	initESClient()
}

func initRedisClient() {
	redisCfg := conf.Cfg.Redis
	opt := redis.Options{
		Addr:         redisCfg.Host,
		DB:           redisCfg.DB,
		MaxRetries:   redisCfg.MaxRetries,
		PoolSize:     redisCfg.PoolSize,
		MinIdleConns: redisCfg.MinIdleConn,
	}
	redisTrace := trace.Cache{
		Name:                  "redis",
		SlowLoggerMillisecond: 500,
		Logger:                logger.GetLogger(),
		AlwaysTrace:           conf.Cfg.App.RunMode == conf.RunModeDev,
	}
	err := cache.InitRedis(conf.DefaultRedisClient, &opt, &redisTrace)
	if err != nil {
		logger.Error("redis init error", zap.Error(err))
		panic("initRedisClient error")
	}
}

func initESClient() {
	// TO DO ...
}

func initMongoClient() {
	// TO DO ...
}
func main() {

}

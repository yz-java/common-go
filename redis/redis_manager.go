package redis

import (
	"stathat.com/c/consistent"
	"time"
	"common-go/log"
	"github.com/go-redis/redis"
)

var (
	redisManager *RedisManager = nil
	RedisConfigs []RedisConfig = nil
)

type RedisManager struct {

	ipPoolmap map[string]*redis.Client //IP与redis pool 映射

	consistent *consistent.Consistent //hash一致性
}

//实例化
func instance() *RedisManager {
	if RedisConfigs == nil {
		log.Logger.Panic("redis configs is not instance")
	}
	manager := &RedisManager{}
	manager.consistent = consistent.New()
	manager.ipPoolmap = make(map[string]*redis.Client)
	for _, conf := range RedisConfigs {
		client := manager.createClient(conf)
		manager.consistent.Add(conf.Host)
		manager.ipPoolmap[conf.Host] = client
	}
	return manager
}

//获取单实例
func GetInstance() *RedisManager {
	if redisManager == nil {
		redisManager = instance()
	}
	return redisManager
}

//生成redisClient
func (redisManager *RedisManager) createClient(config RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:config.Host,
		DB:0,
		DialTimeout:config.IdleTimeout * time.Second,

	})
	pong, err :=client.Ping().Result()
	log.Logger.Warning(pong,err)
	return client
}

//通过缓存key获取redisClient
func (redisManager *RedisManager) GetRedisClient(key string) *redis.Client {
	ele, _ := redisManager.consistent.Get(key)
	client := redisManager.ipPoolmap[ele]
	return client
}

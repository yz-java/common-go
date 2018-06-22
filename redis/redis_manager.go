package redis

import (
	"stathat.com/c/consistent"
	"time"
	"common-go/log"
	"github.com/go-redis/redis"
)

type RedisManager struct {

	RedisConfigs []RedisConfig

	ipPoolmap map[string]*redis.Client //IP与redis pool 映射

	consistent *consistent.Consistent //hash一致性
}

//实例化
func (this *RedisManager)Instance(){
	if this.RedisConfigs == nil {
		log.Logger.Panic("redis configs is not instance")
	}
	this.consistent = consistent.New()
	this.ipPoolmap = make(map[string]*redis.Client)
	for _, conf := range this.RedisConfigs {
		client := this.createClient(conf)
		this.consistent.Add(conf.Host)
		this.ipPoolmap[conf.Host] = client
	}
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

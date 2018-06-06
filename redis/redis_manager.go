package redis

import (
	"github.com/garyburd/redigo/redis"
	"stathat.com/c/consistent"
	"time"
	"../log"
)

var (
	redisManager *RedisManager = nil
	RedisConfigs []RedisConfig = nil
)

type RedisManager struct {

	ipPoolmap map[string]*redis.Pool //IP与redis pool 映射

	consistent *consistent.Consistent //hash一致性
}

//实例化
func instance() *RedisManager {
	if RedisConfigs == nil {
		log.Logger.Panic("redis configs is not instance")
	}
	manager := &RedisManager{}
	manager.consistent = consistent.New()
	manager.ipPoolmap = make(map[string]*redis.Pool)
	for _, conf := range RedisConfigs {
		pool := manager.createPool(conf)
		manager.consistent.Add(conf.Host)
		manager.ipPoolmap[conf.Host] = pool
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

//生成redisPool
func (redisManager *RedisManager) createPool(config RedisConfig) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     config.MaxIdle,
		IdleTimeout: config.IdleTimeout * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", config.Host) },
	}
}

//通过缓存key获取redisPool/通过pool 获取connection
func (redisManager *RedisManager) GetRedisConnection(key string) redis.Conn {
	ele, _ := redisManager.consistent.Get(key)
	pool := redisManager.ipPoolmap[ele]
	return pool.Get()
}

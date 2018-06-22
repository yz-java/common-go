package redis

var (
	CommonRedisManager *RedisManager
)

func InstanceCommonRedisManager(path string) {
	redisConfigs := LoadConfigFile(path)
	CommonRedisManager = &RedisManager{RedisConfigs:redisConfigs}
	CommonRedisManager.Instance()
}

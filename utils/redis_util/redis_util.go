package redis_util

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
	"strconv"
	rd "../../redis"
	"../../log"
)

func Set(key string, value string, time string) bool {

	connection := rd.GetInstance().GetRedisConnection(key)
	defer connection.Close()
	var err error
	if time == "" {
		_, err = connection.Do("SET", key, value)
	} else {
		_, err = connection.Do("SET", key, value, "EX", time)

	}
	if err != nil {
		log.Logger.Error("Redis Failed to HSET with ERROR, key:", key, "value:", value, "--ERROR--", err)
		return false
	} else {
		log.Logger.Notice("Redis Success to HSET, key:", key, "value:", value)
		return true
	}
}

func Get(key string) (string,error) {
	connection := rd.GetInstance().GetRedisConnection(key)
	defer connection.Close()

	result, err := connection.Do("GET", key)
	resString, err := redis.String(result, nil)
	if err != nil {
		log.Logger.Error(err)
	}
	return resString,err
}

func IncrKey(key string) bool {
	connection := rd.GetInstance().GetRedisConnection(key)
	defer connection.Close()
	var err error
	_, err = connection.Do("INCR", key)
	if err != nil {
		log.Logger.Error("Redis Failed to INCR Key with ERROR, key:", key, "--ERROR--", err)
		return false
	} else {
		log.Logger.Notice("Redis Success to INCR Key, key:", key)
		return true
	}
}

func IncrKeyFloat(key string, increament float64) bool {
	connection := rd.GetInstance().GetRedisConnection(key)
	defer connection.Close()
	var err error
	_, err = connection.Do("INCRBYFLOAT", key,increament)
	if err != nil {
		log.Logger.Error("Redis Failed to INCRBYFLOAT Key with ERROR, key:", key, "--ERROR--", err)
		return false
	} else {
		log.Logger.Notice("Redis Success to INCRBYFLOAT Key, key:", key)
		return true
	}
}

func TTL(key string) (int, error) {
	connection := rd.GetInstance().GetRedisConnection(key)
	defer connection.Close()
	result, err := connection.Do("TTL", key)
	if err != nil {
		log.Logger.Error("Redis Failed to TTL with ERROR, key:", key, "--ERROR--", err)
		return -1, err
	}
	intResult, _ := redis.Int(result, err)
	return intResult, nil
}

func ExsistsKey(key string) bool {
	connection := rd.GetInstance().GetRedisConnection(key)
	defer connection.Close()

	result, err := connection.Do("EXISTS", key)
	if err != nil {
		log.Logger.Error("Redis Opt fail",err)
		return false
	}
	existsInt, _ := redis.Int(result, err)
	if existsInt <= 0 {
		log.Logger.Notice("Reids Key is not exists:", "key:", key)
		return false
	}
	log.Logger.Notice("Reids Key exists:", "key:", key)
	return true

}

//Delete
func DeleteKey(key string) bool {

	connection := rd.GetInstance().GetRedisConnection(key)
	defer connection.Close()
	_, err := connection.Do("DEL", key)
	if err != nil {
		log.Logger.Error("Redis Failed to DEL key:", key)
		return false
	}
	return true
}
func HSet(key string, data map[string]interface{}) bool {
	connection := rd.GetInstance().GetRedisConnection(key)
	defer connection.Close()
	return true
}

func GetHash(key string, desc interface{}) error {
	connection := rd.GetInstance().GetRedisConnection(key)
	defer connection.Close()

	result, err := connection.Do("HGETALL", key)
	if err != nil {
		log.Logger.Error("Redis Failed to HGETALL With ERROR, key:", key, "--ERROR--", err)
		return err
	}
	values, err := redis.Values(result, err)
	if err != nil {
		log.Logger.Error("Redis Failed to HGETALL With ERROR, key:", key, "--ERROR--", err)
		return err
	}
	if len(values)%2 != 0 {
		return errors.New("the length of hash is not %2")
	}
	object := reflect.ValueOf(desc).Elem()
	objTyp := reflect.TypeOf(desc).Elem()
	for i := 0; i < len(values); i += 2 {
		filedName := string(values[i].([]byte))
		for j := 0; j < object.NumField(); j++ {
			if objTyp.Field(j).Name == filedName {
				switch object.Field(j).Interface().(type) {
				case string:
					object.Field(j).SetString(string(values[i+1].([]byte)))
					break
				case int64, int, int32:
					intValue, _ := strconv.ParseInt(string(values[i+1].([]byte)), 10, 64)
					object.Field(j).SetInt(intValue)
					break
				case float64, float32:
					intValue, _ := strconv.ParseFloat(string(values[i+1].([]byte)), 64)
					object.Field(j).SetFloat(intValue)
					break
				default:
					return errors.New("unknow type")
				}
			}
		}
	}
	log.Logger.Notice("Redis Success to HGETALL, key:", key)
	return nil
}

func HashIncr(key, filed string, number interface{}) error {
	connection := rd.GetInstance().GetRedisConnection(key)
	defer connection.Close()

	_, err := connection.Do("HINCRBY", key, filed, number)
	if err != nil {
		log.Logger.Error("Redis Failed to update Hash With ERROR, key:", key, "Filed", filed, "--ERROR--", err)
		return err
	}
	log.Logger.Notice("Redis Success to update Hash With, key:", key, "Filed", filed)
	return nil
}


func HashIncrFloat(key, filed string, number float64) error {
	connection := rd.GetInstance().GetRedisConnection(key)
	defer connection.Close()

	fmt.Println(number)
	_, err := connection.Do("HINCRBYFLOAT", key, filed, number)
	if err != nil {
		log.Logger.Error("Redis Failed to update Hash With ERROR, key:", key, "Filed", filed, "--ERROR--", err)
		return err
	}
	log.Logger.Notice("Redis Success to update Hash With, key:", key, "Filed", filed)
	return nil
}

func SetHashSingle(key, field string, value interface{}) error {
	connection := rd.GetInstance().GetRedisConnection(key)
	defer connection.Close()

	_, err := connection.Do("HSET", key, field, value)
	if err != nil {
		log.Logger.Error("Redis Failed to Set Hash Single:", key, "--ERROR--", err)
		return err
	} else {
		log.Logger.Notice("Redis Success to Set Hash Single:", "key:", key)
		return nil
	}
}

func GetSortScore(key, field string) (string, error) {
	connection := rd.GetInstance().GetRedisConnection(key)
	defer connection.Close()

	reply, err := connection.Do("ZSCORE", key, field)
	if err != nil {
		log.Logger.Error("Redis Failed to Get Sort value key:", key, "field", field, "--ERROR--", err)
		return "", err
	}
	result, _ := redis.String(reply, err)
	return result, nil
}

func GetSortHeight(key string) int {
	connection := rd.GetInstance().GetRedisConnection(key)
	defer connection.Close()
	result, err := connection.Do("ZCARD", key)
	if err != nil {
		log.Logger.Error("Redis Faild to get mining detail count:", "key:", key, "--ERROR--", err)
		return -1
	}
	resultInt, _ := redis.Int(result, err)
	return resultInt
}

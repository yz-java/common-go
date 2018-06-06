package reflect_util

import (
	"reflect"
	"strconv"
	"time"
	"errors"
	"common-go/log"
)


//类型转换
func TypeConversion(value string, tp reflect.Type) (reflect.Value, error) {
	ntype:=tp.Name()
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}

	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}

func ToString(data interface{}) string  {

	ntype:=reflect.ValueOf(data).Type().Name()
	str:=""
	if ntype == "string" {
		str = data.(string)
	} else if ntype == "time.Time" {
		timestamp := data.(time.Time).Unix()
		str= strconv.FormatInt(timestamp, 10)
	} else if ntype == "Time" {
		timestamp := data.(time.Time).Unix()
		str= strconv.FormatInt(timestamp, 10)
	} else if ntype == "int" {
		str= strconv.Itoa(data.(int))
	} else if ntype == "int8" {
		str= strconv.Itoa(int(data.(int8)))
	} else if ntype == "int32" {
		str= strconv.Itoa(int(data.(int32)))
	} else if ntype == "int64" {
		str= strconv.FormatInt(data.(int64), 10)
	} else if ntype == "float32" {
		str= strconv.FormatFloat(data.(float32),'f',6,32)
	} else if ntype == "float64" {
		str= strconv.FormatFloat(data.(float64),'f',6,64)
	}
	log.Logger.Warning("Unknown Type")
	return str
}
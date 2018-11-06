package utils


import (
	"reflect"
	"fmt"
	"common-go/utils/reflect_util"
	"strconv"
	"errors"
	"common-go/log"
	"time"
)

//用map填充结构
func MapToStruct(data map[string]interface{}, obj interface{}) error {
	for k, v := range data {
		err := SetField(obj, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

//用map填充结构
func MapStringToStruct(data map[string]string, obj interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()

	for k, v := range data {
		structFieldValue := structValue.FieldByName(k)

		value,err:=stringToType(v,structFieldValue.Interface())
		if err != nil {
			panic(err)
		}

		structFieldValue.Set(reflect.ValueOf(value))
	}
	return nil
}

//用map的值替换结构的值
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()        //结构体属性值
	structFieldValue := structValue.FieldByName(name) //结构体单个属性值

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type() //结构体的类型
	val := reflect.ValueOf(value)              //map值的反射值

	var err error
	if structFieldType != val.Type() {
		val, err = reflect_util.TypeConversion(fmt.Sprintf("%v", value), structFieldValue.Type()) //类型转换
		if err != nil {
			return err
		}
	}

	structFieldValue.Set(val)
	return nil
}

func stringToType(val string, valType interface{}) (interface{}, error) {

	switch valType.(type) {
	case bool:
		return strconv.ParseBool(val)
	case string:
		return val, nil
	case int:
		return strconv.Atoi(val)
	case int8:
		return strconv.ParseInt(val, 10, 8)
	case int16:
		return strconv.ParseInt(val, 10, 16)
	case int32:
		return strconv.ParseInt(val, 10, 32)
	case int64:
		return strconv.ParseInt(val, 10, 64)
	case uint:
		newVal, err := strconv.Atoi(val)
		return uint(newVal), err
	case uint8:
		strconv.ParseUint(val, 10, 8)
	case uint16:
		strconv.ParseUint(val, 10, 16)
	case uint32:
		strconv.ParseUint(val, 10, 32)
	case uint64:
		strconv.ParseUint(val, 10, 64)
	case float32:
		iVal, err := strconv.ParseFloat(val, 32)
		return float32(iVal), err
	case float64:
		iVal, err := strconv.ParseFloat(val, 64)
		return float64(iVal), err
	case time.Time:
		log.Logger.Info(valType)
		return nil,nil
	default:
		return nil, errors.New("Type not handled")
	}
	return nil, errors.New("Not handled")

}
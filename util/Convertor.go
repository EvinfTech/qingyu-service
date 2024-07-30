package util

/**
此工具类是用来做类型转换使用
Author：陈兆年
data：2021年7月5日16:07:00
*/
import (
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"
)

// byte 转 int64
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

// string to int
func StringToInt(value string) (i int) {
	i, _ = strconv.Atoi(value)
	return
}

// int to int64
func IntToInt64(value int) int64 {
	i, _ := strconv.ParseInt(string(value), 10, 64)
	return i
}

// convert any numeric value to int64
// 任意类型转int64
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	case string:
		d, err = strconv.ParseInt(val.String(), 10, 64)
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}

// convert any numeric value to string
// 任意类型转string
func ToString(value interface{}) (d string) {
	return fmt.Sprintf("%v", value)
}

/**
 * 转Int64
 * @param interface{} data 源数据
 * @return (bool, int64)
 */
func Int64(data interface{}) (ok bool, result int64) {
	var err error
	ok = true
	switch data.(type) {
	case string:
		result, err = strconv.ParseInt(data.(string), 10, 64)
		if err != nil {
			ok = false
		}
		break
	case int:
		result = int64(data.(int))
		break
	case int8:
		result = int64(data.(int8))
		break
	case int16:
		result = int64(data.(int16))
		break
	case int32:
		result = int64(data.(int32))
		break
	case int64:
		result = int64(data.(int64))
		break
	case uint8:
		result = int64(data.(uint8))
		break
	case uint16:
		result = int64(data.(uint16))
		break
	case uint32:
		result = int64(data.(uint32))
		break
	case uint64:
		result = int64(data.(uint64))
		break
	case float32:
		result = int64(data.(float32))
		break
	case float64:
		result = int64(data.(float64))
		break
	case []byte:
		//1
		// b_tmp := data.([]byte)
		// b_buf := bytes.NewBuffer(b_tmp)
		// var result_tmp int32
		// err = binary.Read(b_buf, binary.BigEndian, &result_tmp)
		// if err != nil {
		// 	ok = false
		// 	result = 0
		// } else {
		// 	result = int64(result_tmp)
		// }

		//2
		r_str := string(data.([]byte))
		result, err = strconv.ParseInt(r_str, 10, 64)
		if err != nil {
			ok = false
		}
		break
	case bool:
		if data.(bool) {
			result = 1
		} else {
			result = 0
		}
		break
	default:
		ok = false
		result = 0
		break
	}
	return ok, result
}

/**
 * 转Int
 * @param interface{} data 源数据
 * @return (bool, int)
 */
func Int(data interface{}) (bool, int) {
	ok, result := Int64(data)
	return ok, int(result)
}

// int64 转 byte
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

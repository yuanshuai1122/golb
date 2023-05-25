package utils

import (
	"fmt"
	"reflect"
	"strings"
)

/**
* 模块名: common_utils
* 代码描述: 公共的方法
* 作者:lizibin
* 创建时间:2022/12/23 14:37:41
 */

/**
* 入参: count int64, pageSize int64
* 出参: int64
* 函数作用: 计算分页
* 作者:lizibin
* 创建时间:2022/12/23 14:36:50
 */
func ReturnTotalPage(count int64, pageSize int64) int64 {
	if count%pageSize == 0 {
		return int64(count / pageSize)
	} else {
		return int64(count/pageSize) + 1
	}
}

/**
* 入参: Struct
* 出参: map
* 函数作用: 利用反射将结构体转化为map
* 作者:lizibin
* 创建时间:2022/12/23 14:36:50
 */
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

/**
* 入参: 数组
* 出参: 字符串
* 函数作用: 将数组格式化为字符串
* 作者:lizibin
* 创建时间:2022/12/23 14:36:50
 */
func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

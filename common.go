package zze_goutils

import (
	"fmt"
	"strings"
)

//
//  MergeMaps
//  @Description: 将一个 map[interface{}]interface{} 补丁合并到另一个 map[interface{}]interface{}
//  @param dest 合并到的目标 map
//  @param src 合并源 map
//  @return map[interface{}]interface{} 合并后的 map
//
func MergeMaps(dest, src map[interface{}]interface{}) map[interface{}]interface{} {
	out := make(map[interface{}]interface{}, len(dest))
	for k, v := range dest {
		out[k] = v
	}
	for k, v := range src {
		value := v
		if av, ok := out[k]; ok {
			if v, ok := v.(map[interface{}]interface{}); ok {
				if av, ok := av.(map[interface{}]interface{}); ok {
					out[k] = MergeMaps(av, v)
				} else {
					out[k] = v
				}
			} else {
				out[k] = value
			}
		} else {
			out[k] = v
		}
	}
	return out
}

//
//  AppendStringToMapByKeyExpr
//  @Description: 通过键表达式将指定值追加到嵌套的 map 中的指定键的值后
//  @param keyExpr 表达式 a.b.c 表示定位到 map["a"]["b"]["c"]
//  @param value 要追加到指定键的值
//
func AppendStringToMapByKeyExpr(destMap map[interface{}]interface{}, keyExpr, value string) {
	keys := strings.Split(keyExpr, ".")

	var currentMap map[interface{}]interface{}
	currentMap = destMap

	for index, objKey := range keys {
		if index == len(keys)-1 {
			if obj, ok := currentMap[objKey]; ok {
				currentMap[objKey] = fmt.Sprintf("%s%s", obj, value)
			} else {
				currentMap[objKey] = value
			}
		} else {
			if obj, ok := currentMap[objKey]; ok {
				if obj, ok := obj.(map[interface{}]interface{}); ok {
					currentMap = obj
					continue
				}
			}
			currentMap[objKey] = map[interface{}]interface{}{}
		}
	}
}

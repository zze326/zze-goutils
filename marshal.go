package zze_goutils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//
//  UnmarshalYamlFileToMap
//  @Description: 将 yaml 字符串反序列化为 map[interface{}]interface{} 结构的字典
//  @param yamlFilePath yaml 文件路径
//  @return map[interface{}]interface{} 反序列化后的 map 对象
//  @return error 出错时的错误信息
//
func UnmarshalYamlFileToMap(yamlFilePath string) (map[interface{}]interface{}, error) {
	var m map[interface{}]interface{}
	bs, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(bs, &m); err != nil {
		return nil, err
	}
	return m, nil
}

//
//  UnmarshalYamlToMap
//  @Description: 将 yaml 字符串反序列化为 map[interface{}]interface{} 结构的字典
//  @param yamlStr 要反序列化的 yaml 格式字符串
//  @return map[interface{}]interface{} 反序列化后的 map 对象
//  @return error 出错时的错误信息
//
func UnmarshalYamlToMap(yamlStr string) (map[interface{}]interface{}, error) {
	var m map[interface{}]interface{}
	if err := yaml.Unmarshal([]byte(yamlStr), &m); err != nil {
		return nil, err
	}
	return m, nil
}

//
//  MarshalObjectToYamlString
//  @Description: 将对象序列化为 Yaml 格式字符串
//  @param obj 目标对象
//  @return string 序列化后的 Yaml 字符串
//  @return error
//
func MarshalObjectToYamlString(obj interface{}) (string, error) {
	bs, err := yaml.Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("marshal obj [%#v] faild, err: %v", obj, err)
	}
	return string(bs), nil
}

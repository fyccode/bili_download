package videoinfo

import (
	"encoding/json"
	"fmt"
)

type UrlMaps struct {
	Maps interface{}
}

func UnSerialize(data []byte) map[string]interface{} {
	//将数据先反序列化，存入map中
	urlMaps := make(map[string]interface{})
	err := json.Unmarshal(data, &urlMaps)
	if err != nil {
		panic(err)
	}
	return urlMaps
}

// Get 根据输入的多个参数,根据前后的关系找到对应的值
func (u UrlMaps) Get(para ...interface{}) interface{} {
	for i := 0; i < len(para); i++ {
		u.Maps = u.getValue(para[i])
	}
	return u.Maps
}

// getValue 根据输入的key(不限制类型),获取对应的value
func (u UrlMaps) getValue(key interface{}) interface{} {
	if isString(key) {
		return u.getStrValue(key.(string))
	} else {
		return u.getNumValue(key.(int))
	}
}

// isString 判断断言后是否为字符串类型
func isString(i interface{}) bool {
	_, ok := i.(string)
	return ok
}

// isMap 判断断言后是否是map类型
func isMap(i interface{}) bool {
	_, ok := i.(map[string]interface{})
	return ok
}

// isSlice 判断断言后是否是切片类型
func isSlice(i interface{}) bool {
	_, ok := i.([]interface{})
	return ok
}

// getStrValue 根据key取值
func (u UrlMaps) getStrValue(key string) interface{} {
	//如果是map类型，那就返回key对应的value
	if isMap(u.Maps) {
		a := u.Maps.(map[string]interface{})
		for i, v := range a {
			if i == key {
				return v
			}
		}
	}
	return nil
}

// getNumValue 根据索引取出值
func (u UrlMaps) getNumValue(index int) interface{} {
	if isSlice(u.Maps) {
		return u.Maps.([]interface{})[index]
	}
	return nil
}

// Show 打印视频的清晰度列表
func Show(str, num interface{}) {
	fmt.Println()
	if isSlice(str) && isSlice(num) {
		strList := str.([]interface{})
		numList := num.([]interface{})
		for i := 0; i < len(strList); i++ {
			fmt.Printf("%v --- %v \n", strList[i], numList[i])
		}
		fmt.Println("\n---------------")
	}
}

// GetUrl 根据choice获取对应的url
func GetUrl(i interface{}, choice string) string {
	if !isSlice(i) {
		return ""
	}
	urlList := i.([]interface{})
	for _, v := range urlList {
		if !isMap(v) {
			return ""
		}
		urlMap := v.(map[string]interface{})
		id := urlMap["id"]
		if fmt.Sprintf("%v", id) == choice {
			return fmt.Sprintf("%v", urlMap["baseUrl"])
		}
	}
	return ""
}

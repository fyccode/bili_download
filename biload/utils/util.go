package utils

import (
	"fmt"
	"os"
	"strconv"
)

// ConvertTime 转换时间
func ConvertTime(time []byte) {
	var hour, minute, second int
	timeInt, err := strconv.Atoi(string(time))
	if err != nil {
		panic(err)
	}
	//计算时分秒
	if timeInt >= 3600 {
		hour = timeInt / 3600
		if (timeInt - hour*3600) >= 60 {
			minute = (timeInt - hour*3600) / 60
			second = timeInt - hour*3600 - minute*60
		} else {
			second = timeInt - hour*3600
		}
	} else if timeInt >= 60 {
		minute = timeInt / 60
		second = timeInt - minute*60
	} else {
		second = timeInt
	}
	fmt.Printf("时长： %v时%v分%v秒\n", hour, minute, second)
}

// Read 读取文件
func Read(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(content)
}

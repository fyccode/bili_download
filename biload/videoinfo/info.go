package videoinfo

import (
	"regexp"
)

// RegInfo 通过正则表达式匹配出对应的视频信息
func RegInfo(data []byte) (title, duration, urlMaps, page []byte) {
	//数组存放匹配出的信息，切片存放正则表达式
	var arr [4][]byte
	regs := make([]string, 0)
	regs = append(regs, "<title data-vue-meta=\"true\">(.*?)</title>")
	regs = append(regs, "duration\":(.*?),\"minBufferTime")
	regs = append(regs, "window.__playinfo__=(.*?)</script>")
	regs = append(regs, "\\(1/(.*?)\\)")
	for i, v := range regs {
		reg, err := regexp.Compile(v)
		if err != nil {
			panic(err)
		}
		dataSlice := reg.FindSubmatch(data)
		if len(dataSlice) == 0 {
			arr[i] = nil
			continue
		}
		arr[i] = dataSlice[len(dataSlice)-1]
	}
	return arr[0], arr[1], arr[2], arr[3]
}

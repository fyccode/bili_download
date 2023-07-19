package main

import (
	"biload/download"
	"biload/utils"
	"biload/videoinfo"
	"fmt"
	"log"
	"os/user"
	"strconv"
)

func main() {
	// 实例化一个对象
	fmt.Println("请输入视频url: ")
	var url string
	var p int
	_, _ = fmt.Scanln(&url)
	ur, err := user.Current()
	if err != nil {
		log.Fatal("获取系统用户失败")
	}
	cookie := utils.Read(ur.HomeDir + "/desktop/cookie.txt")
	r := videoinfo.ReqInfo{
		Url:     url,
		Cookie:  cookie,
		Referer: "https://www.bilibili.com/",
	}

	// 获取源码
	sourceCode := r.SourceCode()
	title, duration, urlMaps, page := videoinfo.RegInfo(sourceCode)
	fmt.Println("\n-----查询中-----")
	fmt.Println()
	fmt.Println("标题：", string(title))
	utils.ConvertTime(duration)
	if page == nil {
		fmt.Println("列表数：", 1)
	} else {
		p, _ = strconv.Atoi(string(page))
		fmt.Println("列表数：", p)
	}
	// 进行json解析
	var u videoinfo.UrlMaps
	u.Maps = videoinfo.UnSerialize(urlMaps)

	// 打印视频的所有清晰度
	videoinfo.Show(u.Get("data", "accept_description"), u.Get("data", "accept_quality"))
	fmt.Println("\n请选择 ( **输入清晰度对应的编号** )：")

	//获取音频、视频url，下载
	var choice string
	_, _ = fmt.Scan(&choice)
	videoUrl := videoinfo.GetUrl(u.Get("data", "dash", "video"), choice)
	audioUrl := fmt.Sprintf("%v", u.Get("data", "dash", "audio", 0, "baseUrl"))
	download.DownLoad(videoUrl, audioUrl, string(title))
}

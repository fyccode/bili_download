package download

import (
	"biload/utils"
	"biload/videoinfo"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
)

type reader struct {
	io.Reader
	total   int
	current int
}

func (r *reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.current += n
	fmt.Printf("\r进度：%.2f%%  大小：%.1fMB", float64(r.current*10000/r.total)/100, float64(r.total)/1024.0/1024.0)
	return
}

func downVideo(req *http.Request, filename string) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() { _ = res.Body.Close() }()
	u, err := user.Current()
	if err != nil {
		log.Fatal("获取系统用户失败")
	}
	f, err := os.Create(u.HomeDir + "\\desktop\\" + filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()
	r := &reader{
		total:  int(res.ContentLength),
		Reader: res.Body,
	}
	_, _ = io.Copy(f, r)
}

func DownLoad(videoUrl, audioUrl, title string) {
	if videoUrl == "" {
		fmt.Println("暂无该清晰度视频，请确认输入或是cookie是否正确!!!")
		os.Exit(0)
	}
	u, err := user.Current()
	if err != nil {
		log.Fatal("获取系统用户失败")
	}
	cookiePath := u.HomeDir + "\\desktop\\cookie.txt"
	cookie := utils.Read(cookiePath)
	// 下载视频，文件标题写死了
	v := videoinfo.ReqInfo{
		Url:     videoUrl,
		Cookie:  cookie,
		Referer: "https://www.bilibili.com/",
	}
	downVideo(v.MakeRequest(), "1.mp4")
	fmt.Println("\n视频下载完成")
	// 下载音频
	a := videoinfo.ReqInfo{
		Url:     audioUrl,
		Cookie:  cookie,
		Referer: "https://www.bilibili.com/",
	}
	downVideo(a.MakeRequest(), "1.mp3")
	fmt.Println("\n音频下载完成")
	//合成视频
	Merge("1.mp4", "1.mp3", title)
}

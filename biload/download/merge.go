package download

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

// Merge 合并视频与音频,调用shell执行ffmpeg命令
func Merge(video, audio, mergedVideo string) {
	fmt.Println("开始合成")
	// 获取该操作系统的当前用户信息
	u, err := user.Current()
	if err != nil {
		log.Fatal("获取系统用户失败")
	}
	path := u.HomeDir + "\\desktop\\"
	//将名字中的正反斜杠替换为空
	mergedVideo = strings.ReplaceAll(mergedVideo, "\\", "")
	mergedVideo = strings.ReplaceAll(mergedVideo, "/", "")
	//由于Command会自动将空格两端分开并添加双引号，所以将文件名用反引号进行包裹
	mergedVideo = fmt.Sprintf(`%s`, path+mergedVideo+".mp4")
	cmd := exec.Command("cmd", "/C", "ffmpeg", "-i", path+video, "-i", path+audio, "-vcodec", "copy", "-acodec", "copy", mergedVideo)
	res, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(res))
		fmt.Println("合成失败!!! -- 1")
		os.Exit(0)
	}
	isExit(mergedVideo)
	deleteFile()
}

// isExit 判断是否文件存在
func isExit(filename string) {
	_, err := os.Stat(filename)
	if err == nil {
		fmt.Println("合成成功")
	} else {
		fmt.Println("合成失败!!! -- 2")
	}
}

// 删除原来的文件
func deleteFile() {
	u, err := user.Current()
	if err != nil {
		log.Fatal("获取系统用户失败")
	}
	path := u.HomeDir + "\\desktop\\"
	// 可以使用 os.Remove() 来实现删除文件
	cmd := exec.Command("cmd", "/C", "del", path+"1.mp4", path+"1.mp3")
	_ = cmd.Run()
}

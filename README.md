# b站视频下载(含大会员) Go语言版本(Golang) 
新人练手之作，文件目录分得比较乱，求轻喷！！

**运行的前提条件**：

1.golang，当前使用的是1.19，如果版本不同在go.mod界面更改

2.主机上必须安装ffmpeg，并添加全局的环境变量

**开始运行**：

1.必须在桌面创建一个cookie.txt文件存放cookie(你的b站用户信息)

2.运行
直接运行：进入main目录，```go run main.go```就行
编译后运行：进入main目录 go build -o xxxx.exe(自定义的路径文件名) main.go  
           运行该exe文件即可

3.程序运行过程中两处需要手动输入，第一处是视频url，第二处是选择视频清晰度(输入清晰度对应的编号)

package videoinfo

import (
	"io"
	"net/http"
)

// ReqInfo 包含视频的url和所需的cookie
type ReqInfo struct {
	Url     string
	Cookie  string
	Referer string
}

// MakeRequest 构造请求头
func (r *ReqInfo) MakeRequest() *http.Request {
	req, err := http.NewRequest(http.MethodGet, r.Url, nil)
	if err != nil {
		panic(err)
	}
	//添加cookie
	req.Header.Add("cookie", r.Cookie)
	req.Header.Add("referer", r.Referer)
	return req
}

// SourceCode 获取web页面的源码(byte切片)
func (r *ReqInfo) SourceCode() []byte {
	//发送请求
	req, err := http.DefaultClient.Do(r.MakeRequest())
	if err != nil {
		panic(nil)
	}
	defer func() { _ = req.Body.Close() }()
	content, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	return content
}

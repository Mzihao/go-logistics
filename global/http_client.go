package global

import (
	"net/http"
	"net/http/cookiejar"
)

var HttpClient *http.Client

func InitHttpClient() {
	// 自动更新cookie
	var jar, _ = cookiejar.New(nil)
	HttpClient = &http.Client{Jar: jar}
}

package oj

import (
	"net/http"
)

type OjBaseInfo struct {
	Name      string
	LoginUrl  string
	SubmitUrl string
	StatusUrl string
	Language  map[string]int
	WebCookie *http.Cookie
}

type OjInterface interface {
	Login() (*http.Cookie, error)
	Submit(pid string, language string, usercode string) (*int, error)
	GetRemoteRunId(html string) (*int, error)
	QueryResult(remote_run_id int) (*string, error)
}

var OjManager = make(map[string]OjInterface)

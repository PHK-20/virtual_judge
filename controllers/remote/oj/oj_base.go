package oj

import (
	"net/http"
)

type OjBaseInfo struct {
	Name       string
	LoginUrl   string
	SubmitUrl  string
	StatusUrl  string
	ProblemUrl string
	Language   map[string]int
	WebCookie  *http.Cookie
}

type OjInterface interface {
	Login() (*http.Cookie, error)
	Submit(pid, language, usercode *string) (*string, error)
	GetRemoteRunId(html *string) (*int, error)
	QueryResult(remote_run_id *int) (*string, error)
	IsFinalResult(result *string) bool
	GetProblem(problemid *string) (*ProblemInfo, error)
	GetLanguage() *map[string]int
}

type ProblemInfo struct {
	Title        string
	Description  string
	Input        string
	Output       string
	SampleInput  string
	SampleOutput string
	TimeLimit    string
	MemoryLimit  string
	Hint         string
}

var OjManager = make(map[string]OjInterface)

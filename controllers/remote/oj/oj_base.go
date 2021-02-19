package oj

import (
	"net/http"
)

type OjBase struct {
	Name      string
	Language  map[string]int
	Result    map[string]int
	WebCookie *http.Cookie
}

type OjInterface interface {
	Login() (*http.Cookie, error)
	Submit(pid, language, usercode *string) error
	GetRemoteRunId(pid, language *string) (*int, error)
	QueryResult(remote_run_id *int) (*string, error)
	//IsFinalResult(result *string) bool
	GetProblem(problemid *string) (*ProblemInfo, error)
	GetLanguage() *map[string]int
	GetResultCode(*string) int
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

const (
	ALL   = iota
	AC    //accept
	WA    //wrong answer
	TLE   // time limit exceed
	RE    //runtime error
	PE    //presentation error
	MLE   //memory limit exceed
	OLE   //output limit exceed
	CE    //compile error
	SE    //submit error
	WAIT  //wating
	OTHER //other
)

func (oj *OjBase) GetResultCode(result *string) int {
	switch *result {
	case "Queuing":
		return WAIT
	case "Compiling":
		return WAIT
	case "Running":
		return WAIT
	case "Accepted":
		return AC
	case "Presentation Error":
		return PE
	case "Wrong Answer":
		return WA
	case "Runtime Error":
		return RE
	case "Time Limit Exceeded":
		return TLE
	case "Memory Limit Exceeded":
		return MLE
	case "Output Limit Exceeded":
		return OLE
	case "Compilation Error":
		return CE
	}
	return OTHER
}

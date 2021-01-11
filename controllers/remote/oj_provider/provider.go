package oj_provider

type Provider interface {
	Login() error
	Submit(problemid int, language string, usercode string) error
}

type f struct {
	name string
}
type s struct {
	word string
}

func test() {
}

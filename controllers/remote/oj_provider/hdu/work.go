package hdu

import "net/http"

type hdu_work struct {
	Cookie *http.Cookie
}

var work *hdu_work

func GetHduWork() *hdu_work {
	if work == nil {
		work = new(hdu_work)
		work.Login()
	}
	return work
}

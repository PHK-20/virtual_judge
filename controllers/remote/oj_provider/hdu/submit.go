package hdu

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (oj *hdu_work) Submit(problemid int, language string, usercode string) error {
	// login and get cookie
	if oj.Cookie == nil {
		err := oj.Login()
		if err != nil {
			return err
		}
	}

	oj_info := oj.GetOjInfo()
	url_val := make(url.Values)
	url_val.Add("_usercode", base64.RawStdEncoding.EncodeToString([]byte(url.QueryEscape(usercode))))
	url_val.Add("problemid", strconv.Itoa(problemid))
	url_val.Add("language", oj_info.Language[language])

	req, err := http.NewRequest(http.MethodPost, oj_info.Submit_url, strings.NewReader(url_val.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(oj.Cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

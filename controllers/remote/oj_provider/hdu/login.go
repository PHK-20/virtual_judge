package hdu

import (
	"beego_judge/conf/remote_account"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (oj *hdu_work) Login() error {
	config := remote_account.GetConfig()
	url_val := make(url.Values)
	url_val.Add("username", config.Account.Hdu.Accounts[0].Username)
	url_val.Add("userpass", config.Account.Hdu.Accounts[0].Password)
	req, err := http.NewRequest(http.MethodPost, oj.GetOjInfo().Login_url, strings.NewReader(url_val.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	body_str := string(body)
	if strings.Contains(body_str, "action=login") {
		return errors.New("username or userpass wrong , login fail")
	}
	oj.Cookie = resp.Cookies()[0]
	return nil
}

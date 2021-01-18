package oj

import (
	"beego_judge/conf/remote_account"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type HDU struct {
	OjBaseInfo //base_info
}

var hdu *HDU

func init() {
	hdu = &HDU{
		OjBaseInfo{
			Name:      "HDU",
			LoginUrl:  "http://acm.hdu.edu.cn/userloginex.php?action=login",
			SubmitUrl: "http://acm.hdu.edu.cn/submit.php?action=submit",
			StatusUrl: "http://acm.hdu.edu.cn/status.php",
			Language: map[string]int{
				"G++":    0,
				"GCC":    1,
				"C++":    2,
				"C":      3,
				"Pascal": 4,
				"Java":   5,
				"C#":     6,
			},
		},
	}
	hdu.WebCookie, _ = hdu.Login()
	OjManager[hdu.Name] = hdu
}

func (oj *HDU) Login() (*http.Cookie, error) {
	config := remote_account.GetConfig()
	url_val := make(url.Values)
	url_val.Add("username", config.Account.Hdu.Accounts[0].Username)
	url_val.Add("userpass", config.Account.Hdu.Accounts[0].Password)
	req, err := http.NewRequest(http.MethodPost, oj.LoginUrl, strings.NewReader(url_val.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body_str := string(body)
	if strings.Contains(body_str, "action=login") {
		return nil, errors.New("username or userpass wrong , login fail")
	}
	return resp.Cookies()[0], nil
}

func (oj *HDU) Submit(pid, language, usercode *string) (*string, error) {
	// login and get cookie
	url_val := make(url.Values)
	url_val.Add("_usercode", base64.RawStdEncoding.EncodeToString([]byte(url.QueryEscape(*usercode))))
	url_val.Add("problemid", *pid)
	url_val.Add("language", strconv.Itoa(oj.Language[*language]))

	req, err := http.NewRequest(http.MethodPost, oj.SubmitUrl, strings.NewReader(url_val.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(hdu.WebCookie)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if strings.Contains(string(body), "No such problem") {
		return nil, errors.New("No such problem")
	}
	if strings.Contains(string(body), "action=login") {
		oj.WebCookie, err = oj.Login()
		if err != nil {
			return nil, errors.New("relogin fail")
		}
		return oj.Submit(pid, language, usercode)
	}
	html := string(body)
	return &html, nil
}

func (oj *HDU) GetRemoteRunId(html *string) (*int, error) {
	remote_run_id, err := strconv.Atoi(regexp.MustCompile("(<td height=22px>)\\d+").FindString(*html)[16:])
	if err != nil {
		return nil, err
	}
	return &remote_run_id, nil
}

func (oj *HDU) QueryResult(remote_run_id *int) (*string, error) {
	url_str := fmt.Sprintf("%s?first=%d", oj.StatusUrl, remote_run_id)
	req, err := http.NewRequest(http.MethodGet, url_str, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(oj.WebCookie)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	reg_str := fmt.Sprintf(">%d[\\s\\S]*?font.+?>.+?<", remote_run_id)
	str := regexp.MustCompile(reg_str).FindString(string(body))
	result := str[strings.LastIndex(str, ">")+1 : len(str)-1]
	return &result, nil

}

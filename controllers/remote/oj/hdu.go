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
			Name:      "hdu",
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
	OjManager["hdu"] = hdu
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

func (oj *HDU) Submit(pid string, language string, usercode string) error {
	// login and get cookie
	url_val := make(url.Values)
	url_val.Add("_usercode", base64.RawStdEncoding.EncodeToString([]byte(url.QueryEscape(usercode))))
	url_val.Add("problemid", pid)
	url_val.Add("language", strconv.Itoa(oj.Language[language]))

	req, err := http.NewRequest(http.MethodPost, oj.SubmitUrl, strings.NewReader(url_val.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(hdu.WebCookie)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if strings.Contains(string(body), "No such problem") {
		return errors.New("No such problem")
	}
	return nil
}

func (oj *HDU) GetRemoteRunId(pid string, username string) (*int, error) {
	req, err := http.NewRequest(http.MethodGet, oj.StatusUrl, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(oj.WebCookie)
	url_values := req.URL.Query()
	url_values.Add("pid", pid)
	url_values.Add("user", username)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	remote_run_id, err := strconv.Atoi(regexp.MustCompile("(<td height=22px>)\\d+").FindString(string(body))[16:])
	if err != nil {
		return nil, err
	}
	return &remote_run_id, nil
}

func (oj *HDU) QueryResult(remote_run_id int) (*string, error) {
	req, err := http.NewRequest(http.MethodGet, oj.StatusUrl, nil)
	if err != nil {
		return nil, err
	}
	url_values := req.URL.Query()
	url_values.Add("first", strconv.Itoa(remote_run_id))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	str := regexp.MustCompile(">35084216[\\s\n\\S]*?</font>").FindString(string(body))
	fmt.Println(str)
	if err != nil {
		return nil, err
	}
	return &str, nil

}

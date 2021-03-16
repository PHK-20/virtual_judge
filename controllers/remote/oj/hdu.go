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

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
)

type hdu struct {
	OjBase
	LoginUrl   string
	SubmitUrl  string
	StatusUrl  string
	ProblemUrl string
}

var Hdu *hdu

func init() {
	Hdu = &hdu{
		OjBase{
			Name: "HDU",
			Language: map[string]int{
				"ALL":    0,
				"G++":    1,
				"GCC":    2,
				"C++":    3,
				"C":      4,
				"Pascal": 5,
				"Java":   6,
				"C#":     7,
			},
		},
		"http://acm.hdu.edu.cn/userloginex.php?action=login",
		"http://acm.hdu.edu.cn/submit.php?action=submit",
		"http://acm.hdu.edu.cn/status.php",
		"http://acm.hdu.edu.cn/showproblem.php",
	}
	Hdu.WebCookie, _ = Hdu.Login()
}

func (oj *hdu) Login() (*http.Cookie, error) {
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

func (oj *hdu) Submit(pid, language, usercode *string) error {
	// login and get cookie
	url_val := make(url.Values)
	url_val.Add("_usercode", base64.RawStdEncoding.EncodeToString([]byte(url.QueryEscape(*usercode))))
	url_val.Add("problemid", *pid)
	language_int, ok := oj.Language[*language]
	if !ok {
		return errors.New("wrong language")
	}
	fmt.Println(language_int)
	url_val.Add("language", strconv.Itoa(language_int-1))

	req, err := http.NewRequest(http.MethodPost, oj.SubmitUrl, strings.NewReader(url_val.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(Hdu.WebCookie)
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
	if strings.Contains(string(body), "action=login") {
		oj.WebCookie, err = oj.Login()
		if err != nil {
			return errors.New("relogin fail")
		}
		return oj.Submit(pid, language, usercode)
	}
	return nil
}

func (oj *hdu) GetRemoteRunId(pid, lang *string) (*int, error) {
	username := remote_account.GetConfig().Account.Hdu.Accounts[0].Username
	resp, err := http.Get(fmt.Sprintf("%v?pid=%v&user=%v&lang=%v", oj.StatusUrl, *pid, username, oj.Language[*lang]))
	if err != nil {
		return nil, err
	}
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

func (oj *hdu) QueryResult(remote_runid *int) (*ResultInfo, error) {
	url_str := fmt.Sprintf("%s?first=%d", oj.StatusUrl, *remote_runid)
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
	//result string
	reg_str := fmt.Sprintf(">%d[\\s\\S]*?font.+?>.+?<", *remote_runid)
	str := regexp.MustCompile(reg_str).FindString(string(body))
	res := str[strings.LastIndex(str, ">")+1 : len(str)-1]
	if res == "" {
		return nil, fmt.Errorf("%v queryResult fail ,remote_runid: %v", oj.Name, *remote_runid)
	}
	//time cost
	reg_str = fmt.Sprintf(">%d[\\s\\S\n]*?MS", *remote_runid)
	str = regexp.MustCompile(reg_str).FindString(string(body))
	time_cost := str[strings.LastIndex(str, ">")+1:]
	//Mem cost
	reg_str = fmt.Sprintf(">%d[\\s\\S\n]*?K<", *remote_runid)
	str = regexp.MustCompile(reg_str).FindString(string(body))
	mem_cost := str[strings.LastIndex(str, ">")+1 : len(str)-1]

	result := ResultInfo{}
	result.Res = res
	result.TimeCost = time_cost
	result.MemCost = mem_cost
	return &result, nil

}

func (oj *hdu) GetProblem(problemid *string) (*ProblemInfo, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?pid=%s", oj.ProblemUrl, *problemid), nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	html_utf8 := mahonia.NewDecoder("gbk").ConvertString(string(body))
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html_utf8))
	if err != nil {
		return nil, err
	}
	if strings.Contains(html_utf8, "Invalid Parameter.") {
		return nil, errors.New("wrong problemid")
	}
	if strings.Contains(html_utf8, "No such problem") {
		return nil, errors.New("No such problem")
	}
	var info ProblemInfo
	info.Title = doc.Find("h1").Text()
	tmp_str := []string{}
	doc.Find("div.panel_content").Each(func(idx int, s *goquery.Selection) {
		tmp_str = append(tmp_str, s.Text())
	})
	if len(tmp_str) < 5 {
		return nil, errors.New("get problem fail")
	}
	info.Description = tmp_str[0]
	info.Input = tmp_str[1]
	info.Output = tmp_str[2]
	info.SampleInput = tmp_str[3]
	info.SampleOutput = tmp_str[4]
	if strings.Contains(info.SampleOutput, "Hint") {
		info.Hint = info.SampleOutput[strings.LastIndex(info.SampleOutput, "Hint")+4:]
		info.SampleOutput = info.SampleOutput[:strings.LastIndex(info.SampleOutput, "Hint")]
	}

	if strings.Contains(html_utf8, "Source") {
		reg_str := "source[\\s\\S\\n]*?<"
		str := regexp.MustCompile(reg_str).FindString(html_utf8)
		info.Src = str[strings.LastIndex(str, ">")+1 : len(str)-1]
	}
	reg_str := "Time Limit: [\\s\\S]*?\\)"
	str := regexp.MustCompile(reg_str).FindString(html_utf8)
	info.TimeLimit = str[strings.LastIndex(str, ":")+1:]
	reg_str = "Memory Limit: [\\s\\S]*?\\)"
	str = regexp.MustCompile(reg_str).FindString(html_utf8)
	info.MemoryLimit = str[strings.LastIndex(str, ":")+2:]
	return &info, nil
}

func (oj *hdu) GetLanguage() *map[string]int {
	return &oj.Language
}

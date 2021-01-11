package hdu

import (
	"beego_judge/controllers/remote/oj_provider"
)

type hdu_info struct {
	oj_provider.OjBaseInfo //base_info
}

var info *hdu_info

func init() {
	info = &hdu_info{
		oj_provider.OjBaseInfo{
			Login_url:  "http://acm.hdu.edu.cn/userloginex.php?action=login",
			Submit_url: "http://acm.hdu.edu.cn/submit.php?action=submit",
			Language: map[string]string{
				"G++":    "0",
				"GCC":    "1",
				"C++":    "2",
				"C":      "3",
				"Pascal": "4",
				"Java":   "5",
				"C#":     "6",
			},
		},
	}
}

func (*hdu_work) GetOjInfo() *hdu_info {
	return info
}

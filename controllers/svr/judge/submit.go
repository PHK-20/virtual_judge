package judge

import (
	"beego_judge/controllers/remote/oj"
	"beego_judge/controllers/remote/ojmanager"
	"beego_judge/models"
	"encoding/json"
	"fmt"
	"sync/atomic"

	"github.com/astaxie/beego"
)

type SubmitController struct {
	beego.Controller
}

type reqSubmit struct {
	Username  string `json:"username"`
	Oj        string `json:"oj"`
	Problemid string `json:"problemid"`
	Language  string `json:"language"`
	Usercode  string `json:"usercode"`
}

type respSubmit struct {
	Status   string
	ErrorMsg string
	Data     DataSubmit
}

type DataSubmit struct {
	Runid int
}

var max_run_id *int32

func init() {
	max_run_id, _ = models.GetMaxRunId()
	fmt.Printf("max_run_id: %v\n", *max_run_id)
}

func (c *SubmitController) Post() {
	resp := respSubmit{Status: "fail"}
	defer func() {
		c.Data["json"] = &resp
		c.ServeJSON()
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	req := reqSubmit{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		resp.ErrorMsg = "Wrong request parmas"
		return
	}
	if req.Username == "" {
		resp.ErrorMsg = "Username is empty"
		return
	}
	if len(req.Usercode) < 50 {
		resp.ErrorMsg = "Submit code at least 50 characters"
		return
	}
	ojwork, err := ojmanager.GetOj(&req.Oj)
	if err != nil {
		resp.ErrorMsg = err.Error()
		return
	}
	runid := int(atomic.AddInt32(max_run_id, 1))
	item := models.Submit_status{
		RunId:       runid,
		UserName:    req.Username,
		Oj:          req.Oj,
		ProblemId:   req.Problemid,
		Result:      "submiting",
		ResultCode:  oj.WAIT,
		ExecuteTime: 0,
		Memory:      0,
		Language:    req.Language,
		Length:      len(req.Usercode),
	}
	err = item.AddItem()
	if err != nil {
		resp.ErrorMsg = "server error"
		panic(err)
	}
	resp.Status = "success"
	resp.Data.Runid = runid
	go func() {
		err = ojwork.Submit(&req.Problemid, &req.Language, &req.Usercode)
		var code int
		var result string
		if err != nil {
			code = oj.SE
			result = "Submit Fail"
		} else {
			code = oj.WAIT
			result = "submiting"
		}
		if result == "submiting" {
			go ojmanager.Run(&req.Oj, &req.Problemid, &req.Language, &runid)
		} else {
			item := models.Submit_status{
				RunId:      runid,
				Result:     result,
				ResultCode: code,
			}
			_, err := item.Update("Result", "ResultCode")
			if err != nil {
				panic(err)
			}
		}
	}()

}

func (c *SubmitController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

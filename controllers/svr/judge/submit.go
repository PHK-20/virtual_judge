package judge

import (
	"beego_judge/controllers/remote/oj"
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
	req := reqSubmit{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		resp.ErrorMsg = "wrong request parmas"
		return
	}
	if req.Username == "" {
		resp.ErrorMsg = "username is empty"
		return
	}
	fmt.Println(req)
	if len(req.Usercode) < 50 {
		resp.ErrorMsg = "submit code at least 50 characters"
		return
	}
	oj, ok := oj.OjManager[req.Oj]
	if !ok {
		resp.ErrorMsg = "wrong oj"
		return
	}
	html, err := oj.Submit(&req.Problemid, &req.Language, &req.Usercode)
	if err != nil {
		resp.ErrorMsg = err.Error()
		return
	}
	runid := int(atomic.AddInt32(max_run_id, 1))
	go func() {
		remote_runid, err := oj.GetRemoteRunId(html)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		item := models.Submit_status{
			RunId:        runid,
			RemoteRunId:  *remote_runid,
			Username:     req.Username,
			Oj:           req.Oj,
			ProblemId:    req.Problemid,
			Result:       "submited",
			Execute_Time: 0,
			Memory:       0,
			Language:     req.Language,
			Length:       len(req.Usercode),
		}
		err = item.AddItem()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}()

	resp.Status = "success"
	resp.Data.Runid = runid
}

func (c *SubmitController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

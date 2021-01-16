package svr

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
	Username  string
	Oj        string
	Problemid string
	Language  string
	Usercode  string
}

type respSubmit struct {
	Status string
	Msg    string
	RunId  int
}

var max_run_id *int32

func init() {
	max_run_id, _ = models.GetMaxRunId()
	fmt.Println(*max_run_id)
}

func (c *SubmitController) Post() {
	var req reqSubmit
	resp := &respSubmit{Status: "fail"}
	defer func() {
		c.Data["json"] = &resp
		c.ServeJSON()
	}()
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		resp.Msg = "wrong request parmas"
		return
	}
	if len(req.Usercode) < 50 {
		resp.Msg = "submit code at least 50 characters"
		return
	}
	oj := oj.OjManager[req.Oj]
	html, err := oj.Submit(req.Problemid, req.Language, req.Usercode)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	runid := int(atomic.AddInt32(max_run_id, 1))

	go func() {
		oj.GetRemoteRunId(html)
	}()

	resp.Status = "success"
	resp.RunId = runid
}

func (c *SubmitController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

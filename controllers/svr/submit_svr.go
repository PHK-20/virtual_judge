package svr

import (
	"beego_judge/controllers/remote/oj"
	"encoding/json"
	"sync/atomic"

	"github.com/astaxie/beego"
)

type SubmitController struct {
	beego.Controller
}

type reqSubmit struct {
	Oj        string
	Usercode  string
	Language  string
	Problemid string
}

type respSubmit struct {
	Status      string
	Msg         string
	RemoteRunId int
	RunId       int32
}

var runid int32

func init() {
	runid = 0
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
	remote_run_id, err := oj.Submit(req.Problemid, req.Language, req.Usercode)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	if err != nil {
		resp.Msg = err.Error()
	} else {
		resp.Status = "success"
		resp.RemoteRunId = *remote_run_id
		resp.RunId = atomic.AddInt32(&runid, 1)
	}
}

func (c *SubmitController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

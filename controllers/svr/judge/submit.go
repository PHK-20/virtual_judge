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
	oj, ok := oj.OjManager[req.Oj]
	if !ok {
		resp.ErrorMsg = "Wrong oj"
		return
	}
	err = oj.Submit(&req.Problemid, &req.Language, &req.Usercode)
	if err != nil {
		fmt.Println(err.Error())
		resp.ErrorMsg = err.Error()
		return
	}
	runid := int(atomic.AddInt32(max_run_id, 1))
	item := models.Submit_status{
		RunId:        runid,
		Username:     req.Username,
		Oj:           req.Oj,
		ProblemId:    req.Problemid,
		Result:       "submiting",
		Execute_Time: 0,
		Memory:       0,
		Language:     req.Language,
		Length:       len(req.Usercode),
	}
	item.AddItem()
	go func() {
		remote_runid, err := oj.GetRemoteRunId(&req.Problemid, &req.Language)
		fmt.Printf("runid: %d -> remote_runid: %d\n", runid, *remote_runid)
		if err != nil {
			fmt.Printf("get remote_runid fail, runid: %v error: %v\n", runid, err.Error())
			return
		}
		item := models.Submit_status{
			RunId:       runid,
			RemoteRunId: *remote_runid,
			Result:      "submited",
		}
		_, err = item.Update("RunId", "RemoteRunId", "Result")
		if err != nil {
			fmt.Printf("db: submit_status update fail, runid: %v error: %v", runid, err.Error())
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

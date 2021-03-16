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
	Problem  problemInfo `json:"problem"`
	Username string      `json:"username"`
	Usercode string      `json:"usercode"`
}

type problemInfo struct {
	Id       string
	Oj       string
	Language string
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
	_, err = ojmanager.GetOj(&req.Problem.Oj)
	if err != nil {
		resp.ErrorMsg = err.Error()
		return
	}
	runid := int(atomic.AddInt32(max_run_id, 1))
	item := models.Submit_status{
		RunId:       runid,
		UserName:    req.Username,
		Oj:          req.Problem.Oj,
		ProblemId:   req.Problem.Id,
		Result:      "submiting",
		ResultCode:  oj.WAIT,
		ExecuteTime: "0",
		Memory:      "0",
		Language:    req.Problem.Language,
		Length:      len(req.Usercode),
	}
	err = item.AddItem()
	if err != nil {
		resp.ErrorMsg = "server error"
		panic(err)
	}
	resp.Status = "success"
	resp.Data.Runid = runid
	go Submit(&req, &resp)
}

func (c *SubmitController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

func Submit(req *reqSubmit, resp *respSubmit) {
	ojwork, _ := ojmanager.GetOj(&req.Problem.Oj)
	err := ojwork.Submit(&req.Problem.Id, &req.Problem.Language, &req.Usercode)
	if err != nil {
		item := models.Submit_status{
			RunId:      resp.Data.Runid,
			Result:     "Submit Fail",
			ResultCode: oj.SE,
		}
		_, err := item.Update("Result", "ResultCode")
		if err != nil {
			panic(err)
		}
	} else {
		go ojmanager.Run(&req.Problem.Oj, &req.Problem.Id, &req.Problem.Language, &resp.Data.Runid)
	}
}

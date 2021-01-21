package problem

import (
	"beego_judge/controllers/remote/oj"

	"github.com/astaxie/beego"
)

type GetProblemController struct {
	beego.Controller
}
type reqProblem struct {
	Problemid string
	Oj        string
}

type respProblem struct {
	Status   string
	ErrorMsg string
	Data     DataProblem
}

type DataProblem struct {
	ProblemInfo *oj.ProblemInfo
}

func (c *GetProblemController) Get() {
	resp := respProblem{
		Status: "fail",
	}
	defer func() {
		c.Data["json"] = &resp
		c.ServeJSON()
	}()
	problemid := c.GetString("problemid", "")
	if problemid == "" {
		resp.ErrorMsg = "wrong problemid"
		return
	}
	oj_name := c.GetString("oj", "HDU")
	oj, ok := oj.OjManager[oj_name]
	if !ok {
		resp.ErrorMsg = "wrong oj"
		return
	}
	var err error
	resp.Data.ProblemInfo, err = oj.ShowProblem(&problemid)
	if err != nil {
		resp.ErrorMsg = err.Error()
		return
	}
	resp.Status = "success"
}

func (c *GetProblemController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

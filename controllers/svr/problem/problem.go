package problem

import (
	"beego_judge/controllers/remote/oj"
	"beego_judge/controllers/remote/ojmanager"

	"github.com/astaxie/beego"
)

type GetProblemController struct {
	beego.Controller
}

type reqProblem struct {
	Problemid    string
	Oj           string
	NeedLanguage bool
}

type respProblem struct {
	Status   string
	ErrorMsg string
	Data     DataProblem
}
	
type DataProblem struct {
	ProblemInfo *oj.ProblemInfo
	Language    []string
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
		resp.ErrorMsg = "problemid empty"
		return
	}
	oj_name := c.GetString("oj", "HDU")
	oj, err := ojmanager.GetOj(&oj_name)
	if err != nil {
		resp.ErrorMsg = "Wrong oj"
		return
	}
	need_language, err := c.GetBool("needLanguage", true)
	if err != nil {
		resp.ErrorMsg = err.Error()
		return
	}
	if need_language {
		for key := range *oj.GetLanguage() {
			if key != "ALL" {
				resp.Data.Language = append(resp.Data.Language, key)
			}
		}
	}
	resp.Data.ProblemInfo, err = oj.GetProblem(&problemid)
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

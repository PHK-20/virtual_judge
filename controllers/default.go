package controllers

import (
	"beego_judge/controllers/remote/oj_provider"
	"beego_judge/controllers/remote/oj_provider/hdu"
	"encoding/json"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type SubmitController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
	c.Render()
}

type submitParmas struct {
	Usercode  string
	Language  string
	Problemid string
}

func (c *SubmitController) Post() {
	defer c.ServeJSON()
	var req_parmas submitParmas
	json.Unmarshal(c.Ctx.Input.RequestBody, &req_parmas)
	resp_parmas := make(map[string]interface{})
	if len(req_parmas.Usercode) < 50 {
		resp_parmas["status"] = "fail"
		resp_parmas["msg"] = "submit code at least 50 characters"
		c.Data["json"] = resp_parmas
		return
	}

	var oj oj_provider.Provider
	oj = hdu.GetHduWork()
	err := oj.Submit(req_parmas.Problemid, req_parmas.Language, req_parmas.Usercode)
	if err != nil {
		resp_parmas["status"] = "fail"
		resp_parmas["msg"] = err.Error()
	} else {
		resp_parmas["status"] = "success"
	}
	c.Data["json"] = resp_parmas
}

func (c *SubmitController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

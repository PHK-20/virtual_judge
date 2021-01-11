package controllers

import (
	_ "beego_judge/conf/remote_account"

	"beego_judge/controllers/remote/oj_provider"
	"beego_judge/controllers/remote/oj_provider/hdu"
	"fmt"

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

func (c *SubmitController) Post() {
	code := c.GetString("usercode")
	language := c.GetString("language", "G++")
	problemid, err := c.GetInt("problemid", 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	var oj oj_provider.Provider
	oj = hdu.GetHduWork()
	oj.Submit(problemid, language, code)
}

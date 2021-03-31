package user

import (
	"encoding/json"
	"fmt"
	"time"

	"beego_judge/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}
type reqLogin struct {
	UserName string
	Password string
}

type respLogin struct {
	Status   string
	ErrorMsg string
	Data     DataLogin
}

type DataLogin struct {
	Status       bool
	RegisterTime time.Time
}

func (c *LoginController) Post() {
	resp := respLogin{Status: "fail", Data: DataLogin{Status: false}}
	defer func() {
		c.Data["json"] = &resp
		c.ServeJSON()
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	req := reqLogin{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		resp.ErrorMsg = "Wrong request parmas"
		return
	}
	item, err := models.Check(req.UserName, req.Password)
	if err != nil {
		resp.ErrorMsg = err.Error()
		return
	}
	resp.Data.RegisterTime = item.RegisterTime
	resp.Data.Status = true
	resp.Status = "success"
}

func (c *LoginController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

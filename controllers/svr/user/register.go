package user

import (
	"encoding/json"
	"fmt"

	"beego_judge/models"

	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}
type reqRegister struct {
	Username string
	Password string
	Nickname string
}

type respRegister struct {
	Status   string
	ErrorMsg string
	Data     DataRegister
}

type DataRegister struct {
	Status bool
}

func (c *RegisterController) Post() {
	resp := respRegister{Status: "fail", Data: DataRegister{false}}
	defer func() {
		c.Data["json"] = &resp
		c.ServeJSON()
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	req := reqRegister{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		resp.ErrorMsg = "Wrong request parmas"
		return
	}
	item := models.User_info{
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
	}
	err = item.Register()
	if err != nil {
		resp.ErrorMsg = err.Error()
		return
	}
	resp.Data.Status = true
	resp.Status = "success"
}

func (c *RegisterController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

package user

import (
	"encoding/json"
	"log"
	"time"

	"beego_judge/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}
type reqLogin struct {
	Username string
	Password string
}

type respLogin struct {
	Status   string
	ErrorMsg string
	Data     DataLogin
}

type DataLogin struct {
	Username     string
	Nickname     string
	RegisterTime time.Time
}

func (c *LoginController) Post() {
	resp := respLogin{Status: "fail"}
	defer func() {
		c.Data["json"] = &resp
		c.ServeJSON()
	}()
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	username := c.GetSession("username")
	if username == nil {
		req := reqLogin{}
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
		if err != nil {
			resp.ErrorMsg = "Wrong request parmas"
			return
		}
		user := &models.User_info{}
		user, err = user.Check(req.Username, req.Password)
		if err != nil {
			resp.ErrorMsg = err.Error()
			return
		}
		c.SetSession("username", user.Username)
		c.SetSession("nickname", user.Nickname)
		c.SetSession("register_time", user.RegisterTime)
	}
	SetRespData(&resp, c)
}

func SetRespData(resp *respLogin, c *LoginController) {
	resp.Data.Nickname = c.GetSession("nickname").(string)
	resp.Data.Username = c.GetSession("username").(string)
	resp.Data.RegisterTime = c.GetSession("register_time").(time.Time)
	resp.Status = "success"
}
func (c *LoginController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

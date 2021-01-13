package svr

import (
	"beego_judge/controllers/remote/oj"
	"encoding/json"

	"github.com/astaxie/beego"
)

type QueryController struct {
	beego.Controller
}
type reqQuery struct {
	RunId int32
}

type respQuery struct {
	Status string
	Result string
	Msg    string
}

func (c *QueryController) Post() {
	var req reqQuery
	resp := &respQuery{
		Status: "fail",
	}
	defer func() {
		c.Data["json"] = &resp
		c.ServeJSON()
	}()
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		resp.Msg = "wrong request parmas"
		return
	}
	result, err := oj.OjManager["hdu"].QueryResult(35086405)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Status = "success"
	resp.Result = *result
}

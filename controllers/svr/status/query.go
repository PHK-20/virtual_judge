package status

import (
	"beego_judge/controllers/remote/oj"
	"beego_judge/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type QueryController struct {
	beego.Controller
}
type reqQuery struct {
	Offset    int
	PageSize  int
	condition string
}

type respQuery struct {
	Status   string
	ErrorMsg string
	Data     DataQuery
}

type DataQuery struct {
	Submitions []models.Submit_status
	Total      int64
}

type condition struct {
	Username  string
	ProblemId string
	Oj        string
	Result    string
}

func (c *QueryController) Get() {
	resp := respQuery{
		Status: "fail",
	}
	defer func() {
		c.Data["json"] = &resp
		c.ServeJSON()
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var req reqQuery
	var err error
	req.Offset, err = c.GetInt("offset", 0)
	if err != nil {
		resp.ErrorMsg = err.Error()
		panic(err)
	}
	req.PageSize, err = c.GetInt("num", 10)
	if err != nil {
		resp.ErrorMsg = err.Error()
		panic(err)
	}
	req.condition = c.GetString("condition", "")
	con := condition{}
	json.Unmarshal([]byte(req.condition), &con)
	o := orm.NewOrm()
	qs := o.QueryTable("submit_status")
	if con.Username != "" {
		qs = qs.Filter("UserName", con.Username)
	}
	if con.Oj != "" && con.Oj != "ALL" {
		qs = qs.Filter("Oj", con.Oj)
	}
	if con.ProblemId != "" {
		qs = qs.Filter("ProblemId", con.ProblemId)
	}
	if con.Result != "" && con.Result != "ALL" {
		base := oj.OjBase{}
		code := base.GetResultCode(&con.Result)
		qs = qs.Filter("ResultCode", code)
	}
	_, err = qs.Limit(req.PageSize, req.Offset).All(&resp.Data.Submitions)
	if err != nil {
		resp.ErrorMsg = err.Error()
		panic(err)
	}
	resp.Data.Total, err = qs.Count()
	if err != nil {
		resp.ErrorMsg = err.Error()
		panic(err)
	}
	resp.Status = "success"
}

func (c *QueryController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

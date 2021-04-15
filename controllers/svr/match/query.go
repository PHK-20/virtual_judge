package match

import (
	"beego_judge/models"
	"encoding/json"
	"log"

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
	MatchItem []models.Contest
	Total     int64
}

type condition struct {
	Title   string
	Onwer   string
	MatchId int
}

func (c *QueryController) Get() {
	resp := respQuery{Status: "fail"}
	defer func() {
		c.Data["json"] = &resp
		c.ServeJSON()
	}()
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
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
	qs := o.QueryTable("contest")
	qs = qs.OrderBy("-MatchId")
	if con.MatchId != 0 {
		qs = qs.Filter("MatchId", con.MatchId)
	}
	if con.Title != "" {
		qs = qs.Filter("Title__icontains", con.Title)
	}
	if con.Onwer != "" {
		qs = qs.Filter("Onwer__icontains", con.Onwer)
	}
	_, err = qs.Limit(req.PageSize, req.Offset).All(&resp.Data.MatchItem)
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

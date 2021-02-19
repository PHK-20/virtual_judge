package judge

import (
	"beego_judge/models"
	"fmt"

	"github.com/astaxie/beego"
)

type QueryController struct {
	beego.Controller
}
type reqQuery struct {
	RunId int
}

type respQuery struct {
	Status   string
	ErrorMsg string
	Data     DataQuery
}

type DataQuery struct {
	Result        string
	IsFinalResult bool
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
	req.RunId, err = c.GetInt("runid", 1)
	if err != nil {
		resp.ErrorMsg = err.Error()
		panic(err)
	}

	item := models.Submit_status{}
	isFinalRes, result, err := item.QueryResult(&req.RunId)
	if err != nil {
		resp.ErrorMsg = err.Error()
		panic(err)
	}
	resp.Data.Result = *result
	resp.Status = "success"
	resp.Data.IsFinalResult = *isFinalRes
}

func (c *QueryController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

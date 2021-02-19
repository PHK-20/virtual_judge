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
	var req reqQuery
	var err error
	req.RunId, err = c.GetInt("runid", 1)
	if err != nil {
		fmt.Println(err.Error())
		resp.ErrorMsg = err.Error()
		return
	}

	item := models.Submit_status{}
	isFinalRes, result, err := item.QueryResult(&req.RunId)
	if err != nil {
		fmt.Println(err.Error())
		resp.ErrorMsg = err.Error()
		return
	}
	resp.Data.Result = *result
	resp.Status = "success"
	resp.Data.IsFinalResult = *isFinalRes
}

func (c *QueryController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

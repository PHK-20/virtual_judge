package judge

import (
	"beego_judge/models"

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
	data     DataQuery
}

type DataQuery struct {
	Result string
}

func (c *QueryController) Get() {
	resp := &respQuery{
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
		resp.ErrorMsg = err.Error()
		return
	}

	item := models.Submit_status{}
	result, err := item.QueryResult(&req.RunId)
	if err != nil {
		resp.ErrorMsg = err.Error()
		return
	}
	resp.data.Result = *result
	resp.Status = "success"
}

package routers

import (
	"beego_judge/controllers"
	"beego_judge/controllers/svr/judge"
	"beego_judge/controllers/svr/problem"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//filter setting
	var allow_access = func(c *context.Context) {
		c.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://localhost:7001")
		c.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST")
		c.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
		c.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Accept,Accept-Encoding") //header的类型
	}
	beego.InsertFilter("/submit", beego.BeforeRouter, allow_access)
	beego.InsertFilter("/problem", beego.BeforeRouter, allow_access)
	beego.InsertFilter("/queryResult", beego.BeforeRouter, allow_access)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/submit", &judge.SubmitController{})
	beego.Router("/queryResult", &judge.QueryController{})
	beego.Router("/problem", &problem.GetProblemController{})
}

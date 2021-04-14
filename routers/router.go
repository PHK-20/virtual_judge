package routers

import (
	"beego_judge/controllers"
	"beego_judge/controllers/svr/judge"
	"beego_judge/controllers/svr/match"
	"beego_judge/controllers/svr/problem"
	"beego_judge/controllers/svr/status"
	"beego_judge/controllers/svr/user"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//filter setting
	var allow_access = func(c *context.Context) {
		c.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://localhost:7000")
		c.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST")
		c.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
		c.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Accept,Accept-Encoding")
		c.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true") //header的类型
	}
	//跨域请求
	beego.InsertFilter("/submit", beego.BeforeRouter, allow_access)
	beego.InsertFilter("/problem", beego.BeforeRouter, allow_access)
	beego.InsertFilter("/result", beego.BeforeRouter, allow_access)
	beego.InsertFilter("/status", beego.BeforeRouter, allow_access)
	beego.InsertFilter("/login", beego.BeforeRouter, allow_access)
	beego.InsertFilter("/logout", beego.BeforeRouter, allow_access)
	beego.InsertFilter("/register", beego.BeforeRouter, allow_access)
	beego.InsertFilter("/createContest", beego.BeforeRouter, allow_access)
	beego.InsertFilter("/matchList", beego.BeforeRouter, allow_access)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/submit", &judge.SubmitController{})
	beego.Router("/result", &judge.QueryController{})
	beego.Router("/problem", &problem.GetProblemController{})
	beego.Router("/status", &status.QueryController{})
	beego.Router("/login", &user.LoginController{})
	beego.Router("/register", &user.RegisterController{})
	beego.Router("/logout", &user.LogoutController{})
	beego.Router("/createContest", &match.CreateController{})
	beego.Router("/matchList", &match.QueryController{})
}

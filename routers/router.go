package routers

import (
	"beego_judge/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/submit", &controllers.SubmitController{})
}

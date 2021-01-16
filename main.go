package main

import (
	_ "beego_judge/conf/remote_account"
	_ "beego_judge/routers"

	"github.com/astaxie/beego"
)


func main() {
	beego.Run()
}

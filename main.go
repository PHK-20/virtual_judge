package main

import (
	"beego_judge/conf/remote_account"
	_ "beego_judge/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	remote_account.ReadConfig()
	fmt.Println(remote_account.GetConfig())
	beego.Run()
}

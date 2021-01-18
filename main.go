package main

import (
	_ "beego_judge/routers"
	"fmt"

	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("service start")
	beego.Run()
}

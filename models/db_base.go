package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db_user := beego.AppConfig.String("mysql_user")
	db_pw := beego.AppConfig.String("mysql_password")
	db_host := beego.AppConfig.String("mysql_host")
	db_name := beego.AppConfig.String("mysql_dbname")

	err := orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", db_user, db_pw, db_host, db_name))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db connect success")
	}
	orm.Debug = true
}
func GetMaxRunId() (*int32, error) {
	var runid int32
	db := orm.NewOrm()
	err := db.Raw("select max(runid) from submit_status").QueryRow(&runid)
	if err != nil {
		return nil, err
	}
	return &runid, nil
}

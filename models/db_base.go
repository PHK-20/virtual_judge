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

	err := orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", db_user, db_pw, db_host, db_name))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db connect success")
	}
	orm.Debug = false
}

func GetMaxId(table_name, col string) (*int32, error) {
	var id int32
	db := orm.NewOrm()
	err := db.Raw(fmt.Sprintf("select max(%s) from %s", col, table_name)).QueryRow(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

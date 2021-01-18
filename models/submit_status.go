package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Submit_status struct {
	RunId        int       `orm:"column(runid);pk"`
	RemoteRunId  int       `orm:"column(remote_runid)"`
	Username     string    `orm:"column(username)"`
	Oj           string    `orm:"column(oj)"`
	ProblemId    string    `orm:"column(problemid)"`
	Result       string    `orm:"column(result)"`
	Execute_Time int       `orm:"column(execute_time)"`
	Memory       int       `orm:"column(memory)"`
	Language     string    `orm:"column(language)"`
	Length       int       `orm:"column(length)"`
	Submit_Time  time.Time `orm:"column(submit_time)"`
}

func init() {
	orm.RegisterModel(new(Submit_status))
}

func (table *Submit_status) TableName() string {
	return "submit_status"
}

func (table *Submit_status) AddItem() error {
	db := orm.NewOrm()
	_, err := db.Insert(table)
	if err != nil {
		return err
	}
	return nil
}

package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Contest struct {
	MatchId    int       `orm:"column(matchid);pk"`
	Title      string    `orm:"column(title)"`
	Onwer      string    `orm:"column(onwer)"`
	Desc       string    `orm:"column(descr)"`
	Problem    string    `orm:"column(problem)"`
	BeginTime  time.Time `orm:"column(begin_time)"`
	EndTime    time.Time `orm:"column(end_time)"`
	CreateTime time.Time `orm:"column(create_time)"`
}

func init() {
	orm.RegisterModel(new(Contest))
}

func (item *Contest) itemName() string {
	return "contest"
}

func (item *Contest) Create() error {
	db := orm.NewOrm()
	_, err := db.Insert(item)
	fmt.Println(item)
	if err != nil {
		return err
	}
	return nil
}

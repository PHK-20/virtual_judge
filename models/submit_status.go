package models

import (
	"beego_judge/controllers/remote/oj"
	"errors"
	"fmt"
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
	SubmitTime  time.Time `orm:"column(submit_time)"`
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

func (table *Submit_status) SetResult(runid *int, result *string) error {
	db := orm.NewOrm()
	item := Submit_status{}
	item.RunId = *runid
	item.Result = *result
	num, err := db.Update(&item, "result")
	if err != nil || num != 1 {
		return err
	}
	fmt.Printf("runid:%v result:%s\n", *runid, *result)
	return nil
}

func (table *Submit_status) QueryResult(runid *int) (*string, error) {
	db := orm.NewOrm()
	table.RunId = *runid
	err := db.Read(table)
	if err != nil {
		return nil, errors.New("db submit_status QueryResult fail , wrong runid")
	}
	result := &table.Result
	oj := oj.OjManager[table.Oj]
	if !oj.IsFinalResult(result) || *result == "submited" {
		fmt.Println(oj.IsFinalResult(result))
		result, err = oj.QueryResult(&table.RemoteRunId)
		if err != nil {
			return nil, err
		}
		go table.SetResult(&table.RunId, result)
	}
	return result, nil
}

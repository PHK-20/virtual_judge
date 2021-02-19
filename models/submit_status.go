package models

import (
	"beego_judge/controllers/remote/oj"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Submit_status struct {
	RunId       int       `orm:"column(runid);pk"`
	RemoteRunId int       `orm:"column(remote_runid)"`
	UserName    string    `orm:"column(username)"`
	Oj          string    `orm:"column(oj)"`
	ProblemId   string    `orm:"column(problemid)"`
	Result      string    `orm:"column(result)"`
	ResultCode  int       `orm:"column(result_code)"`
	ExecuteTime int       `orm:"column(execute_time)"`
	Memory      int       `orm:"column(memory)"`
	Language    string    `orm:"column(language)"`
	Length      int       `orm:"column(length)"`
	SubmitTime  time.Time `orm:"column(submit_time)"`
}

func init() {
	orm.RegisterModel(new(Submit_status))
}

func (item *Submit_status) itemName() string {
	return "submit_status"
}

func (item *Submit_status) AddItem() error {
	db := orm.NewOrm()
	_, err := db.Insert(item)
	if err != nil {
		return err
	}
	return nil
}

func (item *Submit_status) Update(cols ...string) (*int64, error) {
	db := orm.NewOrm()
	num, err := db.Update(item, cols...)
	if err != nil || num == 0 {
		return nil, err
	}
	return &num, nil
}

func (item *Submit_status) SetResult(runid *int, result *string) error {
	fmt.Printf("db set result,runid:%v result:%v\n", *runid, *result)
	db := orm.NewOrm()
	item.RunId = *runid
	item.Result = *result
	num, err := db.Update(item, "result")
	if err != nil || num != 1 {
		return err
	}
	return nil
}

func (item *Submit_status) QueryResult(runid *int) (*bool, *string, error) {
	is_final_res := bool(true)
	db := orm.NewOrm()
	item.RunId = *runid
	err := db.Read(item)
	if err != nil {
		return nil, nil, err
	}
	if item.ResultCode == oj.WAIT {
		is_final_res = false
	}
	return &is_final_res, &item.Result, nil
}
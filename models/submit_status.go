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
	SubmitTime   time.Time `orm:"column(submit_time)"`
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

func (item *Submit_status) QueryResult(runid *int) (bool, *string, error) {
	is_final_res := true
	db := orm.NewOrm()
	item.RunId = *runid
	err := db.Read(item)
	if err != nil {
		return false, nil, errors.New(fmt.Sprintf("db submit_status QueryResult fail , wrong runid: %v", *runid))
	}
	result := &item.Result
	if *result == "submiting" {
		return false, result, nil
	}
	oj, ok := oj.OjManager[item.Oj]
	if !ok {
		fmt.Println(item)
		return false, nil, errors.New("wrong oj")
	}
	if !oj.IsFinalResult(result) || *result == "submited" {
		result, err = oj.QueryResult(&item.RemoteRunId)
		if !oj.IsFinalResult(result) {
			is_final_res = false
		}
		if err != nil {
			return is_final_res, nil, err
		}
		go item.SetResult(runid, result)
	}

	return is_final_res, result, nil
}

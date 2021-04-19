package models

import (
	"beego_judge/controllers/remote/oj"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Submit_status struct {
	RunId       int       `orm:"column(runid);pk"`
	MatchId     int       `orm:"column(matchid)"`
	MatchIdx    string    `orm:"column(matchidx)"`
	RemoteRunId int       `orm:"column(remote_runid)"`
	UserName    string    `orm:"column(username)"`
	Oj          string    `orm:"column(oj)"`
	ProblemId   string    `orm:"column(problemid)"`
	Result      string    `orm:"column(result)"`
	ResultCode  int       `orm:"column(result_code)"`
	ExecuteTime string    `orm:"column(execute_time)"`
	Memory      string    `orm:"column(memory)"`
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

func (item *Submit_status) QueryResult(runid *int) (*bool, *oj.ResultInfo, error) {
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
	result := oj.ResultInfo{
		Res:      item.Result,
		TimeCost: item.ExecuteTime,
		MemCost:  item.Memory,
	}
	return &is_final_res, &result, nil
}

func (item *Submit_status) QueryMatchSubmit(matchid int) ([]Submit_status, *int64, error) {
	var record []Submit_status
	o := orm.NewOrm()
	qs := o.QueryTable(item.itemName())
	total, err := qs.Filter("matchid", matchid).OrderBy("submit_time").All(&record, "username", "matchidx", "result", "submit_time")
	if err != nil {
		return nil, nil, err
	}
	return record, &total, nil
}

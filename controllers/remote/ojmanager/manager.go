package ojmanager

import (
	"beego_judge/controllers/remote/oj"
	"beego_judge/models"
	"errors"
	"fmt"
)

type ojManager struct {
	ojs map[string]oj.OjInterface
}

var instance ojManager

func init() {
	instance.ojs = make(map[string]oj.OjInterface)
	Register("HDU", oj.Hdu)
}

func GetOj(oj_name *string) (oj.OjInterface, error) {
	oj, ok := instance.ojs[*oj_name]
	if !ok {
		return nil, errors.New("Wrong oj")
	}
	return oj, nil
}

func Register(oj_name string, oj oj.OjInterface) {
	instance.ojs[oj_name] = oj
}

func Run(oj_name, pid, lang *string, runid *int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	oj_work, _ := GetOj(oj_name)
	remote_runid, err := oj_work.GetRemoteRunId(pid, lang)
	if err != nil {
		item := models.Submit_status{
			RunId:      *runid,
			Result:     "Submit Fail",
			ResultCode: oj.SE,
		}
		item.Update("Result", "ResultCode")
		panic(fmt.Sprintf("runid: %v error: %v\n", *runid, err.Error()))
	}
	fmt.Printf("runid: %d -> remote_runid: %d\n", *runid, *remote_runid)
	item := models.Submit_status{
		RunId:       *runid,
		RemoteRunId: *remote_runid,
		Result:      "submited",
		ResultCode:  oj.WAIT,
	}
	_, err = item.Update("RemoteRunId", "Result", "ResultCode")
	if err != nil {
		panic(fmt.Sprintf("runid: %v error: %v\n", *runid, err.Error()))
	}
	go func() {
		for {
			result, err := oj_work.QueryResult(remote_runid)
			if err != nil {
				panic(fmt.Sprintf("runid: %v error: %v\n", *runid, err.Error()))
			}
			if code := oj_work.GetResultCode(&result.Res); code != oj.WAIT {
				item := models.Submit_status{
					RunId:      *runid,
					Result:     result.Res,
					ExecuteTime: result.TimeCost,
					Memory: result.MemCost,
					ResultCode: code,
				}
				fmt.Println(item)
				_, err = item.Update("RemoteRunId", "Result", "ResultCode","ExecuteTime","Memory")
				if err != nil {
					panic(fmt.Sprintf("runid: %v error: %v\n", *runid, err.Error()))
				}
				break
			}
		}
	}()
}

package match

import (
	"beego_judge/models"
	"encoding/json"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/astaxie/beego"
)

type CreateController struct {
	beego.Controller
}

type reqCreate struct {
	Title       string    `json:"title"`
	Desc        string    `json:"desc"`
	ContestTime []string  `json:"contestTime"`
	ProblemSet  []Problem `json:"problem"`
}

type respCreate struct {
	Status   string
	ErrorMsg string
	Data     DataCreate
}

type DataCreate struct {
	MatchId int
}

type Problem struct {
	Pid string `json:"pid"`
	Oj  string `json:"oj"`
}

var max_match_id *int32

func init() {
	var err error
	max_match_id, err = models.GetMaxId("contest", "matchid")
	if err != nil {
		panic(err)
	}
	fmt.Printf("max_match_id: %v\n", *max_match_id)
}

func (c *CreateController) Post() {

	resp := respCreate{Status: "fail"}
	defer func() {
		c.Data["json"] = &resp
		c.ServeJSON()
	}()
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	username := c.GetSession("username")
	// if username == nil {
	// 	resp.ErrorMsg="Login First"
	// 	return
	// }
	username = "string"
	req := reqCreate{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	if err != nil {
		resp.ErrorMsg = "Wrong request parmas"
		return
	}
	matchid := int(atomic.AddInt32(max_match_id, 1))

	go func() {
		var problem_str string
		for i := 0; i < len(req.ProblemSet); i++ {
			problem_str += req.ProblemSet[i].Oj + "-" + req.ProblemSet[i].Pid + ","
		}
		local, _ := time.LoadLocation("Local")
		bt, _ := time.ParseInLocation("2006-01-02 15:04:05", req.ContestTime[0], local)
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", req.ContestTime[1], local)
		item := &models.Contest{
			MatchId:   matchid,
			Title:     req.Title,
			Onwer:     username.(string),
			Desc:      req.Desc,
			Problem:   problem_str,
			BeginTime: bt,
			EndTime:   et,
		}
		err := item.Create()
		if err != nil {
			log.Println(err.Error())
		}
	}()
	resp.Data.MatchId = matchid
	resp.Status = "success"
}

func (c *CreateController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

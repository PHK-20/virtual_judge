package match

import (
	"beego_judge/controllers/remote/oj"
	"beego_judge/models"
	"log"
	"sort"

	"github.com/astaxie/beego"
)

type RankController struct {
	beego.Controller
}

type reqRank struct {
	matchid int
}

type respRank struct {
	Status   string
	ErrorMsg string
	Data     DataRank
}

type DataRank struct {
	Rank []PersonRank
}

type PersonRank struct {
	Name      string
	Problem   map[string]*ProblemStaut
	TotalCost float64
	ACnum     int
}

type ProblemStaut struct {
	ProblemIdx string
	Status     string
	TryTimes   int
	ACTime     string
}

func (c *RankController) Get() {
	resp := respRank{Status: "fail"}
	defer func() {
		c.Data["json"] = &resp
		c.ServeJSON()
	}()
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	req := reqRank{}
	req.matchid, _ = c.GetInt("matchid", 0)
	if req.matchid == 0 {
		resp.ErrorMsg = "Wrong MatchId"
		return
	}
	item := &models.Submit_status{}
	records, _, err := item.QueryMatchSubmit(req.matchid)
	if err != nil {
		resp.ErrorMsg = err.Error()
		panic(err)
	}
	name2rank := make(map[string]*PersonRank)
	contest := &models.Contest{}
	contest, err = contest.QueryContest(req.matchid)
	if err != nil {
		panic(err)
	}
	for _, v := range records {
		log.Println(v)
		name := v.UserName
		idx := v.MatchIdx
		code := v.ResultCode
		res := v.Result
		p, ok := name2rank[name]
		if !ok {
			name2rank[name] = &PersonRank{
				Name:    name,
				ACnum:   0,
				Problem: make(map[string]*ProblemStaut),
			}
			p = name2rank[name]
		}
		ps, ok := p.Problem[idx]
		if !ok {
			p.Problem[idx] = &ProblemStaut{ProblemIdx: idx}
			ps = p.Problem[idx]
		}
		if ps.Status == "AC" {
			continue
		}
		if code == oj.AC {
			ps.Status = "AC"
			ps.ACTime = v.SubmitTime.String()
			p.TotalCost += v.SubmitTime.Sub(contest.BeginTime).Seconds()
			p.TotalCost += float64(ps.TryTimes) * 20 * 60
			p.ACnum++
		} else {
			ps.Status = res
			ps.TryTimes++
		}
	}
	var finalRes RankSlice
	for _, v := range name2rank {
		finalRes = append(finalRes, *v)
	}
	log.Println(finalRes)
	sort.Sort(finalRes)
	resp.Data.Rank = finalRes
	resp.Status = "success"
}

type RankSlice []PersonRank

func (s RankSlice) Len() int {
	return len(s)
}

func (s RankSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s RankSlice) Less(i, j int) bool {
	if s[i].ACnum != s[j].ACnum {
		return s[i].ACnum > s[j].ACnum
	}
	return s[i].TotalCost < s[j].TotalCost
}

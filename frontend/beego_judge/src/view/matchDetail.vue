<template>
  <div class="matchDetail">
    <h1>{{ matchInfo.Title }}</h1>
    <el-row :gutter="1">
      <el-col :span="4">{{
        new Date(matchInfo.BeginTime).toLocaleString()
      }}</el-col>
      <el-col :span="16">
        <el-progress
          :percentage="percentage"
          :show-text="showText"
          :text-inside="true"
          :stroke-width="17"
        ></el-progress>
      </el-col>
      <el-col :span="4">{{
        new Date(matchInfo.EndTime).toLocaleString()
      }}</el-col>
    </el-row>
    <div v-bind:style="{ color: statusColor, 'margin-top': '20px' }">
      {{ status }}
    </div>
    <el-row>
      <el-col :span="22" :offset="1">
        <el-tabs
          type="border-card"
          v-model="activeTab"
          @tab-click="handleClick"
          style="margin-top: 20px"
        >
          <el-tab-pane label="Overview">
            <el-table :data="problemSet" stripe style="width: 100%">
              <el-table-column prop="status" label="" align="center">
              </el-table-column>
              <el-table-column prop="idx" label="#"> </el-table-column>
              <el-table-column label="title" width="1400">
                <template slot-scope="scope">
                  <el-link
                    type="primary"
                    :underline="false"
                    @click="toProblem(scope.row)"
                    >{{ scope.row.title }}</el-link
                  >
                </template>
              </el-table-column>
              ></el-table
            >
          </el-tab-pane>
          <el-tab-pane label="Status">
            <status ref="status" :matchid="matchid"></status>
          </el-tab-pane>
          <el-tab-pane label="Rank">
            <rank
              ref="rank"
              :matchid="matchid"
              :problemSet="problemSet"
              :matchInfo="matchInfo"
            ></rank>
          </el-tab-pane>
        </el-tabs>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import status from "@/components/status";
import rank from "@/components/rank";
export default {
  components: {
    status,
    rank,
  },
  data() {
    return {
      matchid: 0,
      matchInfo: {},
      percentage: 0,
      showText: true,
      status: "",
      statusColor: "red",
      activeTab: "",
      problemSet: [],
    };
  },
  methods: {
    text(percentage) {
      if (percentage >= 100) {
        return "已结束";
      }
      return percentage + "%";
    },
    handleClick(tab, event) {
      if (tab.label == "Status") {
        this.$refs.status.query();
      }
      if (tab.label == "Rank") {
        this.$refs.rank.query();
      }
    },
    calPercentage() {
      let all = this.getSecond(
        new Date(this.matchInfo.BeginTime),
        new Date(this.matchInfo.EndTime)
      );
      let gone = this.getSecond(new Date(this.matchInfo.BeginTime), new Date());
      let res = (gone / all) * 100;
      if (res >= 100) res = 100;
      else if (res < 0) res = -1;
      this.percentage = Number(res.toFixed(1));
      if (this.percentage == 100 || this.percentage == -1) {
        this.showText = false;
      }
      if (this.percentage == 100) {
        this.status = "Ended";
        this.statusColor = "green";
      } else if (this.percentage == -1) {
        this.status = "Pending";
        this.statusColor = "gray";
      } else {
        this.status = "Running";
        this.statusColor = "red";
      }
    },
    getSecond(s, e) {
      return (e.getTime() - s.getTime()) / 1000;
    },
    toProblem(row) {
      let info = {
        idx: row.idx,
        pid: row.pid,
        oj: row.oj,
      };

      let routeUrl = this.$router.resolve({
        name: "matchProblem",
        params: {
          matchid: this.matchid,
          info: JSON.stringify(info),
        },
      });
      window.open(routeUrl.href, "_blank");
    },
  },
  beforeDestroy() {
    clearInterval(this.timer);
  },
  mounted: function () {
    this.matchid = Number(this.$route.params.matchid);
    this.$axios
      .get("/matchList", {
        params: {
          condition: {
            matchid: this.matchid,
          },
        },
      })
      .then((resp) => {
        let data = resp.data;
        if (data.Data.Total == 1) {
          this.matchInfo = data.Data.MatchItem[0];
          let problem_info = [];
          data.Data.MatchItem[0].Problem.split(",").forEach((item) => {
            let tmp = [];
            tmp = item.split("-");
            problem_info.push({
              oj: tmp[0],
              pid: tmp[1],
            });
          });
          data.Data.MatchItem[0].ProblemTitle.split(",").forEach(
            (item, idx) => {
              this.problemSet.push({
                idx: String.fromCharCode("A".charCodeAt() + idx),
                title: item,
                oj: problem_info[0].oj,
                pid: problem_info[0].pid,
              });
            }
          );
          this.calPercentage();
          this.timer = setInterval(() => {
            this.calPercentage();
          }, 5000);
        } else {
          this.$notify.error({
            title: "Error",
            message: "Wrong MatchId",
          });
        }
      })
      .catch((err) => {
        console.error(err);
      });
  },
};
</script>

<style>
</style>
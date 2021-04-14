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
  </div>
</template>

<script>
export default {
  data() {
    return {
      matchid: "",
      matchInfo: {},
      percentage: 0,
      showText: true,
      status: "",
      statusColor: "red",
    };
  },
  methods: {
    text(percentage) {
      if (percentage >= 100) {
        return "已结束";
      }
      return percentage + "%";
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
  },
  beforeDestroy() {
    clearInterval(this.timer);
  },
  mounted: function () {
    this.matchid = this.$route.params.id;
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
        console.log(resp);
        if (data.Data.Total == 1) {
          this.matchInfo = data.Data.MatchItem[0];
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
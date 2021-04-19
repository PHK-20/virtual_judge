<template>
  <div class="rank">
    <el-table :data="tableData" stripe style="width: 100%" v-loading="loading">
      <el-table-column align="center" label="rank" prop="Rank">
      </el-table-column>
      <el-table-column
        align="center"
        label="User"
        prop="Nickname"
      ></el-table-column>
      <el-table-column
        align="center"
        label="Solved"
        prop="ACnum"
      ></el-table-column>
      <el-table-column
        align="center"
        label="Penalty"
        prop="Penalty"
      ></el-table-column>
      <template v-for="item in problemSet">
        <el-table-column align="center" :key="item.idx">
          <template slot="header">
            {{ item.idx }}
          </template>
          <template slot-scope="scope">
            {{ text(scope.row.Problem, item.idx) }}
          </template>
        </el-table-column>
      </template>
    </el-table>
  </div>
</template>

<script>
export default {
  props: {
    matchid: Number,
    problemSet: Array,
    matchInfo: Object,
  },
  data() {
    return {
      tableData: [],
      loading: false,
    };
  },
  methods: {
    text(problem, idx) {
      if (problem[idx]) {
        let times_str = "";
        if (problem[idx].TryTimes != 0) {
          times_str = "(-" + problem[idx].TryTimes + ")";
        }
        if (problem[idx].ACTime != "") {
          var s1 = new Date(this.matchInfo.BeginTime).getTime();
          var s2 = new Date(problem[idx].ACTime).getTime();
          var total = (s2 - s1) / 1000;
          var day = parseInt(total / (24 * 60 * 60)); //计算整数天数
          var afterDay = total - day * 24 * 60 * 60; //取得算出天数后剩余的秒数
          var hour = parseInt(afterDay / (60 * 60)); //计算整数小时数
          var afterHour = total - day * 24 * 60 * 60 - hour * 60 * 60; //取得算出小时数后剩余的秒数
          var min = parseInt(afterHour / 60); //计算整数分
          var afterMin = total - day * 24 * 60 * 60 - hour * 60 * 60 - min * 60; //取得算出分后剩余的秒数
          hour += day * 24;
          return hour + ":" + min + ":" + afterMin + times_str;
        } else {
          return times_str;
        }
      }
      return "";
    },
    query() {
      this.loading = true;
      this.$axios
        .get("/rank", {
          params: {
            matchid: this.matchid,
          },
        })
        .then((resp) => {
          let data = resp.data;
          if (data.Status == "success") {
            this.tableData = data.Data.Rank;
            if (this.tableData) {
              this.tableData.forEach((item) => {
                item.Penalty = Number(item.Penalty.toFixed(1));
              });
            }
            this.loading = false;
          } else {
            this.$notify.error({
              title: "Error",
              message: data.ErrorMsg,
            });
          }
        })
        .catch((err) => {
          console.error(err);
        });
    },
  },
};
</script>
 
<style>
</style>
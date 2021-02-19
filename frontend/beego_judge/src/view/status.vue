<template>
  <div class="status">
    <el-pagination
      background
      layout="sizes, prev, pager, next"
      :total="total"
      :current-page="page"
      :page-sizes="[10, 20, 50, 100]"
      :page-size="pageSize"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    >
    </el-pagination>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column align="center" label="RunId" prop="RunId">
      </el-table-column>
      <el-table-column align="center" prop="UserName">
        <template slot="header" slot-scope="scope">
          <div>Username</div>
          <el-input v-model="condition.username" @change="query()" size="mini" />
        </template>
      </el-table-column>
      <el-table-column align="center" label="OJ" prop="Oj">
        <template slot="header" slot-scope="scope">
          <div>OJ</div>
          <el-select v-model="condition.oj" @change="query()" size="mini">
            <el-option v-for="item in ojArray" :key="item" :value="item">
            </el-option>
          </el-select>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="ProblemId">
        <template slot="header" slot-scope="scope">
          <div>ProblemId</div>
          <el-input v-model="condition.problemid" @change="query()"  size="mini" />
        </template>
      </el-table-column>
      <el-table-column align="center" prop="Result">
        <template slot="header" slot-scope="scope">
          <div>Result</div>
          <el-select v-model="condition.result" @change="query()" size="mini">
            <el-option v-for="item in resultArray" :key="item" :value="item">
            </el-option>
          </el-select>
        </template>
        <template slot-scope="scope">
          <span v-if="scope.row.ResultCode === 1" style="color: green">{{
            scope.row.Result
          }}</span>
          <span v-else-if="scope.row.ResultCode === 10" style="color: gray">{{
            scope.row.Result
          }}</span>
          <span v-else style="color: red">{{ scope.row.Result }}</span>
        </template>
      </el-table-column>
      <el-table-column label="Time(ms)" prop="ExecuteTime"> </el-table-column>
      <el-table-column label="Mem(mb)" prop="Memory"> </el-table-column>
      <el-table-column label="Length" prop="Length"> </el-table-column>
      <el-table-column label="Lang" prop="Language"> </el-table-column>
      <el-table-column label="Submit Time" prop="SubmitTime"> </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      pageSize: 10,
      currentPage: 1,
      total: 0,
      condition: {
        username: "",
        problemid: "",
        oj: "ALL",
        result: "ALL",
      },
      ojArray: ["ALL", "HDU"],
      resultArray: [
        "ALL",
        "Accepted",
        "Wrong Answer",
        "Time Limit Exceeded",
        "Runtime Error",
        "Presentation Error",
        "Memory Limit Exceeded",
        "Output Limit Exceeded",
        "Compilation Error",
        "Submit Error",
        "Wating",
        "Other Error",
      ],
      tableData: [],
    };
  },
  mounted: function () {
    this.query();
  },
  methods: {
    query: function () {
      this.$axios
        .get("/status", {
          params: {
            num: this.pageSize,
            offset: (this.currentPage - 1) * 10,
            condition: this.condition,
          },
        })
        .then((resp) => {
          console.log(resp);
          if (resp.data.Status == "success") {
            this.tableData = [];
            let result = resp.data.Data;
            result.Submitions.forEach((item) => {
              this.tableData.push(item);
            });
            this.total = result.Total;
          }
        })
        .catch((err) => {
          console.error(err);
        });
    },
    handleCurrentChange(page) {
      this.currentPage = page;
      this.query();
    },
    handleSizeChange(size) {
      this.pageSize = size;
      this.query();
    },
  },
};
</script>

<style>
</style>
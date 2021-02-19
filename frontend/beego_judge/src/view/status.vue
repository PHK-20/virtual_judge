<template>
  <div class="status">
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column align="center" label="RunId" prop="runid">
      </el-table-column>
      <el-table-column align="center" prop="username">
        <template slot="header" slot-scope="scope">
          <div>Username</div>
          <el-input v-model="username" size="mini" />
        </template>
      </el-table-column>
      <el-table-column align="center" label="OJ" prop="oj">
        <template slot="header" slot-scope="scope">
          <div>OJ</div>
          <el-select v-model="oj" size="mini">
            <el-option v-for="item in oj_array" :key="item" :value="item">
            </el-option>
          </el-select>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="problemid">
        <template slot="header" slot-scope="scope">
          <div>ProblemId</div>
          <el-input v-model="problemid" size="mini" />
        </template>
      </el-table-column>
      <el-table-column align="center" prop="result">
        <template slot="header" slot-scope="scope">
          <div>Result</div>
          <el-select v-model="result" size="mini">
            <el-option v-for="item in result_array" :key="item" :value="item">
            </el-option>
          </el-select>
        </template>
      </el-table-column>
      <el-table-column label="Time(ms)" prop="time"> </el-table-column>
      <el-table-column label="Mem(mb)" prop="memory"> </el-table-column>
      <el-table-column label="Length" prop="length"> </el-table-column>
      <el-table-column label="Lang" prop="length"> </el-table-column>
      <el-table-column label="Submit Time" prop="submit"> </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      oj: "ALL",
      oj_array: ["ALL", "HDU"],
      result: "ALL",
      result_array: [
        "ALL",
        "Accepted",
        "Wrong Answer",
        "Time Limit Exceed",
        "Runtime Error",
        "Presentation Error",
        "Memory Limit Exceed",
        "Output Limit Exceed",
        "Compile Error",
        "Submit Error",
        "Other Error",
        "Wating",
      ],
      username: "",
      problemid: "",
      tableData: [{ runid: "123", oj: "hdu", username: "LLLLLL0420" }],
    };
  },
  methods: {
    query: function () {},
  },
};
</script>

<style>
</style>
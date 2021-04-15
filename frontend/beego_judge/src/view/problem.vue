<template>
  <div class="problem" style="margin-left: 20px">
    <el-row :gutter="4">
      <el-col :span="4">
        <el-select v-model="oj">
          <el-option
            v-for="item in oj_array"
            :key="item"
            :value="item"
          ></el-option>
        </el-select>
      </el-col>
      <el-col :span="4">
        <el-input v-model="pid" placeholder="problemid"></el-input>
      </el-col>
      <el-col :span="2">
        <el-button type="primary" @click="toProblem()">Query</el-button>
      </el-col>
    </el-row>
    <h1 style="margin-top:100px;">Supports the following online judges</h1>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column label="OJ" prop="oj"> </el-table-column>
      <el-table-column label="Provider" prop="provider"> </el-table-column>
      <el-table-column label="github">
        <template slot-scope="scope">
          <el-link
            type="primary"
            :underline="false"
            @click="toWeb(scope.row.github)"
            >{{ scope.row.github }}</el-link
          >
        </template>
      </el-table-column>
      <el-table-column label="OJ Src">
        <template slot-scope="scope">
          <el-link
            type="primary"
            :underline="false"
            @click="toWeb(scope.row.url)"
            >{{ scope.row.url }}</el-link
          >
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
 
<script>
export default {
  name: "problem",
  data() {
    return {
      pid: "",
      oj: "HDU",
      oj_array: ["HDU"],
      tableData: [
        {
          oj: "HDU",
          provider: "PHK",
          github: "https://github.com/PHK-20",
          url: "http://acm.hdu.edu.cn/",
        },
      ],
    };
  },
  mounted: function () {},
  methods: {
    toWeb(url) {
      window.open(url, "_blank");
    },
    toProblem() {
      let routeUrl = this.$router.resolve({
        path: "/problem/" + this.oj + "/" + this.pid,
      });
      window.open(routeUrl.href, "_blank");
    },
  },
};
</script>

<style>
.el-select {
  width: 100%;
}
.el-button {
  width: 100%;
}
</style>
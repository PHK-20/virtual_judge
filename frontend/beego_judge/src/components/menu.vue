<template>
  <div class="menu">
    <el-row>
      <el-col :span="3">
        <el-menu
          :default-active="index"
          class="el-menu-vertical-demo"
          @select="handleSelect"
        >
          <el-menu-item index="problem">
            <span slot="title">问题</span>
          </el-menu-item>
          <el-menu-item index="status">
            <span slot="title">状态</span>
          </el-menu-item>
          <el-menu-item index="match">
            <span slot="title">比赛</span>
          </el-menu-item>
        </el-menu>
      </el-col>
      <el-col :span="21">
        <status v-show="index === 'status'" ref="status" @toProblem="toProblem"></status>
        <problem v-show="index === 'problem'" ref="problem"></problem>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import status from "@/view/status";
import problem from "@/view/problem";
export default {
  components: {
    status,
    problem,
  },
  data() {
    return {
      index: "problem",
    };
  },
  methods: {
    toProblem(oj, pid) {
      this.index = "problem";
      this.$refs.problem.toProblem(oj, pid);
    },
    handleSelect(index, indexPath) {
      this.index = index;
      if (index == "status") {
          this.$refs.status.query()
      }
    },
  },
};
</script>

<style>
.a {
}
</style>
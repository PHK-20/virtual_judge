<template>
  <div class="main">
    <el-row>
      <el-col :span="3">
        <mainMenu @menuIndex="menuIdx"></mainMenu>
      </el-col>
      <el-col :span="21">
        <status
          v-show="index === 'status'"
          ref="status"
          @toProblem="toProblem"
        ></status>
        <problem v-show="index === 'problem'" ref="problem"></problem>
        <user v-show="index === 'user'" ref="user"></user>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import mainMenu from "@/components/mainMenu";
import status from "@/view/status";
import problem from "@/view/problem";
import user from "@/view/user";
export default {
  components: { mainMenu, status, problem, user },
  data() {
    return {
      index: "problem",
    };
  },
  methods: {
    menuIdx: function (idx) {
      this.index = idx;
      if (idx == "status") {
        this.$refs.status.query();
      }
    },
    toProblem(oj, pid) {
      this.index = "problem";
      this.$refs.problem.toProblem(oj, pid);
    },
  },
};
</script>

<style>
</style>
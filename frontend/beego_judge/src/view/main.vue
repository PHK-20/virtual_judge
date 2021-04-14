<template>
  <div class="main">
    <login @setUser="setUser" ref="login"></login>
    <el-row>
      <el-col :span="3">
        <mainMenu ref="mainMenu" @menuIndex="menuIdx"></mainMenu>
      </el-col>
      <el-col :span="21">
        <status
          v-show="index === 'status'"
          ref="status"
          @toProblem="toProblem"
        ></status>
        <problem
          :name="user.name"
          v-show="index === 'problem'"
          ref="problem"
        ></problem>
        <matchList v-show="index === 'match'" ref="match" :name="user.name">
        </matchList>
        <user
          @logout="logout"
          @setIndex="setIndex"
          v-show="index === 'user'"
          ref="user"
        ></user>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import mainMenu from "@/components/mainMenu";
import status from "@/view/status";
import problem from "@/view/problem";
import user from "@/view/user";
import matchList from "@/view/matchList";
import login from "@/components/login";
export default {
  components: { mainMenu, status, problem, user, login, matchList },
  data() {
    return {
      index: "problem",
      user: {
        name: "",
        nickname: "",
        registerTime: "",
      },
    };
  },
  mounted: function () {
    if (!this.$refs.login.login()) {
      this.logout();
    }
  },
  methods: {
    logout() {
      this.$refs.mainMenu.setUser("login");
    },
    setIndex(idx) {
      this.index = idx;
    },
    menuIdx(idx) {
      console.log(idx);
      if (idx == "login") {
        this.$refs.login.showLogin();
        return;
      }
      this.index = idx;
      if (idx == "status") {
        this.$refs.status.query();
      }
      if (idx == "match") {
        this.$refs.match.query();
      }
    },
    toProblem(oj, pid) {
      this.index = "problem";
      this.$refs.problem.toProblem(oj, pid);
    },
    setUser(user) {
      console.log(user);
      this.user = user;
      this.$refs.mainMenu.setUser(user.name);
      this.$refs.user.setUser(user);
    },
  },
};
</script>

<style>
</style>
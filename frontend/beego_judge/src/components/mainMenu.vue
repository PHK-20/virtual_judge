<template>
  <div class="mainMenu">
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

      <el-menu-item index="user">
        <span slot="title">{{ user }}</span>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script>
export default {
  data() {
    return {
      user: "登录",
      username: "",
      password: "",
    };
  },
  methods: {
    handleSelect(index, indexPath) {
      this.index = index;
      this.$emit("menuIndex", index);
    },
    login() {
      this.axios
        .post("/login", { user: this.username, password: this.password })
        .then((res) => {
          console.log(res);
          this.user = this.username;
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
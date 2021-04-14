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
      <el-menu-item index="login" v-if="title === 'login'">
        <span slot="title">登录</span>
      </el-menu-item>
      <el-menu-item index="user" v-else>
        <span slot="title">{{ title }}</span>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script>
export default {
  data() {
    return {
      title: "",
      index: "",
    };
  },
  methods: {
    handleSelect(index, indexPath) {
      this.index = index;
      this.$emit("menuIndex", index);
    },
    setUser(title) {
      this.title = title;
    },
  },
};
</script>

<style>
</style>
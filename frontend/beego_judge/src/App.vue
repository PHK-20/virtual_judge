<!-- 展示模板 -->
<template>
  <div id="app">
    <submit></submit>
  </div>
</template>
 
<script>
// 导入组件
import submit from "./components/submit";

export default {
  name: "app",
  components: {
    submit,
  },
};
</script>
<!-- 样式代码 -->
<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
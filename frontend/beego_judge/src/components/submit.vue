<template>
  <div class="submit">
    <el-input v-model="problemid" placeholder="problemid"></el-input>
    <el-select v-model="languge" placeholder="请选择语言">
      <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      >
      </el-option>
    </el-select>
    <el-button type="primary" v-on:click="submit">Submit</el-button>
    <el-input
      type="textarea"
      :autosize="{ minRows: 10 }"
      placeholder="请输入内容"
      v-model="code"
    >
    </el-input>
  </div>
</template>
 
<script>
export default {
  name: "submit",
  data() {
    return {
      problemid: "",
      code: "",
      options: [
        {
          value: 0,
          label: "G++",
        },
        {
          value: 1,
          label: "GCC",
        },
        {
          value: 2,
          label: "python",
        },
      ],
      languge: 0,
    };
  },
  methods: {
    submit: function () {
      this.code = "";
      this.$axios({
        method: "post",
        url: "/submit",
        data: {
          problemid: this.problemid,
          code: this.code,
          languge: this.languge,
        },
      });
    },
  },
};
</script>
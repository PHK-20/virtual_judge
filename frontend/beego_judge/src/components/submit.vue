<template>
  <div class="submit">
    <el-input v-model="problemid" placeholder="problemid"></el-input>
    <el-select v-model="language" placeholder="请选择语言">
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
      v-model="usercode"
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
      usercode: "",
      options: [
        {
          value: "G++",
          label: "G++",
        },
        {
          value: "GCC",
          label: "GCC",
        },
        {
          value: "C++",
          label: "C++",
        },
      ],
      language: "G++",
    };
  },
  methods: {
    submit: function () {
      this.$axios({
        method: "post",
        url: "/submit",
        data: {
          problemid: this.problemid,
          usercode: this.usercode,
          language: this.language,
        },
      });
      console.log(this.usercode);
      this.usercode = "";
    },
  },
};
</script>
<template>
  <div class="problem">
    <el-row :gutter="3">
      <el-col :span="4" offset="2">
        <el-input v-model="problemid" placeholder="problemid"></el-input>
      </el-col>
      <el-col :span="4">
        <el-select v-model="oj" placeholder="oj">
          <el-option
            v-for="item in oj_array"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          >
          </el-option>
        </el-select>
      </el-col>
      <el-col :span="3">
        <el-button type="primary" @click="query">Query</el-button>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="20" :offset="2">
        <h1>{{ title }}</h1>
      </el-col>
    </el-row>
    <el-row class="desc">
      <el-col :span="4" offset="2">Description</el-col>
    </el-row>
    <el-row>
      <el-col :span="20" offset="2">
        <el-card class="box-card">
          <p class="content">{{ description }}</p>
        </el-card>
      </el-col>
    </el-row>
    <el-row class="desc">
      <el-col :span="4" offset="2">Input</el-col>
    </el-row>
    <el-row>
      <el-col :span="20" offset="2">
        <el-card class="box-card"> {{ input }}</el-card>
      </el-col>
    </el-row>
    <el-row class="desc">
      <el-col :span="4" offset="2">Output</el-col>
    </el-row>
    <el-row>
      <el-col :span="20" offset="2">
        <el-card class="box-card">{{ output }} </el-card>
      </el-col>
    </el-row>
    <el-row class="desc">
      <el-col :span="4" offset="2">Sample Input</el-col>
    </el-row>
    <el-row>
      <el-col :span="20" offset="2">
        <el-card class="box-card">{{ sampleInput }} </el-card>
      </el-col>
    </el-row>
    <el-row class="desc">
      <el-col :span="4" offset="2">Sample Output</el-col>
    </el-row>
    <el-row>
      <el-col :span="20" offset="2">
        <el-card class="box-card">{{ sampleOutput }} </el-card>
      </el-col>
    </el-row>
  </div>
</template>
 
<script>
export default {
  name: "submit",
  data() {
    return {
      problemid: "1000",
      oj_array: [
        {
          value: "HDU",
        },
      ],
      oj: "HDU",
      title: "A+B Problem",
      description: "123\n213",
      input: "",
      output: "",
      sampleInput: "",
      sampleOutput: "",
    };
  },
  methods: {
    query: function () {
      this.$axios
        .get("/problem", {
          params: {
            problemid: this.problemid,
            oj: this.oj,
          },
        })
        .then((resp) => {
          if (resp.data.Status == "success") {
            console.log(resp.data.Data.ProblemInfo);
            var info = {};
            info = resp.data.Data.ProblemInfo;
            this.title = info.Title;
            this.description = info.Description;
            this.input = info.Input;
            this.output = info.Output;
            this.sampleInput = info.SampleInput;
            this.sampleOutput = info.SampleOutput;
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
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


<style>
.el-select {
  width: 100%;
}

.el-button {
  width: 100%;
}

.box-card {
  margin-top: 10px;
  text-align: left;
}

.content {
  margin: 0px;
  white-space: pre-wrap;
}
.desc {
  margin-top: 20px;
  text-align: left;
}
</style>
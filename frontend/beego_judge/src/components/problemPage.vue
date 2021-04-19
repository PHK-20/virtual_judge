<template>
  <div class="problem">
    <login ref="login"></login>
    <el-row>
      <h1>{{ problem.title }}</h1>
      <div>
        Time Limit: {{ problem.timeLimit }}<br />
        Memory Limit: {{ problem.memoryLimit }}
      </div>
      <div class="desc">Description</div>
      <el-card class="box-card">
        <p class="content">{{ problem.description }}</p>
      </el-card>
      <div class="desc">Input</div>
      <el-card class="box-card">
        <p class="content">{{ problem.input }}</p>
      </el-card>
      <div class="desc">Output</div>
      <el-card class="box-card">
        <p class="content">{{ problem.output }}</p>
      </el-card>
      <div class="desc">Sample Input</div>
      <el-card class="box-card">
        <p class="content">{{ problem.sampleInput }}</p>
      </el-card>
      <div class="desc">Sample Output</div>
      <el-card class="box-card">
        <p class="content">{{ problem.sampleOutput }}</p>
      </el-card>
      <div class="desc" v-if="problem.hint">Hint</div>
      <el-card class="box-card" v-if="problem.hint">
        <p class="content">{{ problem.hint }}</p>
      </el-card>
      <div class="desc" v-if="problem.src">Source</div>
      <el-card class="box-card" v-if="problem.src">
        <p class="content">{{ problem.src }}</p>
      </el-card>
      <el-row :gutter="3" style="margin-top: 30px">
        <el-col :span="4">
          <el-select v-model="problem.language" placeholder="language">
            <el-option v-for="item in lang_array" :key="item" :value="item">
            </el-option>
          </el-select>
        </el-col>
        <el-col :span="3">
          <el-button type="primary" @click="submit">Submit</el-button>
        </el-col>
      </el-row>
      <el-input
        type="textarea"
        :autosize="{ minRows: 20 }"
        placeholder="code"
        v-model="usercode"
        style="margin-top: 20px"
      >
      </el-input>
    </el-row>
  </div>
</template>
 
<script>
import login from "@/components/login";
export default {
  components: {
    login,
  },
  props: {
    oj: String,
    pid: String,
  },
  name: "problemTab",
  data() {
    return {
      oj_array: ["HDU"],
      problem: {
        id: "",
        oj: "",
        language: "",
        title: "HelloWorld",
        description: "",
        input: "",
        output: "",
        sampleInput: "",
        sampleOutput: "",
        timeLimit: "",
        memoryLimit: "",
        hint: "",
        src: "",
        matchid: 0,
        matchidx: "#",
      },
      lang_array: [],
      usercode: "",
      needLanguage: true,
      notifications: {},
      loading: {},
    };
  },
  created: function () {
    this.loading = this.$loading({
      text: "Loading",
      spinner: "el-icon-loading",
    });
    console.log(this.$route);
    if (this.$route.params.matchid) {
      this.problem.matchid = Number(this.$route.params.matchid);
      let info = JSON.parse(this.$route.params.info);
      console.log(info);
      this.problem.id = info.pid;
      this.problem.oj = info.oj;
      this.problem.matchidx = info.idx;
    } else {
      this.problem.id = this.$route.params.pid;
      this.problem.oj = this.$route.params.oj;
    }
    this.queryProblem();
  },
  methods: {
    queryProblem: function () {
      console.log("querProblem " + this.problem.oj + " " + this.problem.id);
      if (this.problem.id == "") {
        this.$notify.error({
          title: "Error",
          message: "problemid empty",
        });
        return;
      }
      this.$axios
        .get("/problem", {
          params: {
            problemid: this.problem.id,
            oj: this.problem.oj,
            needLanguage: this.needLanguage,
          },
        })
        .then((resp) => {
          if (resp.data.Status == "success") {
            let info = {};
            info = resp.data.Data.ProblemInfo;
            this.problem.title = info.Title;
            this.problem.description = info.Description;
            this.problem.input = info.Input;
            this.problem.output = info.Output;
            this.problem.sampleInput = info.SampleInput;
            this.problem.sampleOutput = info.SampleOutput;
            this.problem.hint = info.Hint;
            this.problem.timeLimit = info.TimeLimit;
            this.problem.memoryLimit = info.MemoryLimit;
            this.problem.src = info.Src;
            if (this.needLanguage) {
              this.lang_array = [];
              resp.data.Data.Language.forEach((v) => {
                this.lang_array.push(v);
              });
              this.problem.language = this.lang_array[0];
              this.needLanguage = false;
            }
            this.loading.close();
            this.$emit("title", this.problem.oj + "-" + this.problem.id);
          } else {
            this.$notify.error({
              title: "Error",
              message: resp.data.ErrorMsg,
            });
          }
        })
        .catch((err) => {
          this.notifications[runid].message = "Request Server Error";
          this.notifications[runid].type = "error";
          console.error(err);
        });
    },
    submit: function () {
      if (this.problem.language == "") {
        this.$notify.error({
          title: "Error",
          message: "language is empty",
        });
      }
      this.$axios({
        method: "post",
        url: "/submit",
        data: {
          problem: this.problem,
          username: this.username,
          usercode: this.usercode,
        },
      }).then((resp) => {
        if (resp.data.Status == "fail") {
          if (resp.data.ErrorMsg == "Login Firstly") {
            this.$refs.login.showLogin();
          }
          this.$notify.error({
            title: "Error",
            message: resp.data.ErrorMsg,
          });
        } else {
          let runid = resp.data.Data.Runid;
          this.notifications[runid] = this.$notify({
            title: "Solution",
            message: this.notifyMsg(runid, "submited"),
            status: "submited",
            type: "info",
            duration: 0,
            dangerouslyUseHTMLString: true,
          });
          this.queryResult(runid);
        }
      });
    },
    queryResult: function (runid) {
      console.log(runid + " query");
      this.$axios
        .get("/result", {
          params: {
            runid: runid,
          },
        })
        .then((resp) => {
          if (resp.data.Status == "success") {
            let result = resp.data.Data.Result.Res;
            if (result != this.notifications[runid].status) {
              if (resp.data.Data.IsFinalResult == true) {
                console.log(runid + " query finish");
                if (result == "Accepted") {
                  this.notifications[runid].type = "success";
                } else {
                  this.notifications[runid].type = "error";
                }
              }
              this.notifications[runid].status = result;
              this.notifications[runid].message = this.notifyMsg(runid, result);
            }
            if (!resp.data.Data.IsFinalResult) {
              setTimeout(() => {
                this.queryResult(runid);
              }, 3000);
            }
          } else {
            this.$notify.error({
              title: "Error",
              message: resp.data.ErrorMsg,
            });
          }
        })
        .catch((err) => {
          this.notifications[runid].message = "Request Server Error";
          this.notifications[runid].type = "error";
          console.error(err);
        });
    },
    notifyMsg: function (runid, status) {
      let msg =
        "<span>Runid:" +
        runid +
        "</span><br>" +
        "<span>Status:" +
        status +
        "</span>";
      return msg;
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
.el-notification {
  white-space: pre-wrap !important;
}
</style>
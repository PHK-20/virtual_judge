<template>
  <div class="problem">
    <el-row>
      <el-col :span="20">
        <el-row :gutter="3">
          <el-col :span="4">
            <el-input v-model="problemid" placeholder="problemid"></el-input>
          </el-col>
          <el-col :span="4">
            <el-select
              v-model="oj"
              placeholder="oj"
              @change="needLanguage = true"
            >
              <el-option v-for="item in oj_array" :key="item" :value="item">
              </el-option>
            </el-select>
          </el-col>
          <el-col :span="3">
            <el-button type="primary" @click="queryProblem">Query</el-button>
          </el-col>
        </el-row>
        <h1>{{ title }}</h1>
        <div class="desc">Description</div>
        <el-card class="box-card">
          <p class="content">{{ description }}</p>
        </el-card>
        <div class="desc">Input</div>
        <el-card class="box-card">
          <p class="content">{{ input }}</p>
        </el-card>
        <div class="desc">Output</div>
        <el-card class="box-card">
          <p class="content">{{ output }}</p>
        </el-card>
        <div class="desc">Sample Input</div>
        <el-card class="box-card">
          <p class="content">{{ sampleInput }}</p>
        </el-card>
        <div class="desc">Sample Output</div>
        <el-card class="box-card">
          <p class="content">{{ sampleOutput }}</p>
        </el-card>
        <div class="desc" v-if="hint">Hint</div>
        <el-card class="box-card" v-if="hint">
          <p class="content">{{ hint }}</p>
        </el-card>
        <el-row :gutter="3" style="margin-top: 30px">
          <el-col :span="4">
            <el-select v-model="language" placeholder="language">
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
      </el-col>
    </el-row>
  </div>
</template>
 
<script>
export default {
  name: "problem",
  data() {
    return {
      problemid: "1000",
      oj_array: ["HDU"],
      oj: "HDU",
      title: "Hello World",
      description: "",
      input: "",
      output: "",
      sampleInput: "",
      sampleOutput: "",
      hint: "",
      lang_array: [],
      language: "",
      usercode: "",
      needLanguage: true,
      username: "LLLLLL0420",
      notifications: {},
    };
  },
  methods: {
    queryProblem: function () {
      if (this.problemid == "") {
        this.$notify.error({
          title: "Error",
          message: "problemid empty",
        });
        return;
      }
      this.$axios
        .get("/problem", {
          params: {
            problemid: this.problemid,
            oj: this.oj,
            needLanguage: this.needLanguage,
          },
        })
        .then((resp) => {
          if (resp.data.Status == "success") {
            let info = {};
            info = resp.data.Data.ProblemInfo;
            this.title = info.Title;
            this.description = info.Description;
            this.input = info.Input;
            this.output = info.Output;
            this.sampleInput = info.SampleInput;
            this.sampleOutput = info.SampleOutput;
            this.hint = info.Hint;
            if (this.needLanguage) {
              this.lang_array = [];
              resp.data.Data.Language.forEach((v) => {
                this.lang_array.push(v);
              });
              this.language = this.lang_array[0];
              this.needLanguage = false;
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
    submit: function () {
      if (this.language == "") {
        this.$notify.error({
          title: "Error",
          message: "language is empty",
        });
      }
      this.$axios({
        method: "post",
        url: "/submit",
        data: {
          problemid: this.problemid,
          usercode: this.usercode,
          language: this.language,
          username: this.username,
          oj: this.oj,
        },
      }).then((resp) => {
        if (resp.data.Status == "fail") {
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
            let result = resp.data.Data.Result;
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
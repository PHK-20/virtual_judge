<template>
  <div class="problem">
    <el-row>
      <el-col :span="20" :offset="2">
        <el-tabs v-model="activeTab" type="card">
          <el-tab-pane label="问题" name="problemTab">
            <el-row :gutter="3">
              <el-col :span="4">
                <el-input
                  v-model="problemid"
                  placeholder="problemid"
                ></el-input>
              </el-col>
              <el-col :span="4">
                <el-select
                  v-model="oj"
                  placeholder="oj"
                  @change="needLanguage = true"
                >
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
                <el-button type="primary" @click="queryProblem"
                  >Query</el-button
                >
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
            <div class="desc">Hint</div>
            <el-card class="box-card">
              <p class="content">{{ hint }}</p>
            </el-card>
          </el-tab-pane>
          <el-tab-pane label="提交" name="submitTab">
            <el-row :gutter="3">
              <el-col :span="4">
                <el-select v-model="language" placeholder="language">
                  <el-option
                    v-for="item in lang_array"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                  >
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
          </el-tab-pane>
        </el-tabs>
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
      oj_array: [
        {
          value: "HDU",
        },
        {
          value: "HNUST",
        },
      ],
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
      activeTab: "submitTab",
      needLanguage: true,
      username: "LLLLLL0420",
      notifications: {},
    };
  },
  mounted: function () {
    this.queryProblem();
  },
  methods: {
    queryProblem: function () {
      console.log(this.needLanguage);
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
          console.log(resp);
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
                this.lang_array.push({ value: v });
              });
              this.needLanguage = false;
            }
          } else {
            this.$notify.error({
              title: "Error",
              message: resp.data.ErrorMsg,
            });
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
          username: this.username,
          oj: this.oj,
        },
      }).then((resp) => {
        console.log(resp);
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
      this.$axios
        .get("/queryResult", {
          params: {
            runid: runid,
          },
        })
        .then((resp) => {
          if (resp.data.Status == "success") {
            console.log(resp.data);
            let result = resp.data.Data.Result;
            if (result != this.notifications[runid].status) {
              if (result == "Accepted") {
                this.notifications[runid].type = "success";
              } else if (resp.data.Data.IsFinalResult == true) {
                this.notifications[runid].type = "error";
              }
              this.notifications[runid].status = result;
              this.notifications[runid].message = this.notifyMsg(runid, result);
            }
            if (!resp.data.Data.IsFinalResult) {
              setTimeout(() => {
                this.queryResult(runid);
              }, 1000);
            }
          } else {
            this.$notify.error({
              title: "Error",
              message: resp.data.ErrorMsg,
            });
          }
        })
        .catch((err) => {
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
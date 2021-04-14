<template>
  <div class="matchList" style="margin-left: 10px">
    <el-dialog title="Create Match" :visible.sync="visible" width="40%">
      <el-form ref="form" :model="form" label-width="90px" :rules="rules">
        <el-form-item label="Title" prop="title">
          <el-input v-model="form.title"></el-input>
        </el-form-item>
        <el-form-item label="Description" prop="desc">
          <el-input v-model="form.desc"></el-input>
        </el-form-item>
        <el-form-item label="比赛时间" prop="dataTime">
          <el-date-picker
            v-model="form.contestTime"
            format="yyyy-MM-dd HH:mm:ss"
            value-format="yyyy-MM-dd HH:mm:ss"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
          >
          </el-date-picker>
        </el-form-item>
        <el-form-item label="Add">
          <el-col :span="4">
            <el-button type="primary" @click="addProblem">Problem</el-button>
          </el-col>
        </el-form-item>
        <el-form-item
          v-for="(problem, index) in form.problem"
          :label="'问题' + index"
          :key="index"
          required
        >
          <el-row :gutter="2">
            <el-col :span="6">
              <el-form-item>
                <el-select v-model="problem.oj">
                  <el-option
                    v-for="item in oj"
                    :key="item.value"
                    :value="item.value"
                  >
                  </el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item>
                <el-input
                  @change="checkProblem(index)"
                  v-model="problem.pid"
                  placeholder="problemId"
                >
                </el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-button type="danger" @click="removeProblem(problem)"
                >删除</el-button
              >
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item>
          <el-row>
            <el-col :span="4">
              <el-button :icon="icon" type="primary" @click="submitForm('form')"
                >Confirm</el-button
              >
            </el-col>
          </el-row>
        </el-form-item>
      </el-form>
    </el-dialog>
    <el-row>
      <el-col :span="2" style="margin-bottom: 10px">
        <el-button type="primary" @click="createContest"
          >Create Contest</el-button
        >
      </el-col>
    </el-row>
    <el-table :data="tableData" stripe style="width: 100%" v-loading="loading">
      <el-table-column align="center" label="MatchId" prop="MatchId">
      </el-table-column>
      <el-table-column align="center" prop="Title">
        <template slot="header" slot-scope="scope">
          <div>Title</div>
          <el-input v-model="condition.title" @change="query()" size="mini" />
        </template>
        <template slot-scope="scope">
          <el-link
            type="primary"
            :underline="false"
            @click="toMatch(scope.row.MatchId)"
            >{{ scope.row.Title }}</el-link
          >
        </template>
      </el-table-column>
      <el-table-column align="center" label="Begin Time" prop="BeginTime">
      </el-table-column>
      <el-table-column align="center" label="EndTime" prop="EndTime">
      </el-table-column>
      <el-table-column align="center" prop="Onwer">
        <template slot="header" slot-scope="scope">
          <div>Onwer</div>
          <el-input v-model="condition.onwer" @change="query()" size="mini" />
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      style="margin-top: 20px"
      background
      layout="sizes, prev, pager, next"
      :total="total"
      :current-page="currentPage"
      :page-sizes="[10, 20, 50, 100]"
      :page-size="pageSize"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    >
    </el-pagination>
  </div>
</template>

<script>
export default {
  props: {
    name: String,
  },
  data() {
    return {
      icon: "el-icon-check",
      visible: false,
      loading: false,
      pageSize: 10,
      currentPage: 1,
      total: 0,
      condition: {
        title: "",
        onwer: "",
      },
      tableData: [],
      oj: [
        {
          value: "HDU",
        },
      ],
      form: {
        problem: [
          {
            oj: "HDU",
            pid: "",
            res: false,
          },
        ],
        desc: "",
        title: "",
        contestTime: "",
      },
      rules: {
        title: [{ required: true, trigger: "blur", max: 32 }],
        desc: [{ trigger: "blur", max: 128 }],
      },
    };
  },
  methods: {
    createContest() {
      if (!this.name) {
        this.$notify.error({
          title: "Error",
          message: "Login First",
        });
      } else {
        this.visible = true;
      }
    },
    removeProblem(item) {
      var index = this.form.problem.indexOf(item);
      if (index !== -1) {
        this.form.problem.splice(index, 1);
      }
    },
    addProblem() {
      this.form.problem.push({ oj: "HDU", pid: "" });
    },
    handleCurrentChange(page) {
      this.currentPage = page;
      this.query();
    },
    handleSizeChange(size) {
      this.pageSize = size;
      this.query();
    },
    query() {
      this.loading = true;
      console.log("querMatch");
      this.$axios
        .get("/matchList", {
          params: {
            num: this.pageSize,
            offset: (this.currentPage - 1) * 10,
            condition: this.condition,
          },
        })
        .then((resp) => {
          if (resp.data.Status == "success") {
            this.tableData = [];
            let result = resp.data.Data;
            result.MatchItem.forEach((item) => {
              item.BeginTime = new Date(item.BeginTime).toLocaleString();
              item.EndTime = new Date(item.EndTime).toLocaleString();
              this.tableData.push(item);
            });
            this.total = result.Total;
          }
          this.loading = false;
        })
        .catch((err) => {
          console.error(err);
        });
    },
    toMatch(id) {
      let routeUrl = this.$router.resolve({
        path: "/match/" + id,
      });
      window.open(routeUrl.href, "_blank");
    },
    checkProblem(idx) {
      this.icon = "el-icon-loading";
      this.$axios
        .get("/problem", {
          params: {
            problemid: this.form.problem[idx].pid,
            oj: this.form.problem[idx].oj,
          },
        })
        .then((resp) => {
          let data = resp.data;
          console.log(resp);
          if (data.Status == "fail") {
            this.$notify.error({
              title: "Error",
              message: "问题" + idx + " 不存在",
            });
          } else {
            this.form.problem[idx].res = true;
          }
          this.icon = "el-icon-check";
        })
        .catch((err) => {
          console.error(err);
        });
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          if (!this.form.contestTime) {
            this.$notify.error({
              title: "Error",
              message: "contest time is empty",
            });
            return;
          }
          let flag = true;
          console.log(this.form.problem);
          this.form.problem.forEach((item, idx) => {
            if (!item.res) {
              this.$notify.error({
                title: "Error",
                message: "问题" + idx + " 不存在",
              });
              flag = false;
            }
          });
          if (flag) {
            console.log(this.form);
            this.$axios
              .post("/createContest", this.form)
              .then((resp) => {
                console.log(resp);
                let data = resp.data;
                if (data.Status == "success") {
                  this.$notify.success({
                    title: "Success",
                    message: "Create Contest Success",
                  });
                  this.visible = false;
                } else {
                  this.$notify.error({
                    title: "Error",
                    message: data.ErrorMsg,
                  });
                }
              })
              .catch((err) => {
                console.error(err);
              });
          }
        }
      });
    },
  },
};
</script>

<style>
</style>
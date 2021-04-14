<template>
  <div class="login">
    <el-dialog title="Login" :visible.sync="loginVisible" width="30%">
      <el-form
        style="text-align: left"
        ref="loginForm"
        :model="loginForm"
        :rules="loginRules"
        status-icon
      >
        <div style="margin-top: -20px">
          <el-form-item prop="username">
            <el-input
              v-model="loginForm.username"
              placeholder="username"
            ></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              type="password"
              v-model="loginForm.password"
              placeholder="password"
            ></el-input>
          </el-form-item>
        </div>
      </el-form>
      <el-row :gutter="10">
        <el-col :span="4">
          <el-button type="primary" @click="login">登录</el-button>
        </el-col>
        <el-col :span="4">
          <el-button @click="showRegister">注册</el-button>
        </el-col>
      </el-row>
    </el-dialog>
    <el-dialog title="Register" :visible.sync="registerVisible" width="30%">
      <div style="margin-top: -20px">
        <el-form
          ref="registerForm"
          :model="registerForm"
          :rules="registerRules"
          status-icon
        >
          <el-form-item prop="username">
            <el-input
              v-model="registerForm.username"
              placeholder="username (can not change)"
            ></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              type="password"
              v-model="registerForm.password"
              placeholder="password"
            ></el-input>
          </el-form-item>
          <el-form-item prop="repeatPw">
            <el-input
              type="password"
              v-model="registerForm.repeatPw"
              placeholder="repeat password"
            ></el-input>
          </el-form-item>
          <el-form-item prop="nickname">
            <el-input
              v-model="registerForm.nickname"
              placeholder="nickname"
            ></el-input>
          </el-form-item>
        </el-form>
      </div>
      <el-button type="primary" @click="register">注册</el-button>
    </el-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      loginVisible: false,
      registerVisible: false,
      user: {
        name: "",
        nickname: "",
        registerTime: "",
      },
      loginForm: {
        username: "",
        password: "",
      },
      registerForm: {
        username: "",
        password: "",
        repeatPw: "",
        nickname: "",
      },
      registerRules: {
        username: [{ required: true, trigger: "blur", max: 20 }],
        nickname: [{ required: true, trigger: "blur", max: 20 }],
        password: [{ required: true, trigger: "blur", min: 8, max: 32 }],
        repeatPw: [{ required: true, trigger: "blur", min: 8, max: 32 }],
      },
      loginRules: {
        username: [{ required: true, trigger: "blur" }],
        password: [{ required: true, trigger: "blur" }],
      },
    };
  },
  methods: {
    showRegister() {
      this.loginVisible = false;
      this.registerVisible = true;
    },
    showLogin() {
      this.loginVisible = true;
    },
    login() {
      this.$axios
        .post("/login", this.loginForm)
        .then((resp) => {
          let data = resp.data;
          if (data.Status == "success") {
            this.loginVisible = false;
            this.user.name = data.Data.Username;
            this.user.nickname = data.Data.Nickname;
            this.user.registerTime = data.Data.RegisterTime;
            this.setUser(this.user);
            return true;
          } else {
            if (this.loginForm.username) {
              this.$notify.error({
                title: "Error",
                message: data.ErrorMsg,
              });
            }
            return false;
          }
        })
        .catch((err) => {
          console.error(err);
        });
    },
    register() {
      if (this.registerForm.password != this.registerForm.repeatPw) {
        this.$notify.error({
          title: "Error",
          message: "密码与二次密码不一致",
        });
        return;
      }
      this.$axios
        .post("/register", this.registerForm)
        .then((resp) => {
          let data = resp.data;
          if (data.Status == "fail") {
            this.$notify.error({
              title: "Error",
              message: data.ErrorMsg,
            });
          } else {
            this.$notify.success({
              title: "Success",
              message: "Register Success",
            });
            this.registerVisible = false;
            this.loginForm.username = this.registerForm.username;
            this.loginForm.password = this.registerForm.password;
            this.login();
          }
        })
        .catch((err) => {
          console.error(err);
        });
    },
    setUser(user) {
      this.$emit("setUser", user);
    },
  },
};
</script>

<style>
.el-form {
  text-align: left;
}
</style>
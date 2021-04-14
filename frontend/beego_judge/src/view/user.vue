<template>
  <div class="user">
    <el-row>
      <el-col :offset="1" :span="22">
        <el-card class="box-card">
          <div class="card_head">{{ user.name }}</div>
          <el-divider></el-divider>
          <div class="card_item">nickname: {{ user.nickname }}</div>
          <div class="card_item">RegisterTime: {{ user.registerTime }}</div>
          <el-row>
            <el-col :span="2" style="margin-top: 10px">
              <el-button type="danger" @click="logout">Logout</el-button>
            </el-col>
          </el-row>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      user: {
        name: "",
        nickname: "",
        registerTime: "",
      },
    };
  },
  methods: {
    logout() {
      this.$axios
        .get("/logout", {})
        .then((res) => {
          console.log(res);
          this.$emit("logout");
          this.$emit("setIndex", "problem");
        })
        .catch((err) => {
          console.error(err);
        });
    },
    setUser(user) {
      this.user = user;
    },
  },
};
</script>

<style>
.card_item {
  font-size: 24px;
  margin-top: 10px;
}
.card_head {
  font-size: 30px;
}
</style>
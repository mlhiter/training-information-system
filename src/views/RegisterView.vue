<template>
  <div class="home">
    <div class="bg"></div>
    <div class="card_reg">
      <div class="top">注册</div>
      <div class="mid">
        <el-input
          prefix-icon="el-icon-user"
          v-model="input3"
          placeholder="账号"
          class="in"
        ></el-input>
        <el-input
          prefix-icon="el-icon-key"
          v-model="input4"
          placeholder="密码"
          class="in"
          show-password
        ></el-input>
        <el-input
          prefix-icon="el-icon-key"
          v-model="input5"
          placeholder="确认密码"
          class="in"
          show-password
        ></el-input>
        <el-button type="primary" class="but" @click="checkreg">注册</el-button>
        <el-button class="but2" @click="back">返回</el-button>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
export default {
  name: 'RegisterView',
  components: {},
  data() {
    return {
      input3: '',
      input4: '',
      input5: '',
    }
  },
  methods: {
    //检查注册信息是否合法，并返回
    checkreg() {
      if (this.input3.length == 0) this.$message.error('用户名不能为空')
      else if (this.input4.length == 0) this.$message.error('请输入密码')
      else if (this.input5.length == 0) this.$message.error('请确认密码')
      else if (this.input4 != this.input5) this.$message.error('密码不一致')
      else {
        //向服务器请求注册
        this.postreg()
        this.input3 = ''
        this.input4 = ''
        this.input5 = ''
        this.$message({
          message: '注册成功',
          type: 'success',
        })
        this.$router.push('/')
      }
    },
    //注册信息传给服务器
    postreg() {
      this.$axios({
        url: '/login/normal/register',
        method: 'post',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        data: {
          username: this.input3,
          password: this.input4,
        },
        transformRequest: [
          function (dat) {
            let ret = ''
            for (let it in dat) {
              ret +=
                encodeURIComponent(it) + '=' + encodeURIComponent(dat[it]) + '&'
            }
            ret = ret.substring(0, ret.lastIndexOf('&'))
            return ret
          },
        ],
      })
        .then((res) => {
          console.log(res)
        })
        .catch((error) => {
          this.$message.error(error)
        })
    },
    //按下返回键时触发
    back() {
      this.$router.push('/')
    },
  },
}
</script>

<style lang="scss">
.card_reg {
  display: block;
  width: 430px;
  height: 420px;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate3d(-50%, -50%, 0);
  z-index: 20;
  background-color: #fff;
  border: 1px rgb(200, 200, 200);
  border-radius: 10px;
  overflow: visible;
}
.but2 {
  width: 100%;
  text-align: center;
  margin-top: 20px !important;
  margin-left: 0px !important;
}
</style>

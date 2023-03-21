<template>
  <div class="home">
    <div class="bg"></div>
    <div class="card">
      <div class="top">登录</div>
      <div class="mid">
        <el-input
          prefix-icon="el-icon-user"
          v-model="input1"
          placeholder="账号"
          class="in"
        ></el-input>
        <el-input
          prefix-icon="el-icon-key"
          v-model="input2"
          placeholder="密码"
          class="in"
          show-password
        ></el-input>
        <span class="tip">
          没有账号？
          <el-link
            type="primary"
            :underline="false"
            @click="$router.push('register')"
          >
            立即注册
          </el-link>
        </span>
        <el-button type="primary" class="but" @click="confirmlogin">
          登录
        </el-button>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
export default {
  name: 'AuthView',
  components: {},
  data() {
    return {
      input1: '',
      input2: '',
    }
  },
  methods: {
    //检查注册信息是否合法
    confirmlogin() {
      if (this.input1.length == 0) this.$message.error('请输入用户名')
      else if (this.input2.length == 0) this.$message.error('请输入密码')
      else {
        this.postlogininfo()
      }
    },
    //请求登录
    postlogininfo() {
      this.$axios({
        url: '/login/normal/login',
        method: 'post',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        data: {
          username: this.input1,
          password: this.input2,
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
          if (res.msg == '成功') {
            this.$message.success('登录成功')
            //更新token
            this.$store.commit('updateToken', res.data)
            this.$store.commit('Token')
            //跳转页面
            this.$router.push('/manage/place')
          } else this.$message.error('登录失败')
        })
        .catch((error) => {
          this.$message.error(error)
        })
    },
  },
}
</script>

<style lang="scss">
.home {
  margin: 0px;
  padding: 0px;
}

.bg {
  display: block;
  position: fixed;
  left: 0px;
  top: 0px;
  width: 100%;
  height: 100%;
  background-image: url('../assets/bg.svg');
  background-repeat: no-repeat;
  background-size: cover;
  z-index: 2;
}

.card {
  display: block;
  width: 430px;
  height: 330px;
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

.top {
  text-align: center;
  font-size: 20px;
  margin-top: 30px;
  margin-bottom: 20px;
}

.mid {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  width: 80%;
}

.in {
  margin-bottom: 20px;
}

.tip {
  display: inline-flex;
  width: 100%;
  text-align: left;
  font-size: 14px;
}

.but {
  width: 100%;
  text-align: center;
  margin-top: 20px !important;
}
</style>

<template>
  <div class="main_container">
    <div class="navigation_container">
      <!-- 标题 -->
      <el-card class="title-box">
        <i class="el-icon-cloudy" id="topIco">&nbsp;&nbsp;管理后台</i>
      </el-card>
      <!-- 导航栏 -->
      <el-container height="66px" direction="horizontal" class="menu-container">
        <el-menu :default-active="activeIndex" mode="horizontal" router>
          <el-menu-item class="blank"></el-menu-item>
          <el-menu-item index="place" style="font-size: 17px; width: 200px">
            地点
          </el-menu-item>
          <el-menu-item index="state" style="font-size: 17px; width: 200px">
            用户
          </el-menu-item>
          <el-menu-item index="chat" style="font-size: 17px; width: 200px">
            消息
          </el-menu-item>
          <el-menu-item index="about" style="font-size: 17px; width: 200px">
            关于我们
          </el-menu-item>
          <el-menu-item @click="logout()" style="font-size: 17px; width: 200px">
            退出登录
          </el-menu-item>
        </el-menu>
      </el-container>
    </div>
    <!-- 内容栏：嵌套路由 -->
    <router-view @activeIndex="getActiveIndex"></router-view>
  </div>
</template>

<script>
export default {
  name: 'MainView',
  data() {
    return {
      activeIndex: '', //elementUi规定是字符串类型
    }
  },
  methods: {
    // 获取子页面路由：处理在子路由下刷新页面但是导航栏改变的bug
    getActiveIndex(index) {
      this.activeIndex = index
    },
    // 退出登录
    logout() {
      this.$confirm('确认退出登录吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }).then(() => {
        this.$store.commit('updateToken', '')
        this.$message.info('成功退出登录')
        this.$router.push('/')
      })
    },
  },
}
</script>
<style lang="scss">
.navigation_container {
  display: flex;
  flex-direction: row;
  .title-box {
    background-color: rgb(48, 65, 86) !important;
    width: 295px;
    margin-right: 4px;
  }
  .menu-container {
    padding: 0px !important;
    box-shadow: 0 2px 12px 0 rgb(0 0 0 / 10%);
    border: 1px solid #ebeef5;
    .blank {
      width: 1px;
    }
  }
  #topIco {
    font-size: 23px;
    color: #fff;
  }
}
</style>

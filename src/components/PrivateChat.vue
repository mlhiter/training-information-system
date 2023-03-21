<template>
  <el-container>
    <!-- 左侧用户列表 -->
    <el-aside class="aside">
      <el-menu @select="handleSelect">
        <el-menu-item
          v-for="item in userlist"
          :key="item.userId"
          :index="String(item.userId)"
        >
          <el-avatar class="avatarClass">{{ item.userId }}</el-avatar>
          {{ item.userWxName }}
        </el-menu-item>
      </el-menu>
      <el-pagination
        hide-on-single-page="true"
        class="pagination_userlist"
        @current-change="handleCurrentChange"
        :current-page="currentPage"
        :page-size="pageSize"
        layout="total, prev, pager, next"
        :total="totalPage"
      ></el-pagination>
    </el-aside>
    <!-- 右侧功能区 -->
    <el-container>
      <!-- 消息框 -->
      <el-main id="messageContain" ref="box">
        <div v-for="item in messageList" :key="item.id">
          <div v-if="item.senderId == selectId" class="messageReceive">
            <div class="receiveBubble">
              {{ item.mess }}
            </div>
            <div class="receiveTime">
              {{ item.time }}
            </div>
          </div>
          <div v-if="item.receiverId == selectId" class="messageSend">
            <div class="sendBubble">
              {{ item.mess }}
            </div>
            <div class="sendTime">
              {{ item.time }}
            </div>
          </div>
        </div>
        <div v-if="messageList.length == 0">
          <el-empty description="暂无内容"></el-empty>
        </div>
      </el-main>
      <!-- 输入框 -->
      <el-footer class="inputArea">
        <el-input
          class="message"
          v-model="content"
          clearable
          autosize
          type="textarea"
          v-on:keyup.enter="confirmSend"
        ></el-input>
        <el-button
          class="sendMessage"
          type="primary"
          plain
          @click="confirmSend"
        >
          发送
        </el-button>
      </el-footer>
    </el-container>
  </el-container>
</template>
<script>
export default {
  name: 'PrivateMessage',
  data() {
    return {
      userlist: [],
      messageList: [],
      currentPage: 1,
      pageSize: 10,
      totalPage: 0,
      content: '',
      selectId: -1,
    }
  },
  sockets: {
    connecting() {
      console.log('Socket 正在连接')
    },
    disconnect() {
      console.log('Socket 断开')
    },
    connect_error() {
      console.log('Socket 连接失败')
    },
    connect() {
      console.log('Socket 连接成功')
    },
  },
  mounted() {
    this.upUserList()
    setInterval(() => {
      this.upMessageList()
    }, 1000)
    // // 打开socket链接
    // this.$socket.open()
  },
  beforeDestroy() {
    this.$socket.close()
  },
  methods: {
    upUserList() {
      this.$axios({
        url: '/user/get/all/user',
        method: 'get',

        params: {
          num: this.currentPage,
          size: this.pageSize,
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
          console.log('服务器获取用户列表', res)
          this.userlist = res
          this.totalPage = parseInt(res.msg)
          // 滚动条放到底部（这里要异步执行）
          this.messagesContainerTimer = setTimeout(() => {
            this.$refs.box.scrollTop = 300
            // 清理定时器
            clearTimeout(this.messagesContainerTimer)
          }, 100)
        })
        .catch((error) => {
          console.error(error)
        })
    },
    handleCurrentChange(event) {
      this.currentPage = event
      this.upUserList()
    },
    handleSelect(event) {
      this.selectId = event
      this.upMessageList()
    },
    upMessageList() {
      if (this.selectId < 0) return
      this.$axios({
        url: '/message/get',
        method: 'post',

        params: {
          oid: this.selectId,
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
          // console.log('服务器获取消息列表', res)
          this.messageList = res.data
          this.messageList.reverse()
        })
        .catch((error) => {
          console.error(error)
        })
    },
    confirmSend() {
      this.$axios({
        url: '/message/send',
        method: 'post',

        params: {
          rid: this.selectId,
          message: this.content,
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
        .then(() => {
          this.content = ''
          setTimeout(() => {
            var pt = this.$el.querySelector('#messageContain')
            pt.scrollTop = pt.scrollHeight
            this.upMessageList()
          }, 300)
        })
        .catch((error) => {
          console.error(error)
        })
    },
  },
}
</script>
<style scoped>
.aside {
  margin-right: 0px;
}
.avatarClass {
  float: left;
  margin-top: 8px;
  background-color: rgb(123, 175, 214);
}
.pagination_userlist {
  position: absolute;
  left: 310px;
  top: 660px;
}
.message {
  display: inline-block;
  margin-right: 20px;
  width: 500px;
}
.sendMessage {
  display: inline-block;
}
.inputArea {
  padding: 10px;
  border-top: 2px solid #ddd;
}
.messageReceive {
  float: left;
  display: block;
  clear: both;
  margin-bottom: 20px;
  font-family: Tahoma, '仿宋';
}
.messageSend {
  float: right;
  display: block;
  clear: both;
  margin-bottom: 20px;
  font-family: Tahoma, '仿宋';
}
.receiveBubble {
  background-color: #eee;
  min-height: 40px;
  max-width: 500px;
  padding: 15px;
  font-size: 25px;
  text-align: left;
  border-radius: 5px;
  word-break: break-all;
  word-wrap: break-word;
}
.sendBubble {
  background-color: #6d6;
  min-height: 40px;
  max-width: 500px;
  padding: 15px;
  font-size: 25px;
  text-align: left;
  border-radius: 5px;
  word-break: break-all;
  word-wrap: break-word;
}
.receiveTime {
  float: left;
}
.sendTime {
  float: right;
}
</style>

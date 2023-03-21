<template>
  <div>
    <el-container direction="vertical" class="table_container">
      <!-- 数据表 -->
      <el-table :data="tableData" style="width: 100%">
        <el-table-column
          prop="userId"
          label="用户id"
          width="500px"
          class="test"
        ></el-table-column>
        <el-table-column
          prop="userWxName"
          label="用户名"
          width="600px"
        ></el-table-column>
        <el-table-column prop="userState" label="核酸状态" width="500px">
          <template slot-scope="scope">
            <el-tag size="large" :type="userStateTag(scope.row.userState)">{{
              dealState(scope.row.userState)
            }}</el-tag>
          </template></el-table-column
        >
        <el-table-column label="操作">
          <template v-slot="scope">
            <el-button
              @click="handleEdit(scope.$index, scope.row)"
              type="primary"
              size="medium"
              icon="el-icon-edit"
              native-type="submit"
              plain
              >修改
            </el-button>
          </template>
        </el-table-column>
        <!-- 空数据处理 -->
        <div slot="empty" style="text-align: left">
          <el-empty> </el-empty>
        </div>
      </el-table>
      <!-- 修改框 -->
      <el-dialog title="修改核酸状态" :visible.sync="dialogVisible" width="30%">
        <h2>正在修改id为{{ idChanged }}的用户危险等级</h2>
        <el-select v-model="stateChanged" placeholder="危险等级">
          <el-option label="0(安全)" :value="0"></el-option>
          <el-option label="1" :value="1"></el-option>
          <el-option label="2" :value="2"></el-option>
          <el-option label="3(确诊)" :value="3"></el-option>
        </el-select>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="stateSubmit">确 定</el-button>
        </span>
      </el-dialog>
      <!-- 底部页码部分 -->
      <el-pagination
        class="pagination"
        :hide-on-single-page="true"
        :current-page="currentPage"
        @current-change="handleCurrentChange"
        :page-size="pageSize"
        :total="totalPersons"
        layout="total, prev, pager, next, jumper"
        style="margin: 15px"
      ></el-pagination>
    </el-container>
  </div>
</template>

<script>
import { mixins } from '../libs/provideIndex.js'
export default {
  name: 'StateView',
  mixins: [mixins],
  data() {
    return {
      currentPage: 1,
      pageSize: 10,
      tableData: [],
      idChanged: -1,
      stateChanged: -1,
      dialogVisible: false,
      totalPersons: 0,
      activeIndex: 'state',
    }
  },
  mounted() {
    // 初始化页面后立即更新数据
    this.upUserData()
  },
  methods: {
    //状态标签赋值
    userStateTag(state) {
      switch (state) {
        case 0:
          return 'success'
        case 3:
          return 'danger'
        default:
          return 'info'
      }
    },
    // 状态文字赋值
    dealState(state) {
      switch (state) {
        case 0:
          return '安全'
        case 3:
          return '危险'
        default:
          return '风险'
      }
    },
    // 页码更改时更新数据
    handleCurrentChange(val) {
      this.currentPage = val
      this.upUserData()
    },
    //请求用户数据
    upUserData() {
      this.$axios({
        url: '/user/get/all/user',
        method: 'get',
        headers: {
          authToken: localStorage.getItem('token'),
        },
        params: {
          num: this.currentPage,
          size: this.pageSize,
        },
        // transformRequest不能用于GET
      })
        .then((resp) => {
          this.tableData = resp.data
          this.totalPersons = parseInt(resp.msg)
        })
        .catch(() => {
          this.$message.error('获取用户数据失败，请联系工程人员')
        })
    },
    //操作核酸状态
    handleEdit(index, row) {
      this.idChanged = row.userId
      this.stateChanged = row.userState
      this.dialogVisible = true
    },
    //请求修改核酸状态
    stateSubmit() {
      this.$axios({
        url: '/judge/setacidstate',
        method: 'POST',
        data: {
          state: this.stateChanged,
          userid: this.idChanged,
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
          this.$message.success('修改成功')
          this.upUserData()
          this.dialogVisible = false
        })
        .catch(() => {
          this.$message.error('修改失败，请联系工程人员')
          this.dialogVisible = false
        })
    },
  },
}
</script>

<style lang="scss">
.table_container {
  width: 100%;
  margin-top: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04);
}
.pagination {
  margin: 0 auto;
  display: block;
}
</style>

<template>
  <el-menu @select="handleSelect">
    <el-menu-item-group>
      <div v-for="item in fc" :key="item.areaId">
        <el-menu-item :index="String(item.areaId)">
          [{{ item.areaId }}]{{ item.areaName }}
          <i
            class="el-icon-delete"
            @click="delcat(item.areaId)"
            style="float: right; margin-top: 15px"
          ></i>
        </el-menu-item>
        <el-dialog
          title="删除地点"
          :visible="dialogDelCatVisible"
          width="400px"
          :destroy-on-close="true"
        >
          是否要删除这个地点？此操作无法撤销！
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="erasecatConfirm()">
              删 除
            </el-button>
            <el-button @click="dialogDelCatVisible = false">取 消</el-button>
          </div>
        </el-dialog>
      </div>
    </el-menu-item-group>
  </el-menu>
</template>

<script>
export default {
  name: 'PlaceAdd',
  props: ['fc'],
  //该组件被MainView使用，目的是显示所有的地点
  data() {
    return {
      //是否显示删除地点对话框
      dialogDelCatVisible: false,
      //待删除的地点
      dele: null,
    }
  },
  methods: {
    //选中某一项时触发的事件
    handleSelect(key, keyPath) {
      console.log(key, keyPath)
      this.$emit('lookup', key)
    },
    //按下删除按钮之后，弹出对话框
    delcat(deldata) {
      console.log('删除地点', deldata)
      this.dialogDelCatVisible = true
      this.dele = deldata
    },
    //确认删除，向服务器提交删除请求
    erasecatConfirm() {
      console.log('正在删除地点', this.dele)
      this.$axios({
        url: '/area/set/delete/area',
        method: 'post',

        data: {
          areaid: this.dele,
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
          if (res.code == 10000) {
            this.$message.success('删除成功')
            setTimeout(() => {
              this.$emit('UPDtype')
              this.$emit('deltable')
            }, 200)
          } else this.$message.error('删除失败')
          this.dialogDelCatVisible = false
        })
        .catch((error) => {
          this.$message.error(error)
        })
    },
  },
}
</script>

<style></style>

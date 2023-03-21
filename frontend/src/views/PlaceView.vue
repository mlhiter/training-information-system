<template>
  <div>
    <!-- 主体：左侧边栏和右侧表格 -->
    <el-container style="height: 500px; left: 0; margin-top: 10px">
      <!--左侧边栏：重点地点的查看和添加 -->
      <el-aside width="300px">
        <i @click="open" class="el-icon-plus addicon"></i>
        <el-menu default-active="2">
          <el-submenu index="1">
            <template v-slot:title>
              <i class="el-icon-location"></i>
              <span>重要地点</span>
            </template>
            <PlaceAdd
              :fc="this.form.class"
              @UPDtype="UPDtype"
              @lookup="lookup"
              @deltable="deltable"
            />
          </el-submenu>
        </el-menu>
      </el-aside>
      <el-dialog
        title="新建地点"
        :visible="dialogFormVisible"
        width="600px"
        :destroy-on-close="true"
      >
        <el-form :model="form">
          <el-form-item label="地点名称" label-width="75px">
            <el-input v-model="word" autocomplete="off"></el-input>
          </el-form-item>
          <div id="mapcontainer"></div>
          <el-form-item label="经度" label-width="40px" class="positem">
            <el-input
              v-model="pointlongitude"
              autocomplete="off"
              class="pos"
            ></el-input>
          </el-form-item>
          <el-form-item label="纬度" label-width="40px" class="positem">
            <el-input
              v-model="pointlatitude"
              autocomplete="off"
              class="pos"
            ></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="close">取 消</el-button>
          <el-button type="primary" @click="createcat">创 建</el-button>
        </div>
      </el-dialog>
      <!-- 中间表格 -->
      <el-container direction="vertical">
        <div class="buttons">
          <el-button @click="show" type="primary" plain>
            <i class="el-icon-plus">&nbsp;新增</i>
          </el-button>
          &nbsp;&nbsp;&nbsp;
          <!-- 添加坐标弹出框 -->
          <el-dialog
            title="添加坐标"
            v-model:visible="dialogAddVisible"
            width="1000px"
            :destroy-on-close="true"
            :close-on-click-modal="false"
            center
          >
            <el-form :inline="true">
              <el-form-item label="地点ID" label-width="70px" prop="category">
                <el-select
                  v-model="addareaId"
                  placeholder="请选择侧边栏中的地点"
                  disabled
                >
                  <el-option
                    v-for="item in form.class"
                    :key="item.areaId"
                    :label="item.areaId"
                    :value="item.areaId"
                  ></el-option>
                </el-select>
              </el-form-item>
              <div id="mapcontainer2"></div>
              <el-form-item label="经度" label-width="50px" class="positem">
                <el-input
                  v-model="pointlongitude2"
                  autocomplete="off"
                  class="pos"
                  :disabled="true"
                ></el-input>
              </el-form-item>
              <el-form-item label="纬度" label-width="40px" class="positem">
                <el-input
                  v-model="pointlatitude2"
                  autocomplete="off"
                  class="pos"
                  :disabled="true"
                ></el-input>
              </el-form-item>
              <el-form-item label="半径" label-width="40px" class="positem">
                <el-input
                  v-model="radius2"
                  autocomplete="off"
                  class="pos"
                  :disabled="true"
                ></el-input>
              </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
              <el-button type="primary" @click="createitem">上 传</el-button>
              <el-button @click="dialogAddVisible = false">取 消</el-button>
            </div>
          </el-dialog>
          <!-- <el-button type="info" plain @click="modify"
            ><i class="el-icon-edit">&nbsp;修改</i></el-button
          >&nbsp;&nbsp;&nbsp; -->
          <el-dialog
            title="修改坐标"
            v-model:visible="dialogModVisible"
            width="1000px"
            :destroy-on-close="true"
            :close-on-click-modal="false"
          >
            <!--在此处添加修改坐标对话框的内容-->
            <div slot="footer" class="dialog-footer">
              <el-button type="primary" @click="modiConfirm">修 改</el-button>
              <el-button @click="modiCancel">取 消</el-button>
            </div>
          </el-dialog>
          <!-- <el-button type="danger" plain @click="eraseit"
            ><i class="el-icon-delete">&nbsp;删除</i></el-button
          >&nbsp;&nbsp;&nbsp; -->
          <el-dialog
            title="删除坐标"
            v-model:visible="dialogDelVisible"
            width="400px"
            :destroy-on-close="true"
          >
            <!--在此处添加删除坐标对话框的内容-->
            <div slot="footer" class="dialog-footer">
              <el-button type="primary" @click="eraseConfirm">删 除</el-button>
              <el-button @click="dialogDelVisible = false">取 消</el-button>
            </div>
          </el-dialog>
          <el-button type="success" plain @click="lookup(addareaId)">
            <i class="el-icon-refresh-left">&nbsp;刷新</i>
          </el-button>
        </div>
        <MainTable
          :tableData="this.tableData"
          :currentPage="this.currentPage"
          :pageSize="this.pageSize"
          :fc="this.form.class"
          @Trans="Trans"
          @eraseSingle="eraseSingle"
        />
      </el-container>
      <el-container>
        <div id="mapcontainer3"></div>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import { mixins } from '../libs/provideIndex.js'
import PlaceAdd from '../components/PlaceAdd.vue'
import MainTable from '../components/MainTable.vue'
import AMapLoader from '@amap/amap-jsapi-loader'
// @ is an alias to /src
export default {
  name: 'PlaceView',
  components: {
    PlaceAdd,
    MainTable,
  },
  mixins: [mixins],
  mounted() {
    //启动时先更新地点列表
    this.UPDtype()
    AMapLoader.load({
      key: '1953d1d7aa2b0b62ca8087efc9bfd9d8', //此处填入我们注册账号后获取的Key
      version: '2.0', //指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
      plugins: [''], //需要使用的的插件列表，如比例尺'AMap.Scale'等
    })
      .then((AMap) => {
        this.AMap3 = AMap
        this.map3 = new AMap.Map('mapcontainer3', {
          //设置地图容器id
          viewMode: '2D', //是否为3D地图模式
          zoom: 15, //初始化地图级别
          center: [122.082, 37.533], //初始化地图中心点位置
        })
      })
      .catch((e) => {
        console.log(e)
      })
  },
  data() {
    return {
      activeIndex: 'place',
      //每页显示的坐标数量
      pageSize: 7,
      //目前的页码
      currentPage: 1,
      //“新建地点”对话框中的“地点名称”
      word: '',
      //要显示的坐标列表
      tableData: [],
      //是否显示“添加地点”对话框
      dialogFormVisible: false,
      //是否显示“添加坐标”对话框
      dialogAddVisible: false,
      //是否显示“修改坐标”对话框
      dialogModVisible: false,
      //是否显示“删除坐标”对话框
      dialogDelVisible: false,
      //要显示的地点列表
      form: {
        class: [],
      },
      //新建坐标时待上传的areaId
      addareaId: null,
      //显示在“修改坐标”对话框中的内容
      moditemp: {},
      //多选框状态
      multipleSelection: [],
      //地图组件相关数据
      map: null,
      AMap: null,
      //“添加地点”对话框中的经纬度
      pointlongitude: 122.082,
      pointlatitude: 37.533,
      //“添加坐标点”对话框中的经纬度及半径
      pointlongitude2: 122.082,
      pointlatitude2: 37.533,
      radius2: 15,
      circleEditor: null,
      circle: null,
      map3: null,
      AMap3: null,
    }
  },
  methods: {
    //更新地点列表
    UPDtype() {
      console.log('type is being updated。')
      this.$axios({
        url: '/area/get/only/area',
        method: 'get',
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
          this.form.class = res.data
          console.log(this.form.class)
        })
        .catch((error) => {
          this.$message.error(error)
        })
    },
    //修改当前页时执行的函数
    handleCurrentChange(e) {
      this.currentPage = e
    },
    //点击位于“添加地点”中的地图组件时执行的函数
    ClickMap(e) {
      this.pointlongitude = e.lnglat.getLng()
      this.pointlatitude = e.lnglat.getLat()
    },
    //点击位于“添加坐标”中的地图组件时执行的函数
    ClickMap2(e) {
      this.circleEditor.close()
      this.circle.setCenter(e.lnglat)
      this.pointlongitude2 = e.lnglat.getLng()
      this.pointlatitude2 = e.lnglat.getLat()
      clearTimeout(this.timer)
      this.timer = setTimeout(() => {
        this.circleEditor.open()
      }, 100)
    },
    //“添加地点”对话框打开时执行的函数
    open() {
      this.dialogFormVisible = true
      console.log('调用地图初始化')
      ;(this.pointlongitude = 122.082),
        (this.pointlatitude = 37.533),
        AMapLoader.load({
          key: '1953d1d7aa2b0b62ca8087efc9bfd9d8', //此处填入我们注册账号后获取的Key
          version: '2.0', //指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
          plugins: ['AMap.CircleEditor'], //需要使用的的插件列表，如比例尺'AMap.Scale'等
        })
          .then((AMap) => {
            this.AMap = AMap
            this.map = new AMap.Map('mapcontainer', {
              //设置地图容器id
              viewMode: '2D', //是否为3D地图模式
              zoom: 16, //初始化地图级别
              center: [122.082, 37.533], //初始化地图中心点位置
            })
            this.map.on('click', this.ClickMap)
          })
          .catch((e) => {
            console.log(e)
          })
    },
    //“添加地点”对话框关闭时执行的函数
    close() {
      this.dialogFormVisible = false
      this.word = ''
    },
    //“添加坐标”对话框打开时执行的函数
    show() {
      var that = this
      this.dialogAddVisible = true
      console.log('调用地图初始化2')
      ;(this.pointlongitude2 = 122.082),
        (this.pointlatitude2 = 37.533),
        AMapLoader.load({
          key: '1953d1d7aa2b0b62ca8087efc9bfd9d8', //此处填入我们注册账号后获取的Key
          version: '2.0', //指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
          plugins: ['AMap.CircleEditor'], //需要使用的的插件列表，如比例尺'AMap.Scale'等
        })
          .then((AMap) => {
            this.AMap = AMap
            this.map = new AMap.Map('mapcontainer2', {
              //设置地图容器id
              viewMode: '2D', //是否为3D地图模式
              zoom: 16, //初始化地图级别
              center: [122.082, 37.533], //初始化地图中心点位置
            })
            this.map.on('click', this.ClickMap2)
            that.updcircle()
            if (this.addareaId != null) {
              this.$axios({
                url: '/area/get/coor/by/area/id',
                method: 'post',
                headers: {
                  'Content-Type': 'application/x-www-form-urlencoded',
                  authToken: localStorage.getItem('token'),
                },
                data: {
                  areaId: this.addareaId,
                },
                transformRequest: [
                  function (dat) {
                    let ret = ''
                    for (let it in dat) {
                      ret +=
                        encodeURIComponent(it) +
                        '=' +
                        encodeURIComponent(dat[it]) +
                        '&'
                    }
                    ret = ret.substring(0, ret.lastIndexOf('&'))
                    return ret
                  },
                ],
              })
                .then((res) => {
                  console.log(res)
                  if (res.code == '10000') {
                    this.tableData = res.data
                    for (var i = 0; i < res.data.length; i++) {
                      var tempcircle = new this.AMap.Circle({
                        center: new this.AMap.LngLat(
                          res.data[i].longitude,
                          res.data[i].latitude
                        ), // 圆心位置
                        radius: res.data[i].radius * 6300000, //半径
                        strokeColor: '#22b14c', //线颜色
                        strokeOpacity: 1, //线透明度
                        strokeWeight: 3, //线粗细度
                        fillColor: '#B5E61D', //填充颜色
                        fillOpacity: 0.35, //填充透明度
                      })
                      this.map.add(tempcircle)
                    }
                  } else this.$message.error(res.errorMsg)
                })
                .catch((error) => {
                  this.$message.error(error)
                })
            }
          })
          .catch((e) => {
            console.log(e)
          })
    },
    //确认添加坐标时执行的函数
    createitem() {
      if (this.addareaId == null) {
        this.$message.error('必填项不完整')
        return
      }
      var that = this
      this.dialogAddVisible = false
      let coord = {
        id: null,
        longitude: that.pointlongitude2,
        latitude: that.pointlatitude2,
        radius: that.radius2,
      }
      this.$axios({
        url: '/area/set/coor',
        method: 'post',
        headers: {
          'Content-Type': 'application/json',
          authToken: localStorage.getItem('token'),
        },
        data: {
          areaId: that.addareaId,
          areaCoordinate: coord,
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
            this.$message.success('成功创建坐标')
            setTimeout(() => {
              this.lookup(this.addareaId)
            }, 200)
          } else this.$message.error('无法创建坐标')
        })
        .catch((error) => {
          this.$message.error(error)
        })
    },
    //确认添加地点时执行的函数
    createcat() {
      if (this.word.length == 0) {
        this.$message.error('必填项不完整')
        return
      }
      var that = this
      this.dialogFormVisible = false
      this.$axios({
        url: '/area/set/one',
        method: 'post',
        headers: {
          'Content-Type': 'application/json',
          authToken: localStorage.getItem('token'),
        },
        data: {
          name: that.word,
          latitude: that.pointlatitude,
          longitude: that.pointlongitude,
          personNumber: 0,
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
            this.$message.success('成功创建地点')
            this.word = ''
          } else this.$message.error('无法创建地点')
        })
        .catch((error) => {
          this.$message.error(error)
        })
      setTimeout(() => {
        this.UPDtype()
      }, 100)
    },
    //按下修改按钮时执行的函数
    modify() {
      this.$message.warning('功能暂未开放')
      //this.dialogModVisible = true;
    },
    //按下坐标右侧修改按钮时执行的函数
    modifySingle(idx) {
      console.log(idx)
      this.$message.warning('功能暂未开放')
      //this.dialogModVisible = true;
    },
    //确认修改执行
    modiConfirm() {
      this.dialogModVisible = false
    },
    //取消修改执行
    modiCancel() {
      this.moditemp = {}
      this.dialogModVisible = false
    },
    //按下删除按钮执行
    eraseit() {
      this.$message.warning('功能暂未开放，请使用坐标右侧的删除按钮')
      //this.dialogDelVisible = true;
    },
    //按下坐标右侧删除按钮执行
    eraseSingle(idx) {
      console.log('正在删除', idx)
      this.$axios({
        url: '/area/set/delete/coo',
        method: 'post',

        data: {
          cooid: idx,
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
              this.lookup(this.addareaId)
            }, 200)
          } else this.$message.error('删除失败')
        })
        .catch((error) => {
          this.$message.error(error)
        })
    },
    //确认删除执行
    eraseConfirm() {
      console.log(this.multipleSelection)
      this.dialogDelVisible = false
    },
    lookup(key) {
      if (key == null) {
        this.$message.warning('请选择一个地点再更新')
        return
      }
      this.addareaId = key
      this.$axios({
        url: '/area/get/coor/by/area/id',
        method: 'post',

        data: {
          areaId: key,
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
          if (res.code == '10000') {
            this.tableData = res.data
            this.map3.clearMap()
            for (var i = 0; i < res.data.length; i++) {
              var tempcircle = new this.AMap3.Circle({
                center: new this.AMap3.LngLat(
                  res.data[i].longitude,
                  res.data[i].latitude
                ), // 圆心位置
                radius: res.data[i].radius * 6300000, //半径
                strokeColor: '#2d9bfc', //线颜色
                strokeOpacity: 1, //线透明度
                strokeWeight: 3, //线粗细度
                fillColor: '#87c5f8', //填充颜色
                fillOpacity: 0.35, //填充透明度
              })
              this.map3.add(tempcircle)
            }
          } else this.$message.error(res.errorMsg)
        })
        .catch((error) => {
          this.$message.error(error)
        })
    },
    //相应子组件MainTable提示的函数
    Trans(value) {
      this.multipleSelection = value
    },
    //“添加地点”中的地图更新标点
    updmark() {
      this.map.clearMap()
      var ico = new this.AMap.Icon({
        size: new this.AMap.Size(32, 40), // 图标尺寸
        image:
          'https://a.amap.com/jsapi_demos/static/demo-center/icons/poi-marker-default.png', // Icon的图像
        imageSize: new this.AMap.Size(32, 40), // 根据所设置的大小拉伸或压缩图片
      })

      var marker = new this.AMap.Marker({
        icon: ico,
        position: [this.pointlongitude, this.pointlatitude],
        anchor: 'bottom-center',
      })
      this.map.add(marker)
    },
    //“添加坐标”中的地图更新标点
    updcircle() {
      this.map.clearMap()
      this.circle = new this.AMap.Circle({
        center: new this.AMap.LngLat(this.pointlongitude2, this.pointlatitude2), // 圆心位置
        radius: this.radius2, //半径
        strokeColor: '#2d9bfc', //线颜色
        strokeOpacity: 1, //线透明度
        strokeWeight: 3, //线粗细度
        fillColor: '#87c5f8', //填充颜色
        fillOpacity: 0.35, //填充透明度
      })
      this.map.add(this.circle)
      this.map.setFitView([this.circle])
      this.circleEditor = new this.AMap.CircleEditor(this.map, this.circle)
      this.circleEditor.open()
      this.circleEditor.on('move', this.circleposchange)
      this.circleEditor.on('adjust', this.circleradiuschange)
    },
    circleposchange(event) {
      this.pointlongitude2 = event.lnglat.KL
      this.pointlatitude2 = event.lnglat.kT
    },
    circleradiuschange(event) {
      this.radius2 = event.radius
    },
    updsmallmap() {
      this.$axios({
        url: '/area/get/coor/by/area/id',
        method: 'post',

        data: {
          areaId: this.addareaId,
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
          this.updcircle()
          if (res.code == '10000') {
            this.tableData = res.data
            for (var i = 0; i < res.data.length; i++) {
              var tempcircle = new this.AMap.Circle({
                center: new this.AMap.LngLat(
                  res.data[i].longitude,
                  res.data[i].latitude
                ), // 圆心位置
                radius: res.data[i].radius * 6300000, //半径
                strokeColor: '#22b14c', //线颜色
                strokeOpacity: 1, //线透明度
                strokeWeight: 3, //线粗细度
                fillColor: '#B5E61D', //填充颜色
                fillOpacity: 0.35, //填充透明度
              })
              this.map.add(tempcircle)
            }
          } else this.$message.error(res.errorMsg)
        })
        .catch((error) => {
          this.$message.error(error)
        })
    },
    deltable() {
      this.tableData = []
    },
  },
  watch: {
    //监控变量值的变化，便于做出及时响应
    pointlongitude(now, pre) {
      console.log('longitude变化。', pre, '->', now)
      this.updmark()
    },
    pointlatitude(now, pre) {
      console.log('latitude变化。', pre, '->', now)
      this.updmark()
    },
    addareaId(now, pre) {
      console.log('addareaId变化', pre, '->', now)
      if (now == null) return
      else if (this.AMap == null) return
      else this.updsmallmap()
    },
  },
}
</script>

<style>
.buttons {
  text-align: left;
}
.Tit {
  display: block;
  font-size: 20px;
  font-weight: bold;
  margin-top: 4px;
}
.searchidx {
  width: 200px;
}
.funct {
  margin-left: 100px;
  margin-right: 150px;
  float: right;
}
.funct i {
  display: inline-block;
  text-align: center;
  font-size: 23px;
  height: 25px;
  width: 25px;
  border-radius: 10px;
  border: rgb(240, 240, 240) 2px solid;
  margin-right: 40px;
  cursor: pointer;
}
.el-pagination {
  float: right;
  margin-right: 80px;
}
.el-dialog {
  border-radius: 10px !important;
  cursor: default !important;
}
.el-aside {
  border-right: solid 1px #e6e6e6;
  margin-right: 20px;
}
#mapcontainer {
  width: 560px;
  height: 300px;
  margin-bottom: 10px;
  text-align: left;
}
#mapcontainer2 {
  width: 800px;
  height: 370px;
  margin-bottom: 10px;
  text-align: left;
}
#mapcontainer3 {
  width: 600px;
  height: 535px;
  text-align: left;
}
.pos {
  width: 200px !important;
}
.positem {
  display: inline-block !important;
  width: 260px;
}
.addicon {
  position: absolute;
  top: 103px;
  left: 240px;
  cursor: pointer;
  z-index: 9989;
}
</style>

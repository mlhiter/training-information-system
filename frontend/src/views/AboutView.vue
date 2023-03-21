<template>
  <div>
    <el-row :gutter="20">
      <el-col :span="12">
        <img src="../assets/logo.jpg" />
      </el-col>
      <el-col :span="12">
        <div class="abouttextcontain">
          <div class="abouttext">
            我们的小程序是一个能在在传染病爆发时期，为未登录用户展示某一聚集地点的实时人流量，经用户登录授权后能为用户展示自己与感染者的轨迹重合之处并能让用户和管理员通畅交流的小程序。
          </div>
          <div class="abouttext">
            我们的小程序的优势在于简便实用，非常符合当下遏制农村疫情的需要，但还不够成熟，虽然有解决方案但是由于时间关系还没能完成。网页端作为管理员端也非常简洁易用，我们小程序投入使用需要与政府和基层组织合作。
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mixins } from '../libs/provideIndex.js'
export default {
  name: 'AboutView',
  mixins: [mixins],
  data() {
    return {
      activeIndex: 'about',
    }
  },
  methods: {},
}
</script>

<style>
.abouttext {
  font-family: 楷体, 仿宋, 宋体;
  text-align: left;
  text-indent: 2em;
  font-size: 30px;
  line-height: 40px;
  margin-top: 50px;
}
.abouttextcontain {
  margin-top: 160px;
}
</style>

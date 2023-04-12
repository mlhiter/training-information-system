import type { RouteRecordRaw } from 'vue-router'

import MainLayout from '@/layouts/index.vue'

// 前台路由表
const frontRoutes: RouteRecordRaw[] = [
  {
    path: '/frontstage/enroll', // 报名
    component: () =>
      import(
        /* webpackChunkName: "enroll" */ '@/pages/frontstage/enroll/index.vue'
      ),
    meta: {
      title: '报名',
      role: 'user',
    },
  },
  {
    path: '/frontstage/inplace', // 签到
    component: () =>
      import(
        /* webpackChunkName: "inplace" */ '@/pages/frontstage/inplace/index.vue'
      ),
    meta: {
      title: '签到',
      role: 'user',
    },
  },
  {
    path: '/frontstage/questionnaire', // 调查问卷
    component: () =>
      import(
        /* webpackChunkName: "inplace" */ '@/pages/frontstage/questionnaire/index.vue'
      ),
    meta: {
      title: '调查问卷',
      role: 'user',
    },
  },
]

// 后台路由表

//执行人路由
const executorRoutes: RouteRecordRaw[] = [
  {
    path: '/backstage/executor/enroll', // 个人报名审核
    component: () =>
      import(
        /* webpackChunkName: "enroll" */ '@/pages/backstage/executor/enroll/index.vue'
      ),
    meta: {
      title: '个人报名审核',
      role: 'executor',
    },
  },
  {
    path: '/backstage/executor/lecturer', // 讲师管理
    component: () =>
      import(
        /* webpackChunkName: "lecturer" */ '@/pages/backstage/executor/lecturer/index.vue'
      ),
    meta: {
      title: '讲师管理',
      role: 'executor',
    },
  },
  {
    path: '/backstage/executor/questionnaire', // 问卷管理
    component: () =>
      import(
        /* webpackChunkName: "questionnaire" */ '@/pages/backstage/executor/questionnaire/index.vue'
      ),
    meta: {
      title: '问卷管理',
      role: 'executor',
    },
  },
  {
    path: '/backstage/executor/student', // 学员管理
    component: () =>
      import(
        /* webpackChunkName: "student" */ '@/pages/backstage/executor/student/index.vue'
      ),
    meta: {
      title: '学员管理',
      role: 'executor',
    },
  },
  {
    path: '/backstage/executor/training', // 培训信息管理
    component: () =>
      import(
        /* webpackChunkName: "training" */ '@/pages/backstage/executor/training/index.vue'
      ),
    meta: {
      title: '培训信息管理',
      role: 'executor',
    },
  },
]
// 经理路由
const managerRoutes: RouteRecordRaw[] = [
  {
    path: '/backstage/manager/training', // 培训审核
    component: () =>
      import(
        /* webpackChunkName: "training" */ '@/pages/backstage/manager/training/index.vue'
      ),
    meta: {
      title: '培训审核',
      role: 'manager',
    },
  },
  {
    path: '/backstage/manager/report', // 汇总报表
    component: () =>
      import(
        /* webpackChunkName: "report" */ '@/pages/backstage/manager/report/index.vue'
      ),
    meta: {
      title: '汇总报表',
      role: 'manager',
    },
  },
  {
    path: '/backstage/manager/executor', // 执行人管理
    component: () =>
      import(
        /* webpackChunkName: "executor" */ '@/pages/backstage/manager/executor/index.vue'
      ),
    meta: {
      title: '执行人管理',
      role: 'manager',
    },
  },
]
//现场工作人员路由
const operatorRoutes: RouteRecordRaw[] = [
  {
    path: '/backstage/operator', //现场工作
    component: () =>
      import(
        /* webpackChunkName: "operator" */ '@/pages/backstage/operator/index.vue'
      ),
    meta: {
      title: '现场管理',
      role: 'operator',
    },
  },
]
const backRoutes: RouteRecordRaw[] = [
  ...executorRoutes,
  ...managerRoutes,
  ...operatorRoutes,
]

export const constRoutes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Layout',
    component: MainLayout,
    children: [...frontRoutes, ...backRoutes],
  },
]

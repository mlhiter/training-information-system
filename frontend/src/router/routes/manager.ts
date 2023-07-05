import { RouteRecordRaw } from 'vue-router'

// 经理路由
export const managerRoutes: RouteRecordRaw[] = [
  {
    path: '/backstage/manager/report', // 汇总报表
    component: () =>
      import(
        /* webpackChunkName: "report" */ '@/pages/backstage/manager/report/index.vue'
      ),
    meta: { requiredAuth: true, title: '汇总报表', role: 'manager' },
  },
  {
    path: '/backstage/manager/executor', // 执行人管理
    component: () =>
      import(
        /* webpackChunkName: "executor" */ '@/pages/backstage/manager/executor/index.vue'
      ),
    meta: { requiredAuth: true, title: '执行人管理', role: 'manager' },
  },
  {
    path: '/backstage/manager/user', // 用户管理
    component: () =>
      import(
        /* webpackChunkName: "executor" */ '@/pages/backstage/manager/user/index.vue'
      ),
    meta: { requiredAuth: true, title: '用户管理', role: 'manager' },
  },
]

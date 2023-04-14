import type { RouteRecordRaw } from 'vue-router'
import frontstageLayout from '@/layouts/frontstage/index.vue'
import backstageLayout from '@/layouts/backstage/index.vue'
import { executorRoutes } from './executor'
import { managerRoutes } from './manager'
import { operatorRoutes } from './operator'
import { userRoutes } from './user'

//登录路由
const authRoutes: RouteRecordRaw[] = [
  {
    path: '/auth',
    component: () =>
      import(/* webpackChunkName: "auth" */ '@/pages/auth/enroll/index.vue'),
    meta: {
      title: '权限认证',
      requiredAuth: false,
    },
  },
]
// 前台路由表
const frontRoutes: RouteRecordRaw[] = [
  {
    path: '/', // 前台layout
    name: 'FrontLayout',
    component: frontstageLayout,
    redirect: '/frontstage/enroll',
    meta: { title: '前台' },
    children: userRoutes,
  },
]

// 后台路由表
const backRoutes: RouteRecordRaw[] = [
  {
    path: '/backstage', // 后台layout
    component: backstageLayout,
    meta: { title: '后台' },
    children: [...executorRoutes, ...managerRoutes, ...operatorRoutes],
  },
]

export const constRoutes: RouteRecordRaw[] = [
  ...frontRoutes,
  ...backRoutes,
  ...authRoutes,
]

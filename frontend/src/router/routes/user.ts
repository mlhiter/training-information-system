import { RouteRecordRaw } from 'vue-router'

export const userRoutes: RouteRecordRaw[] = [
  {
    path: '/frontstage/enroll', // 报名
    component: () =>
      import(
        /* webpackChunkName: "enroll" */ '@/pages/frontstage/enroll/index.vue'
      ),
    meta: { requiredAuth: true, title: '报名', role: 'user' },
  },
  {
    path: '/frontstage/questionnaire', // 调查问卷
    component: () =>
      import(
        /* webpackChunkName: "inplace" */ '@/pages/frontstage/questionnaire/index.vue'
      ),
    meta: { requiredAuth: true, title: '调查问卷', role: 'user' },
  },
  {
    path: '/frontstage/profile', // 个人信息
    component: () =>
      import(
        /* webpackChunkName: "profile" */ '@/pages/frontstage/profile/index.vue'
      ),
    meta: { requiredAuth: true, title: '个人信息', role: 'user' },
  },
]

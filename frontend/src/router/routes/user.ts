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
    path: '/frontstage/inplace', // 签到
    component: () =>
      import(
        /* webpackChunkName: "inplace" */ '@/pages/frontstage/inplace/index.vue'
      ),
    meta: { requiredAuth: true, title: '签到', role: 'user' },
  },
  {
    path: '/frontstage/questionnaire', // 调查问卷
    component: () =>
      import(
        /* webpackChunkName: "inplace" */ '@/pages/frontstage/questionnaire/index.vue'
      ),
    meta: { requiredAuth: true, title: '调查问卷', role: 'user' },
  },
]

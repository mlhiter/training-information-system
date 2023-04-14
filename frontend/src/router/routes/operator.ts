import { RouteRecordRaw } from 'vue-router'

//现场工作人员路由
export const operatorRoutes: RouteRecordRaw[] = [
  {
    path: '/backstage/operator', //现场工作
    component: () =>
      import(
        /* webpackChunkName: "operator" */ '@/pages/backstage/operator/index.vue'
      ),
    meta: { requiredAuth: true, title: '现场管理', role: 'operator' },
  },
]

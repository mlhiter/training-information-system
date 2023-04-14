import { RouteRecordRaw } from 'vue-router';

//执行人路由
export const executorRoutes: RouteRecordRaw[] = [
  {
    path: '/backstage/executor/enroll', // 个人报名审核
    component: () =>
      import(
        /* webpackChunkName: "enroll" */ '@/pages/backstage/executor/enroll/index.vue'
      ),
    meta: { requiredAuth: true, title: '个人报名审核', role: 'executor' },
  },
  {
    path: '/backstage/executor/lecturer', // 讲师管理
    component: () =>
      import(
        /* webpackChunkName: "lecturer" */ '@/pages/backstage/executor/lecturer/index.vue'
      ),
    meta: { requiredAuth: true, title: '讲师管理', role: 'executor' },
  },
  {
    path: '/backstage/executor/questionnaire', // 问卷管理
    component: () =>
      import(
        /* webpackChunkName: "questionnaire" */ '@/pages/backstage/executor/questionnaire/index.vue'
      ),
    meta: { requiredAuth: true, title: '问卷管理', role: 'executor' },
  },
  {
    path: '/backstage/executor/student', // 学员管理
    component: () =>
      import(
        /* webpackChunkName: "student" */ '@/pages/backstage/executor/student/index.vue'
      ),
    meta: { requiredAuth: true, title: '学员管理', role: 'executor' },
  },
  {
    path: '/backstage/executor/training', // 培训信息管理
    component: () =>
      import(
        /* webpackChunkName: "training" */ '@/pages/backstage/executor/training/index.vue'
      ),
    meta: { requiredAuth: true, title: '培训信息管理', role: 'executor' },
  },
]

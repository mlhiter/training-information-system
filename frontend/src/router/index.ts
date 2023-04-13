import { createRouter, createWebHistory } from 'vue-router'
import { constRoutes } from './routes'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: constRoutes,
  scrollBehavior: () => ({ left: 0, top: 0 }),
})

//路由守卫
router.beforeEach(async (to) => {
  // 更新标题栏title
  const subtitle = to.meta.title
  if (subtitle) {
    document.title = subtitle
  }
})
// 添加路由守卫
// router.beforeEach((to, from, next) => {
//   const role = to.meta.role
//   const currentUser = getAuth().currentUser
//   if (currentUser) {
//     // 判断角色是否匹配
//     if (role === 'executor' && currentUser.role === 'executor') {
//       next()
//     } else if (role === 'manager' && currentUser.role === 'manager') {
//       next()
//     } else if (role === 'operator' && currentUser.role === 'operator') {
//       next()
//     } else {
//       next('/login')
//     }
//   } else {
//     next({
//       path: '/login',
//       query: { redirect: to.fullPath },
//     })
//   }
// })

export default router

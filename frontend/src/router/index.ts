import { createRouter, createWebHistory } from 'vue-router'
import { constRoutes } from './routes'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: constRoutes,
  scrollBehavior: () => ({ left: 0, top: 0 }),
})

// 添加路由守卫
router.beforeEach((to, from, next) => {
  // 更新标题栏title
  const subtitle = to.meta.title
  if (subtitle) {
    document.title = subtitle
  }
  next()
})

export default router

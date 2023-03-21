import Vue from 'vue'
import VueRouter from 'vue-router'
import AuthView from '../views/AuthView.vue'
import MainView from '../layouts/MainView.vue'
import RegisterView from '../views/RegisterView.vue'
import PlaceView from '../views/PlaceView.vue'
import StateView from '../views/StateView.vue'
import ChatView from '../views/ChatView.vue'
import AboutView from '../views/AboutView.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'auth',
    component: AuthView,
    meta:{
      title:"登录界面"
    }
  },
  {
    path: '/manage',
    name: 'main',
    component: MainView,
    children:[
      {
        path:'place',
        component: PlaceView,
        meta:{
          title:"管理界面"
        },
      },
      {
        path:'state',
        component:StateView,
        meta:{
          title:"管理界面"
        },
      },
      {
        path:'chat',
        component:ChatView,
        meta:{
            title:"管理界面"
        },
      },
      {
        path:'about',
        component:AboutView,
        meta:{
            title:"管理界面"
        },
      }
    ]
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterView,
    meta:{
      title:"注册界面"
    }
  }
]

const router = new VueRouter({
  mode: 'hash',
  base: process.env.BASE_URL,
  routes
})

router.afterEach((to) => {
  document.title = to.meta.title //在全局后置守卫中获取路由元信息设置title
})

export default router

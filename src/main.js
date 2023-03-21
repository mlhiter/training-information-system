import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import axios from './libs/axios'
import VueCookies from 'vue-cookies'
import 'element-ui/lib/theme-chalk/index.css'
import './router/guard.js'
import socket from './libs/socket'

Vue.use(ElementUI)
Vue.use(VueCookies)
Vue.use(socket)

Vue.config.productionTip = false

// 挂载axios
Vue.prototype.$axios = axios

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app')

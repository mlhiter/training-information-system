import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex) //vue的插件机制

//Vuex.Store 构造器选项
const store = new Vuex.Store({
  state: {
    //存放状态
    token: localStorage.getItem('token') ? localStorage.getItem('token') : '',
  },
  mutations: {
    updateToken(state, payload) {
      state.token = payload
      localStorage.setItem('token', payload)
    },
    Token(state) {
      console.log(state.token)
    },
    isToken(state) {
      if (state.token.length > 0) return true
      else return false
    },
  },
}) //this.$store.commit('函数名'，除state外的参数...)
export default store

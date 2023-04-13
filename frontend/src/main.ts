import './styles'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/axios'

const app = createApp(App)

app.use(router)
app.use(store)
app.mount('#app')

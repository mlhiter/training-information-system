import axios from 'axios'
import qs from 'qs'

const axiosInstance = axios.create({
  baseURL: 'https://www.coordinatemanager.xyz:441',
  headers: {
    'Content-Type': 'application/x-www-form-urlencoded',
    authToken: localStorage.getItem('token'),
  },
  timeout: 15000,
  // if(methods="PUT")
})
// FIXME:axios拦截器失效，会造成报错
// axiosInstance.interceptors.request.use(
//   (config) => {
//     config.data = qs.stringify(config.data) // 转为formdata数据格式
//     return config
//   },
//   (error) => Promise.error(error)
// )

// 将包裹data层去除
axiosInstance.interceptors.response.use((resp) => resp.data)
axiosInstance.defaults.withCredentials = false
export default axiosInstance

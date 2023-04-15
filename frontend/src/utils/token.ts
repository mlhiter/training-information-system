// 获取登录状态
export const getAuth = () => {
  return localStorage.getItem('authentication')
}

// 设置登录状态
// accepted和rejected
export const setAuth = (value: string) => {
  localStorage.setItem('authentication', value)
}

// 设置角色
export const setRole = (value: string) => {
  localStorage.setItem('role', value)
}
//获取角色
export const getRole = () => {
  return localStorage.getItem('role')
}
//设置用户名
export const setUser = (value: string) => {
  localStorage.setItem('user', value)
}
//获取用户名
export const getUser = () => {
  return localStorage.getItem('user')
}
// 清除登录信息
// export const clearToken = () => {
//   localStorage.removeItem(ACCESS_TOKEN_KEY)
//   localStorage.removeItem(REFRESH_TOKEN_KEY)
// }

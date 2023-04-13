// 得到登录信息
export const getAuth = (tokenKey) => {
  return localStorage.getItem(tokenKey)
}

// 设置登录信息
export const setAuth = (token: string, tokenKey = ACCESS_TOKEN_KEY) => {
  localStorage.setItem(tokenKey, token)
}

// 清除登录信息
export const clearToken = () => {
  localStorage.removeItem(ACCESS_TOKEN_KEY)
  localStorage.removeItem(REFRESH_TOKEN_KEY)
}

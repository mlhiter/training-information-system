// 该服务为 vercel serve跨域处理
import { createProxyMiddleware } from 'http-proxy-middleware'

export default function (req, res, next) {
  let target = ''

  // 代理目标地址
  // 这里使用 backend 主要用于区分 vercel serverless 的 api 路径
  if (req.url.startsWith('/backend')) {
    target = 'http://mrc.vipgz1.91tunnel.com/api'
  }
  // 创建代理对象并转发请求
  // 安装@types/express和express提供类型支持，不然报错
  createProxyMiddleware({
    target,
    changeOrigin: true,
    pathRewrite: {
      // 通过路径重写，去除请求路径中的 `/backend`
      // 例如 /backend/user/login 将被转发到 http://backend-api.com/user/login
      '^/backend/': '/',
    },
  })(req, res, next)
}

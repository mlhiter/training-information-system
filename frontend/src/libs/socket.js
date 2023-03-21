import SocketIO from 'vue-socket.io'
import ClientSocketIO from 'socket.io-client'
const socketInstance = new SocketIO({
  debug: false, //开启调试模式
  connection: ClientSocketIO.connect('http://39.106.249.239:443', {
    transports: ['websocket'], //默认使用的请求方式
    autoConnect: false, //是否自动连接
  }),
})

export default socketInstance

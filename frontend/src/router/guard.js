import router from "./index.js";

//定义一个列表存放不需登录就能查看的页面地址（白名单）
const whiteRouter = ['/register', '/'];

router.beforeEach((to, from, next) => {
    if (localStorage.getItem('token')) {
        if (to.path == '/') {
            // this.$store.commit('updateToken', '');//清除token
            next();
        }
        next();
    } else {
        if (whiteRouter.indexOf(to.path) != -1) {
            //如果白名单中有就放行
            next();
        } else {
            // 当不满足登陆条件就跳转到登录页面
            alert("请先登录");
            next('/');
        }
    }
})

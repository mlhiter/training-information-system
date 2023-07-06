# 浩奇培训报名管理系统——前端部分

大二时期的系统分析大作业。

## 前台+后台：

**系统结构和功能划分：**

前台软件公司和学员报名、缴费、签到和问卷填写，后台经理和执行人管理培训信息以及现场工作人员进行学生名单管理，包括执行人（报名审核、讲师管理、课程管理、学员管理、问卷管理）、经理（执行人管理、审核软件公司培训申请、汇总报表）、现场工作人员（获取现场学生名单及状态）

**特色和亮点：**

使用 Vite 框架，

1. 更快的加载速度：由于 Vite 可以将应用程序分成更小、更快的代码块，因此在生产环境下，您的应用程序可以更快地加载，提高应用程序的整体性能和用户体验。
2. 更高的缓存效率：由于 Vite 的历史版本控制技术，您可以安全地将静态资源缓存在浏览器中。因此，当用户发生页面切换时，这些资源可以非常快地加载，提高应用程序的整体加载速度和性能。

总结地说，Vite 相比其他框架能够使页面加载和切换更加快速。

使用 Vue3 框架：

1. 对 TS 支持完善，更好的类型检查带来更快的编码速度和错误率，进而提高项目的安全性和功能。
2. 渲染速度更加快速。

使用 Naive UI：

1. 界面更加简约轻松，符合大众审美，带来更好的用户体验。
2. 对 TS 语言的支持更加完善。

**技术栈：**

框架：Vue3 + NaiveUI + Typescript + Vite

语言：html + typescript + vue + css

工具： pnpm（包管理工具）、windicss（CSS 代码原子化工具）、unplugin 系列（vite 精简代码工具）、axios（API 请求包）、pinia（全局数据存储包）、vue-router（vue 路由工具）、sass(CSS 代码预处理器工具包)

## 相关链接：

APIFOX 接口管理部分:https://www.apifox.cn/apidoc/shared-c7f60f96-1de2-4f96-88c3-cc5763809a73

项目 github 地址：https://github.com/mlhiter/training-information-system

项目前端预览地址（vercel 部署，需要科学上网）：https://training-information-system.vercel.app/

## 项目启动

```sh
pnpm install #安装依赖包
pnpm run dev #启动本地服务器
```

浏览器会弹出页面，如果没有弹出界面进入 http://localhost:5173/ 可看到。

## 项目统计

代码量：4867 行代码，45 个文件。

![image-20230415234518545](https://raw.githubusercontent.com/mlhiter/typora-images/master/202304152345674.png)

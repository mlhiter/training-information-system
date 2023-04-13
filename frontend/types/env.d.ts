/// <reference types="vite/client" />
// 声明一些vite特定的全局变量，使Typescript能够解析这些对象的类型

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}

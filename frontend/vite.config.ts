import { defineConfig } from 'vite'
import { resolve } from 'pathe'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import WindiCSS from 'vite-plugin-windicss'
import Icons from 'unplugin-icons/vite'
import IconsResolver from 'unplugin-icons/resolver'
import Components from 'unplugin-vue-components/vite'
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'
// https://vitejs.dev/config/
export default defineConfig({
  //配置代理
  server: {
    open: true, //自动在浏览器启动应用
    host: true,
    proxy: {
      '/api': {
        target: 'http://101.43.219.115:8001',
        changeOrigin: true,
      },
    },
  },
  plugins: [
    vue(),
    //https://github.com/windicss/vite-plugin-windicss
    WindiCSS(),
    //https://github.com/antfu/unplugin-auto-import
    AutoImport({
      imports: [
        'vue',
        'vue-router',
        'pinia',
        '@vueuse/core',
        {
          axios: [
            // default imports
            ['default', 'axios'], // import { default as axios } from 'axios',
          ],
        },
      ],
      dts: 'types/auto-imports.d.ts',
    }),
    //https://github.com/antfu/unplugin-vue-components
    Components({
      extensions: ['vue'],
      dts: 'types/components.d.ts',
      //全局注册组件不需要导入，手动注册类型
      types: [
        {
          from: 'vue-router',
          names: ['RouterLink', 'RouterView'],
        },
      ],
      resolvers: [
        NaiveUiResolver(),
        // icon auto import: {prefix}-{collection}-{icon}
        IconsResolver({
          prefix: 'icon',
          // enabledCollections: ['bx'],
          // customCollections: ['custom'], custom icon: Icons/customCollections
        }),
      ],
    }),
    //https://github.com/antfu/unplugin-icons
    Icons({
      scale: 1.125,
      compiler: 'vue3',
    }),
  ],
  // 简化文件引用路径
  //想使用resolve，要安装pathe
  resolve: {
    alias: [
      {
        find: '~',
        replacement: resolve(__dirname, './'),
      },
      {
        find: '@',
        //要使用__dirname，要安装@types/node
        replacement: resolve(__dirname, './src'),
      },
      {
        find: '#',
        replacement: resolve(__dirname, './types'),
      },
    ],
  },
})


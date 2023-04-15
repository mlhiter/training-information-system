// vite.config.ts
import { defineConfig } from "file:///C:/Users/30764/Desktop/%E4%B8%AA%E4%BA%BA%E9%A1%B9%E7%9B%AE/%E5%BC%80%E5%8F%91%E4%B8%AD/training-information-system/frontend/node_modules/.pnpm/vite@4.2.1_hghbulspu73jfdazs4i6yiqype/node_modules/vite/dist/node/index.js";
import { resolve } from "file:///C:/Users/30764/Desktop/%E4%B8%AA%E4%BA%BA%E9%A1%B9%E7%9B%AE/%E5%BC%80%E5%8F%91%E4%B8%AD/training-information-system/frontend/node_modules/.pnpm/pathe@1.1.0/node_modules/pathe/dist/index.mjs";
import vue from "file:///C:/Users/30764/Desktop/%E4%B8%AA%E4%BA%BA%E9%A1%B9%E7%9B%AE/%E5%BC%80%E5%8F%91%E4%B8%AD/training-information-system/frontend/node_modules/.pnpm/@vitejs+plugin-vue@4.1.0_vite@4.2.1+vue@3.2.47/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import AutoImport from "file:///C:/Users/30764/Desktop/%E4%B8%AA%E4%BA%BA%E9%A1%B9%E7%9B%AE/%E5%BC%80%E5%8F%91%E4%B8%AD/training-information-system/frontend/node_modules/.pnpm/unplugin-auto-import@0.15.2_@vueuse+core@9.13.0/node_modules/unplugin-auto-import/dist/vite.js";
import WindiCSS from "file:///C:/Users/30764/Desktop/%E4%B8%AA%E4%BA%BA%E9%A1%B9%E7%9B%AE/%E5%BC%80%E5%8F%91%E4%B8%AD/training-information-system/frontend/node_modules/.pnpm/vite-plugin-windicss@1.8.10_vite@4.2.1/node_modules/vite-plugin-windicss/dist/index.mjs";
import Icons from "file:///C:/Users/30764/Desktop/%E4%B8%AA%E4%BA%BA%E9%A1%B9%E7%9B%AE/%E5%BC%80%E5%8F%91%E4%B8%AD/training-information-system/frontend/node_modules/.pnpm/unplugin-icons@0.16.1/node_modules/unplugin-icons/dist/vite.mjs";
import IconsResolver from "file:///C:/Users/30764/Desktop/%E4%B8%AA%E4%BA%BA%E9%A1%B9%E7%9B%AE/%E5%BC%80%E5%8F%91%E4%B8%AD/training-information-system/frontend/node_modules/.pnpm/unplugin-icons@0.16.1/node_modules/unplugin-icons/dist/resolver.mjs";
import Components from "file:///C:/Users/30764/Desktop/%E4%B8%AA%E4%BA%BA%E9%A1%B9%E7%9B%AE/%E5%BC%80%E5%8F%91%E4%B8%AD/training-information-system/frontend/node_modules/.pnpm/unplugin-vue-components@0.24.1_vue@3.2.47/node_modules/unplugin-vue-components/dist/vite.mjs";
import { NaiveUiResolver } from "file:///C:/Users/30764/Desktop/%E4%B8%AA%E4%BA%BA%E9%A1%B9%E7%9B%AE/%E5%BC%80%E5%8F%91%E4%B8%AD/training-information-system/frontend/node_modules/.pnpm/unplugin-vue-components@0.24.1_vue@3.2.47/node_modules/unplugin-vue-components/dist/resolvers.mjs";
var __vite_injected_original_dirname = "C:\\Users\\30764\\Desktop\\\u4E2A\u4EBA\u9879\u76EE\\\u5F00\u53D1\u4E2D\\training-information-system\\frontend";
var vite_config_default = defineConfig({
  //配置代理
  server: {
    open: true,
    //自动在浏览器启动应用
    host: true,
    proxy: {
      "/api": {
        target: "http://101.43.177.191:8081",
        changeOrigin: true
      }
    }
  },
  plugins: [
    vue(),
    //https://github.com/windicss/vite-plugin-windicss
    WindiCSS(),
    //https://github.com/antfu/unplugin-auto-import
    AutoImport({
      imports: [
        "vue",
        "vue-router",
        "pinia",
        "@vueuse/core",
        {
          axios: [
            // default imports
            ["default", "axios"]
            // import { default as axios } from 'axios',
          ]
        }
      ],
      dts: "types/auto-imports.d.ts"
    }),
    //https://github.com/antfu/unplugin-vue-components
    Components({
      extensions: ["vue"],
      dts: "types/components.d.ts",
      //全局注册组件不需要导入，手动注册类型
      types: [
        {
          from: "vue-router",
          names: ["RouterLink", "RouterView"]
        }
      ],
      resolvers: [
        NaiveUiResolver(),
        // icon auto import: {prefix}-{collection}-{icon}
        IconsResolver({
          prefix: "icon"
          // enabledCollections: ['bx'],
          // customCollections: ['custom'], custom icon: Icons/customCollections
        })
      ]
    }),
    //https://github.com/antfu/unplugin-icons
    Icons({
      scale: 1.125,
      compiler: "vue3"
    })
  ],
  // 简化文件引用路径
  //想使用resolve，要安装pathe
  resolve: {
    alias: [
      {
        find: "~",
        replacement: resolve(__vite_injected_original_dirname, "./")
      },
      {
        find: "@",
        //要使用__dirname，要安装@types/node
        replacement: resolve(__vite_injected_original_dirname, "./src")
      },
      {
        find: "#",
        replacement: resolve(__vite_injected_original_dirname, "./types")
      }
    ]
  }
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCJDOlxcXFxVc2Vyc1xcXFwzMDc2NFxcXFxEZXNrdG9wXFxcXFx1NEUyQVx1NEVCQVx1OTg3OVx1NzZFRVxcXFxcdTVGMDBcdTUzRDFcdTRFMkRcXFxcdHJhaW5pbmctaW5mb3JtYXRpb24tc3lzdGVtXFxcXGZyb250ZW5kXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ZpbGVuYW1lID0gXCJDOlxcXFxVc2Vyc1xcXFwzMDc2NFxcXFxEZXNrdG9wXFxcXFx1NEUyQVx1NEVCQVx1OTg3OVx1NzZFRVxcXFxcdTVGMDBcdTUzRDFcdTRFMkRcXFxcdHJhaW5pbmctaW5mb3JtYXRpb24tc3lzdGVtXFxcXGZyb250ZW5kXFxcXHZpdGUuY29uZmlnLnRzXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ltcG9ydF9tZXRhX3VybCA9IFwiZmlsZTovLy9DOi9Vc2Vycy8zMDc2NC9EZXNrdG9wLyVFNCVCOCVBQSVFNCVCQSVCQSVFOSVBMSVCOSVFNyU5QiVBRS8lRTUlQkMlODAlRTUlOEYlOTElRTQlQjglQUQvdHJhaW5pbmctaW5mb3JtYXRpb24tc3lzdGVtL2Zyb250ZW5kL3ZpdGUuY29uZmlnLnRzXCI7aW1wb3J0IHsgZGVmaW5lQ29uZmlnIH0gZnJvbSAndml0ZSdcbmltcG9ydCB7IHJlc29sdmUgfSBmcm9tICdwYXRoZSdcbmltcG9ydCB2dWUgZnJvbSAnQHZpdGVqcy9wbHVnaW4tdnVlJ1xuaW1wb3J0IEF1dG9JbXBvcnQgZnJvbSAndW5wbHVnaW4tYXV0by1pbXBvcnQvdml0ZSdcbmltcG9ydCBXaW5kaUNTUyBmcm9tICd2aXRlLXBsdWdpbi13aW5kaWNzcydcbmltcG9ydCBJY29ucyBmcm9tICd1bnBsdWdpbi1pY29ucy92aXRlJ1xuaW1wb3J0IEljb25zUmVzb2x2ZXIgZnJvbSAndW5wbHVnaW4taWNvbnMvcmVzb2x2ZXInXG5pbXBvcnQgQ29tcG9uZW50cyBmcm9tICd1bnBsdWdpbi12dWUtY29tcG9uZW50cy92aXRlJ1xuaW1wb3J0IHsgTmFpdmVVaVJlc29sdmVyIH0gZnJvbSAndW5wbHVnaW4tdnVlLWNvbXBvbmVudHMvcmVzb2x2ZXJzJ1xuLy8gaHR0cHM6Ly92aXRlanMuZGV2L2NvbmZpZy9cbmV4cG9ydCBkZWZhdWx0IGRlZmluZUNvbmZpZyh7XG4gIC8vXHU5MTREXHU3RjZFXHU0RUUzXHU3NDA2XG4gIHNlcnZlcjoge1xuICAgIG9wZW46IHRydWUsIC8vXHU4MUVBXHU1MkE4XHU1NzI4XHU2RDRGXHU4OUM4XHU1NjY4XHU1NDJGXHU1MkE4XHU1RTk0XHU3NTI4XG4gICAgaG9zdDogdHJ1ZSxcbiAgICBwcm94eToge1xuICAgICAgJy9hcGknOiB7XG4gICAgICAgIHRhcmdldDogJ2h0dHA6Ly8xMDEuNDMuMTc3LjE5MTo4MDgxJyxcbiAgICAgICAgY2hhbmdlT3JpZ2luOiB0cnVlLFxuICAgICAgfSxcbiAgICB9LFxuICB9LFxuICBwbHVnaW5zOiBbXG4gICAgdnVlKCksXG4gICAgLy9odHRwczovL2dpdGh1Yi5jb20vd2luZGljc3Mvdml0ZS1wbHVnaW4td2luZGljc3NcbiAgICBXaW5kaUNTUygpLFxuICAgIC8vaHR0cHM6Ly9naXRodWIuY29tL2FudGZ1L3VucGx1Z2luLWF1dG8taW1wb3J0XG4gICAgQXV0b0ltcG9ydCh7XG4gICAgICBpbXBvcnRzOiBbXG4gICAgICAgICd2dWUnLFxuICAgICAgICAndnVlLXJvdXRlcicsXG4gICAgICAgICdwaW5pYScsXG4gICAgICAgICdAdnVldXNlL2NvcmUnLFxuICAgICAgICB7XG4gICAgICAgICAgYXhpb3M6IFtcbiAgICAgICAgICAgIC8vIGRlZmF1bHQgaW1wb3J0c1xuICAgICAgICAgICAgWydkZWZhdWx0JywgJ2F4aW9zJ10sIC8vIGltcG9ydCB7IGRlZmF1bHQgYXMgYXhpb3MgfSBmcm9tICdheGlvcycsXG4gICAgICAgICAgXSxcbiAgICAgICAgfSxcbiAgICAgIF0sXG4gICAgICBkdHM6ICd0eXBlcy9hdXRvLWltcG9ydHMuZC50cycsXG4gICAgfSksXG4gICAgLy9odHRwczovL2dpdGh1Yi5jb20vYW50ZnUvdW5wbHVnaW4tdnVlLWNvbXBvbmVudHNcbiAgICBDb21wb25lbnRzKHtcbiAgICAgIGV4dGVuc2lvbnM6IFsndnVlJ10sXG4gICAgICBkdHM6ICd0eXBlcy9jb21wb25lbnRzLmQudHMnLFxuICAgICAgLy9cdTUxNjhcdTVDNDBcdTZDRThcdTUxOENcdTdFQzRcdTRFRjZcdTRFMERcdTk3MDBcdTg5ODFcdTVCRkNcdTUxNjVcdUZGMENcdTYyNEJcdTUyQThcdTZDRThcdTUxOENcdTdDN0JcdTU3OEJcbiAgICAgIHR5cGVzOiBbXG4gICAgICAgIHtcbiAgICAgICAgICBmcm9tOiAndnVlLXJvdXRlcicsXG4gICAgICAgICAgbmFtZXM6IFsnUm91dGVyTGluaycsICdSb3V0ZXJWaWV3J10sXG4gICAgICAgIH0sXG4gICAgICBdLFxuICAgICAgcmVzb2x2ZXJzOiBbXG4gICAgICAgIE5haXZlVWlSZXNvbHZlcigpLFxuICAgICAgICAvLyBpY29uIGF1dG8gaW1wb3J0OiB7cHJlZml4fS17Y29sbGVjdGlvbn0te2ljb259XG4gICAgICAgIEljb25zUmVzb2x2ZXIoe1xuICAgICAgICAgIHByZWZpeDogJ2ljb24nLFxuICAgICAgICAgIC8vIGVuYWJsZWRDb2xsZWN0aW9uczogWydieCddLFxuICAgICAgICAgIC8vIGN1c3RvbUNvbGxlY3Rpb25zOiBbJ2N1c3RvbSddLCBjdXN0b20gaWNvbjogSWNvbnMvY3VzdG9tQ29sbGVjdGlvbnNcbiAgICAgICAgfSksXG4gICAgICBdLFxuICAgIH0pLFxuICAgIC8vaHR0cHM6Ly9naXRodWIuY29tL2FudGZ1L3VucGx1Z2luLWljb25zXG4gICAgSWNvbnMoe1xuICAgICAgc2NhbGU6IDEuMTI1LFxuICAgICAgY29tcGlsZXI6ICd2dWUzJyxcbiAgICB9KSxcbiAgXSxcbiAgLy8gXHU3QjgwXHU1MzE2XHU2NTg3XHU0RUY2XHU1RjE1XHU3NTI4XHU4REVGXHU1Rjg0XG4gIC8vXHU2MEYzXHU0RjdGXHU3NTI4cmVzb2x2ZVx1RkYwQ1x1ODk4MVx1NUI4OVx1ODhDNXBhdGhlXG4gIHJlc29sdmU6IHtcbiAgICBhbGlhczogW1xuICAgICAge1xuICAgICAgICBmaW5kOiAnficsXG4gICAgICAgIHJlcGxhY2VtZW50OiByZXNvbHZlKF9fZGlybmFtZSwgJy4vJyksXG4gICAgICB9LFxuICAgICAge1xuICAgICAgICBmaW5kOiAnQCcsXG4gICAgICAgIC8vXHU4OTgxXHU0RjdGXHU3NTI4X19kaXJuYW1lXHVGRjBDXHU4OTgxXHU1Qjg5XHU4OEM1QHR5cGVzL25vZGVcbiAgICAgICAgcmVwbGFjZW1lbnQ6IHJlc29sdmUoX19kaXJuYW1lLCAnLi9zcmMnKSxcbiAgICAgIH0sXG4gICAgICB7XG4gICAgICAgIGZpbmQ6ICcjJyxcbiAgICAgICAgcmVwbGFjZW1lbnQ6IHJlc29sdmUoX19kaXJuYW1lLCAnLi90eXBlcycpLFxuICAgICAgfSxcbiAgICBdLFxuICB9LFxufSlcbiJdLAogICJtYXBwaW5ncyI6ICI7QUFBc2MsU0FBUyxvQkFBb0I7QUFDbmUsU0FBUyxlQUFlO0FBQ3hCLE9BQU8sU0FBUztBQUNoQixPQUFPLGdCQUFnQjtBQUN2QixPQUFPLGNBQWM7QUFDckIsT0FBTyxXQUFXO0FBQ2xCLE9BQU8sbUJBQW1CO0FBQzFCLE9BQU8sZ0JBQWdCO0FBQ3ZCLFNBQVMsdUJBQXVCO0FBUmhDLElBQU0sbUNBQW1DO0FBVXpDLElBQU8sc0JBQVEsYUFBYTtBQUFBO0FBQUEsRUFFMUIsUUFBUTtBQUFBLElBQ04sTUFBTTtBQUFBO0FBQUEsSUFDTixNQUFNO0FBQUEsSUFDTixPQUFPO0FBQUEsTUFDTCxRQUFRO0FBQUEsUUFDTixRQUFRO0FBQUEsUUFDUixjQUFjO0FBQUEsTUFDaEI7QUFBQSxJQUNGO0FBQUEsRUFDRjtBQUFBLEVBQ0EsU0FBUztBQUFBLElBQ1AsSUFBSTtBQUFBO0FBQUEsSUFFSixTQUFTO0FBQUE7QUFBQSxJQUVULFdBQVc7QUFBQSxNQUNULFNBQVM7QUFBQSxRQUNQO0FBQUEsUUFDQTtBQUFBLFFBQ0E7QUFBQSxRQUNBO0FBQUEsUUFDQTtBQUFBLFVBQ0UsT0FBTztBQUFBO0FBQUEsWUFFTCxDQUFDLFdBQVcsT0FBTztBQUFBO0FBQUEsVUFDckI7QUFBQSxRQUNGO0FBQUEsTUFDRjtBQUFBLE1BQ0EsS0FBSztBQUFBLElBQ1AsQ0FBQztBQUFBO0FBQUEsSUFFRCxXQUFXO0FBQUEsTUFDVCxZQUFZLENBQUMsS0FBSztBQUFBLE1BQ2xCLEtBQUs7QUFBQTtBQUFBLE1BRUwsT0FBTztBQUFBLFFBQ0w7QUFBQSxVQUNFLE1BQU07QUFBQSxVQUNOLE9BQU8sQ0FBQyxjQUFjLFlBQVk7QUFBQSxRQUNwQztBQUFBLE1BQ0Y7QUFBQSxNQUNBLFdBQVc7QUFBQSxRQUNULGdCQUFnQjtBQUFBO0FBQUEsUUFFaEIsY0FBYztBQUFBLFVBQ1osUUFBUTtBQUFBO0FBQUE7QUFBQSxRQUdWLENBQUM7QUFBQSxNQUNIO0FBQUEsSUFDRixDQUFDO0FBQUE7QUFBQSxJQUVELE1BQU07QUFBQSxNQUNKLE9BQU87QUFBQSxNQUNQLFVBQVU7QUFBQSxJQUNaLENBQUM7QUFBQSxFQUNIO0FBQUE7QUFBQTtBQUFBLEVBR0EsU0FBUztBQUFBLElBQ1AsT0FBTztBQUFBLE1BQ0w7QUFBQSxRQUNFLE1BQU07QUFBQSxRQUNOLGFBQWEsUUFBUSxrQ0FBVyxJQUFJO0FBQUEsTUFDdEM7QUFBQSxNQUNBO0FBQUEsUUFDRSxNQUFNO0FBQUE7QUFBQSxRQUVOLGFBQWEsUUFBUSxrQ0FBVyxPQUFPO0FBQUEsTUFDekM7QUFBQSxNQUNBO0FBQUEsUUFDRSxNQUFNO0FBQUEsUUFDTixhQUFhLFFBQVEsa0NBQVcsU0FBUztBQUFBLE1BQzNDO0FBQUEsSUFDRjtBQUFBLEVBQ0Y7QUFDRixDQUFDOyIsCiAgIm5hbWVzIjogW10KfQo=

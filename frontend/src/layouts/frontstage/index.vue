<template>
  <div class="h-100vh w-100vw bg-#dee2e6 grid place-items-center">
    <n-layout
      content-style="padding: 24px;"
      class="w-4/5 h-4/5 shadow-md rounded-2xl">
      <n-layout-header bordered class="h-60px flex items-center">
        <n-image src="/favicon.svg" class="w-150px"></n-image>
        <div class="text-3xl font-mono">浩奇培训报名管理系统——前台</div>
        <div class="ml-540px"><User /></div>
      </n-layout-header>
      <n-layout has-sider class="h-4/5">
        <n-layout-sider bordered content-style="padding: 24px;">
          <n-menu :options="menuOptions" v-model:value="activeMenuRef" />
        </n-layout-sider>
        <n-layout-content
          content-style="padding: 24px;"
          :native-scrollbar="false">
          <router-view v-slot="{ Component }">
            <transition :duration="200" name="fade-top" mode="out-in">
              <keep-alive>
                <component :is="Component" />
              </keep-alive>
            </transition>
          </router-view>
        </n-layout-content>
      </n-layout>
    </n-layout>
  </div>
</template>

<script lang="ts" setup>
import { MenuOption, NIcon } from 'naive-ui'
import { RouterLink } from 'vue-router'
import { BookOutline as BookIcon } from '@vicons/ionicons5'
import User from '../components/user.vue'
function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) })
}
const menuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: { path: '/frontstage/enroll' },
        },
        { default: () => '报名' }
      ),
    key: 'enroll',
    icon: renderIcon(BookIcon),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: { path: '/frontstage/questionnaire' },
        },
        { default: () => '调查问卷' }
      ),
    key: 'questionnaire',
    icon: renderIcon(BookIcon),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: { path: '/frontstage/profile' },
        },
        { default: () => '个人信息' }
      ),
    key: 'profile',
    icon: renderIcon(BookIcon),
  },
]
//FIXME: 路由刷新后菜单项的活跃路由问题
// const route = useRoute()
const activeMenuRef = ref('enroll')
// onMounted(() => {
//   const activeMenu = menuOptions.find(
//     (option: MenuOption) => route.path === option.label().props.to.path
//   )
//   // 设置 activeMenu 值
//   activeMenuRef.value = activeMenu?.key as string
// })
</script>

<style lang="sass" scoped></style>

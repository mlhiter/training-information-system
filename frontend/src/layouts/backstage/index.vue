<template>
  <div class="h-100vh w-100vw bg-#dee2e6 grid place-items-center">
    <n-layout
      content-style="padding: 24px;"
      class="w-4/5 h-4/5 shadow-md rounded-2xl">
      <n-layout-header bordered class="h-60px flex items-center">
        <n-image src="/favicon.svg" class="w-150px"></n-image>
        <div class="text-3xl font-mono">浩奇培训信息管理系统——后台</div>
        <div class="ml-540px"><User /></div>
      </n-layout-header>
      <n-layout has-sider class="h-4/5">
        <n-layout-sider bordered content-style="padding: 24px;">
          <n-menu :options="menuOptions" default-value="index" />
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
import { getRole } from '@/utils/token'
import { useRouter } from 'vue-router'
import User from '../components/user.vue'

const role = getRole()
const router = useRouter()
// const route = useRoute()
// const activeMenukey = ref('')
function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) })
}
let menuOptions: MenuOption[] = []
const executorMenuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: { path: '/backstage/executor/enroll' },
        },
        { default: () => '个人报名审核' }
      ),
    key: 'index',
    icon: renderIcon(BookIcon),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: { path: '/backstage/executor/lecturer' },
        },
        { default: () => '讲师管理' }
      ),
    key: 'lecturer',
    icon: renderIcon(BookIcon),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: { path: '/backstage/executor/questionnaire' },
        },
        { default: () => '问卷管理' }
      ),
    key: 'questionnaire',
    icon: renderIcon(BookIcon),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: { path: '/backstage/executor/student' },
        },
        { default: () => '学员管理' }
      ),
    key: 'student',
    icon: renderIcon(BookIcon),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: { path: '/backstage/executor/training' },
        },
        { default: () => '课程管理' }
      ),
    key: 'training',
    icon: renderIcon(BookIcon),
  },
]
const managerMenuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: { path: '/backstage/manager/report' },
        },
        { default: () => '汇总报表' }
      ),
    key: 'report',
    icon: renderIcon(BookIcon),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: { path: '/backstage/manager/executor' },
        },
        { default: () => '执行人管理' }
      ),
    key: 'executor',
    icon: renderIcon(BookIcon),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: { path: '/backstage/manager/user' },
        },
        { default: () => '用户管理' }
      ),
    key: 'user',
    icon: renderIcon(BookIcon),
  },
]
const operatorMenuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: { path: '/backstage/operator' },
        },
        { default: () => '学生名单' }
      ),
    key: 'index',
    icon: renderIcon(BookIcon),
  },
]
onBeforeMount(() => {
  console.log(role)
  switch (role) {
    case 'executor':
      menuOptions = executorMenuOptions
      router.push('/backstage/executor/enroll')
      break
    case 'manager':
      menuOptions = managerMenuOptions
      router.push('/backstage/manager/report')
      break
    case 'operator':
      menuOptions = operatorMenuOptions
      router.push('/backstage/operator')
      break
    default:
      break
  }
})
// onMounted(() => {
//    menuOptions.forEach((menuOption) => {
//     if (menuOption.label().props.to.path == route.path) {
//       activeMenukey.value = menuOption.key as string
//     }
//   })
// })
</script>

<style lang="sass" scoped></style>

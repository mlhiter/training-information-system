<template>
  <n-dropdown :options="options" class="font-serif text-3xl" @select="logout">
    <div class="flex-row flex justify-center items-center">
      <n-icon size="30">
        <icon-bx-user></icon-bx-user>
      </n-icon>
      <div>{{ username }}</div>
    </div>
  </n-dropdown>
</template>

<script lang="ts" setup>
import { NIcon, useMessage } from 'naive-ui'
import { LogOutOutline as LogoutIcon } from '@vicons/ionicons5'
import router from '@/router'
import { getUser } from '@/utils/token'
const message = useMessage()
const username = getUser()
const renderIcon = (icon: Component) => {
  return () => {
    return h(NIcon, null, {
      default: () => h(icon),
    })
  }
}
const options = [
  {
    label: '退出登录',
    key: 'logout',
    icon: renderIcon(LogoutIcon),
  },
]
function logout(key: string) {
  message.success('退出登录成功！')
  router.push('/')
}
</script>

<style lang="sass" scoped></style>

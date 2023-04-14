<template>
  <div>
    <!-- 占位div -->
    <div class="w-20 h-30"></div>
    <n-card class="w-120" style="margin: 0 auto">
      <n-space class="mb-10" align="center" justify="center">
        <n-image width="50" src="/favicon.svg"></n-image>
        <div class="text-3xl font-serif text-center">浩奇培训信息管理系统</div>
      </n-space>
      <n-tabs
        class="pl-4px"
        default-value="signin"
        size="large"
        animated
        style="margin: 0 -4px"
        pane-style="padding-left: 4px; padding-right: 4px; box-sizing: border-box;">
        <n-tab-pane name="signin" tab="登录">
          <n-form>
            <n-form-item-row label="用户名">
              <n-input placeholder="请输入用户名" v-model:value="username" />
            </n-form-item-row>
            <n-form-item-row label="密码">
              <n-input placeholder="请输入密码" v-model:value="password" />
            </n-form-item-row>
            <n-form-item-row label="角色">
              <n-radio-group v-model:value="role" name="radiogroup">
                <n-space>
                  <n-radio
                    v-for="role in roles"
                    :key="role.value"
                    :value="role.value">
                    {{ role.label }}
                  </n-radio>
                </n-space>
              </n-radio-group>
            </n-form-item-row>
          </n-form>
          <n-button type="primary" block secondary strong @click="login">
            登录
          </n-button>
        </n-tab-pane>
        <n-tab-pane name="signup" tab="注册">
          <n-form>
            <n-form-item-row label="用户名">
              <n-input placeholder="请输入用户名" />
            </n-form-item-row>
            <n-form-item-row label="密码">
              <n-input placeholder="请输入密码" />
            </n-form-item-row>
            <n-form-item-row label="重复密码">
              <n-input placeholder="请再次输入密码" />
            </n-form-item-row>
          </n-form>
          <n-button type="primary" block secondary strong>注册</n-button>
        </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import { setRole } from '@/utils/token'
import { useMessage } from 'naive-ui'
import { useRouter } from 'vue-router'
const router = useRouter()
const message = useMessage()
//登录
const roles = [
  {
    label: '用户',
    value: 'user',
  },
  {
    label: '执行人',
    value: 'executor',
  },
  {
    label: '经理',
    value: 'manager',
  },
  {
    label: '现场工作人员',
    value: 'operator',
  },
]
const role = ref<string>('')
const username = ref<string>('')
const password = ref<string>('')

const login = async () => {
  try {
    const res = await axios.post('/api/login', {
      username: username.value,
      password: password.value,
      role: role.value,
    })
    if (res.data.msg === '登录成功') {
      setRole(role.value)
      if (role.value === 'user') {
        router.push('/frontstage')
      } else {
        router.push('/backstage')
      }
    }
  } catch (error) {
    message.error('登录失败！')
  }
}
</script>

<style lang="sass" scoped></style>

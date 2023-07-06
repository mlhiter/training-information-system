<template>
  <div>
    <n-card title="个人信息" v-if="change">
      姓名:{{ profile.name }}
      <br />
      性别:{{ profile.sex }}
      <br />
      邮箱:{{ profile.contact }}
      <br />
      公司:{{ profile.company }}
      <br />
      职位:{{ profile.post }}
      <br />
      技能水平:{{ profile.level }}
      <br />
      <template #action>
        <n-button attr-type="button" type="success" @click="change = false">
          修改个人信息
        </n-button>
      </template>
    </n-card>
    <n-card title="请填写个人信息：" style="margin-bottom: 16px" v-else>
      <n-form ref="formRef" :label-width="80" :model="individualFormValue">
        <n-form-item label="姓名" path="name">
          <n-input
            v-model:value="individualFormValue.name"
            placeholder="输入姓名" />
        </n-form-item>
        <n-form-item label="性别" path="sex">
          <n-input
            v-model:value="individualFormValue.sex"
            placeholder="请输入性别" />
        </n-form-item>
        <n-form-item label="公司" path="company">
          <n-input
            v-model:value="individualFormValue.company"
            placeholder="请输入公司" />
        </n-form-item>
        <n-form-item label="职位" path="post">
          <n-input
            v-model:value="individualFormValue.post"
            placeholder="请输入职位" />
        </n-form-item>
        <n-form-item label="技能水平" path="level">
          <n-input
            v-model:value="individualFormValue.level"
            placeholder="请输入技能水平" />
        </n-form-item>
        <n-form-item label="邮箱" path="contact">
          <n-input
            v-model:value="individualFormValue.contact"
            placeholder="请输入邮箱" />
        </n-form-item>
        <n-form-item>
          <n-button attr-type="button" type="success" @click="handleProfile">
            提交个人信息
          </n-button>
        </n-form-item>
      </n-form>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import { getUser } from '@/utils/token'
import { useMessage } from 'naive-ui'
const message = useMessage()
const change = ref(false)
const profile = ref({
  name: '',
  sex: '',
  company: '',
  post: '',
  level: '',
  contact: '',
})
const individualFormValue = ref({
  name: '',
  sex: '',
  company: '',
  post: '',
  level: '',
  contact: '',
})
const username = getUser()
const handleProfile = async () => {
  try {
    const res = await axios.post(`/backend/user/info`, {
      username,
      ...individualFormValue.value,
    })
    if (res.data.msg == '提交成功') {
      fetchProfileData()
      message.success('已提交个人信息')
      change.value = true
    }
  } catch (error) {
    console.log(error)
    message.error('提交失败')
    change.value = true
  }
}
const fetchProfileData = async () => {
  try {
    const res = await axios.get(`/backend/user/info/get?username=${username}`)
    if (res.data.code == 200) {
      change.value = true
      profile.value = res.data.data
    }
  } catch (error) {
    // 拿不到数据就显示填写个人信息界面
    change.value = false
    console.log(error)
  }
}
fetchProfileData()
onDeactivated(() => {
  change.value = true
})
</script>

<style lang="sass" scoped></style>

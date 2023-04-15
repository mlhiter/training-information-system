<template>
  <div>
    <n-card
      title="请选择你想要报名的种类："
      style="margin-bottom: 16px"
      v-if="change">
      <n-tabs type="line" animated>
        <n-tab-pane name="individual" tab="个人">
          <n-form ref="formRef" :label-width="80" :model="individualFormValue">
            <n-form-item label="课程名称" path="courseName">
              <n-input
                v-model:value="individualFormValue.courseName"
                placeholder="输入课程名称" />
            </n-form-item>
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
              <n-button
                attr-type="button"
                type="success"
                @click="handleEnrollIndividual">
                提交报名信息
              </n-button>
            </n-form-item>
          </n-form>
        </n-tab-pane>
        <n-tab-pane name="organization" tab="组织">
          <n-form :label-width="80" :model="organizationFormValue">
            <n-form-item label="公司名称" path="applicant">
              <n-input
                v-model:value="organizationFormValue.applicant"
                placeholder="输入公司名称" />
            </n-form-item>
            <n-form-item label="申请时间" path="trainingDate">
              <n-input
                v-model:value="organizationFormValue.trainingDate"
                placeholder="输入申请时间" />
            </n-form-item>
            <n-form-item label="申请课程" path="trainingContext">
              <n-input
                v-model:value="organizationFormValue.trainingContext"
                placeholder="请输入申请课程" />
            </n-form-item>
            <n-form-item label="申请人数" path="numberOfPeople">
              <n-input
                v-model:value="organizationFormValue.numberOfPeople"
                placeholder="请输入申请人数" />
            </n-form-item>
            <n-form-item>
              <n-button
                attr-type="button"
                type="success"
                @click="handleEnrollOrganization">
                提交报名信息
              </n-button>
            </n-form-item>
          </n-form>
        </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router'
const change = ref(true)
const router = useRouter()
const individualFormValue = ref({
  courseName: '',
  name: '',
  sex: '',
  company: '',
  post: '',
  level: '',
  contact: '',
  status: 'unhandled',
})
const organizationFormValue = ref({
  applicant: '',
  trainingDate: '',
  trainingContext: '',
  numberOfPeople: null,
  status: 'unhandled',
})
const handleEnrollIndividual = async () => {
  try {
    const res = await axios.post(
      '/api/enroll/individual',
      individualFormValue.value
    )
    if (res.data.msg == 'success') {
      change.value = false
      router.push('/frontstage/pay')
    }
  } catch (error) {
    console.log(error)
  }
}
const handleEnrollOrganization = async () => {
  try {
    const res = await axios.post(
      '/api/enroll/organisation',
      organizationFormValue.value
    )
    if (res.data.msg == 'success') {
      change.value = false
      router.push('/frontstage/pay')
    }
  } catch (error) {
    console.log(error)
  }
}
onDeactivated(() => {
  change.value = true
})
</script>

<style lang="sass" scoped></style>

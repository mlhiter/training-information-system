<template>
  <div>
    <n-card title="调查问卷" style="margin-bottom: 16px" v-if="change">
      <n-form ref="formRef" :label-width="80" :model="formValue">
        <n-form-item label="姓名" path="studentName">
          <n-input
            v-model:value="formValue.studentName"
            placeholder="输入姓名" />
        </n-form-item>
        <n-form-item label="课程名" path="courseName">
          <n-input
            v-model:value="formValue.courseName"
            placeholder="请输入课程名" />
        </n-form-item>
        <n-form-item label="满意度" path="satisfaction">
          <n-input-number
            v-model:value="formValue.satisfaction"
            placeholder="请输入满意度（0-5）" />
        </n-form-item>
        <n-form-item label="评价" path="comment">
          <n-input v-model:value="formValue.comment" placeholder="请输入评价" />
        </n-form-item>
        <n-form-item label="建议" path="suggestion">
          <n-input
            v-model:value="formValue.suggestion"
            placeholder="请输入建议" />
        </n-form-item>
        <n-form-item>
          <n-button attr-type="button" type="success" @click="handleSubmit">
            提交调查问卷
          </n-button>
        </n-form-item>
      </n-form>
    </n-card>
    <n-alert title="Success" type="success" v-else>问卷提交成功！</n-alert>
  </div>
</template>

<script lang="ts" setup>
const formValue = ref({
  studentName: '',
  courseName: '',
  satisfaction: null,
  comment: '',
  suggestion: '',
})
const change = ref(true)
const handleSubmit = async () => {
  try {
    const res = await axios.post('/backend/questionnaire', formValue.value)
    if (res.data.msg == 'success') {
      change.value = false
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

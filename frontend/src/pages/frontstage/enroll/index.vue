<template>
  <div>
    <n-card style="margin-bottom: 16px" v-if="change">
      <n-tabs type="line" animated>
        <n-tab-pane name="individual" tab="个人">
          <n-form ref="formRef" :label-width="80">
            <n-form-item label="课程名称" path="courseName">
              <n-select v-model:value="courseName" :options="options" />
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
      </n-tabs>
    </n-card>
    <n-alert title="Success" type="success" v-else>
      报名成功，请前往缴费页面进行缴费。
    </n-alert>
  </div>
</template>

<script lang="ts" setup>
const change = ref(true)
const courseName = ref('')
const options = ref([{ label: '', value: '' }])
const handleEnrollIndividual = async () => {
  try {
    const res = await axios.post('/backend/enroll/individual', courseName.value)
    if (res.data.msg == 'success') {
      change.value = false
    }
  } catch (error) {
    console.log(error)
  }
}
const fetchUnselectedCourseList = async () => {
  try {
    const res = await axios.get('/backend/course')
    if (res.data.msg == 'success') {
      options.value = Array.from(res.data.data, (item) => {
        return { label: item as string, value: item as string }
      })
    }
  } catch (error) {
    console.log(error)
  }
}
onMounted(() => {
  fetchUnselectedCourseList()
})
onDeactivated(() => {
  change.value = true
})
</script>

<style lang="sass" scoped></style>

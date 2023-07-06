<template>
  <div>
    <n-card style="margin-bottom: 16px">
      <n-tabs type="line" animated>
        <n-tab-pane name="individual" tab="个人">
          <n-card title="已选课程">
            <template #header-extra>注意：先填写个人信息再选课</template>
            <n-tag
              class="mr-3"
              v-for="(item, index) in selectedCourses"
              type="success"
              :key="index">
              {{ item }}
            </n-tag>
            <template #action>
              <n-button type="primary" @click="showModal = true">
                报名新课程
              </n-button>
            </template>
          </n-card>
        </n-tab-pane>
      </n-tabs>
    </n-card>
    <!-- <n-alert title="Success" type="success" v-else>
      报名成功，请前往缴费页面进行缴费。
    </n-alert> -->
    <n-modal
      v-model:show="showModal"
      style="width: 600px"
      class="custom-card"
      preset="card"
      title="新增课程"
      size="huge"
      :bordered="false">
      <n-form ref="formRef" :label-width="80">
        <n-form-item label="课程名称" path="courseName">
          <n-select v-model:value="courseName" :options="options" />
        </n-form-item>
      </n-form>
      <template #action>
        <n-button
          attr-type="button"
          type="success"
          @click="handleEnrollIndividual">
          提交
        </n-button>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { useMessage } from 'naive-ui'

const showModal = ref(false)
const change = ref(true)
const courseName = ref('')
const message = useMessage()
const options = ref([{ label: '', value: '' }])
const selectedCourses = ref(['算法设计与实现', 'C语言', 'C++'])
const handleEnrollIndividual = async () => {
  try {
    const res = await axios.post('/backend/enroll/individual', courseName.value)
    if (res.data.msg == 'success') {
      change.value = false
      showModal.value = false
    }
  } catch (error) {
    message.error('提交失败，请先填写个人信息！')
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

<template>
  <div>
    <n-card
      title="问卷列表"
      :bordered="false"
      header-style="font-size:25px"
      header-extra-style="margin-right:23px">
      <template #header-extra>
        <n-button
          strong
          secondary
          type="primary"
          size="small"
          @click="showModal = true">
          新增问卷
        </n-button>
      </template>
      <n-data-table
        :columns="columns"
        :data="renderData"
        :pagination="pagination"
        :bordered="false" />
    </n-card>
    <n-modal
      v-model:show="showModal"
      style="width: 600px"
      class="custom-card"
      preset="card"
      title="新增问卷"
      size="huge"
      :bordered="false">
      <n-form ref="formRef" :label-width="80" :model="formValue">
        <n-form-item label="学生姓名" path="studentName">
          <n-input
            v-model:value="formValue.studentName"
            placeholder="输入学生姓名" />
        </n-form-item>
        <n-form-item label="课程名称" path="courseName">
          <n-input
            v-model:value="formValue.courseName"
            placeholder="请输入课程名称" />
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
      </n-form>
      <template #action>
        <n-button type="primary" @click="addQuestionnaire">提交</n-button>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { DataTableColumns, NButton, useMessage } from 'naive-ui'
const message = useMessage()
const pagination = {
  pageSize: 10,
}
const columns: DataTableColumns = [
  {
    title: '问卷ID',
    key: 'ID',
  },
  {
    title: '课程ID',
    key: 'courseId',
  },
  {
    title: '学生ID',
    key: 'studentId',
  },
  {
    title: '满意度',
    key: 'satisfaction',
  },
  {
    title: '评价',
    key: 'comment',
  },
  {
    title: '建议',
    key: 'suggestion',
  },
  {
    title: '操作',
    key: 'actions',
    render(row: any) {
      return h(
        NButton,
        {
          strong: true,
          size: 'small',
          secondary: true,
          type: 'error',
          onClick: () => deleteQuestionnaire(row.ID),
        },
        { default: () => '删除问卷' }
      )
    },
  },
]
const renderData = ref([])
const fetchRenderData = async () => {
  try {
    const res = await axios.get('/api/questionnaire/list')
    renderData.value = res.data.data
  } catch (error) {
    console.log(error)
  }
}
fetchRenderData()

const deleteQuestionnaire = async (id: number) => {
  try {
    const res = await axios.request({
      url: '/api/questionnaire/delete',
      method: 'delete',
      data: {
        id,
      },
    })
    if (res.data.msg === 'success') {
      message.success('删除成功！')
      fetchRenderData()
    }
  } catch (error) {
    console.log(error)
  }
}
const showModal = ref(false)
const addQuestionnaire = async () => {
  try {
    const res = await axios.post('/api/questionnaire/add', formValue.value)
    if (res.data.msg === 'success') {
      showModal.value = false
      message.success('新增成功！')
      fetchRenderData()
    }
  } catch (error) {
    message.error('没有这个学生！')
    console.log(error)
  }
}
const formValue = ref({
  studentName: '',
  courseName: '',
  satisfaction: '',
  comment: '',
  suggestion: '',
})
</script>

<style lang="sass" scoped></style>

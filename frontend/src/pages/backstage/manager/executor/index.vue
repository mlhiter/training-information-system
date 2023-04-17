<template>
  <div>
    <n-card
      title="执行人列表"
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
          新增执行人
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
      title="新增执行人"
      size="huge"
      :bordered="false">
      <n-form ref="formRef" :label-width="80" :model="formValue">
        <n-form-item label="执行人姓名" path="name">
          <n-input
            v-model:value="formValue.name"
            placeholder="输入执行人姓名" />
        </n-form-item>
        <n-form-item label="管理课程" path="course">
          <n-input v-model:value="formValue.course" placeholder="请输入职称" />
        </n-form-item>
        <n-form-item label="工作状况" path="condition">
          <n-input
            v-model:value="formValue.condition"
            placeholder="请输入职称" />
        </n-form-item>
      </n-form>
      <template #action>
        <n-button type="primary" @click="addExecutor">提交</n-button>
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
    title: '执行人ID',
    key: 'ID',
  },
  {
    title: '姓名',
    key: 'name',
  },
  {
    title: '管理课程',
    key: 'course',
  },
  {
    title: '工作状况',
    key: 'condition',
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
          onClick: () => deleteExecutor(row.ID),
        },
        { default: () => '删除执行人' }
      )
    },
  },
]
const renderData = ref([])
const fetchRenderData = async () => {
  try {
    const res = await axios.get('/backend/executor/list')
    renderData.value = res.data.data
  } catch (error) {
    console.log(error)
  }
}
fetchRenderData()

const deleteExecutor = async (id: number) => {
  try {
    const res = await axios.request({
      url: '/backend/executor/delete',
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
const addExecutor = async () => {
  try {
    const res = await axios.post('/backend/executor/add', formValue.value)
    if (res.data.msg === 'success') {
      showModal.value = false
      message.success('新增成功！')
      fetchRenderData()
    }
  } catch (error) {
    console.log(error)
  }
}
const formValue = ref({
  name: '',
  course: '',
  condition: '',
})
</script>

<style lang="sass" scoped></style>

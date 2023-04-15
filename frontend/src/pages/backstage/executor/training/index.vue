<template>
  <div>
    <n-card
      title="课程列表"
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
          新增课程
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
      title="新增课程"
      size="huge"
      :bordered="false">
      <n-form ref="formRef" :label-width="80" :model="formValue">
        <n-form-item label="课程名" path="name">
          <n-input v-model:value="formValue.name" placeholder="输入课程名" />
        </n-form-item>
        <n-form-item label="时间" path="time">
          <n-input v-model:value="formValue.time" placeholder="请输入时间" />
        </n-form-item>
        <n-form-item label="地点" path="place">
          <n-input v-model:value="formValue.place" placeholder="请输入地点" />
        </n-form-item>
        <n-form-item label="讲师id" path="lecturerId">
          <n-input
            v-model:value="formValue.lecturerId"
            placeholder="请输入讲师id" />
        </n-form-item>
        <n-form-item label="讲师名" path="lecturer">
          <n-input
            v-model:value="formValue.lecturer"
            placeholder="请输入讲师" />
        </n-form-item>
        <n-form-item label="课程价格" path="price">
          <n-input-number
            v-model:value="formValue.price"
            placeholder="请输入课程价格"
            clearable />
        </n-form-item>
      </n-form>
      <template #action>
        <n-button type="primary" @click="addCourse">提交</n-button>
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
    title: '课程ID',
    key: 'ID',
  },
  {
    title: '课程名',
    key: 'name',
  },
  {
    title: '时间',
    key: 'time',
  },
  {
    title: '地点',
    key: 'place',
  },

  {
    title: '讲师ID',
    key: 'lecturerId',
  },
  {
    title: '讲师名称',
    key: 'lecturer',
  },
  {
    title: '价格',
    key: 'price',
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
          onClick: () => deleteCourse(row.ID),
        },
        { default: () => '删除课程' }
      )
    },
  },
]
const renderData = ref([])
const fetchRenderData = async () => {
  try {
    const res = await axios.get('/api/training/list')
    renderData.value = res.data.data
  } catch (error) {
    console.log(error)
  }
}
fetchRenderData()

const deleteCourse = async (id: number) => {
  try {
    const res = await axios.request({
      url: '/api/training/delete',
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
//FIXME:新增api有问题
const addCourse = async () => {
  try {
    const res = await axios.post('/api/training/add', formValue.value)
    if (res.data.msg === 'success') {
      showModal.value = false
      message.success('新增成功！')
      fetchRenderData()
    }
  } catch (error) {
    console.log(error)
    showModal.value = false
    message.success('新增成功！')
  }
}
const formValue = ref({
  name: '',
  time: '',
  place: '',
  lecturerId: '',
  lecturer: '',
  price: null,
})
</script>

<style lang="sass" scoped></style>

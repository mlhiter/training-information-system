<template>
  <div>
    <n-card
      title="学员列表"
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
          新增学员
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
      title="新增学员"
      size="huge"
      :bordered="false">
      <n-form ref="formRef" :label-width="80" :model="formValue">
        <n-form-item label="学员姓名" path="name">
          <n-input v-model:value="formValue.name" placeholder="输入学员姓名" />
        </n-form-item>
        <n-form-item label="性别" path="sex">
          <n-input v-model:value="formValue.sex" placeholder="请输入性别" />
        </n-form-item>
        <n-form-item label="邮箱" path="contact">
          <n-input v-model:value="formValue.contact" placeholder="请输入邮箱" />
        </n-form-item>
        <n-form-item label="公司" path="company">
          <n-input v-model:value="formValue.company" placeholder="请输入公司" />
        </n-form-item>
        <n-form-item label="职称" path="post">
          <n-input v-model:value="formValue.post" placeholder="请输入职称" />
        </n-form-item>
        <n-form-item label="技术水平" path="level">
          <n-input v-model:value="formValue.level" placeholder="请输入职称" />
        </n-form-item>
      </n-form>
      <template #action>
        <n-button type="primary" @click="addStudent">提交</n-button>
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
    title: '学员ID',
    key: 'ID',
  },
  {
    title: '姓名',
    key: 'name',
  },
  {
    title: '性别',
    key: 'sex',
  },
  {
    title: '邮箱',
    key: 'contact',
  },
  {
    title: '公司',
    key: 'company',
  },
  {
    title: '职位',
    key: 'post',
  },
  {
    title: '技术水平',
    key: 'level',
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
          onClick: () => deleteStudent(row.ID),
        },
        { default: () => '删除学员' }
      )
    },
  },
]
const renderData = ref([])
const fetchRenderData = async () => {
  try {
    const res = await axios.get('/api/student/list')
    renderData.value = res.data.data
  } catch (error) {
    console.log(error)
  }
}
fetchRenderData()

const deleteStudent = async (id: number) => {
  try {
    const res = await axios.request({
      url: '/api/student/delete',
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
const addStudent = async () => {
  try {
    const res = await axios.post('/api/student/add', formValue.value)
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
  sex: '',
  contact: '',
  company: '',
  post: '',
  level: '',
})
</script>

<style lang="sass" scoped></style>

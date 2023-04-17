<template>
  <div>
    <n-card
      title="讲师列表"
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
          新增讲师
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
      title="新增讲师"
      size="huge"
      :bordered="false">
      <n-form ref="formRef" :label-width="80" :model="formValue">
        <n-form-item label="讲师姓名" path="name">
          <n-input v-model:value="formValue.name" placeholder="输入讲师姓名" />
        </n-form-item>
        <n-form-item label="职称" path="profession">
          <n-input
            v-model:value="formValue.profession"
            placeholder="请输入职称" />
        </n-form-item>
        <n-form-item label="领域名" path="field">
          <n-input v-model:value="formValue.field" placeholder="请输入领域名" />
        </n-form-item>
        <n-form-item label="邮箱" path="email">
          <n-input v-model:value="formValue.email" placeholder="请输入邮箱" />
        </n-form-item>
        <n-form-item label="手机号" path="telephone">
          <n-input
            v-model:value="formValue.telephone"
            placeholder="请输入手机号" />
        </n-form-item>
      </n-form>
      <template #action>
        <n-button type="primary" @click="addLecturer">提交</n-button>
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
    title: '讲师ID',
    key: 'ID',
  },
  {
    title: '姓名',
    key: 'name',
  },
  {
    title: '职称',
    key: 'profession',
  },
  {
    title: '领域',
    key: 'field',
  },
  {
    title: '邮箱',
    key: 'email',
  },
  {
    title: '手机',
    key: 'telephone',
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
          onClick: () => deleteLecturer(row.ID),
        },
        { default: () => '删除讲师' }
      )
    },
  },
]
const renderData = ref([])
const fetchRenderData = async () => {
  try {
    const res = await axios.get('/backend/lecturer/list')
    renderData.value = res.data.data
  } catch (error) {
    console.log(error)
  }
}
fetchRenderData()

const deleteLecturer = async (id: number) => {
  try {
    const res = await axios.request({
      url: '/backend/lecturer/delete',
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
const addLecturer = async () => {
  try {
    const res = await axios.post('/backend/lecturer/add', formValue.value)
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
  profession: '',
  field: '',
  email: '',
  telephone: '',
})
</script>

<style lang="sass" scoped></style>

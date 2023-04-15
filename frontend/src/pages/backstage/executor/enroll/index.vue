<template>
  <div>
    <n-data-table
      :columns="columns"
      :data="renderData"
      :pagination="pagination"
      :bordered="false" />
  </div>
</template>

<script lang="ts" setup>
import { DataTableColumns, NButton, NDropdown, useMessage } from 'naive-ui'
const message = useMessage()
const pagination = {
  pageSize: 10,
}
const selectOptions = [
  { label: '通过', key: 'passed' },
  { label: '拒绝', key: 'rejected' },
]
const columns: DataTableColumns = [
  {
    title: '报名ID',
    key: 'ID',
  },
  {
    title: '姓名',
    key: 'name',
  },
  {
    title: '课程名',
    key: 'courseName',
  },
  {
    title: '性别',
    key: 'sex',
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
    title: '邮箱',
    key: 'contact',
  },
  {
    title: '审核状态',
    key: 'status',
  },
  {
    title: '操作',
    key: 'actions',
    render(row: any) {
      return h(
        NDropdown,
        {
          trigger: 'click',
          options: selectOptions,
          onSelect: (key: string) => changeStatus(row.ID, key),
        },
        [
          h(
            NButton,
            {
              strong: true,
              size: 'small',
              secondary: true,
              type: 'success',
            },
            { default: () => '更改审核状态' }
          ),
        ]
      )
    },
  },
]
const renderData = ref([])
const fetchRenderData = async () => {
  try {
    const res = await axios.get('/api/enroll/list')
    renderData.value = res.data.data
  } catch (error) {
    console.log(error)
  }
}
fetchRenderData()

const changeStatus = async (id: number, status: string) => {
  try {
    const res = await axios.put('/api/enroll/review', {
      id: id,
      status: status,
    })
    if (res.data.msg === 'success') {
      message.success('修改成功')
      fetchRenderData()
    }
  } catch (error) {
    console.log(error)
  }
}
</script>

<style lang="sass" scoped></style>

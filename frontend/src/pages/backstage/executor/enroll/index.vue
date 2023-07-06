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
const statusMapping = {
  passed:'通过',
  rejected:'拒绝',
} as any;
const columns: DataTableColumns = [
  {
    title: '报名ID',
    key: 'ID',
  },
  {
    title: '课程名',
    key: 'courseName',
  },
  {
    title: '审核状态',
    key: 'status',
    render(row: any) {
      const chineseStatus = statusMapping[row.status]|| row.status; // 获取对应的中文角色，如果映射表中不存在，则保持原值
      return h(
        'span', // 使用一个span标签来包裹显示的文本
        chineseStatus // 中文角色
      );
    },
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
              type: 'info',
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
    const res = await axios.get('/backend/enroll/list')
    renderData.value = res.data.data
  } catch (error) {
    console.log(error)
  }
}
fetchRenderData()

const changeStatus = async (id: number, status: string) => {
  try {
    const res = await axios.put('/backend/enroll/review', {
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

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
const selectOptions = () => [
  { label: '用户', key: 'user' },
  { label: '执行人', key: 'executor' },
  { label: '经理', key: 'manager' },
]
const roleMapping = {
  user: '用户',
  executor: '执行人',
  manager: '经理',
} as any;
const columns: DataTableColumns = [
  {
    title: '用户ID',
    key: 'ID',
  },
  {
    title: '用户名',
    key: 'username',
  },
  {
    title: '用户权限',
    key: 'role',
    render(row: any) {
      const chineseRole = roleMapping[row.role]|| row.role; // 获取对应的中文角色，如果映射表中不存在，则保持原值
      return h(
        'span', // 使用一个span标签来包裹显示的文本
        chineseRole // 中文角色
      );
    },
  },
  {
    title: '创建时间',
    key: 'CreatedAt',
  },
  {
    title: '更新时间',
    key: 'UpdatedAt',
  },
  {
    title: '操作',
    key: 'actions',
    render(row: any) {
      return h(
        NDropdown,
        {
          trigger: 'click',
          options: selectOptions(),
          onSelect: (key: string) => changeRole(row.username, key),
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
            { default: () => '更改用户权限' }
          ),
        ]
      )
    },
  },
]
const renderData = ref([])
const fetchRenderData = async () => {
  try {
    const res = await axios.get('/backend/manager/getuser')
    renderData.value = res.data.data
  } catch (error) {
    console.log(error)
  }
}
fetchRenderData()

const changeRole = async (username: string, role: string) => {
  try {
    const res = await axios.post('/backend/manage/changerole', {
      username: username,
      role: role,
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

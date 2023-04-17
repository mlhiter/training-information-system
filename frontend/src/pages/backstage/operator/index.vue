<template>
  <div>
    <n-card v-if="showCard">
      <n-form ref="formRef" :label-width="80" :model="formValue">
        <n-form-item label="课程姓名" path="courseName">
          <n-input
            v-model:value="formValue.courseName"
            placeholder="输入课程姓名" />
        </n-form-item>
      </n-form>
      <n-button type="primary" @click="fetchRenderData">确认</n-button>
    </n-card>
    <n-card
      title="学员名单"
      :bordered="false"
      header-style="font-size:25px"
      v-else>
      <n-data-table
        :columns="columns"
        :data="renderData"
        :pagination="pagination"
        :bordered="false" />
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import { DataTableColumns, NTag } from 'naive-ui'
const pagination = {
  pageSize: 10,
}
const showCard = ref(true)
const columns: DataTableColumns = [
  {
    title: '名单ID',
    key: 'ID',
  },
  {
    title: '课程ID',
    key: 'courseId',
  },
  {
    title: '课程名称',
    key: 'courseName',
  },
  {
    title: '学生ID',
    key: 'studentId',
  },
  {
    title: '学生名称',
    key: 'studentName',
  },
  {
    title: '缴费状态',
    key: 'isPaid',
    render(row) {
      return h('div', [
        row.isPaid === '已缴费'
          ? h(
              NTag,
              { type: 'success', bordered: false },
              { default: () => '已缴费' }
            )
          : h(
              NTag,
              { type: 'error', bordered: false },
              { default: () => '未缴费' }
            ),
      ])
    },
  },
  {
    title: '签到状态',
    key: 'isSigned',
    render(row) {
      return h('div', [
        row.isSigned === '已签到'
          ? h(
              NTag,
              { type: 'success', bordered: false },
              { default: () => '已签到' }
            )
          : h(
              NTag,
              { type: 'error', bordered: false },
              { default: () => '未签到' }
            ),
      ])
    },
  },
]
const formValue = ref({
  courseName: '',
})
const renderData = ref([])
const fetchRenderData = async () => {
  try {
    showCard.value = false
    const res = await axios.get('/backend/scene/list', {
      params: formValue.value,
    })
    renderData.value = res.data.data.map((item: any) => {
      const isPaid = Math.random() >= 0.5 ? '已缴费' : '未缴费'
      const isSigned = Math.random() >= 0.5 ? '已签到' : '未签到'
      return {
        ...item,
        isPaid,
        isSigned,
      }
    })
  } catch (error) {
    console.log(error)
  }
}
</script>

<style lang="sass" scoped></style>

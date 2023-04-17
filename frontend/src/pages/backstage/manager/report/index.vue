<template>
  <div>
    <n-card title="汇总报表" style="margin-bottom: 16px">
      <n-tabs type="line" animated>
        <n-tab-pane name="income" tab="课程收入">
          <n-data-table
            :columns="incomeColumns"
            :data="incomeData"
            :pagination="pagination"
            :bordered="false" />
        </n-tab-pane>
        <n-tab-pane name="course" tab="课程">
          <n-data-table
            :columns="courseColumns"
            :data="courseData"
            :pagination="pagination"
            :bordered="false" />
        </n-tab-pane>
        <n-tab-pane name="student" tab="学员">
          <n-data-table
            :columns="studentColumns"
            :data="studentData"
            :pagination="pagination"
            :bordered="false" />
        </n-tab-pane>
        <n-tab-pane name="lecturer" tab="讲师">
          <n-data-table
            :columns="lecturerColumns"
            :data="lecturerData"
            :pagination="pagination"
            :bordered="false" />
        </n-tab-pane>
        <n-tab-pane name="questionnaire" tab="问卷">
          <n-data-table
            :columns="questionnaireColumns"
            :data="questionnaireData"
            :pagination="pagination"
            :bordered="false" />
        </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import { DataTableColumns } from 'naive-ui'
const incomeData = ref([])
const courseData = ref([])
const studentData = ref([])
const lecturerData = ref([])
const questionnaireData = ref([])

const pagination = {
  pageSize: 10,
}
const incomeColumns: DataTableColumns = [
  {
    title: '收入ID',
    key: 'ID',
  },
  {
    title: '课程名',
    key: 'name',
  },
  {
    title: '选课人数',
    key: 'selectedNumber',
  },
  {
    title: '价格',
    key: 'price',
  },
  {
    title: '收入',
    key: 'income',
  },
]
const courseColumns: DataTableColumns = [
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
]
const studentColumns: DataTableColumns = [
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
]
const lecturerColumns: DataTableColumns = [
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
]
const questionnaireColumns: DataTableColumns = [
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
]
const fetchIncomeData = async () => {
  try {
    const res = await axios.get('/backend/income/list')
    incomeData.value = res.data.data
  } catch (error) {
    console.log(error)
  }
}
const fetchCourseData = async () => {
  try {
    const res = await axios.get('/backend/training/list')
    courseData.value = res.data.data
  } catch (error) {
    console.log(error)
  }
}
const fetchStudentData = async () => {
  try {
    const res = await axios.get('/backend/student/list')
    studentData.value = res.data.data
  } catch (error) {
    console.log(error)
  }
}
const fetchLecturerData = async () => {
  try {
    const res = await axios.get('/backend/lecturer/list')
    lecturerData.value = res.data.data
  } catch (error) {
    console.log(error)
  }
}
const fetchQuestionnaireData = async () => {
  try {
    const res = await axios.get('/backend/questionnaire/list')
    questionnaireData.value = res.data.data
  } catch (error) {
    console.log(error)
  }
}
fetchIncomeData()
fetchCourseData()
fetchStudentData()
fetchLecturerData()
fetchQuestionnaireData()
</script>

<style lang="sass" scoped></style>

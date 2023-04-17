<template>
  <n-card title="缴费账单" v-if="change">
    <template #header-extra>{{ check.paymentTime }}</template>
    账单id:{{ check.ID }}
    <br />
    学生id:{{ check.studentId }}
    <br />
    学生姓名:{{ check.studentName }}
    <br />
    课程id:{{ check.courseId }}
    <br />
    课程名称:{{ check.courseName }}
    <br />
    课程费用:{{ check.paymentAmount }}
    <br />
    收款人:{{ check.payee }}
    <br />
    收款账户:{{ check.payAccount }}
    <br />
    <template #action>
      <n-button attr-type="button" type="success" @click="handlePayment">
        缴费
      </n-button>
    </template>
  </n-card>
  <n-alert title="Success" type="success" v-else>缴费成功！</n-alert>
</template>

<script lang="ts" setup>
const change = ref(true)
const check = ref({
  ID: 1,
  studentId: 1,
  studentName: '李晨洋',
  courseId: 8,
  courseName: 'C语言',
  paymentAmount: 10000,
  payee: '李致知',
  payAccount: '6222031614003722522',
  paymentTime: '2023-04-14 21:06:23',
})
const studentName = '李致知'
const fetchCheckList = async () => {
  try {
    const res = await axios.get('/backend/pay/individual', {
      params: { studentName: studentName },
    })
    let tempCheck = res.data.data[res.data.data.length - 1]
    check.value = { ...check.value, ...tempCheck }
    console.log(check.value)
  } catch (error) {
    console.log(error)
  }
}
function handlePayment() {
  change.value = false
}
onActivated(async () => {
  fetchCheckList()
})
onDeactivated(() => {
  change.value = true
})
</script>

<style lang="sass" scoped></style>

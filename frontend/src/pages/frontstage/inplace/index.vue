<template>
  <div>
    <n-card title="签到" v-if="change">
      <!-- <template #header-extra>#header-extra</template> -->
      签到人：李晨洋
      <br />
      签到课程：C语言
      <template #action>
        <n-button attr-type="button" type="success" @click="handleInplace">
          签到
        </n-button>
      </template>
    </n-card>
    <n-alert title="Success" type="success" v-else>签到成功！</n-alert>
  </div>
</template>

<script lang="ts" setup>
const change = ref(true)

function handleInplace() {
  change.value = false
}
onDeactivated(() => {
  change.value = true
})
</script>

<style lang="sass" scoped></style>

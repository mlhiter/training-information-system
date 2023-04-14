<template>
  <div>后台模板</div>
  <router-view v-slot="{ Component }">
    <transition :duration="200" name="fade-top" mode="out-in">
      <keep-alive>
        <component :is="Component" />
      </keep-alive>
    </transition>
  </router-view>
</template>

<script lang="ts" setup>
import {} from 'vue'
</script>

<style lang="sass" scoped></style>

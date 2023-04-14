<template>
  <n-card class="w-120">
    <n-space class="mb-10" align="center" justify="center">
      <n-image width="50" src="/favicon.svg"></n-image>
      <div class="text-3xl font-serif text-center">浩奇培训信息管理系统</div>
    </n-space>
    <n-tabs
      class="pl-4px"
      default-value="signin"
      size="large"
      animated
      style="margin: 0 -4px"
      pane-style="padding-left: 4px; padding-right: 4px; box-sizing: border-box;">
      <n-tab-pane name="signin" tab="登录">
        <n-form>
          <n-form-item-row label="用户名">
            <n-input placeholder="请输入用户名" />
          </n-form-item-row>
          <n-form-item-row label="密码">
            <n-input placeholder="请输入密码" />
          </n-form-item-row>
        </n-form>
        <n-button type="primary" block secondary strong>登录</n-button>
      </n-tab-pane>
      <n-tab-pane name="signup" tab="注册">
        <n-form>
          <n-form-item-row label="用户名">
            <n-input placeholder="请输入用户名" />
          </n-form-item-row>
          <n-form-item-row label="密码">
            <n-input placeholder="请输入密码" />
          </n-form-item-row>
          <n-form-item-row label="重复密码">
            <n-input placeholder="请再次输入密码" />
          </n-form-item-row>
        </n-form>
        <n-button type="primary" block secondary strong>注册</n-button>
      </n-tab-pane>
    </n-tabs>
  </n-card>
</template>

<script lang="ts" setup>
import {} from 'vue'
</script>

<style lang="sass" scoped></style>

<script setup>
const { type,title } = defineProps(['type','title'])
</script>
<template>
  <div class="card-cont">
    <div class="card-text">
      {{ title }}
    </div>
    <div class="card-out">
      <slot></slot>
    </div>
  </div>
</template>
<style>
.card-cont{
  background: var(--messageBox-bg);
  border-radius: 8px;
  padding: 16px;
  backdrop-filter: blur(10px);
  transition: 0.2s;
  margin-bottom: 12px;
  display: flex;
}
.card-text{
  user-select: none;
  flex: 1;
  height: 30px;
  line-height: 30px;
}
.card-left{
  height: 30px;
  line-height: 30px;
}
</style>

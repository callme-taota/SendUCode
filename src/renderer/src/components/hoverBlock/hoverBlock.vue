<script setup>
const props = defineProps(['type','icon','title'])
</script>
<template>
  <div class="hoverBlock" :class="{ 'hoverBlock-buttonType' : props.type == 'button' }">
    <slot></slot>
  </div>
</template>

<style>
.hoverBlock {
  background: var(--hoverBlock-button-bg);
  backdrop-filter: blur(12px);
  transition: 0.3s;
}
.hoverBlock-buttonType{
  width: 40px;
  height: 40px;
  border-radius: 12px;
}
.hoverBlock-buttonType:hover{
  background: var(--hoverBlock-button-bg-hover);
}
</style>

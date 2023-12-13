<script setup>
import { defineEmits } from 'vue'
const props = defineProps(['options'])
const emit = defineEmits(['change'])
const clickChild=(value)=>{
  emit('change',value)
}
</script>
<template>
  <div class="option-cont">
    <input type="text" />
  </div>
</template>
<style>
.option-cont{
  overflow: hidden;
  height: 30px;
  border-radius: 8px;
  display: inline-flex;
  flex-direction: row;
}
</style>

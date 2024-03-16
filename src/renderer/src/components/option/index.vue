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
    <div class="options" v-for="o in props.options" :class="{'options-active' : o.active}" @click="clickChild(o.value)">
      {{ o.text }}
    </div>
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
.options{
  height: 30px;
  margin: 0 1px;
  line-height: 30px;
  font-size:small;
  background-color: rgba(150, 150, 150, 0.1);
  padding: 0 8px;
  user-select: none;
  transition: 0.3s;
}
.options-active{
  background-color: rgba(150, 150, 150, 0.2);
}
</style>

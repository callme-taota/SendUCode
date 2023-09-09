<script setup>
import { ref } from 'vue'
import hoverBlock from '../components/hoverBlock/hoverBlock.vue'
const buttonList = ref([
  { active: true, index: 1, icon: '1' },
  { active: false,index: 2,  icon: '2' },
  { active: false,index: 3,  icon: '3' },
  { active: false,index: 4,  icon: '4' }
])
const changeActive = (index) => {
  let list = buttonList.value
  for(let i = 0 ; i< list.length ; i++){
    if(list[i].index == index){
      list[i].active = true
    }else{
      list[i].active = false
    }
  }
  buttonList.value = list
}
</script>
<template>
  <div class="bar">
    <div class="mediaPlayer"></div>
    <div class="buttonGroups">
      <hoverBlock type="button" v-for="btn in buttonList" :key="btn.icon" :active="btn.active" @click="changeActive(btn.index)">{{ btn.icon }}</hoverBlock>
      <hoverBlock type="user"></hoverBlock>
    </div>
  </div>
</template>
<style>
.bar {
  background-color: rgba(0, 0, 0, 0.1);
  height: 70px;
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
  display: flex;
  padding: 0 30px;
}
.mediaPlayer {
  flex: 1;
}
.buttonGroups {
  padding-top: 15px;
  display: flex;
  justify-content: space-around;
}
</style>

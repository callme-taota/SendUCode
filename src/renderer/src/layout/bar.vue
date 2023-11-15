<script setup>
import { ref, markRaw, onMounted , computed } from 'vue'
import { useAppStore } from '../store/modules/app'

const AppStore = useAppStore()

import router from '../router'
import hoverBlock from '../components/hoverBlock/hoverBlock.vue'
import icon from '../components/icon/icon.vue';

import iHome from '../components/icon/icons/iHome.vue';
import iHomeFill from '../components/icon/icons/iHomeFill.vue';
import iSearch from '../components/icon/icons/iSearch.vue'
import iTool from '../components/icon/icons/iTool.vue'
import iToolFill from '../components/icon/icons/iToolFill.vue'
import iSetting from '../components/icon/icons/iSetting.vue';
import iSettingFill from '../components/icon/icons/iSettingFill.vue';

const matchMedia = window.matchMedia('(prefers-color-scheme: dark)')
matchMedia.addEventListener("change", function() {
  if (this.matches) {
    //dark
    isDark.value = true
  } else {
    //light
    isDark.value = false
  }
})

onMounted(() => {
  router.beforeEach((to, from) => {
    let toPath = to.path
    let list = buttonList.value
    for (let i = 0; i < list.length; i++) {
      if (list[i].components == toPath) {
        list[i].active = true
      } else {
        list[i].active = false
      }
    }
    buttonList.value = list
    if (toPath == '/user') {
      userBtnActive.value = true
    } else {
      userBtnActive.value = false
    }
  })
})
const isDark = ref(AppStore.isDarkTheme)
const userBtnActive = ref(false)
const buttonList = ref([
  { active: true, index: 1, icon: markRaw(iHome), activeIcon: markRaw(iHomeFill), components: "/" },
  { active: false, index: 2, icon: markRaw(iSearch), activeIcon: markRaw(iSearch), components: "/find" },
  { active: false, index: 3, icon: markRaw(iTool), activeIcon: markRaw(iToolFill), components: "/tool" },
  { active: false, index: 4, icon: markRaw(iSetting), activeIcon: markRaw(iSettingFill), components: "/setting" }
])
const userObj = { index: 5, components: "/user" }
const changeActive = (btn) => {
  router.push(btn.components)
}
const iconType = (isActive)=>{
  if(isDark.value==true){
    if(isActive){
      return "#ffffff"
    }else{
      return "#CCCCCC"
    }
  }else{
    if(isActive){
      return "#222222"
    }else{
      return "#000000"
    }
  }
}
</script>
<template>
  <div class="bar">
    <div class="mediaPlayer"></div>
    <div class="buttonGroups">
      <hoverBlock type="button" v-for="btn in buttonList" :key="btn.icon" :active="btn.active" @click="changeActive(btn)">
        <icon :icon="btn.active ? btn.activeIcon : btn.icon" height="20" width="20"
          :color="iconType(btn.active)"></icon>
      </hoverBlock>
      <hoverBlock type="user" @click="changeActive(userObj)" :active="userBtnActive"></hoverBlock>
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

<script setup>
import { ref, markRaw, onMounted , computed } from 'vue'
import { storeToRefs } from 'pinia';
import { useAppStore } from '../store/modules/app'

const appStore = useAppStore()
const { btnColor } = storeToRefs(appStore)

import router from '../router'
import hoverBlock from '../components/hoverBlock/hoverBlock.vue'
import icon from '../components/icon/icon.vue';

import iHome from '../components/icon/icons/iHome.vue';
import iHomeFill from '../components/icon/icons/iHomeFill.vue';
import iSetting from '../components/icon/icons/iSetting.vue';
import iSettingFill from '../components/icon/icons/iSettingFill.vue';

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
  })
})
const buttonList = ref([
  { active: true, index: 1, icon: markRaw(iHome), activeIcon: markRaw(iHomeFill), components: "/" },
  { active: false, index: 4, icon: markRaw(iSetting), activeIcon: markRaw(iSettingFill), components: "/setting" }
])
const changeActive = (btn) => {
  router.push(btn.components)
}
</script>
<template>
  <div class="bar">
    <div class="mediaPlayer"></div>
    <div class="buttonGroups">
      <hoverBlock type="button" v-for="btn in buttonList" :key="btn.icon" :active="btn.active" @click="changeActive(btn)">
        <icon :icon="btn.active ? btn.activeIcon : btn.icon" height="20" width="20"
          :color="btnColor"></icon>
      </hoverBlock>
    </div>
  </div>
</template>
<style>
.bar {
  position: fixed;
  right: 30px;
  bottom: 30px;
  display: flex;
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

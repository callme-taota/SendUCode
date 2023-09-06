<script setup>
import {useAppStore} from "./store/modules/app.js"
import hoverBlock from "./components/hoverBlock/hoverBlock.vue"
const AppStore = useAppStore()
console.log(AppStore.platform)
</script>

<template>
  <div class="background">
    <div class="container">
      <div class="main">
        <hoverBlock type="button">
        </hoverBlock>
      </div>
      <div class="bar"></div>
    </div>
  </div>
</template>

<style>
.background {
  width: 100%;
  height: 100%;
  /* background-color: rgba(255, 255, 255, 0.55);
  background-color: rgba(0, 0, 0, 0.65); */
}
.container{
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
}
.main{
  flex: 1;
}
.bar{
  background-color:  rgba(0, 0, 0, 0.1);
  height: 70px;
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
}
</style>

<script setup>
import bar from "./bar.vue"
import bb from "../components/bb/index.vue"
</script>

<template>
  <div class="background">
    <div class="container">
      <div class="main">
        <router-view v-slot="{ Component }">
          <transition name="fade">
            <component :is="Component" />
          </transition>
        </router-view>
      </div>
      <bar></bar>
    </div>
  </div>
  <bb></bb>
</template>

<style>
.background {
  width: 100%;
  height: 100%;
  background-color: var(--background);
}

.container {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
}

.main {
  overflow-y: scroll;
  padding-left: 8px;
  flex: 1;
}

/* WebKit 浏览器样式 */
/* 滚动条整体样式 */
.main::-webkit-scrollbar {
  width: 8px; /* 设置滚动条宽度 */
}

/* 滚动条轨道样式 */
.main::-webkit-scrollbar-track {
  background-color: rgba(0, 0, 0, 0); /* 设置滚动条轨道的背景色 */
}

/* 滚动条滑块样式 */
.main::-webkit-scrollbar-thumb {
  background-color: var(--scroll-track); /* 设置滚动条滑块的背景色 */
  border-radius: 4px; /* 设置滚动条滑块的圆角 */
}

/* 当鼠标悬停在滚动条上时的样式 */
.main::-webkit-scrollbar-thumb:hover {
  background-color: #555;
}

</style>

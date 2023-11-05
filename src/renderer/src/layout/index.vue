<script setup>
import bar from "./bar.vue"
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
  flex: 1;
}
</style>

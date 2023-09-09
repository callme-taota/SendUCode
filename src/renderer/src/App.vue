<script setup>
import {useAppStore} from "./store/modules/app.js"
import Layout from "./layout/index.vue"
const AppStore = useAppStore()
AppStore.SetThemeAuto()
</script>

<template>
  <Layout></Layout>
</template>

<style>

</style>
<style src="./assets/css/light.css">
</style>
<style src="./assets/css/dark.css">
</style>
<style src="./assets/css/index.css">
</style>

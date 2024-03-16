<template>
  <component :width="props.width" :height="props.height" :is="props.icon" :color="color"></component>
</template>

<script setup>
const props = defineProps(['icon','width','height','color'])
</script>
<style>

</style>

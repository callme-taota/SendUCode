<script setup>
const props = defineProps(['type', 'active'])
</script>
<template>
  <div class="hoverBlock" :class="{
    'hoverBlock-buttonType': props.type == 'button',
    'hoverBlock-userType': props.type == 'user'
  }">

    <!-- Button -->
    <div v-if="props.type == 'button'" class="hoverBlock-buttonType-active"
      :class="{ 'hoverBlock-buttonType-active-on': props.active == true && props.type == 'button' }"></div>
    <slot v-if="props.type !== 'user'"></slot>

    <!-- User -->
    <div v-if="props.type == 'user'" class="hoverBlock-userType-active"
      :class="{ 'hoverBlock-userType-active-on': props.active == true && props.type == 'user' }"></div>
    <div v-if="props.type == 'user'" class="user-cont">
      <div class="user-img"></div>
      <div class="user-name">Tao_ta</div>
    </div>


  </div>
</template>

<style scope>
.hoverBlock {
  background: var(--hoverBlock-button-bg);
  backdrop-filter: blur(12px);
  transition: 0.2s;
  overflow: hidden;
  user-select: none;
  text-align: center;
}

.hoverBlock-buttonType {
  position: relative;
  width: 40px;
  height: 36px;
  padding-top: 4px;
  line-height: 40px;
  border-radius: 8px;
}

.hoverBlock-buttonType:hover {
  background: var(--hoverBlock-button-bg-hover);
}

.hoverBlock-buttonType:active {
  background: var(--hoverBlock-button-bg-active);
}

.hoverBlock-buttonType-active {
  display: block;
  width: 22px;
  height: 5px;
  position: absolute;
  background: var(--hoverBlock-button-bg-select);
  bottom: 0px;
  left: 9px;
  border-radius: 3px;
  transition: 0.2s;
  opacity: 0;
}

.hoverBlock-buttonType-active-on {
  opacity: 1;
}

.hoverBlock-userType {
  width: 100px;
  height: 40px;
  line-height: 40px;
  border-radius: 8px;
}

.hoverBlock-userType:hover {
  background: var(--hoverBlock-button-bg-hover);
}

.hoverBlock-userType:active {
  background: var(--hoverBlock-button-bg-active);
}

.hoverBlock-userType-active {
  display: block;
  width: 6px;
  height: 22px;
  position: absolute;
  background: var(--hoverBlock-button-bg-select);
  right: 4px;
  border-radius: 3px;
  transition: 0.2s;
  top: 9px;
  opacity: 0;
}

.hoverBlock-userType-active-on {
  opacity: 1;
}

.user-cont {
  display: flex;
  justify-content: center;
}

.user-img {
  margin-top: 5px;
  width: 30px;
  height: 30px;
  background-color: azure;
  border-radius: 15px;
}

.user-name {
  width: 50px;
  overflow: hidden;
  font-size: small;
  margin-left: 4px;
}
</style>

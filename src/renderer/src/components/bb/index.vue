<script setup>
import { useUserStore } from '../../store/modules/user';
import { storeToRefs } from 'pinia';

const userStore = useUserStore()
const { noticeList } = storeToRefs(userStore)

</script>
<template>
  <div class="bb-cont">
    <div class="bb-text" v-for="i in noticeList">
      {{ i.text }}
    </div>
  </div>
</template>
<style>
.bb-cont{
  top: 0;
  left: 50%;
  transform: translate(-50%);
  position: absolute;
}
.bb-text{
  margin-top: 8px;
  margin-bottom: 8px;
  padding: 8px;
  border-radius: 6px;
  text-align: center;
  width: 80px;
  background: rgba(222, 222, 222,0.5);
}
</style>

<script setup>
import { ref,onMounted } from 'vue';
import { storeToRefs } from 'pinia';
import { useUserStore  } from '../../store/modules/user';
import Message from '../../components/message/index.vue'

const userStore = useUserStore()
const { msgList } = storeToRefs(userStore)

onMounted(async()=>{
  let flag = await userStore.checkSession()
  if(flag == false) return
  userStore.getMsgs()
  userStore.listiner()
})

</script>
<template>
  <div class="home-cont">
    <div class="msg-cont">
      <Message v-for="i in msgList" :device="i.device" :detail="i.detail" :time="i.time"></Message>
    </div>
  </div>
</template>
<style>
.home-cont{
  width: 100%;
  height: 100%;
  position: relative;
}
.msg-cont{
  padding: 16px;
}
</style>

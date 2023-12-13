<script setup>
import { ref } from "vue";
import { useAppStore } from "../../store/modules/app";
import { useUserStore } from "../../store/modules/user";
import { storeToRefs } from "pinia";
import card from "../../components/card/index.vue";
import options from "../../components/option/index.vue";
import vueQr from 'vue-qr/src/packages/vue-qr.vue'
import icon from "../../components/icon/icon.vue";
import iClose from "../../components/icon/icons/iClose.vue";

const userStore = useUserStore()
const appStore = useAppStore()
const { session,userID } = storeToRefs(userStore)
const { qrColor_fill,qrColor_blank,btnColor } = storeToRefs(appStore)

const list = ref([
  { text: "浅色模式", value: 1, active: false },
  { text: "深色模式", value: 2, active: false },
  { text: "跟随系统", value: 3, active: true },
])
const qrCode = ref([
  { text: "二维码", value: 1, active: false }
])
const qrShow = ref(false)

const changeTheme = (value) => {
  appStore.SetTheme(value)
  let l = list.value;
  for (let i = 0; i < l.length; i++) {
    if (l[i].value == value) {
      l[i].active = true;
      continue
    }
    l[i].active = false;
  }
  list.value = l;
}
const userClick = () => {
  qrShow.value = true;
}
const closeQr = () => {
  qrShow.value = false;
}
</script>
<template>
  <div class="set-cont">
    <div class="set-inner-cont">
      <p class="set-title">设置</p>
      <card title="主题">
        <options @change="changeTheme" :options="list"></options>
      </card>
      <card title="用户名">
        <div class="card-left">{{ userID }}</div>
      </card>
      <card title="连接">
        <options @change="userClick" :options="qrCode"></options>
      </card>
      <div class="qr-cont" v-show="qrShow">
        <vue-qr :text="session" :size="200" :margin="0" :colorLight="qrColor_blank" :colorDark="qrColor_fill"/>
        <div class="qr-close" @click="closeQr">
          <icon width="20" height="20" :icon="iClose" :color="btnColor"></icon>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.set-cont {
  width: 100%;
  height: 100%;
  position: relative;
}

.set-inner-cont {
  padding: 16px;
}

.set-title {
  user-select: none;
  font-size: large;
  font-weight: bold;
  padding-left: 10px;
  margin: 12px 0;
}

.qr-cont {
  background-color: rgba(0, 0, 0, 0.07);
  border-radius: 12px;
  position: absolute;
  left: 50%;
  top: 50%;
  width: 200px;
  height: 200px;
  padding: 20px;
  transition: 0.3s;
  transform: translate(-50%,-50%);
}
.qr-close{
  position: relative;
  width: 20px;
  height: 20px;
  padding: 10px;
  top: 30px;
  left: 50%;
  transform: translateX(-50%);
  border-radius: 20px;
  background-color: rgba(0, 0, 0, 0.07);
}
</style>

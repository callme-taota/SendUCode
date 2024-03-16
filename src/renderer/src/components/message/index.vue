<script setup>
import { storeToRefs } from "pinia";
import icon from "../icon/icon.vue"
import iCopy from "../icon/icons/iCopy.vue"
import { useAppStore } from '../../store/modules/app'
import { useUserStore } from "../../store/modules/user";
import { CopyToClipboard } from "../../utils/index.js"

const props = defineProps(['detail', 'device', 'time'])

const userStore = useUserStore()
const appStore = useAppStore()
const { btnColor } = storeToRefs(appStore)

const copyMsg = () => {
  let msg = props.detail
  CopyToClipboard(msg)
  userStore.addNotice("复制成功")
}
</script>
<template>
  <div class="msg-box">
    <div class="msg-detail">
      {{ props.detail }}
    </div>
    <div class="msg-info">
      <div class="msg-device">
        {{ props.device }}
      </div>
      <div class="msg-time">
        {{ props.time }}
      </div>
      <div class="msg-btn-cont" @click="copyMsg">
        <icon width="14" height="14" :color="btnColor" :icon="iCopy"></icon>
      </div>
    </div>
  </div>
</template>
<style>
.msg-box {
  background: var(--messageBox-bg);
  width: calc(100% - 40px);
  height: auto;
  border-radius: 12px;
  padding: 20px;
  backdrop-filter: blur(10px);
  transition: 0.2s;
  margin-bottom: 12px;
}

.msg-box:hover {
  background: var(--messageBox-bg-hover);
}

.msg-info {
  font-size: small;
  user-select: none;
  margin-top: 10px;
  display: flex;
  flex-direction: row;
}

.msg-device {
  flex-grow: 1;
  line-height: 28px;
  flex: 1;
}

.msg-time {
  line-height: 28px;
}

.msg-btn-cont {
  border-radius: 4px;
  width: 28px;
  height: 28px;
  display: flex;
  justify-content: center;
  align-items: center;
  /* margin-left: -6px; */
}

.msg-btn-cont:hover {
  background-color: var(--messageBox-bg-hover);
}

.msg-btn-cont:active {
  background-color: var(--messageBox-bg-active);
}</style>

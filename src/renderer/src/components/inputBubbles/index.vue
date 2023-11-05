<script setup>

</script>
<template>
  <div class="input-bubble-cont">
    <textarea class="input-bubble-textarea" />
  </div>
</template>
<style>
.input-bubble-cont {
  position: absolute;
  left: 10%;
  bottom: 30px;
  width: 80%;
  height: 60px;
  background-color: gray;
  border-radius: 20px;
  overflow: hidden;
}

.input-bubble-textarea {
  line-height: 60px;
  position: relative;
  width: calc(100% - 20px);
  padding: 10px;
  font-size: 16px;
  border: 0;
  resize: none;
  overflow: hidden;
  box-shadow: none;
}
.input-bubble-textarea:focus{
  outline: none;
}
</style>

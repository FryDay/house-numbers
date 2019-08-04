<template>
  <div class="colorBox" :style="{'background-color':hex}">
    <h1 :style="{'color':oppHex}">{{start}} - {{end}}</h1>
  </div>
</template>

<script>
export default {
  name: 'colorBox',
  props: {
    hex: {
      type: String,
      required: true
    },
    oppHex: {
      type: String,
      required: true
    },
    start: {
      type: String,
      required: true
    },
    end: {
      type: String,
      required: true
    }
  }
}
</script>

<style lang="less" scoped>
h1 {
  margin: auto;
}
.colorBox {
  align-items: center;
  display: flex;
  flex-grow: 1;
  height: 100px;
  justify-items: center;
  text-align: center;
  width: 100%;
}
</style>

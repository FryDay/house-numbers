<template>
  <div class="colorBox" :style="{'background-color':color}"></div>
</template>

<script>
export default {
  name: 'colorBox',
  props: {
    color: {
      type: String,
      required: true
    }
  }
}
</script>

<style lang="less" scoped>
.colorBox {
  display: inline-block;
  height: 100px;
  outline: #000000 solid 2px;
  margin: 60px;
  width: 100px;
}
</style>

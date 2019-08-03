<template>
  <div class="colorBox" :style="{'background-color':color}">
    <h2 :style="{'color':oppColor}">{{start}} - {{end}}</h2>
  </div>
</template>

<script>
export default {
  name: 'colorBox',
  props: {
    color: {
      type: String,
      required: true
    },
    oppColor: {
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
h2 {
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

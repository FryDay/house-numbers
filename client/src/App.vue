<template>
  <div id="app">
    <ColorBox
      :color="currentColor"
      :oppColor="oppositeColor"
      :start="currentStartTime"
      :end="currentEndTime"
    />
    <color-picker v-bind="hsl" @change="changeColor" @input="updateHSL"></color-picker>
  </div>
</template>

<script>
import ColorBox from './components/ColorBox.vue'
import ColorPicker from '@radial-color-picker/vue-color-picker'
import { RepositoryFactory } from './repositoryFactory'

const ColorRepository = RepositoryFactory.get('color')
const TimeRepository = RepositoryFactory.get('time')

export default {
  name: 'app',
  components: {
    ColorBox,
    ColorPicker
  },
  created() {
    this.fetchColor()
    this.fetchTime()
  },
  data() {
    return {
      currentColor: '',
      oppositeColor: '',
      currentStartTime: '',
      currentEndTime: '',
      hsl: {
        hue: 0,
        variant: 'persistent'
      }
    }
  },
  methods: {
    async fetchColor() {
      const { data } = await ColorRepository.get()

      this.currentColor = data.hex
      this.oppositeColor = data.xeh
    },
    async fetchTime() {
      const { data } = await TimeRepository.get()

      this.currentStartTime = data.start
      this.currentEndTime = data.end
    },
    updateHSL(hue) {
      this.hsl.hue = hue
    },
    changeColor() {
      this.updateValue(this.hsl.hue)
    },
    async updateValue(value) {
      console.log(value.toFixed(0))
      const { data } = await ColorRepository.post({ hue: value.toFixed(0) })

      this.currentColor = data.hex
      this.oppositeColor = data.xeh
    }
  }
}
</script>

<style lang="less">
@import '~@radial-color-picker/vue-color-picker/dist/vue-color-picker.min.css';

html {
  height: 100% !important;
  margin: 0 !important;
  width: 100% !important;
}
#body {
  height: 100% !important;
  margin: 0 !important;
  width: 100% !important;
}
#app {
  align-items: center;
  background-color: #a9a9a9;
  color: #2c3e50;
  display: flex;
  flex-flow: column;
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  height: 100% !important;
  justify-items: center;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  margin: 0 !important;
  text-align: center;
  width: 100% !important;
}
</style>

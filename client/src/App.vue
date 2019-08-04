<template>
  <div id="app" :style="{'background-color':currentHex}">
    <ColorBox :hex="currentHex" :oppHex="oppositeHex" :start="currentStart" :end="currentEnd" />
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
      currentHex: '',
      oppositeHex: '',
      currentStart: '',
      currentEnd: '',
      hsl: {
        hue: 0,
        variant: 'persistent'
      }
    }
  },
  methods: {
    async fetchColor() {
      const { data } = await ColorRepository.get()

      this.hsl.hue = data.hue
      this.currentHex = data.hex
      this.oppositeHex = data.xeh
    },
    async fetchTime() {
      const { data } = await TimeRepository.get()

      this.currentStart = data.start
      this.currentEnd = data.end
    },
    updateHSL(hue) {
      this.hsl.hue = hue
    },
    changeColor() {
      this.updateHue(this.hsl.hue)
    },
    async updateHue(hue) {
      const { data } = await ColorRepository.post({ hue: hue.toFixed(0) })

      this.currentHex = data.hex
      this.oppositeHex = data.xeh
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
.rcp {
  margin-bottom: 6px;
}
</style>

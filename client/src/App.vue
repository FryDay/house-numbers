<template>
  <div id="app">
    <ColorBox :color="currentColor" :start="currentStartTime" :end="currentEndTime" />
    <chrome-picker @input="updateValue" :value="currentColor" />
  </div>
</template>

<script>
import ColorBox from './components/ColorBox.vue'
import { Chrome } from 'vue-color'
import { RepositoryFactory } from './repositoryFactory'

const ColorRepository = RepositoryFactory.get('color')
const TimeRepository = RepositoryFactory.get('time')

export default {
  name: 'app',
  components: {
    ColorBox,
    'chrome-picker': Chrome
  },
  created() {
    this.fetchColor()
    this.fetchTime()
    // this.postColor('#00ff00')
  },
  data() {
    return {
      currentColor: '',
      currentStartTime: '',
      currentEndTime: ''
    }
  },
  methods: {
    async fetchColor() {
      const { data } = await ColorRepository.get()

      this.currentColor = data.hex
    },
    async fetchTime() {
      const { data } = await TimeRepository.get()

      this.currentStartTime = data.start
      this.currentEndTime = data.end
    },
    async postColor(hex) {
      await ColorRepository.post({ hex: hex })
    },
    updateValue(value) {
      this.currentColor = value.hex
      this.postColor(this.currentColor)
    }
  }
}
</script>

<style lang="less">
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
.vc-chrome {
  background-color: #a9a9a9 !important;
  box-shadow: none !important;
  width: 100% !important;
  .vc-chrome-body {
    padding: 0px !important;
  }
  .vc-alpha {
    display: none !important;
  }
  .vc-chrome-color-wrap {
    display: none !important;
  }
  .vc-chrome-saturation-wrap {
    display: none !important;
  }
  .vc-chrome-fields-wrap {
    display: none !important;
  }
  .vc-chrome-hue-wrap {
    height: 100% !important;
  }
  .vc-hue-picker {
    margin-top: 2px !important;
    height: 50px !important;
  }
  .vc-chrome-sliders {
    height: 50px !important;
  }
}
</style>

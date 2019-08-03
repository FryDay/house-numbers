<template>
  <div id="app">
    <ColorBox :color="currentColor" />
    <h2>{{currentStartTime}} - {{currentEndTime}}</h2>
  </div>
</template>

<script>
import ColorBox from './components/ColorBox.vue'
import { RepositoryFactory } from './repositoryFactory'

const ColorRepository = RepositoryFactory.get('color')
const TimeRepository = RepositoryFactory.get('time')

export default {
  name: 'app',
  components: {
    ColorBox
  },
  created() {
    this.fetchColor()
    this.fetchTime()
    this.postColor('#00ff00')
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
  background-color: #a9a9a9;
  color: #2c3e50;
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  height: 100% !important;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  margin: 0 !important;
  text-align: center;
  width: 100% !important;
}
</style>

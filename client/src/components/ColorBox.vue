<template>
  <div class="colorBox" :style="{'background-color':hex}">
    <h1 :style="{'color':oppHex}">{{start}} - {{end}}</h1>
  </div>
</template>

<script>
import { RepositoryFactory } from '../repositoryFactory'

const TimeRepository = RepositoryFactory.get('time')

export default {
  name: 'colorBox',
  components: {},
  created() {
    this.fetchTime()
  },
  data() {
    return {
      start: '',
      end: ''
    }
  },
  props: {
    hex: {
      type: String,
      required: true
    },
    oppHex: {
      type: String,
      required: true
    }
  },
  methods: {
    async fetchTime() {
      const { data } = await TimeRepository.get()

      this.start = data.start
      this.end = data.end
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

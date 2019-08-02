import Repository from './repository'

const resource = '/color'
export default {
  get() {
    return Repository.get(`${resource}`)
  },
  post(payload) {
    return Repository.post(`${resource}`, payload)
  }
}

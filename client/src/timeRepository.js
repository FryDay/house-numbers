import Repository from './repository'

const resource = '/time'
export default {
  get() {
    return Repository.get(`${resource}`)
  },
  post(payload) {
    return Repository.post(`${resource}`, payload)
  }
}

import Repository from './repository'

const resource = '/force'
export default {
  get() {
    return Repository.get(`${resource}`)
  },
  post(payload) {
    return Repository.post(`${resource}`, payload)
  }
}

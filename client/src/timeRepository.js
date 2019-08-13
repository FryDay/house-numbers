import Repository from './repository'

const resource = '/time'
export default {
  get() {
    return Repository.get(`${resource}`)
  }
}

import axios from 'axios'

export default axios.create({
  baseURL: 'http://localhost:8040',
  timeout: 1000
})

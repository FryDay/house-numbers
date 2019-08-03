import axios from 'axios'

export default axios.create({
  baseURL: 'http://localhost:8040',
  headers: {
    'content-type': 'application/x-www-form-urlencoded'
  },
  timeout: 1000
})

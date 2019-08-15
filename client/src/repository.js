import axios from 'axios'

export default axios.create({
  baseURL: process.env.VUE_APP_API_URL,
  headers: {
    'content-type': 'application/x-www-form-urlencoded'
  },
  timeout: 1000
})

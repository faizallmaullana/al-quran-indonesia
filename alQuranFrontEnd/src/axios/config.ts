import axios from 'axios'

const API = axios.create({
  baseURL: "/api/v1/quran/resources",
  timeout: 5000,
  withCredentials: true
})

export { API as axios }

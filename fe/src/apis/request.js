import axios from 'axios'

const $http = axios.create({
  timeout: 3 * 60 * 1000, // 3 minutes
  responseType: 'json',
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json',
  },
})

$http.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    if (token) {
      Object.defineProperty(config.headers, 'Authorization', {
        value: `Bearer ${token}`,
        configurable: true,
        enumerable: true,
        writable: true,
      })
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

$http.interceptors.response.use(
  (response) => response.data,
  (error) => {
    if (error.response.status === 401) {
      localStorage.setItem('token', '')
      location.href = '/'
    }

    return Promise.reject(error)
  }
)

export default $http

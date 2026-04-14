import axios, { type AxiosInstance, type InternalAxiosRequestConfig, type AxiosResponse } from 'axios'

const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'

// ─── Axios Instance ──────────────────────────────────────
const apiClient: AxiosInstance = axios.create({
  baseURL: BASE_URL,
  timeout: 15000,
  headers: {
    'Content-Type': 'application/json',
  },
})

// ─── Token Helpers ───────────────────────────────────────
export const getAccessToken  = (): string | null => localStorage.getItem('access_token')
export const getRefreshToken = (): string | null => localStorage.getItem('refresh_token')
export const setTokens = (access: string, refresh: string): void => {
  localStorage.setItem('access_token', access)
  localStorage.setItem('refresh_token', refresh)
}
export const clearTokens = (): void => {
  localStorage.removeItem('access_token')
  localStorage.removeItem('refresh_token')
}

// ─── Request Interceptor: Attach Access Token ────────────
apiClient.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const token = getAccessToken()
    if (token && config.headers) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => Promise.reject(error)
)

// ─── Response Interceptor: Auto Refresh on 401 ──────────
let isRefreshing = false
let failedQueue: Array<{
  resolve: (value: unknown) => void
  reject: (reason?: unknown) => void
}> = []

const processQueue = (error: unknown, token: string | null = null): void => {
  failedQueue.forEach(({ resolve, reject }) => {
    if (error) {
      reject(error)
    } else {
      resolve(token)
    }
  })
  failedQueue = []
}

apiClient.interceptors.response.use(
  (response: AxiosResponse) => response,
  async (error) => {
    const originalRequest = error.config

    // Only handle 401 (Unauthorized)
    if (error.response?.status !== 401 || originalRequest._retry) {
      return Promise.reject(error)
    }

    // If already refreshing, queue the request
    if (isRefreshing) {
      return new Promise((resolve, reject) => {
        failedQueue.push({ resolve, reject })
      }).then((token) => {
        if (originalRequest.headers) {
          originalRequest.headers.Authorization = `Bearer ${token}`
        }
        return apiClient(originalRequest)
      })
    }

    originalRequest._retry = true
    isRefreshing = true

    try {
      const refreshToken = getRefreshToken()
      if (!refreshToken) throw new Error('No refresh token')

      // Call refresh endpoint
      const { data } = await axios.post(`${BASE_URL}/auth/refresh`, {
        refreshToken,
      })

      const newAccess: string = data.accessToken
      const newRefresh: string = data.refreshToken || refreshToken

      setTokens(newAccess, newRefresh)
      processQueue(null, newAccess)

      // Retry original request
      if (originalRequest.headers) {
        originalRequest.headers.Authorization = `Bearer ${newAccess}`
      }
      return apiClient(originalRequest)
    } catch (refreshError) {
      processQueue(refreshError, null)
      clearTokens()
      // Redirect to login
      window.location.href = '/login'
      return Promise.reject(refreshError)
    } finally {
      isRefreshing = false
    }
  }
)

export default apiClient

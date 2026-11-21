import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { User, UserRole, LoginPayload, RegisterPayload } from '@/types'
import { login as apiLogin, register as apiRegister, logout as apiLogout } from '@/api/auth'
import { clearTokens, getAccessToken } from '@/api/client'

export const useAuthStore = defineStore('auth', () => {
  // ─── State ──────────────────────────────────────────────
  const user = ref<User | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  // ─── Getters ────────────────────────────────────────────
  const isAuthenticated = computed(() => !!user.value && !!getAccessToken())
  const userRole = computed<UserRole | null>(() => user.value?.role ?? null)
  const userName = computed(() => user.value?.name ?? '')
  const userAvatar = computed(() =>
    user.value?.avatar ||
    `https://ui-avatars.com/api/?name=${encodeURIComponent(userName.value)}&background=e0e7ff&color=4f46e5&bold=true`
  )

  // ─── RBAC Helpers ───────────────────────────────────────
  const isAdmin   = computed(() => userRole.value === 'admin')
  const isTeacher = computed(() => userRole.value === 'teacher')
  const isStudent = computed(() => userRole.value === 'student')
  const isParent  = computed(() => userRole.value === 'parent')

  /** Check if user has any of the given roles */
  const hasRole = (...roles: UserRole[]): boolean => {
    return !!userRole.value && roles.includes(userRole.value)
  }

  // ─── Actions ────────────────────────────────────────────
  async function login(payload: LoginPayload): Promise<void> {
    isLoading.value = true
    error.value = null
    try {
      const response = await apiLogin(payload)
      user.value = response.user
      localStorage.setItem('user', JSON.stringify(response.user))
    } catch (e: any) {
      error.value = e?.response?.data?.message || 'Login failed'
      throw e
    } finally {
      isLoading.value = false
    }
  }

  async function register(payload: RegisterPayload): Promise<void> {
    isLoading.value = true
    error.value = null
    try {
      const response = await apiRegister(payload)
      user.value = response.user
      localStorage.setItem('user', JSON.stringify(response.user))
    } catch (e: any) {
      error.value = e?.response?.data?.message || 'Registration failed'
      throw e
    } finally {
      isLoading.value = false
    }
  }

  async function logout(): Promise<void> {
    await apiLogout()
    user.value = null
    clearTokens()
    localStorage.removeItem('user')
  }

  /** Restore session from localStorage on app init */
  function restoreSession(): void {
    const stored = localStorage.getItem('user')
    const token = getAccessToken()
    if (stored && token) {
      user.value = JSON.parse(stored) as User
    }
  }

  return {
    // State
    user,
    isLoading,
    error,
    // Getters
    isAuthenticated,
    userRole,
    userName,
    userAvatar,
    isAdmin,
    isTeacher,
    isStudent,
    isParent,
    hasRole,
    // Actions
    login,
    register,
    logout,
    restoreSession,
  }
})

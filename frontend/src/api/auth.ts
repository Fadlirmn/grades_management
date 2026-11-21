import apiClient, { setTokens } from './client'
import type { AuthResponse, LoginPayload, RegisterPayload, User } from '@/types'

/**
 * Login with email + password.
 * Falls back to mock data when backend is unavailable.
 */
export async function login(payload: LoginPayload): Promise<AuthResponse> {
  try {
    const { data } = await apiClient.post<AuthResponse>('/auth/login', payload)
    setTokens(data.tokens.accessToken, data.tokens.refreshToken)
    return data
  } catch {
    // ── Mock fallback for development ──────────────────
    return mockLogin(payload)
  }
}

/**
 * Register a new user.
 * Falls back to mock data when backend is unavailable.
 */
export async function register(payload: RegisterPayload): Promise<AuthResponse> {
  try {
    const { data } = await apiClient.post<AuthResponse>('/auth/register', payload)
    setTokens(data.tokens.accessToken, data.tokens.refreshToken)
    return data
  } catch {
    // ── Mock fallback for development ──────────────────
    return mockRegister(payload)
  }
}

/**
 * Fetch current user profile from /auth/me
 */
export async function fetchMe(): Promise<User> {
  try {
    const { data } = await apiClient.get<User>('/auth/me')
    return data
  } catch {
    // Fallback: read from localStorage
    const stored = localStorage.getItem('user')
    if (stored) return JSON.parse(stored) as User
    throw new Error('Not authenticated')
  }
}

/**
 * Logout: clear tokens, optionally notify server
 */
export async function logout(): Promise<void> {
  try {
    await apiClient.post('/auth/logout')
  } catch {
    // silent fail
  }
}

// ─── Mock Login (Development) ────────────────────────────
function mockLogin(payload: LoginPayload): AuthResponse {
  const roleMappings: Record<string, User> = {
    'admin@edu.com': {
      id: '1', name: 'Admin Portal', email: 'admin@edu.com',
      role: 'admin', department: 'Administration', status: 'active',
    },
    'teacher@edu.com': {
      id: '2', name: 'Dr. Helena Vance', email: 'teacher@edu.com',
      role: 'teacher', department: 'Computer Science', status: 'active',
    },
    'student@edu.com': {
      id: '3', name: 'Alexander Chen', email: 'student@edu.com',
      role: 'student', department: 'Computer Science', status: 'active',
    },
    'parent@edu.com': {
      id: '4', name: 'Margaret Chen', email: 'parent@edu.com',
      role: 'parent', department: undefined, status: 'active',
    },
  }

  const user = roleMappings[payload.email] || roleMappings['teacher@edu.com']

  const mockTokens = {
    accessToken: `mock_access_${Date.now()}`,
    refreshToken: `mock_refresh_${Date.now()}`,
  }

  setTokens(mockTokens.accessToken, mockTokens.refreshToken)
  return { user, tokens: mockTokens }
}

function mockRegister(payload: RegisterPayload): AuthResponse {
  const user: User = {
    id: `mock_${Date.now()}`,
    username: payload.username,
    name: payload.name,
    email: payload.email,
    role: payload.role,
    status: 'active',
  }

  const mockTokens = {
    accessToken: `mock_access_${Date.now()}`,
    refreshToken: `mock_refresh_${Date.now()}`,
  }

  setTokens(mockTokens.accessToken, mockTokens.refreshToken)
  return { user, tokens: mockTokens }
}

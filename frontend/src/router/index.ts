import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import type { UserRole } from '@/types'
import LoginView from '@/views/LoginView.vue'
import RegistrationView from '@/views/RegistrationView.vue'
import DashboardLayout from '@/layouts/DashboardLayout.vue'
import DashboardView from '@/views/DashboardView.vue'
import StudentDetailView from '@/views/StudentDetailView.vue'
import SubjectsView from '@/views/SubjectsView.vue'
import AssignmentsView from '@/views/AssignmentsView.vue'
import ProgressView from '@/views/ProgressView.vue'
import ReportsView from '@/views/ReportsView.vue'
import AdminDataView from '@/views/AdminDataView.vue'
import ProfileEdit from '@/views/ProfileEdit.vue'

// Extend route meta with allowed roles
declare module 'vue-router' {
  interface RouteMeta {
    requiresAuth?: boolean
    roles?: UserRole[]
  }
}

const routes: RouteRecordRaw[] = [
  { path: '/', redirect: '/login' },
  { path: '/login', name: 'Login', component: LoginView },
  { path: '/register', name: 'Register', component: RegistrationView },
  {
    path: '/dashboard',
    component: DashboardLayout,
    meta: { requiresAuth: true },
    children: [
      {
        path: '',
        name: 'Dashboard',
        component: DashboardView,
        meta: { roles: ['admin', 'teacher', 'student', 'parent'] },
      },
      {
        path: 'students',
        name: 'Students',
        component: StudentDetailView,
        meta: { roles: ['admin', 'teacher', 'parent'] },
      },
      {
        path: 'subjects',
        name: 'Subjects',
        component: SubjectsView,
        meta: { roles: ['admin', 'teacher'] },
      },
      {
        path: 'assignments',
        name: 'Assignments',
        component: AssignmentsView,
        meta: { roles: ['admin', 'teacher', 'student'] },
      },
      {
        path: 'progress',
        name: 'Progress',
        component: ProgressView,
        meta: { roles: ['admin', 'teacher', 'parent'] },
      },
      {
        path: 'reports',
        name: 'Reports',
        component: ReportsView,
        meta: { roles: ['admin', 'teacher'] },
      },
      {
        path: 'admin-data',
        name: 'AdminData',
        component: AdminDataView,
        meta: { roles: ['admin'] },
      },
      {
        path: 'profile',
        name: 'Profile',
        component: ProfileEdit,
        meta: { roles: ['admin', 'teacher', 'student', 'parent'] },
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

// ─── Navigation Guard ────────────────────────────────────
router.beforeEach((to, _from, next) => {
  const token = localStorage.getItem('access_token')
  const userJson = localStorage.getItem('user')

  // Require auth check
  if (to.meta.requiresAuth || to.matched.some(r => r.meta.requiresAuth)) {
    if (!token || !userJson) {
      return next('/login')
    }

    // RBAC check
    const allowedRoles = to.meta.roles
    if (allowedRoles && allowedRoles.length > 0) {
      try {
        const user = JSON.parse(userJson)
        if (!allowedRoles.includes(user.role)) {
          return next('/dashboard') // Redirect to dashboard if role not allowed
        }
      } catch {
        return next('/login')
      }
    }
  }

  // Redirect authenticated users away from login/register
  if ((to.path === '/login' || to.path === '/register') && token && userJson) {
    return next('/dashboard')
  }

  next()
})

export default router

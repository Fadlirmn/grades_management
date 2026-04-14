<template>
  <v-layout class="min-h-screen">
    <!-- Sidebar -->
    <v-navigation-drawer
      v-model="drawer"
      :temporary="isMobile"
      color="surface"
      elevation="0"
      width="240"
      class="border-e-thin"
    >
      <!-- Logo -->
      <div class="d-flex align-center gap-3 pa-5 pb-6">
        <div class="bg-primary pa-2 rounded-lg d-flex">
          <v-icon color="white" size="20">mdi-school</v-icon>
        </div>
        <div>
          <div class="font-weight-bold text-body-2 text-high-emphasis" style="line-height:1.2;">Grades</div>
          <div class="font-weight-bold text-body-2 text-high-emphasis" style="line-height:1.2;">Management</div>
          <div class="text-uppercase font-weight-bold text-disabled mt-1" style="font-size:0.6rem;letter-spacing:0.12em;">Academic Atelier</div>
        </div>
      </div>

      <!-- Menu -->
      <v-list density="compact" nav class="px-3" bg-color="transparent">
        <template v-for="item in visibleMenuItems" :key="item.to">
          <v-list-item
            :to="item.to"
            :prepend-icon="item.icon"
            :title="item.title"
            :active="isActive(item.to)"
            rounded="lg"
            class="mb-1 text-body-2 font-weight-medium"
            active-class="sidebar-active"
            :exact="item.to === '/dashboard'"
          ></v-list-item>
        </template>
      </v-list>

      <!-- AI Insight Footer -->
      <template v-slot:append>
        <div class="mx-4 mb-4 rounded-xl pa-3 bg-surface-variant">
          <div class="d-flex align-center gap-1 mb-2">
            <v-icon size="16" color="primary">mdi-sparkles</v-icon>
            <span class="text-primary font-weight-bold text-uppercase" style="font-size:0.6rem;letter-spacing:0.08em;">AI Insight</span>
          </div>
          <p class="text-body-2 ma-0 text-medium-emphasis" style="line-height:1.5;">Class performance is up 12% this week in Mathematics. Review Section 4 for details.</p>
        </div>

        <v-list-item
          class="cursor-pointer border-t-thin"
          @click="$router.push('/dashboard/profile')"
        >
          <template v-slot:prepend>
            <v-avatar size="36">
              <v-img :src="authStore.userAvatar"></v-img>
            </v-avatar>
          </template>
          <v-list-item-title class="text-body-2 font-weight-bold text-high-emphasis">{{ authStore.userName }}</v-list-item-title>
          <v-list-item-subtitle class="text-caption text-medium-emphasis text-capitalize">{{ authStore.userRole }}</v-list-item-subtitle>
          <template v-slot:append>
            <v-btn icon variant="text" size="x-small" @click.stop="handleLogout" title="Logout">
              <v-icon size="18" color="error">mdi-logout</v-icon>
            </v-btn>
          </template>
        </v-list-item>
      </template>
    </v-navigation-drawer>

    <!-- Main -->
    <v-main class="bg-background">
      <!-- Top Bar -->
      <v-app-bar flat color="surface" density="comfortable" class="border-b-thin">
        <!-- Hamburger for mobile -->
        <v-app-bar-nav-icon class="d-md-none" @click="drawer = !drawer"></v-app-bar-nav-icon>

        <!-- Search Bar -->
        <div class="d-none d-md-flex align-center rounded-xl px-4 py-1 ml-4 bg-surface-variant" style="width:380px;">
          <v-icon size="18" class="text-medium-emphasis">mdi-magnify</v-icon>
          <input type="text" placeholder="Search student records or subject..." class="ml-2 text-body-2 w-full text-high-emphasis" style="outline:none;border:none;background:transparent;width:100%;">
        </div>

        <!-- Mobile search icon -->
        <v-btn icon variant="text" class="d-md-none ml-1">
          <v-icon>mdi-magnify</v-icon>
        </v-btn>

        <v-spacer></v-spacer>

        <!-- Dark Mode Toggle -->
        <v-btn icon variant="text" @click="toggleTheme" :title="isDark ? 'Switch to Light Mode' : 'Switch to Dark Mode'">
          <v-icon :color="isDark ? 'yellow-darken-2' : 'grey-darken-1'">{{ isDark ? 'mdi-white-balance-sunny' : 'mdi-weather-night' }}</v-icon>
        </v-btn>

        <!-- Notification -->
        <v-btn icon variant="text" class="ml-1">
          <v-badge dot color="error" overlap>
            <v-icon>mdi-bell-outline</v-icon>
          </v-badge>
        </v-btn>

        <!-- User Profile -->
        <div class="d-flex align-center gap-3 ml-3 mr-4 cursor-pointer" @click="$router.push('/dashboard/profile')">
          <div class="text-right d-none d-sm-block">
            <div class="text-body-2 font-weight-bold text-high-emphasis">{{ authStore.userName }}</div>
            <div class="text-caption text-medium-emphasis text-capitalize">{{ authStore.userRole }}</div>
          </div>
          <v-avatar size="40">
            <v-img :src="authStore.userAvatar"></v-img>
          </v-avatar>
        </div>
      </v-app-bar>

      <!-- Router Content -->
      <v-container fluid class="pa-4 pa-md-8">
        <router-view></router-view>
      </v-container>
    </v-main>
  </v-layout>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useTheme } from 'vuetify'
import { useAuthStore } from '@/stores/auth'
import type { UserRole } from '@/types'

const route = useRoute()
const router = useRouter()
const theme = useTheme()
const authStore = useAuthStore()

// Theme
const isDark = computed(() => theme.global.name.value === 'dark')
const toggleTheme = (): void => {
  theme.global.name.value = isDark.value ? 'light' : 'dark'
}

// Mobile responsive
const windowWidth = ref(window.innerWidth)
const isMobile = computed(() => windowWidth.value < 960)
const drawer = ref(!isMobile.value)

const onResize = (): void => {
  windowWidth.value = window.innerWidth
  drawer.value = !isMobile.value
}

onMounted(() => window.addEventListener('resize', onResize))
onBeforeUnmount(() => window.removeEventListener('resize', onResize))

// Menu items with RBAC
interface MenuItem {
  title: string
  icon: string
  to: string
  roles: UserRole[]
}

const menuItems: MenuItem[] = [
  { title: 'Dashboard', icon: 'mdi-view-grid-outline', to: '/dashboard', roles: ['admin', 'teacher', 'student', 'parent'] },
  { title: 'Students', icon: 'mdi-account-group-outline', to: '/dashboard/students', roles: ['admin', 'teacher', 'parent'] },
  { title: 'Subjects', icon: 'mdi-book-open-variant', to: '/dashboard/subjects', roles: ['admin', 'teacher'] },
  { title: 'Assignments', icon: 'mdi-clipboard-text-outline', to: '/dashboard/assignments', roles: ['admin', 'teacher', 'student'] },
  { title: 'Progress Tracking', icon: 'mdi-chart-line-variant', to: '/dashboard/progress', roles: ['admin', 'teacher', 'parent'] },
  { title: 'Reports', icon: 'mdi-chart-bar', to: '/dashboard/reports', roles: ['admin', 'teacher'] },
  { title: 'Data Management', icon: 'mdi-database-cog', to: '/dashboard/admin-data', roles: ['admin'] },
]

const visibleMenuItems = computed(() => {
  const role = authStore.userRole
  if (!role) return menuItems
  return menuItems.filter(item => item.roles.includes(role))
})

// Active menu highlighting
const isActive = (to: string): boolean => {
  if (to === '/dashboard') return route.path === '/dashboard'
  return route.path.startsWith(to)
}

// Logout
const handleLogout = async (): Promise<void> => {
  await authStore.logout()
  router.push('/login')
}
</script>

<style scoped>
.sidebar-active {
  background: rgb(var(--v-theme-primary), 0.08) !important;
  color: rgb(var(--v-theme-primary)) !important;
}
.sidebar-active :deep(.v-list-item__prepend .v-icon) {
  color: rgb(var(--v-theme-primary)) !important;
}
</style>

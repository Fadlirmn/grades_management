<template>
  <v-container fluid class="pa-0 h-screen">
    <v-row no-gutters class="h-100">
      <!-- Left Blue Panel (same as Login) -->
      <v-col cols="6" class="hidden-sm-and-down bg-blue-darken-1 d-flex align-center justify-center position-relative overflow-hidden">
        <div class="position-absolute w-100 h-100" style="background: radial-gradient(circle at 30% 60%, rgba(59,130,246,0.5) 0%, transparent 50%), radial-gradient(circle at 70% 30%, rgba(30,64,175,0.6) 0%, transparent 50%);"></div>
        <div class="position-relative z-1 text-center pa-8" style="max-width:400px;">
          <v-icon size="64" color="white" class="mb-6">mdi-account-plus</v-icon>
          <h1 class="text-white text-h3 font-weight-bold mb-4" style="line-height:1.2;">Create your<br>Academic Account.</h1>
          <p class="text-blue-lighten-3 text-body-1">Join the Grades Management System. Register your account and start managing academic records.</p>
          <div class="d-flex justify-center gap-4 mt-12">
            <div class="rounded-pill" style="width:64px;height:4px;opacity:0.3;background:white;"></div>
            <div class="rounded-pill bg-white" style="width:64px;height:4px;"></div>
            <div class="rounded-pill" style="width:64px;height:4px;opacity:0.3;background:white;"></div>
          </div>
        </div>
      </v-col>

      <!-- Right Register Panel -->
      <v-col cols="12" md="6" class="bg-grey-lighten-4 d-flex align-center justify-center pa-8">
        <div style="width:100%;max-width:420px;">
          <!-- Brand -->
          <div class="d-flex align-center justify-center gap-2 mb-6">
            <div class="bg-primary pa-2 rounded-lg d-flex">
              <v-icon color="white">mdi-school</v-icon>
            </div>
            <span class="text-primary font-weight-bold text-subtitle-1">Grades Management System</span>
          </div>

          <h2 class="text-h4 font-weight-bold text-center text-high-emphasis mb-1">Sign Up</h2>
          <p class="text-medium-emphasis text-body-2 text-center mb-6">Create a new account to get started</p>

          <!-- Role Selection (Student / Parent only) -->
          <label class="text-body-2 font-weight-bold text-high-emphasis d-block mb-2">Register as</label>
          <div class="d-flex gap-2 mb-5">
            <v-btn
              v-for="r in roles"
              :key="r.value"
              :variant="selectedRole === r.value ? 'flat' : 'outlined'"
              :color="selectedRole === r.value ? 'primary' : 'grey-darken-1'"
              :prepend-icon="r.icon"
              rounded="lg"
              class="text-none font-weight-bold flex-1"
              size="small"
              @click="selectedRole = r.value"
            >{{ r.label }}</v-btn>
          </div>

          <v-form @submit.prevent="handleRegister">
            <!-- Username -->
            <label class="text-body-2 font-weight-bold text-high-emphasis d-block mb-2">Username</label>
            <v-text-field
              v-model="username"
              placeholder="johndoe"
              prepend-inner-icon="mdi-account-outline"
              variant="solo"
              density="comfortable"
              bg-color="grey-lighten-3"
              flat
              hide-details
              rounded="lg"
              class="mb-4"
            ></v-text-field>

            <!-- Full Name -->
            <label class="text-body-2 font-weight-bold text-high-emphasis d-block mb-2">Full Name</label>
            <v-text-field
              v-model="name"
              placeholder="John Doe"
              prepend-inner-icon="mdi-badge-account-horizontal-outline"
              variant="solo"
              density="comfortable"
              bg-color="grey-lighten-3"
              flat
              hide-details
              rounded="lg"
              class="mb-4"
            ></v-text-field>

            <!-- Email -->
            <label class="text-body-2 font-weight-bold text-high-emphasis d-block mb-2">Email Address</label>
            <v-text-field
              v-model="email"
              placeholder="john@edu.com"
              append-inner-icon="mdi-at"
              variant="solo"
              density="comfortable"
              bg-color="grey-lighten-3"
              flat
              hide-details
              rounded="lg"
              class="mb-4"
            ></v-text-field>

            <!-- Password -->
            <label class="text-body-2 font-weight-bold text-high-emphasis d-block mb-2">Password</label>
            <v-text-field
              v-model="password"
              :type="showPass ? 'text' : 'password'"
              placeholder="••••••••"
              :append-inner-icon="showPass ? 'mdi-eye-off' : 'mdi-lock'"
              @click:append-inner="showPass = !showPass"
              variant="solo"
              density="comfortable"
              bg-color="grey-lighten-3"
              flat
              hide-details
              rounded="lg"
              class="mb-4"
            ></v-text-field>

            <!-- Confirm Password -->
            <label class="text-body-2 font-weight-bold text-high-emphasis d-block mb-2">Confirm Password</label>
            <v-text-field
              v-model="confirmPassword"
              :type="showConfirm ? 'text' : 'password'"
              placeholder="••••••••"
              :append-inner-icon="showConfirm ? 'mdi-eye-off' : 'mdi-lock-check'"
              @click:append-inner="showConfirm = !showConfirm"
              variant="solo"
              density="comfortable"
              bg-color="grey-lighten-3"
              flat
              hide-details
              rounded="lg"
              class="mb-5"
            ></v-text-field>

            <!-- Validation Error -->
            <v-alert v-if="validationError" type="warning" variant="tonal" density="compact" class="mb-4" rounded="lg">
              {{ validationError }}
            </v-alert>

            <!-- Store Error -->
            <v-alert v-if="authStore.error" type="error" variant="tonal" density="compact" class="mb-4" rounded="lg">
              {{ authStore.error }}
            </v-alert>

            <v-btn
              type="submit"
              color="primary"
              size="x-large"
              block
              flat
              rounded="lg"
              append-icon="mdi-arrow-right"
              class="text-none font-weight-bold mb-5"
              :loading="authStore.isLoading"
            >
              Create Account
            </v-btn>
          </v-form>

          <div class="text-center text-body-2 text-medium-emphasis bg-grey-lighten-3 pa-4 rounded-xl">
            Already have an account? <router-link to="/login" class="text-primary font-weight-bold text-decoration-none">Sign In</router-link>
          </div>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import type { UserRole } from '@/types'

const router = useRouter()
const authStore = useAuthStore()

const username = ref('')
const name = ref('')
const email = ref('')
const password = ref('')
const confirmPassword = ref('')
const showPass = ref(false)
const showConfirm = ref(false)
const selectedRole = ref<UserRole>('student')
const validationError = ref<string | null>(null)

interface RoleOption {
  value: UserRole
  label: string
  icon: string
}

const roles: RoleOption[] = [
  { value: 'student', label: 'Student', icon: 'mdi-school' },
  { value: 'parent', label: 'Parent', icon: 'mdi-account-child' },
]

const handleRegister = async (): Promise<void> => {
  validationError.value = null

  // Basic validation
  if (!username.value.trim() || !name.value.trim() || !email.value.trim() || !password.value) {
    validationError.value = 'All fields are required'
    return
  }
  if (password.value.length < 6) {
    validationError.value = 'Password must be at least 6 characters'
    return
  }
  if (password.value !== confirmPassword.value) {
    validationError.value = 'Passwords do not match'
    return
  }

  try {
    await authStore.register({
      username: username.value.trim(),
      name: name.value.trim(),
      email: email.value.trim(),
      password: password.value,
      role: selectedRole.value,
    })
    router.push('/dashboard')
  } catch {
    // Error handled by store
  }
}
</script>

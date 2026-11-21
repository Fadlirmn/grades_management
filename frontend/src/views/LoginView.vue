<template>
  <v-container fluid class="pa-0 h-screen">
    <v-row no-gutters class="h-100">
      <!-- Left Blue Panel -->
      <v-col cols="6" class="d-none d-md-flex bg-blue-darken-1 align-center justify-center position-relative overflow-hidden">
        <div class="position-absolute w-100 h-100" style="background: radial-gradient(circle at 30% 60%, rgba(59,130,246,0.5) 0%, transparent 50%), radial-gradient(circle at 70% 30%, rgba(30,64,175,0.6) 0%, transparent 50%);"></div>
        <div class="position-relative z-1 text-center pa-8" style="max-width:400px;">
          <v-icon size="64" color="white" class="mb-6">mdi-book-open-page-variant</v-icon>
          <h1 class="text-white text-h3 font-weight-bold mb-4" style="line-height:1.2;">Welcome to your<br>Academic Atelier.</h1>
          <p class="text-blue-lighten-3 text-body-1">Empowering education with editorial intelligence. A curated workspace for the modern educator.</p>
          <div class="d-flex justify-center gap-4 mt-12">
            <div class="rounded-pill bg-white" style="width:64px;height:4px;"></div>
            <div class="rounded-pill" style="width:64px;height:4px;opacity:0.3;background:white;"></div>
            <div class="rounded-pill" style="width:64px;height:4px;opacity:0.3;background:white;"></div>
          </div>
        </div>
      </v-col>

      <!-- Right Login Panel -->
      <v-col cols="12" md="6" class="bg-grey-lighten-4 d-flex align-center justify-center pa-8">
        <div style="width:100%;max-width:420px;">
          <!-- Brand -->
          <div class="d-flex align-center justify-center gap-2 mb-6">
            <div class="bg-primary pa-2 rounded-lg d-flex">
              <v-icon color="white">mdi-school</v-icon>
            </div>
            <span class="text-primary font-weight-bold text-subtitle-1">Grades Management System</span>
          </div>

          <h2 class="text-h4 font-weight-bold text-center text-high-emphasis mb-1">Sign In</h2>
          <p class="text-medium-emphasis text-body-2 text-center mb-6">Enter your credentials to access your account</p>

          <v-form @submit.prevent="handleLogin">
            <label class="text-body-2 font-weight-bold text-high-emphasis d-block mb-2">Email or Username</label>
            <v-text-field
              v-model="email"
              placeholder="your@email.com"
              append-inner-icon="mdi-at"
              variant="solo"
              density="comfortable"
              bg-color="grey-lighten-3"
              flat
              hide-details
              rounded="lg"
              class="mb-5"
            ></v-text-field>

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

            <div class="d-flex align-center justify-space-between mb-6">
              <v-checkbox v-model="keepLogged" label="Keep me logged in" color="primary" hide-details density="compact"></v-checkbox>
              <a href="#" class="text-body-2 font-weight-bold text-primary text-decoration-none">Forgot Password?</a>
            </div>

            <!-- Error -->
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
              class="text-none font-weight-bold mb-6"
              :loading="authStore.isLoading"
            >
              Sign In
            </v-btn>
          </v-form>

          <div class="text-center text-body-2 text-medium-emphasis bg-grey-lighten-3 pa-4 rounded-xl mb-4">
            Don't have an account? <router-link to="/register" class="text-primary font-weight-bold text-decoration-none">Sign Up</router-link> or Contact administrator
          </div>

          <v-alert type="info" variant="tonal" density="compact" rounded="lg" class="text-body-2">
            <strong>Demo mode:</strong> any password works. Try
            <code>admin@edu.com</code>, <code>teacher@edu.com</code>,
            <code>student@edu.com</code>, or <code>parent@edu.com</code>.
          </v-alert>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const email = ref('')
const password = ref('')
const showPass = ref(false)
const keepLogged = ref(false)

const handleLogin = async (): Promise<void> => {
  try {
    await authStore.login({ email: email.value, password: password.value })
    router.push('/dashboard')
  } catch {
    // Error is handled by the store
  }
}
</script>

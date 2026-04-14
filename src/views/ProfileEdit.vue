<template>
  <div>
    <div class="mb-7">
      <h1 class="text-h5 text-md-h4 font-weight-bold text-high-emphasis">Edit Profile</h1>
      <p class="text-body-2 mt-1 text-medium-emphasis">Update your personal information and preferences.</p>
    </div>

    <div class="grid grid-cols-4 gap-5 gap-md-6">
      <v-card color="surface" variant="flat" rounded="xl" class="pa-7 col-span-1 border-thin">
        <div class="d-flex flex-column align-center mb-6 pb-6 border-b-thin">
          <v-avatar size="100" class="mb-4">
            <v-img :src="authStore.userAvatar"></v-img>
          </v-avatar>
          <v-btn variant="outlined" color="primary" size="small" rounded="lg" class="text-none font-weight-bold" prepend-icon="mdi-camera">Change Photo</v-btn>
        </div>
        <div class="d-flex flex-column gap-4">
          <div class="d-flex justify-space-between align-center">
            <span class="text-body-2 text-medium-emphasis">Role</span>
            <v-chip color="primary" size="x-small" variant="tonal" class="font-weight-bold text-capitalize">{{ authStore.userRole }}</v-chip>
          </div>
          <div class="d-flex justify-space-between">
            <span class="text-body-2 text-medium-emphasis">Department</span>
            <span class="text-body-2 font-weight-bold text-high-emphasis">{{ authStore.user?.department || '—' }}</span>
          </div>
          <div class="d-flex justify-space-between">
            <span class="text-body-2 text-medium-emphasis">Joined</span>
            <span class="text-body-2 font-weight-bold text-high-emphasis">Aug 2021</span>
          </div>
          <div class="d-flex justify-space-between align-center">
            <span class="text-body-2 text-medium-emphasis">Status</span>
            <div class="d-flex align-center gap-1">
              <div class="rounded-circle bg-success" style="width:8px;height:8px;"></div>
              <span class="text-body-2 font-weight-bold text-high-emphasis">Active</span>
            </div>
          </div>
        </div>
      </v-card>

      <v-card color="surface" variant="flat" rounded="xl" class="pa-7 col-span-3 border-thin">
        <h3 class="text-body-1 font-weight-bold text-high-emphasis mb-4">Personal Information</h3>
        <div class="grid grid-cols-2 gap-4 mb-8">
          <div>
            <label class="text-body-2 font-weight-bold text-high-emphasis d-block mb-2">Full Name</label>
            <v-text-field v-model="form.name" variant="solo" density="comfortable" bg-color="surface-variant" flat hide-details rounded="lg"></v-text-field>
          </div>
          <div>
            <label class="text-body-2 font-weight-bold text-high-emphasis d-block mb-2">Email Address</label>
            <v-text-field v-model="form.email" variant="solo" density="comfortable" bg-color="surface-variant" flat hide-details rounded="lg"></v-text-field>
          </div>
          <div>
            <label class="text-body-2 font-weight-bold text-high-emphasis d-block mb-2">Phone Number</label>
            <v-text-field v-model="form.phone" variant="solo" density="comfortable" bg-color="surface-variant" flat hide-details rounded="lg"></v-text-field>
          </div>
          <div>
            <label class="text-body-2 font-weight-bold text-high-emphasis d-block mb-2">Username</label>
            <v-text-field v-model="form.username" variant="solo" density="comfortable" bg-color="surface-variant" flat hide-details rounded="lg"></v-text-field>
          </div>
        </div>

        <h3 class="text-body-1 font-weight-bold text-high-emphasis mb-4">Security</h3>
        <div class="grid grid-cols-2 gap-4 mb-8">
          <div>
            <label class="text-body-2 font-weight-bold text-high-emphasis d-block mb-2">Current Password</label>
            <v-text-field v-model="form.currentPassword" type="password" variant="solo" density="comfortable" bg-color="surface-variant" flat hide-details rounded="lg"></v-text-field>
          </div>
          <div>
            <label class="text-body-2 font-weight-bold text-high-emphasis d-block mb-2">New Password</label>
            <v-text-field v-model="form.newPassword" type="password" variant="solo" density="comfortable" bg-color="surface-variant" flat hide-details rounded="lg"></v-text-field>
          </div>
        </div>

        <h3 class="text-body-1 font-weight-bold text-high-emphasis mb-4">Preferences</h3>
        <div class="d-flex flex-column">
          <div class="d-flex justify-space-between align-center py-3 border-b-thin">
            <div>
              <div class="text-body-2 font-weight-bold text-high-emphasis">Dark Mode</div>
              <div class="text-caption text-medium-emphasis">Switch between light and dark theme</div>
            </div>
            <v-switch v-model="darkMode" color="primary" hide-details @update:model-value="toggleTheme"></v-switch>
          </div>
          <div class="d-flex justify-space-between align-center py-3 border-b-thin">
            <div>
              <div class="text-body-2 font-weight-bold text-high-emphasis">Email Notifications</div>
              <div class="text-caption text-medium-emphasis">Receive email updates about student progress</div>
            </div>
            <v-switch v-model="emailNotif" color="primary" hide-details></v-switch>
          </div>
          <div class="d-flex justify-space-between align-center py-3">
            <div>
              <div class="text-body-2 font-weight-bold text-high-emphasis">AI Suggestions</div>
              <div class="text-caption text-medium-emphasis">Show AI-powered recommendations on dashboard</div>
            </div>
            <v-switch v-model="aiSuggestions" color="primary" hide-details></v-switch>
          </div>
        </div>

        <div class="d-flex flex-column flex-sm-row justify-end gap-3 mt-7 pt-5 border-t-thin">
          <v-btn variant="outlined" rounded="lg" class="text-none font-weight-bold text-medium-emphasis">Cancel</v-btn>
          <v-btn color="primary" elevation="0" rounded="lg" class="text-none font-weight-bold">Save Changes</v-btn>
        </div>
      </v-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useTheme } from 'vuetify'
import { useAuthStore } from '@/stores/auth'

const theme = useTheme()
const authStore = useAuthStore()

const darkMode = ref<boolean>(theme.global.name.value === 'dark')
const emailNotif = ref(true)
const aiSuggestions = ref(true)

interface ProfileForm {
  name: string; email: string; phone: string; username: string
  currentPassword: string; newPassword: string
}

const form = ref<ProfileForm>({
  name: authStore.userName || 'Dr. Helena Vance',
  email: authStore.user?.email || 'helena.vance@academy.edu',
  phone: '+62 812 3456 7890',
  username: authStore.user?.email?.split('@')[0] || 'helena.vance',
  currentPassword: '',
  newPassword: '',
})

const toggleTheme = (val: boolean | null): void => {
  theme.global.name.value = val ? 'dark' : 'light'
}
</script>

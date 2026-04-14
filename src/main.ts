import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router/index'
import vuetify from './plugins/vuetify'
import { useAuthStore } from './stores/auth'
import 'uno.css'
import './style.css'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)
app.use(vuetify)

// Restore user session before mounting
const authStore = useAuthStore()
authStore.restoreSession()

app.mount('#app')

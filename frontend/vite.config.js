import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vuetify from 'vite-plugin-vuetify'
import UnoCSS from 'unocss/vite'
import path from 'path'

export default defineConfig({
  // GH_PAGES=true set by .github/workflows/deploy-pages.yml for the project-page base path
  base: process.env.GH_PAGES ? '/grades_management/' : '/',
  plugins: [
    vue(),
    vuetify({ autoImport: true }),
    UnoCSS(),
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src')
    }
  }
})

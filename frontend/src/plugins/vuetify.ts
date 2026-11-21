import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

export default createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: 'light',
    themes: {
      light: {
        dark: false,
        colors: {
          primary: '#3B82F6',
          secondary: '#6177A5',
          tertiary: '#D16900',
          background: '#F8FAFC',
          surface: '#FFFFFF',
          'surface-variant': '#F1F5F9',
          'on-background': '#121212',
          'on-surface': '#121212',
          error: '#DC2626',
          success: '#16A34A',
          warning: '#D16900',
          info: '#3B82F6',
        },
      },
      dark: {
        dark: true,
        colors: {
          primary: '#3B82F6',
          secondary: '#6177A5',
          tertiary: '#D16900',
          background: '#0F1117',
          surface: '#1A1D27',
          'surface-variant': '#252836',
          'on-background': '#E2E8F0',
          'on-surface': '#E2E8F0',
          error: '#EF4444',
          success: '#22C55E',
          warning: '#D16900',
          info: '#3B82F6',
        },
      },
    },
  },
})

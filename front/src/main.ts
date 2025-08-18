// main.ts - Configuration du thème personnalisé
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config'
import { definePreset } from '@primevue/themes'
import Aura from '@primevue/themes/aura'
import App from './App.vue'
import router from './router'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import Ripple from 'primevue/ripple'

const blackOrangePreset = definePreset(Aura, {
  semantic: {
    primary: {
      50: '#fff7ed',
      100: '#ffedd5', 
      200: '#fed7aa',
      300: '#fdba74',
      400: '#fb923c',
      500: '#f97316',
      600: '#ea580c',
      700: '#c2410c',
      800: '#9a3412',
      900: '#7c2d12',
      950: '#431407'
    },
    colorScheme: {
      light: {
        primary: {
          color: '{primary.600}',
          contrastColor: '#ffffff',
          hoverColor: '{primary.700}',
          activeColor: '{primary.800}'
        },
        surface: {
          0: '#ffffff',
          50: '#f8fafc',
          100: '#f1f5f9',
          200: '#e2e8f0',
          300: '#cbd5e1',
          400: '#94a3b8',
          500: '#64748b',
          600: '#475569',
          700: '#334155',
          800: '#1e293b',
          900: '#0f172a', 
          950: '#020617' 
        }
      },
      dark: {
        primary: {
          color: '{primary.500}',
          contrastColor: '{surface.900}',
          hoverColor: '{primary.400}',
          activeColor: '{primary.300}'
        },
        surface: {
          0: '#0f172a',
          50: '#1e293b',
          100: '#334155',
          200: '#475569',
          300: '#64748b',
          400: '#94a3b8',
          500: '#cbd5e1',
          600: '#e2e8f0',
          700: '#f1f5f9',
          800: '#f8fafc',
          900: '#ffffff',
          950: '#ffffff'
        }
      }
    }
  },
  components: {
    menubar: {
      root: {
        background: '{surface.0}',
        borderColor: 'transparent',
        color: '{surface.0}',
        borderRadius: '0',
        gap: '0',
        padding: '0.75rem 1rem',
        borderBottom: '2px solid {primary.600}'
      },
      item: {
        focusBackground: '{primary.600}',
        activeBackground: '{primary.700}',
        color: '#ffffff',
        focusColor: '#ffffff',
        activeColor: '#ffffff',
        padding: '0.75rem 1rem',
        borderRadius: '0.375rem',
        fontWeight: '500',
        transition: 'all 0.3s ease'
      },
      submenu: {
        background: '{surface.50}',
        borderColor: '{primary.600}',
        color: '#ffffff',
        shadow: '0 4px 12px rgba(249, 115, 22, 0.2)',
        borderRadius: '0.5rem'
      },
      submenuitem: {
        focusBackground: '{primary.600}',
        activeBackground: '{primary.700}',
        color: '#ffffff',
        focusColor: '#ffffff',
        activeColor: '#ffffff'
      }
    },
    button: {
      root: {
        background: '{primary.600}',
        hoverBackground: '{primary.700}',
        activeBackground: '{primary.800}',
        borderColor: '{primary.600}',
        color: '#ffffff',
        borderRadius: '0.5rem',
        fontWeight: '600',
        transition: 'all 0.3s ease',
        boxShadow: '0 2px 4px rgba(249, 115, 22, 0.2)'
      },
      outlined: {
        background: 'transparent',
        hoverBackground: '{primary.600}',
        activeBackground: '{primary.700}',
        borderColor: '{primary.600}',
        color: '{primary.600}',
        hoverColor: '#ffffff',
        activeColor: '#ffffff'
      },
      text: {
        primary: {
          background: 'transparent',
          hoverBackground: 'rgba(249, 115, 22, 0.1)',
          activeBackground: 'rgba(249, 115, 22, 0.2)',
          color: '{primary.600}',
          hoverColor: '{primary.700}',
          activeColor: '{primary.800}'
        }
      }
    },
    card: {
      root: {
        background: '{surface.0}',
        borderColor: '{surface.100}',
        color: '#ffffff',
        borderRadius: '0.75rem',
        shadow: '0 4px 12px rgba(0, 0, 0, 0.3)'
      },
      body: {
        padding: '1.25rem'
      },
      title: {
        color: '#ffffff',
        fontWeight: '700'
      },
      subtitle: {
        color: '{surface.400}'
      }
    }
  }
})

const app = createApp(App)

app.use(PrimeVue, {
  theme: {
    preset: blackOrangePreset,
    options: {
      darkModeSelector: 'system',
      cssLayer: false
    }
  }
})

const pinia = createPinia();

app.use(pinia)
app.directive('ripple', Ripple)
pinia.use(piniaPluginPersistedstate)
app.use(router)
app.mount('#app')
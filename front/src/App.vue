
<script setup lang="ts">
import { ref } from 'vue'
import WalletConnect from './components/WalletConnect.vue'

import Menubar from 'primevue/menubar'
import Avatar from 'primevue/avatar'
import Badge from 'primevue/badge'

import 'primeicons/primeicons.css'

const items = ref([
  {
    label: 'Home',
    icon: 'pi pi-home',
    route: '/'
  },
  {
    label: 'Matchs en direct',
    icon: 'pi pi-trophy',
    route: '/matches'
  },
  {
    label: 'Tournois',
    icon: 'pi pi-building-columns',
    items: [
      {
        label: 'CSGO',
        icon: '',
        route: 'tournois/csgo'
      }
    ]
  },
  {
    label: 'Statistiques',
    icon: 'pi pi-chart-bar',
    route: '/stats'
  },
  {
    label: 'Historique',
    icon: 'pi pi-tablet',
    route: '/history'
  }
])
</script>

<template>
  <div class="app-container">
    <Menubar :model="items">
      <template #start>
      </template>

      <template #item="{ item, props, hasSubmenu, root }">
        <router-link v-if="item.route" v-slot="{ href, navigate }" :to="item.route" custom>
          <a v-ripple :href="href" @click="navigate" class="flex items-center" v-bind="props.action">
            <span :class="item.icon" />
            <span class="ml-2">{{ item.label }}</span>
            <Badge v-if="item.badge" :class="{ 'ml-auto': !root, 'ml-2': root }" :value="item.badge" />
            <span v-if="item.shortcut" class="ml-auto border border-surface rounded bg-emphasis text-muted-color text-xs p-1">{{ item.shortcut }}</span>
          </a>
        </router-link>
        <a v-else v-ripple class="flex items-center" v-bind="props.action">
          <span :class="item.icon" />
          <span class="ml-2">{{ item.label }}</span>
          <Badge v-if="item.badge" :class="{ 'ml-auto': !root, 'ml-2': root }" :value="item.badge" />
          <span v-if="item.shortcut" class="ml-auto border border-surface rounded bg-emphasis text-muted-color text-xs p-1">{{ item.shortcut }}</span>
          <i v-if="hasSubmenu" :class="['pi pi-angle-down ml-auto', { 'pi-angle-down': root, 'pi-angle-right': !root }]"></i>
        </a>
      </template>
      <template #end>
        <div style="display: flex; align-items: center; gap: 30px;">
          <WalletConnect style="display: inline-block;" />
          <router-link to="/profil">
            <a>
              <Avatar icon="pi pi-user" size="xlarge" shape="circle"/>
            </a>
          </router-link>
        </div>
      </template>
    </Menubar>
    
    <main class="main-content">
      <RouterView v-slot="{ Component }">
        <Transition name="page-opacity" mode="out-in">
          <component :is="Component"/>
        </Transition>
      </RouterView>
    </main>
  </div>
</template>

<style scoped>


.page-opacity-enter-active,
.page-opacity-leave-active {
  transition: 600ms ease all;
}

.page-opacity-enter-from,
.page-opacity-leave-to {
  opacity: 0;
}

.app-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #0f172a 0%, #1e293b 100%);
  color: white;
}

.custom-menubar {
  border-bottom: 2px solid #ea580c;
  backdrop-filter: blur(10px);
  background: rgba(15, 23, 42, 0.95) !important;
}

.logo-container {
  display: flex;
  align-items: center;
  padding: 0.5rem 0;
}

.main-content {
  min-height: calc(100vh - 120px);
  padding: 2rem;
}

.app-footer {
  background: rgba(15, 23, 42, 0.8);
  border-top: 1px solid #334155;
  margin-top: auto;
}

.login-btn {
  background: linear-gradient(45deg, #ea580c, #f97316) !important;
  border: none !important;
  font-weight: 600;
  transition: all 0.3s ease;
}

.login-btn:hover {
  background: linear-gradient(45deg, #c2410c, #ea580c) !important;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(249, 115, 22, 0.4);
}

/* Styles pour les éléments du menu */
:deep(.p-menubar-root-list > .p-menuitem > .p-menuitem-link) {
  color: white !important;
  transition: all 0.3s ease;
}

:deep(.p-menubar-root-list > .p-menuitem > .p-menuitem-link:hover) {
  background: linear-gradient(45deg, #ea580c, #f97316) !important;
  color: white !important;
  transform: translateY(-1px);
  box-shadow: 0 2px 8px rgba(249, 115, 22, 0.3);
}

:deep(.p-menubar-submenu) {
  background: #1e293b !important;
  border: 1px solid #ea580c !important;
  box-shadow: 0 4px 12px rgba(249, 115, 22, 0.2) !important;
}

:deep(.p-menubar-submenu .p-menuitem-link) {
  color: white !important;
}

:deep(.p-menubar-submenu .p-menuitem-link:hover) {
  background: #ea580c !important;
  color: white !important;
}

/* Animation d'apparition */
.app-container {
  animation: fadeIn 0.5s ease-in-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
import Home from '@/Home.vue'
import NotFound from '@/views/NotFound.vue'
import Profil from '@/views/Profil.vue'
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {path: '/', name: 'Home', component: Home},
  {path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound},
  {path: '/profil', name: 'Profil', component: Profil}
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router

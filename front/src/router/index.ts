import Home from '@/MainHome.vue'
import MatchesCurrent from '@/views/Matches/MatchesCurrent.vue'
import MatchesSelector from '@/views/Matches/MatchesSelector.vue'
import MatchesPast from '@/views/Matches/MatchesPast.vue'
import MatchesUpcoming from '@/views/Matches/MatchesUpcoming.vue'
import NotFound from '@/views/NotFound.vue'
import Profil from '@/views/ProfilClient.vue'
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {path: '/', name: 'Home', component: Home},
  {path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound},
  {path: '/profil', name: 'Profil', component: Profil},
  {path: '/matches', name: 'Matches', component: MatchesSelector},
  {path: '/matches/current/:teamId?', name: 'MatchesCurrent', component: MatchesCurrent, props: true},
  {path: '/matches/past/:teamId?', name: 'MatchesPast', component: MatchesPast, props: true},
  {path: '/matches/upcoming/:teamId?', name: 'MatchesUpcoming', component: MatchesUpcoming, props: true},
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
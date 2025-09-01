import Home from '@/MainHome.vue'
import MatchesCurrent from '@/views/Matches/MatchesCurrent.vue'
import MatchesSelector from '@/views/Matches/MatchesSelector.vue'
import MatchesPast from '@/views/Matches/MatchesPast.vue'
import MatchesUpcoming from '@/views/Matches/MatchesUpcoming.vue'
import NotFound from '@/views/NotFound.vue'
import Profil from '@/views/ProfilClient.vue'
import { createRouter, createWebHistory } from 'vue-router'
import MatchesBets from '@/views/Matches/MatchesBets.vue'

const routes = [
  {path: '/', name: 'Home', component: Home},
  {path: '/profil', name: 'Profil', component: Profil},

  {path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound},

  {path: '/matches', name: 'Matches', component: MatchesSelector},
  {path: '/matches/current/', name: 'MatchesCurrent', component: MatchesCurrent, props: true},
  {path: '/matches/past/', name: 'MatchesPast', component: MatchesPast, props: true},
  {path: '/matches/upcoming/', name: 'MatchesUpcoming', component: MatchesUpcoming, props: true},
  {path: '/bets', name: 'Bets', component: MatchesBets},
  {path: '/bets/current/:id', name: 'MatchesCurrentBets', component: MatchesBets, props: true},
  {path: '/bets/upcoming/:id', name: 'MatchesUpcomingBets', component: MatchesBets, props: true}
]

/**
 * Pour la route Bets -> format : /bets/current/id (du match)
 * Pour upcoming: /bets/upcoming/id (du match)
 */
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
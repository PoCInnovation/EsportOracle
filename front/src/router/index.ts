import Home from '@/MainHome.vue'
import MatchesCurrent from '@/views/Matches/MatchesCurrent.vue'
import MatchesSelector from '@/views/Matches/MatchesSelector.vue'
import MatchesPast from '@/views/Matches/MatchesPast.vue'
import MatchesUpcoming from '@/views/Matches/MatchesUpcoming.vue'
import NotFound from '@/views/NotFound.vue'
import Profil from '@/views/ProfilClient.vue'
import { createRouter, createWebHistory } from 'vue-router'
import MatchesBetsTeams from '@/views/Matches/MatchesBetsTeams.vue'
import MatchesSelectorBets from '@/views/Matches/MatchesSelectorBets.vue'
import MatchesHistoryBets from '@/views/Matches/MatchesHistoryBets.vue'
import { matchStore } from '@/stores/matchStore'

const mapView = (v?: string) =>
  v === 'history' ? 'past' : (v as 'current'|'upcoming'|'past' | undefined) ?? 'current'

const routes = [
  {path: '/', name: 'Home', component: Home},
  {path: '/profil', name: 'Profil', component: Profil},

  {path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound},

  {path: '/matches', name: 'Matches', component: MatchesSelector},
  
  {path: '/matches/current/', name: 'MatchesCurrent', component: MatchesCurrent, props: true},
  {path: '/matches/past/', name: 'MatchesPast', component: MatchesPast, props: true},
  {path: '/matches/upcoming/', name: 'MatchesUpcoming', component: MatchesUpcoming, props: true},
  
  {path: '/bets', name: 'Bets', component: MatchesSelectorBets, meta: { view: null }},
  
  { path: '/bets/current',  name: 'BetsCurrent',  component: MatchesSelectorBets, 
    async beforeEnter(to) {
      console.log('[guard] to:', to.fullPath, 'view=', to.params.view);
      const store = matchStore()
      const TypeMatches = (to.meta.view as string) ?? 'current'
      const Url = store.createUrlMatches("current", "")
      await store.fetchMatches(Url, mapView(TypeMatches))
      console.log("Succesfully fetched matches");
      return true
    }, meta: { view: 'current' } },
    
  { path: '/bets/upcoming', name: 'BetsUpcoming', component: MatchesSelectorBets,
    async beforeEnter(to) {
      console.log('[guard] to:', to.fullPath, 'view=', to.params.view);
      const store = matchStore()
      const TypeMatches = (to.meta.view as string) ?? 'upcoming'
      const Url = store.createUrlMatches("upcoming", "")
      await store.fetchMatches(Url, mapView(TypeMatches))
      console.log("Succesfully fetched matches");
      return true
    }, meta: { view: 'upcoming' } },

  { path: '/bets/history',  name: 'BetsHistory',  component: MatchesHistoryBets, meta: { view: 'history' } },


  {path: '/bets/current/:id(\\d+)', name: 'MatchesCurrentBets', component: MatchesBetsTeams, props: true},
  {path: '/bets/upcoming/:id(\\d+)', name: 'MatchesUpcomingBets', component: MatchesBetsTeams, props: true},
  
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
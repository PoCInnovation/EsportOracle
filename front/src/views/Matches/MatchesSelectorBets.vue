<template>
    <div>
        <nav class="nav-desktop">
            <router-link
            v-for="item in navigation"
            :key="item.name"
            :to="item.to">
            <span class="page-title">{{ item.name }}</span>
            </router-link>
        </nav>
        <div class="bets-container">
            <div>
              <MatchOfTheDay  v-if="specificMatch" :match="specificMatch"/>
            </div>
          <div class="matches-grid" v-if="!isLoading && !combinedError && matchesWithActiveBets.length">
            <MatchCard
              v-for="match in matchesWithActiveBets"
              :key="match.id"
              :match="match"
              class="match-item"
            />
         </div>
         <div v-else-if="combinedError" class="bets-feedback error">Une erreur est survenue lors du chargement des paris : {{ combinedError }}</div>
         <div v-else-if="isLoading" class="bets-feedback">Chargement des paris...</div>
         <div v-else class="bets-feedback">Aucun pari disponible pour cette catégorie pour le moment.</div>
        </div>
    </div>
</template>

<script setup lang="ts">
import MatchOfTheDay from '@/components/MatchOfTheDay.vue';
import MatchCard from '@/components/MatchCard.vue'
import { matchStore } from '@/stores/matchStore';
import { betStore } from '@/stores/betStore';

import { storeToRefs } from 'pinia';
import { useRoute } from 'vue-router';
import { computed, onMounted, onUnmounted, watch } from 'vue';

const route = useRoute();


const isValidViewType = (value: any): value is ViewType => {
  return ['current', 'upcoming'].includes(value);
};

type ViewType = 'current' | 'upcoming'
type ApiViewType = 'current' | 'upcoming' | 'past';
const view = computed<ViewType | null>(() => {
  return isValidViewType(route.meta.view) ? route.meta.view : null;
});


const viewMapper: Record<ViewType, ApiViewType> = {
  'current': 'current',
  'upcoming': 'upcoming',
};

const matchesStore = matchStore()
const betsStore = betStore()

const { matches, currentMatches, upcomingMatches, loading: matchesLoading, error: matchesError } = storeToRefs(matchesStore)
const { openBets, loading: betsLoading, error: betsError } = storeToRefs(betsStore)

let Url = view.value === "current" || view.value === "upcoming" ? matchesStore.createUrlMatches(view.value, "") : null
let autoRefreshInterval: NodeJS.Timeout | null = null

const specificMatch = computed(() => {
  return matches.value.filter(match =>
  ["s", "a"].includes(match.tournament?.tier ?? "")
  )[0] ?? null;
})

const navigation = [
  { name: 'Home',     to: '/bets' },
  { name: 'Live',     to: '/bets/current' },
  { name: 'Upcoming', to: '/bets/upcoming' },
  { name: 'History',  to: '/bets/history' },
]

const matchesFromView = computed(() => {
  if (view.value === 'current') {
    return currentMatches.value
  }
  if (view.value === 'upcoming') {
    return upcomingMatches.value
  }
  return matches.value
})

const activeMatchIds = computed(() => new Set(openBets.value.map(bet => Number(bet.bet.matchId))))

const matchesWithActiveBets = computed(() => {
  return matchesFromView.value.filter(match => activeMatchIds.value.has(Number(match.id)))
})

const isLoading = computed(() => matchesLoading.value || betsLoading.value)
const combinedError = computed(() => matchesError.value || betsError.value)


const startAutoRefresh = () => {
  if (autoRefreshInterval) {
    clearInterval(autoRefreshInterval)
  }
  
  autoRefreshInterval = setInterval(async () => {
    if (!matchesLoading.value) {
      if (view.value !=  null && Url) {
        await matchesStore.fetchMatches(Url, viewMapper[view.value])
      } else {
        const currentUrl = matchesStore.createUrlMatches('current', '')
        const upcomingUrl = matchesStore.createUrlMatches('upcoming', '')
        await matchesStore.fetchMatches(currentUrl, 'current')
        await matchesStore.fetchMatches(upcomingUrl, 'upcoming')
      }
    }
    await betsStore.refresh()
  }, 30000)
}

const stopAutoRefresh = () => {
  if (autoRefreshInterval) {
    clearInterval(autoRefreshInterval)
    autoRefreshInterval = null
  }
}

/**
 * Component lifecycle - fetch matches on mount
 */
onMounted(async () => {
  if (view.value !=  null && Url) {
    await matchesStore.fetchMatches(Url, view.value)
  } else {
    const currentUrl = matchesStore.createUrlMatches('current', '')
    const upcomingUrl = matchesStore.createUrlMatches('upcoming', '')
    await matchesStore.fetchMatches(currentUrl, 'current')
    await matchesStore.fetchMatches(upcomingUrl, 'upcoming')
  }
  
  // Set up auto-refresh every 30 seconds for live updates
  await betsStore.fetchBets()
  startAutoRefresh()
})

onUnmounted(() => {
  stopAutoRefresh()
})

watch((view), async (newView) => {
  if (newView === "current" || newView === "upcoming") {
    Url = matchesStore.createUrlMatches(newView, "")
    stopAutoRefresh()

    await matchesStore.fetchMatches(Url, newView)
    await betsStore.refresh()
    startAutoRefresh()
  } else {
    console.log("Back to /home", newView)
  }
})

</script>

<style lang="css" scoped>

@import "../../styles/matches.css";
@import "../../styles/navbar.css";

/** Container center */

.bets-container {
  max-width: 900px;
  margin: 0 auto;
  padding: 2rem;
  min-height: calc(100vh - 200px);
  position: relative;
  overflow: hidden;
}

/* Floating Background Elements */
.bets-container::before {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: 
    radial-gradient(circle at 20% 20%, rgba(249, 115, 22, 0.1) 0%, transparent 50%),
    radial-gradient(circle at 80% 80%, rgba(251, 146, 60, 0.08) 0%, transparent 50%),
    radial-gradient(circle at 40% 60%, rgba(234, 88, 12, 0.05) 0%, transparent 50%);
  animation: float-bg 20s ease-in-out infinite;
  pointer-events: none;
  z-index: -1;
}

.bets-container::after {
  content: '';
  position: fixed;
  top: 10%;
  right: 5%;
  width: 100px;
  height: 100px;
  background: linear-gradient(135deg, rgba(249, 115, 22, 0.15), rgba(251, 146, 60, 0.1));
  border-radius: 50%;
  filter: blur(40px);
  animation: float-right 15s ease-in-out infinite;
  pointer-events: none;
  z-index: -1;
}

.bets-feedback {
  text-align: center;
  color: rgba(255, 255, 255, 0.7);
  font-size: 0.95rem;
  padding: 2rem 0;
}

.bets-feedback.error {
  color: #f87171;
}

/* Responsive adjustments for floating elements */

@media (max-width: 768px) {
  .bets-container {
    padding: 1rem;
  }

  .nav-desktop {
    flex-direction: column;
    gap: 1rem;
    text-align: center;
    padding: 1.5rem;
  }
}

@media (max-width: 480px) {
  .bets-container {
    padding: 0.75rem;
  }

  .nav-desktop {
    padding: 1rem;
  }

  page-title {
    font-size: clamp(1.5rem, 5vw, 2rem);
  }
}

</style>

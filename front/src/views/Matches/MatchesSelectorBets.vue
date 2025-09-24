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
          <div  v-if="view" class="matches-grid">
            <MatchCard
              v-for="match in ChoiceMatchType(view)"
              :key="match.id"
              :match="match"
              class="match-item"
            />
         </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import MatchOfTheDay from '@/components/MatchOfTheDay.vue';
import MatchCard from '@/components/MatchCard.vue'
import { matchStore } from '@/stores/matchStore';

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


const { matches } = storeToRefs(matchesStore)

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

const ChoiceMatchType = (viewType: ViewType): typeof matchesStore.matches => {
  switch (viewType) {
    case "current":
      return matchesStore.currentMatches
    case "upcoming":
      return matchesStore.upcomingMatches
    case null:
      console.log("Error null value")
      return [];
    default:
      console.log(`We are out of ${viewType}`)
      return [];
  }
}

//Sur le Home, mettre que des matchs qui ont des rank S, A, pas plus. Des matches avec de l'importance

const startAutoRefresh = () => {
  if (autoRefreshInterval) {
    clearInterval(autoRefreshInterval)
  }
  
  autoRefreshInterval = setInterval(async () => {
    if (!matchesStore.loading) {
      if (view.value !=  null) {
        await matchesStore.fetchMatches(Url, viewMapper[view.value])
      }
    }
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
  if (view.value !=  null) {
    await matchesStore.fetchMatches(Url, view.value)
  }
  
  // Set up auto-refresh every 30 seconds for live updates
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
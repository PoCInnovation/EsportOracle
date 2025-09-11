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
import { computed, onMounted, onUnmounted, ref, watch } from 'vue';

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
  ["s"].includes(match.tournament?.tier ?? "")
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
  console.log(`LE URL == ${Url}`)
  
  // Set up auto-refresh every 30 seconds for live updates
  startAutoRefresh()
})

onUnmounted(() => {
  stopAutoRefresh()
})

watch((view), async (newView) => {
  console.log('[watch] fetching matches for', newView)
  if (newView === "current" || newView === "upcoming") {
    Url = matchesStore.createUrlMatches(newView, "")
    console.log(`Voir le nouveau View == ${newView}`)
    stopAutoRefresh()

    await matchesStore.fetchMatches(Url, newView)
    startAutoRefresh()
  } else {
    console.log("/home => ", newView)
  }
})

</script>

<style lang="css" scoped>

@import "../../components/matches.css";

/** NavBar */

.nav-desktop {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
  padding: 1.5rem 1.5rem;
  background: rgba(26, 26, 26, 0.85);
  border: 1px solid rgba(249, 115, 22, 0.3);
  border-radius: 2rem;
  backdrop-filter: blur(25px);
  max-width: 1000px;
  margin: 0 auto 1.5rem;
  box-shadow: 
    0 8px 32px rgba(249, 115, 22, 0.15),
    0 0 0 1px rgba(255, 255, 255, 0.05),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
  position: relative;
  overflow: hidden;
  transition: all 0.4s ease;
  z-index: 10;
}

.nav-desktop:hover {
  transform: translateY(-2px);
  box-shadow: 
    0 12px 40px rgba(249, 115, 22, 0.2),
    0 0 0 1px rgba(249, 115, 22, 0.2),
    inset 0 1px 0 rgba(255, 255, 255, 0.15);
}

.nav-desktop::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(249, 115, 22, 0.08) 0%, transparent 50%, rgba(251, 146, 60, 0.08) 100%);
  pointer-events: none;
}

.nav-desktop::after {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: conic-gradient(from 0deg, transparent, rgba(249, 115, 22, 0.1), transparent);
  animation: rotate-slow 20s linear infinite;
  pointer-events: none;
  opacity: 0.5;
}

.nav-desktop a {
  text-decoration: none;
  color: inherit;
}

.page-title {
  display: flex;
  align-items: center;
  gap: 1rem;
  font-size: 1.5rem;
  font-weight: 800;
  margin: 0;
  background: linear-gradient(135deg, #f97316, #fb923c, #fbbf24);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  position: relative;
  z-index: 2;
  animation: glow-text 3s ease-in-out infinite alternate;
}

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
<template>
    <div class="matches-container">

    <div class="matches-header">
      <h1 class="page-title">
        <i class="pi pi-calendar-plus"></i>
        Upcoming Matches
      </h1>
      <button @click="refreshMatches" class="refresh-button" :disabled="MatchesStore.loading">
        <i class="pi pi-refresh" :class="{ 'pi-spin': MatchesStore.loading }"></i>
        {{ MatchesStore.loading ? 'Loading...' : 'Refresh' }}
      </button>
      <form @submit="retrieveTeams" class="refresh-button">
    <div>
      <MultiSelect
        v-model="selectedTeams"
        :options="MatchesStore.AcronymIdTeams"
        optionLabel="name" 
        optionValue="id"
        filter 
        placeholder="Select Teams" 
        :maxSelectedLabels="0"
        class="custom-multiselect" 
      >Teams</MultiSelect>
    </div>
    <Button type="submit" severity="secondary" label="Submit" class="refresh-button">
      Submit
    </Button>
  </form>
    </div>

    <div v-if="MatchesStore.loading && MatchesStore.upcomingMatches.length === 0" class="loading-container">
      <div class="loading-spinner"></div>
      <p class="loading-text">Fetching upcoming matches...</p>
    </div>

    <!-- Error state -->
    <div v-else-if="MatchesStore.error" class="error-container">
      <i class="pi pi-exclamation-triangle"></i>
      <h3>Loading error</h3>
      <p>{{ MatchesStore.error }}</p>
      <button @click="refreshMatches" class="retry-button">
        <i class="pi pi-refresh"></i>
        Retry
      </button>
    </div>

    <!-- Empty state -->
    <div v-else-if="MatchesStore.upcomingMatches.length === 0 && !MatchesStore.loading" class="empty-container">
      <i class="pi pi-calendar-times"></i>
      <h3>No upcoming matches</h3>
      <p>There are currently no scheduled matches. Come back later!</p>
      <button @click="refreshMatches" class="refresh-button">
        <i class="pi pi-refresh"></i>
        Refresh
      </button>
    </div>

    <!-- Matches list -->
    <div v-else class="matches-grid">
      <MatchCard 
        v-for="match in MatchesStore.upcomingMatches" 
        :key="match.id" 
        :match="match"
        :ref="el => matchCardRefs[match.id] = el"
        class="match-item"
      />
    </div>

    <div v-if="MatchesStore.matches.length > 0" class="matches-footer">
      <div class="stats">
        <span class="stat-item">
          <i class="pi pi-chart-bar"></i>
          {{ MatchesStore.upcomingMatches.length }} match{{ MatchesStore.upcomingMatches.length > 1 ? 'es' : '' }} trouvé{{ MatchesStore.upcomingMatches.length > 1 ? 's' : '' }}
        </span>
        <span class="stat-item">
          <i class="pi pi-clock"></i>
          Dernière mise à jour : {{ MatchesStore.lastUpdated }}
        </span>
      </div>
    </div>
    </div>
</template>


<script setup lang="ts">
import { ref, onMounted, watch, onUnmounted, nextTick } from 'vue'
import MatchCard from '@/components/MatchCard.vue'
import { matchStore } from '@/stores/matchStore'
import Button from 'primevue/button'
import MultiSelect from 'primevue/multiselect'
import { useRoute, useRouter } from 'vue-router'
import 'primeicons/primeicons.css'

const route = useRoute()
const router = useRouter()
const teamId = ref(route.query.teamId)
const MatchesStore = matchStore()
const matchCardRefs = ref<Record<string, any>>({})

const valueTeamId: string = teamId.value as string
let Url = MatchesStore.createUrlMatches("upcoming", valueTeamId)
let autoRefreshInterval: NodeJS.Timeout | null = null;
const selectedTeams = ref<number[]>([])

let isUpdatingUrl = false

const retrieveTeams = async (event: Event) => {
  event.preventDefault()
  
  const selectedTeamsObjects = MatchesStore.AcronymIdTeams.filter(
    team => selectedTeams.value.includes(team.id)
  )
  console.log('Objet complete:', selectedTeamsObjects)
  
  isUpdatingUrl = true
  
  const newQuery = { ...route.query }
  
  if (selectedTeams.value.length > 0) {
    newQuery.teamId = selectedTeams.value.join(',')
  } else {
    delete newQuery.teamId
  }
  
  await router.replace({ query: newQuery })
  
  const newTeamId = newQuery.teamId as string
  const newUrl = MatchesStore.createUrlMatches("upcoming", newTeamId)
  
  stopAutoRefresh()
  await MatchesStore.fetchMatches(newUrl, "upcoming")
  MatchesStore.retrieveIdAndNamesTeams(MatchesStore.upcomingMatches)
  
  Url = newUrl
  startAutoRefresh()
  
  isUpdatingUrl = false
}

const startAutoRefresh = () => {
  if (autoRefreshInterval) {
    clearInterval(autoRefreshInterval)
  }
  
  autoRefreshInterval = setInterval(async () => {
    if (!MatchesStore.loading && MatchesStore.currentMatchType === "upcoming") {
      await MatchesStore.fetchMatches(Url, "upcoming")
    }
  }, 30000)
}

const stopAutoRefresh = () => {
  if (autoRefreshInterval) {
    clearInterval(autoRefreshInterval)
    autoRefreshInterval = null
  }
}

const refreshMatches = async (): Promise<void> => {
  await MatchesStore.fetchMatches(Url, "upcoming")
  MatchesStore.retrieveIdAndNamesTeams(MatchesStore.upcomingMatches)
}

// Fonction pour détecter et ouvrir la popup depuis le hash
const checkAndOpenMatchFromHash = async () => {
  const hash = window.location.hash
  if (hash.startsWith('#match-')) {
    const matchId = hash.replace('#match-', '')
    
    // Attendre que les matches soient chargés et les refs créées
    await nextTick()
    
    // Chercher le match correspondant
    const matchCard = matchCardRefs.value[matchId]
    if (matchCard && matchCard.openPopupFromHash) {
      matchCard.openPopupFromHash()
    }
  }
}

// Écouter les changements de hash
const handleHashChange = () => {
  checkAndOpenMatchFromHash()
}

onMounted(async () => {
  // Ajouter l'écouteur de changement de hash
  window.addEventListener('hashchange', handleHashChange)
})

/**
 * Component lifecycle - fetch matches on mount
 */
onMounted(async () => {
  await MatchesStore.fetchMatches(Url, "upcoming")
  MatchesStore.retrieveIdAndNamesTeams(MatchesStore.upcomingMatches)
  
  // Vérifier s'il y a un hash à l'ouverture de la page
  await checkAndOpenMatchFromHash()
  
  // Set up auto-refresh every 30 seconds for live updates
  startAutoRefresh()
})

onUnmounted(() => {
  stopAutoRefresh()
  // Supprimer l'écouteur de hash
  window.removeEventListener('hashchange', handleHashChange)
})
watch(() => route.params.teamId, async (newTeamId) => {
  teamId.value = newTeamId
  const newValueTeamId: string = newTeamId as string
  Url = MatchesStore.createUrlMatches("upcoming", newValueTeamId)
  
  stopAutoRefresh()
  
  // Charger les nouveaux matches
  await MatchesStore.fetchMatches(Url, "upcoming")
  MatchesStore.retrieveIdAndNamesTeams(MatchesStore.upcomingMatches)
  
  startAutoRefresh()
})

</script>

<style scoped>

@import "../../components/matches.css";
</style>

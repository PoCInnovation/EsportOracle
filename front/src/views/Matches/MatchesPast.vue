<template>
    <div class="matches-container">

    <div class="matches-header">
      <h1 class="page-title">
        <i class="pi pi-history"></i>
        Past Matches
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

    <div v-if="MatchesStore.loading && MatchesStore.pastMatches.length === 0" class="loading-container">
      <div class="loading-spinner"></div>
      <p class="loading-text">Fetching past matches...</p>
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
    <div v-else-if="MatchesStore.pastMatches.length === 0 && !MatchesStore.loading" class="empty-container">
      <i class="pi pi-calendar-times"></i>
      <h3>No past matches</h3>
      <p>There are currently no finished matches in the history.</p>
      <button @click="refreshMatches" class="refresh-button">
        <i class="pi pi-refresh"></i>
        Refresh
      </button>
    </div>

    <!-- Matches list -->
    <div v-else class="matches-grid">
      <MatchCard 
        v-for="match in MatchesStore.pastMatches" 
        :key="match.id" 
        :match="match"
        class="match-item"
      />
    </div>

    <div v-if="MatchesStore.pastMatches.length > 0" class="matches-footer">
      <div class="stats">
        <span class="stat-item">
          <i class="pi pi-chart-bar"></i>
          {{ MatchesStore.pastMatches.length }} match{{ MatchesStore.pastMatches.length > 1 ? 'es' : '' }} trouvé{{ MatchesStore.pastMatches.length > 1 ? 's' : '' }}
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

import { ref, onMounted, watch, onUnmounted } from 'vue'
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

const valueTeamId: string = teamId.value as string
let Url = MatchesStore.createUrlMatches("past", valueTeamId)
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
  const newUrl = MatchesStore.createUrlMatches("past", newTeamId)
  
  stopAutoRefresh()
  await MatchesStore.fetchMatches(newUrl, "past")
  MatchesStore.retrieveIdAndNamesTeams(MatchesStore.pastMatches)
  
  Url = newUrl
  startAutoRefresh()
  
  isUpdatingUrl = false
}

const startAutoRefresh = () => {
  if (autoRefreshInterval) {
    clearInterval(autoRefreshInterval)
  }
  
  autoRefreshInterval = setInterval(async () => {
    if (!MatchesStore.loading) {
      await MatchesStore.fetchMatches(Url, "past")
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
  await MatchesStore.fetchMatches(Url, "past")
  MatchesStore.retrieveIdAndNamesTeams(MatchesStore.pastMatches)
}

/**
 * Component lifecycle - fetch matches on mount
 */
onMounted(async () => {
  await MatchesStore.fetchMatches(Url, "past")
  MatchesStore.retrieveIdAndNamesTeams(MatchesStore.pastMatches)
  
  // Set up auto-refresh every 30 seconds for live updates
  startAutoRefresh()
})

onUnmounted(() => {
  stopAutoRefresh()
})

watch(() => route.params.teamId, async (newTeamId) => {
  teamId.value = newTeamId
  const newValueTeamId: string = newTeamId as string
  Url = MatchesStore.createUrlMatches("past", newValueTeamId)
  
  stopAutoRefresh()
  
  // Loading new matches
  await MatchesStore.fetchMatches(Url, "past")
  MatchesStore.retrieveIdAndNamesTeams(MatchesStore.pastMatches)
  
  startAutoRefresh()
})

</script>

<style scoped>

@import "../../components/matches.css";
</style>

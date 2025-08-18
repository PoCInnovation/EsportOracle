<template>
    <div class="matches-container">
    <div class="matches-header">
      <h1 class="page-title">
        <i class="pi pi-play-circle"></i>
        Matches à venir
      </h1>
      <button @click="refreshMatches" class="refresh-button" :disabled="MatchesStore.loading">
        <i class="pi pi-refresh" :class="{ 'pi-spin': MatchesStore.loading }"></i>
        {{ MatchesStore.loading ? 'Chargement...' : 'Actualiser' }}
      </button>
    </div>

    <div v-if="MatchesStore.loading && MatchesStore.matches.length === 0" class="loading-container">
      <div class="loading-spinner"></div>
      <p class="loading-text">Récupération des matches en cours...</p>
    </div>

    <!-- Error state -->
    <div v-else-if="MatchesStore.error" class="error-container">
      <i class="pi pi-exclamation-triangle"></i>
      <h3>Erreur lors du chargement</h3>
      <p>{{ MatchesStore.error }}</p>
      <button @click="refreshMatches" class="retry-button">
        <i class="pi pi-refresh"></i>
        Réessayer
      </button>
    </div>

    <!-- Empty state -->
    <div v-else-if="MatchesStore.matches.length === 0 && !MatchesStore.loading" class="empty-container">
      <i class="pi pi-calendar-times"></i>
      <h3>Aucun match en cours</h3>
      <p>Il n'y a actuellement aucun match en direct. Revenez plus tard !</p>
      <button @click="refreshMatches" class="refresh-button">
        <i class="pi pi-refresh"></i>
        Actualiser
      </button>
    </div>

    <!-- Matches list -->
    <div v-else class="matches-grid">
      <MatchCard 
        v-for="match in MatchesStore.matches" 
        :key="match.id" 
        :match="match"
        class="match-item"
      />
    </div>

    <div v-if="MatchesStore.matches.length > 0" class="matches-footer">
      <div class="stats">
        <span class="stat-item">
          <i class="pi pi-chart-bar"></i>
          {{ MatchesStore.matches.length }} match{{ MatchesStore.matches.length > 1 ? 'es' : '' }} trouvé{{ MatchesStore.matches.length > 1 ? 's' : '' }}
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
import { matchStore } from '@/stores/matchStore';
import { useRoute } from 'vue-router'

const route = useRoute()
const teamId = ref(route.params.teamId)
const MatchesStore = matchStore()

const valueTeamId: string = teamId.value as string
let Url = MatchesStore.createUrlMatches("upcoming", valueTeamId)
let autoRefreshInterval: NodeJS.Timeout | null = null;

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
}

/**
 * Component lifecycle - fetch matches on mount
 */
onMounted(async () => {
  await MatchesStore.fetchMatches(Url, "upcoming")
  
  // Set up auto-refresh every 30 seconds for live updates
  startAutoRefresh()
})

onUnmounted(() => {
  stopAutoRefresh()
})
watch(() => route.params.teamId, async (newTeamId) => {
  teamId.value = newTeamId
  const newValueTeamId: string = newTeamId as string
  Url = MatchesStore.createUrlMatches("upcoming", newValueTeamId)
  
  stopAutoRefresh()
  
  // Charger les nouveaux matches
  await MatchesStore.fetchMatches(Url, "upcoming")
  
  startAutoRefresh()
})

</script>

<style scoped>

@import "../../components/matches.css";
</style>

<template>
  <div class="matches-container">
    <div class="matches-header">
      <h1 class="page-title">
        <i class="pi pi-play-circle"></i>
        Matches en Cours
      </h1>
      <button @click="refreshMatches" class="refresh-button" :disabled="loading">
        <i class="pi pi-refresh" :class="{ 'pi-spin': loading }"></i>
        {{ loading ? 'Chargement...' : 'Actualiser' }}
      </button>
    </div>

    <!-- Loading state -->
    <div v-if="loading && matches.length === 0" class="loading-container">
      <div class="loading-spinner"></div>
      <p class="loading-text">Récupération des matches en cours...</p>
    </div>

    <!-- Error state -->
    <div v-else-if="error" class="error-container">
      <i class="pi pi-exclamation-triangle"></i>
      <h3>Erreur lors du chargement</h3>
      <p>{{ error }}</p>
      <button @click="refreshMatches" class="retry-button">
        <i class="pi pi-refresh"></i>
        Réessayer
      </button>
    </div>

    <!-- Empty state -->
    <div v-else-if="matches.length === 0 && !loading" class="empty-container">
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
        v-for="match in matches" 
        :key="match.id" 
        :match="match"
        class="match-item"
      />
    </div>

    <!-- Stats footer -->
    <div v-if="matches.length > 0" class="matches-footer">
      <div class="stats">
        <span class="stat-item">
          <i class="pi pi-chart-bar"></i>
          {{ matches.length }} match{{ matches.length > 1 ? 'es' : '' }} trouvé{{ matches.length > 1 ? 's' : '' }}
        </span>
        <span class="stat-item">
          <i class="pi pi-clock"></i>
          Dernière mise à jour : {{ lastUpdated }}
        </span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import MatchCard from '@/components/MatchCard.vue'

/**
 * Interface for match data structure from backend API
 */
interface Match {
  id: number
  status: string
  begin_at?: string
  scheduled_at?: string
  match_type?: string
  videogame?: {
    name: string
  }
  opponents?: Array<{
    opponent: {
      id: number
      name: string
      acronym: string
      image_url?: string
    }
  }>
  league?: {
    name: string
  }
  tournament?: {
    name: string
  }
}

/**
 * Reactive state management
 */
const matches = ref<Match[]>([])
const loading = ref(false)
const error = ref<string>('')
const lastUpdate = ref<Date | null>(null)

/**
 * Computed property for formatted last update time
 */
const lastUpdated = computed(() => {
  if (!lastUpdate.value) return 'Jamais'
  
  return lastUpdate.value.toLocaleString('fr-FR', {
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit'
  })
})

/**
 * Fetches current matches from the backend API
 * Makes a call to your Go backend endpoint /CurrentMatches
 */
const fetchMatches = async (): Promise<void> => {
  try {
    loading.value = true
    error.value = ''
    
    // Adjust this URL to match your backend server address and port
    const response = await fetch('http://localhost:8080/CurrentMatches', {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      }
    })
    
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    
    const data: Match[] = await response.json()
    
    // Sort matches by status priority (live first, then upcoming, then finished)
    const sortedMatches = data.sort((a, b) => {
      const statusPriority = { 'running': 0, 'live': 0, 'not_started': 1, 'upcoming': 1, 'finished': 2 }
      const aPriority = statusPriority[a.status?.toLowerCase() as keyof typeof statusPriority] ?? 3
      const bPriority = statusPriority[b.status?.toLowerCase() as keyof typeof statusPriority] ?? 3
      
      return aPriority - bPriority
    })
    
    matches.value = sortedMatches
    lastUpdate.value = new Date()
    
    console.log(`Successfully loaded ${data.length} matches`)
    
  } catch (err) {
    console.error('Error fetching matches:', err)
    error.value = err instanceof Error ? err.message : 'Une erreur inconnue est survenue'
  } finally {
    loading.value = false
  }
}

/**
 * Refreshes the matches list
 * Can be called manually by user or automatically
 */
const refreshMatches = async (): Promise<void> => {
  await fetchMatches()
}

/**
 * Component lifecycle - fetch matches on mount
 */
onMounted(async () => {
  await fetchMatches()
  
  // Set up auto-refresh every 30 seconds for live updates
  setInterval(async () => {
    if (!loading.value) {
      await fetchMatches()
    }
  }, 30000)
})
</script>

<style scoped>
.matches-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
  min-height: 100vh;
}

.matches-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  padding-bottom: 1rem;
  border-bottom: 2px solid rgba(255, 255, 255, 0.1);
}

.page-title {
  color: #ffffff;
  font-size: 2.5rem;
  font-weight: 800;
  margin: 0;
  display: flex;
  align-items: center;
  gap: 1rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.page-title i {
  color: #667eea;
  -webkit-text-fill-color: #667eea;
}

.refresh-button {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #ffffff;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  font-weight: 600;
  font-size: 0.875rem;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.refresh-button:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

.refresh-button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}

.loading-container,
.error-container,
.empty-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 4rem 2rem;
  text-align: center;
  color: #9ca3af;
}

.loading-spinner {
  width: 48px;
  height: 48px;
  border: 4px solid rgba(102, 126, 234, 0.2);
  border-left: 4px solid #667eea;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 1rem;
}

.loading-text {
  font-size: 1.125rem;
  font-weight: 500;
  margin: 0;
}

.error-container i,
.empty-container i {
  font-size: 3rem;
  color: #ef4444;
  margin-bottom: 1rem;
}

.empty-container i {
  color: #6b7280;
}

.error-container h3,
.empty-container h3 {
  color: #ffffff;
  font-size: 1.5rem;
  font-weight: 600;
  margin: 0 0 0.5rem 0;
}

.error-container p,
.empty-container p {
  font-size: 1rem;
  margin: 0 0 1.5rem 0;
  max-width: 400px;
}

.retry-button {
  background: #ef4444;
  color: #ffffff;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  font-weight: 600;
  font-size: 0.875rem;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.retry-button:hover {
  background: #dc2626;
  transform: translateY(-1px);
}

.matches-grid {
  display: grid;
  gap: 1.5rem;
  grid-template-columns: repeat(auto-fill, minmax(500px, 1fr));
}

.match-item {
  animation: fadeInUp 0.5s ease-out;
}

.matches-footer {
  margin-top: 3rem;
  padding-top: 2rem;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.stats {
  display: flex;
  justify-content: center;
  gap: 2rem;
  flex-wrap: wrap;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: #9ca3af;
  font-size: 0.875rem;
  font-weight: 500;
}

.stat-item i {
  color: #667eea;
}

/* Responsive design */
@media (max-width: 768px) {
  .matches-container {
    padding: 1rem;
  }
  
  .matches-header {
    flex-direction: column;
    gap: 1rem;
    text-align: center;
  }
  
  .page-title {
    font-size: 2rem;
  }
  
  .matches-grid {
    grid-template-columns: 1fr;
    gap: 1rem;
  }
  
  .stats {
    flex-direction: column;
    gap: 1rem;
  }
}

/* Animations */
@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.pi-spin {
  animation: spin 1s linear infinite;
}
</style>

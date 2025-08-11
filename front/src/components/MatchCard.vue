<template>
  <div class="match-card">
    <div class="match-header">
      <div class="match-status">
        <span class="status-indicator" :class="getStatusClass(match.status)"></span>
        <span class="status-text">{{ getStatusText(match.status) }}</span>
      </div>
      <div class="match-time">
        {{ formatDate(match.begin_at || match.scheduled_at) }}
      </div>
    </div>

    <div class="match-content">
      <div class="team" v-if="match.opponents && match.opponents[0]">
        <img 
          :src="match.opponents[0].opponent.image_url || '/default-team.png'" 
          :alt="match.opponents[0].opponent.name"
          class="team-logo"
          @error="handleImageError"
        >
        <div class="team-info">
          <h3 class="team-name">{{ match.opponents[0].opponent.name }}</h3>
          <p class="team-acronym">{{ match.opponents[0].opponent.acronym }}</p>
        </div>
      </div>

      <div class="vs-section">
        <span class="vs-text">VS</span>
        <div class="match-details">
          <p class="game-type">{{ match.videogame?.name || 'CS:GO' }}</p>
          <p class="match-type">{{ match.match_type || 'Best of 3' }}</p>
        </div>
      </div>

      <div class="team" v-if="match.opponents && match.opponents[1]">
        <img 
          :src="match.opponents[1].opponent.image_url || '/default-team.png'" 
          :alt="match.opponents[1].opponent.name"
          class="team-logo"
          @error="handleImageError"
        >
        <div class="team-info">
          <h3 class="team-name">{{ match.opponents[1].opponent.name }}</h3>
          <p class="team-acronym">{{ match.opponents[1].opponent.acronym }}</p>
        </div>
      </div>
    </div>

    <div class="match-footer" v-if="match.league || match.tournament">
      <div class="tournament-info">
        <i class="pi pi-trophy"></i>
        <span>{{ match.league?.name || match.tournament?.name || 'Tournament' }}</span>
      </div>
      <div class="match-id">
        #{{ match.id }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">

/**
 * Interface for match data structure from PandaScore API
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
 * Component props
 */
const props = defineProps<{
  match: Match
}>()

/**
 * Formats the date string into a readable format
 * @param dateString - ISO date string from API
 * @returns Formatted date and time string
 */
const formatDate = (dateString?: string): string => {
  if (!dateString) return 'Time TBD'
  
  const date = new Date(dateString)
  return date.toLocaleString('fr-FR', {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

/**
 * Returns the appropriate CSS class for match status
 * @param status - Match status from API
 * @returns CSS class name
 */
const getStatusClass = (status: string): string => {
  switch (status?.toLowerCase()) {
    case 'running':
    case 'live':
      return 'status-live'
    case 'finished':
      return 'status-finished'
    case 'not_started':
    case 'upcoming':
      return 'status-upcoming'
    default:
      return 'status-default'
  }
}

/**
 * Returns human-readable status text
 * @param status - Match status from API
 * @returns Formatted status text
 */
const getStatusText = (status: string): string => {
  switch (status?.toLowerCase()) {
    case 'running':
    case 'live':
      return 'EN DIRECT'
    case 'finished':
      return 'TERMINÉ'
    case 'not_started':
    case 'upcoming':
      return 'À VENIR'
    default:
      return status?.toUpperCase() || 'INCONNU'
  }
}

/**
 * Handles image loading errors by setting a default image
 * @param event - Image error event
 */
const handleImageError = (event: Event) => {
  const img = event.target as HTMLImageElement
  img.src = '/default-team.png'
}
</script>

<style scoped>
.match-card {
  background: linear-gradient(135deg, #1a1d29 0%, #2d3748 100%);
  border-radius: 16px;
  padding: 1.5rem;
  margin: 1rem 0;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
  border: 1px solid rgba(255, 255, 255, 0.1);
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.match-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: linear-gradient(90deg, #667eea 0%, #764ba2 100%);
}

.match-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 48px rgba(102, 126, 234, 0.2);
}

.match-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}

.match-status {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.status-indicator {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  animation: pulse 2s infinite;
}

.status-live {
  background-color: #ef4444;
  box-shadow: 0 0 8px rgba(239, 68, 68, 0.5);
}

.status-finished {
  background-color: #6b7280;
}

.status-upcoming {
  background-color: #3b82f6;
}

.status-default {
  background-color: #9ca3af;
}

.status-text {
  color: #ffffff;
  font-weight: 600;
  font-size: 0.75rem;
  letter-spacing: 0.05em;
}

.match-time {
  color: #9ca3af;
  font-size: 0.875rem;
  font-weight: 500;
}

.match-content {
  display: grid;
  grid-template-columns: 1fr auto 1fr;
  gap: 1.5rem;
  align-items: center;
  margin-bottom: 1.5rem;
}

.team {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.team:last-child {
  flex-direction: row-reverse;
  text-align: right;
}

.team-logo {
  width: 48px;
  height: 48px;
  border-radius: 8px;
  object-fit: cover;
  border: 2px solid rgba(255, 255, 255, 0.1);
}

.team-info {
  flex: 1;
}

.team-name {
  color: #ffffff;
  font-size: 1.125rem;
  font-weight: 700;
  margin: 0;
  line-height: 1.2;
}

.team-acronym {
  color: #9ca3af;
  font-size: 0.875rem;
  font-weight: 500;
  margin: 0.25rem 0 0 0;
}

.vs-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
}

.vs-text {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #ffffff;
  padding: 0.5rem 1rem;
  border-radius: 20px;
  font-weight: 700;
  font-size: 0.875rem;
  letter-spacing: 0.1em;
  margin-bottom: 0.5rem;
}

.match-details {
  color: #9ca3af;
  font-size: 0.75rem;
}

.game-type {
  margin: 0;
  font-weight: 600;
}

.match-type {
  margin: 0.25rem 0 0 0;
  font-weight: 400;
}

.match-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 1rem;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.tournament-info {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: #9ca3af;
  font-size: 0.875rem;
  font-weight: 500;
}

.tournament-info i {
  color: #fbbf24;
}

.match-id {
  color: #6b7280;
  font-size: 0.75rem;
  font-weight: 600;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
}

/* Responsive design */
@media (max-width: 768px) {
  .match-card {
    padding: 1rem;
    margin: 0.5rem 0;
  }
  
  .match-content {
    grid-template-columns: 1fr;
    gap: 1rem;
    text-align: center;
  }
  
  .team {
    justify-content: center;
    flex-direction: column;
    text-align: center;
  }
  
  .team:last-child {
    flex-direction: column;
    text-align: center;
  }
  
  .team-logo {
    width: 40px;
    height: 40px;
  }
  
  .team-name {
    font-size: 1rem;
  }
}

/* Animation for live status */
@keyframes pulse {
  0% {
    transform: scale(1);
    opacity: 1;
  }
  50% {
    transform: scale(1.2);
    opacity: 0.8;
  }
  100% {
    transform: scale(1);
    opacity: 1;
  }
}
</style>

<template>
  <article class="match-card">
    <div class="status-bar">
      <div class="status-badge" :class="getStatusClass(match.status)">
        <span class="status-icon"></span>
        {{ getStatusText(match.status) }}
      </div>
      <div class="match-time">
        {{ formatDate(match.begin_at || match.scheduled_at) }}
      </div>
    </div>

    <div class="teams-section">
      <div class="team" v-if="match.opponents && match.opponents[0]">
        <div class="team-logo-container">
          <img
            v-if="getTeamImageUrl(match.opponents[0])"
            :src="getTeamImageUrl(match.opponents[0])" 
            :alt="match.opponents[0].opponent.name"
            class="team-logo"
            @error="handleImageError"
            @load="handleImageLoad"
          >
          <div v-else class="team-fallback">
            {{ getTeamInitials(match.opponents[0].opponent.name) }}
          </div>
        </div>
        <div class="team-info">
          <h3 class="team-name">{{ match.opponents[0].opponent.name }}</h3>
          <span class="team-tag">{{ match.opponents[0].opponent.acronym }}</span>
        </div>
      </div>

      <div class="vs-divider">VS</div>

      <div class="team" v-if="match.opponents && match.opponents[1]">
        <div class="team-logo-container">
          <img 
            v-if="getTeamImageUrl(match.opponents[1])"
            :src="getTeamImageUrl(match.opponents[1])" 
            :alt="match.opponents[1].opponent.name"
            class="team-logo"
            @error="handleImageError"
            @load="handleImageLoad"
          >
          <div v-else class="team-fallback">
            {{ getTeamInitials(match.opponents[1].opponent.name) }}
          </div>
        </div>
        <div class="team-info">
          <h3 class="team-name">{{ match.opponents[1].opponent.name }}</h3>
          <span class="team-tag">{{ match.opponents[1].opponent.acronym }}</span>
        </div>
      </div>
    </div>

    <div class="match-bet" v-if="match.status === 'running' && match && match.opponents ">
        <button @click="PushMatchBets(match.id)" class="match-buttonBet">{{ match.opponents[0].opponent.name }}</button>
        <button @click="PushMatchBets(match.id)" class="match-buttonBet">{{ match.opponents[1].opponent.name }}</button>
    </div>
    <div v-else-if="match.status === 'not_started'">
      <button class="match-buttonBet-upcoming">Ouverture le {{ formatDate(match.begin_at || match.scheduled_at) }}</button>
    </div>

    <div class="tournament-section" v-if="match.league || match.tournament">
      <div class="tournament-info">
        <span class="tournament-icon">🏆</span>
        <span class="tournament-name">{{ match.league?.name || match.tournament?.name }}</span>
      </div>
      <span class="match-id">#{{ match.id }}</span>
    </div>
  </article>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { matchStore } from '@/stores/matchStore'

const router = useRouter()
const matchesStore = matchStore()

interface Match {
  id: number
  status: string
  begin_at?: string
  scheduled_at?: string
  opponents?: Array<{
    opponent: {
      id: number
      name: string
      acronym: string
      image_url?: string | null
    }
  }>
  league?: { name: string }
  tournament?: { name: string }
}

const PushMatchBets = (matchId: number) => {
  return router.push(`/bets/current/${matchId}`)
}

const props = defineProps<{ match: Match }>()
const failedImages = ref<Set<string>>(new Set())

const getTeamImageUrl = (opponent: any): string | null => {
  if (!opponent?.opponent?.image_url) return null
  const url = opponent.opponent.image_url.trim()
  if (!url || url === 'null' || failedImages.value.has(url)) return null
  try { new URL(url); return url } catch { return null }
}

const getTeamInitials = (teamName: string): string => {
  if (!teamName) return '?'
  return teamName.split(' ').map(word => word.charAt(0).toUpperCase()).join('').substring(0, 3)
}

const formatDate = (dateString?: string): string => {
  if (!dateString) return 'Time TBD'
  const date = new Date(dateString)
  return date.toLocaleString('fr-FR', {
    day: '2-digit', month: '2-digit', year: 'numeric', hour: '2-digit', minute: '2-digit'
  })
}

const getStatusClass = (status: string): string => {
  switch (status?.toLowerCase()) {
    case 'running': case 'live': return 'status-live'
    case 'finished': return 'status-finished'
    case 'not_started': case 'upcoming': return 'status-upcoming'
    default: return 'status-default'
  }
}

const getStatusText = (status: string): string => {
  switch (status?.toLowerCase()) {
    case 'running': case 'live': return 'LIVE'
    case 'finished': return 'FINISHED'
    case 'not_started': case 'upcoming': return 'UPCOMING'
    default: return status?.toUpperCase() || 'UNKNOWN'
  }
}

const handleImageError = (event: Event) => {
  const img = event.target as HTMLImageElement
  if (img.src) failedImages.value.add(img.src)
  img.src = ''
}

const handleImageLoad = (event: Event) => {
  (event.target as HTMLImageElement).style.opacity = '1'
}
</script>

<style scoped>

@import "../../src/components/team.css";

.match-bet {
  margin-left: calc(48px + 12px);
  margin-bottom: 0.5rem;
  display: flex;
  gap: 10rem;
}

.match-buttonBet {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
  padding: 0.75rem 0.8rem;
  border-radius: 0.5rem;
  font-size: 0.75rem;
  font-weight: 700;
  text-transform: capitalize;
  letter-spacing: 0.3px;
  backdrop-filter: blur(10px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
  background: rgba(249, 115, 22, 0.3);
  border: none;
  outline: none;
  cursor: pointer;

  min-width: 100px;
  max-width: 160px;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.match-buttonBet:focus {
  outline: none;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2), 0 0 0 3px #f97316;
}

.match-buttonBet:hover {
  background:  #f97316;
   transform: scale(1.05);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.3);
  backdrop-filter: blur(15px);
}

.match-buttonBet-upcoming {
  display: flex;
  margin-bottom: 0.5rem;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
  padding: 0.75rem 1rem;
  border-radius: 0.5rem;
  font-size: 0.8rem;
  font-weight: 700;
  text-transform: capitalize;
  letter-spacing: 0.5px;
  backdrop-filter: blur(10px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
  transition: all 0.3s ease;
  background: rgba(107, 114, 128, 0.25);
  border: none;
  outline: none;
  cursor: pointer;
  color: white;
  width: 100%;
}

.match-buttonBet-upcoming:hover {
  background: rgba(107, 114, 128, 0.4);
  transform: translateY(-1px);
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.25);
}

.match-card {
  background: rgba(26, 26, 26, 0.95);
  border: 1px solid rgba(249, 115, 22, 0.3);
  border-radius: 2rem;
  padding: 2rem;
  transition: all 0.5s cubic-bezier(0.4, 0, 0.2, 1);
  overflow: hidden;
  position: relative;
  backdrop-filter: blur(25px);
  box-shadow: 
    0 8px 32px rgba(0, 0, 0, 0.2),
    0 0 0 1px rgba(255, 255, 255, 0.05),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
  position: relative;
  padding-bottom: 3rem;
}

.match-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(249, 115, 22, 0.08) 0%, transparent 50%, rgba(251, 146, 60, 0.08) 100%);
  opacity: 0;
  transition: opacity 0.5s ease;
  pointer-events: none;
}

.match-card::after {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: conic-gradient(from 0deg, transparent, rgba(249, 115, 22, 0.1), transparent);
  animation: rotate-slow 25s linear infinite;
  pointer-events: none;
  opacity: 0;
  transition: opacity 0.5s ease;
}

.match-card:hover {
  transform: translateY(-12px) scale(1.03);
  box-shadow: 
    0 10px 20px rgba(249, 115, 22, 0.1),
    0 0 0 1px rgba(249, 115, 22, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
  border-color: #fb923c64;
}

.match-card:hover::before {
  opacity: 1;
}

.match-card:hover::after {
  opacity: 0.3;
}

.status-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid rgba(255, 255, 255, 0.15);
  position: relative;
  z-index: 2;
}

.status-badge {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.75rem 1.25rem;
  border-radius: 2rem;
  font-size: 0.8rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  backdrop-filter: blur(10px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.status-badge::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.1), transparent);
  transition: left 0.5s ease;
}

.status-badge:hover::before {
  left: 100%;
}

.status-live { 
  background: rgba(239, 68, 68, 0.15); 
  color: #ef4444; 
  border: 1px solid rgba(239, 68, 68, 0.3);
  box-shadow: 0 4px 20px rgba(239, 68, 68, 0.2);
}

.status-upcoming { 
  background: rgba(59, 130, 246, 0.15); 
  color: #60a5fa; 
  border: 1px solid rgba(59, 130, 246, 0.3);
  box-shadow: 0 4px 20px rgba(59, 130, 246, 0.2);
}

.status-finished { 
  background: rgba(107, 114, 128, 0.15); 
  color: #9ca3af; 
  border: 1px solid rgba(107, 114, 128, 0.3);
  box-shadow: 0 4px 20px rgba(107, 114, 128, 0.1);
}

.status-default { 
  background: rgba(156, 163, 175, 0.15); 
  color: #9ca3af; 
  border: 1px solid rgba(156, 163, 175, 0.3);
  box-shadow: 0 4px 20px rgba(156, 163, 175, 0.1);
}

.status-icon {
  width: 0.6rem;
  height: 0.6rem;
  border-radius: 50%;
  background: currentColor;
  box-shadow: 0 0 10px currentColor;
}

.status-live .status-icon { 
  animation: pulse-bright 1.5s infinite;
  box-shadow: 0 0 15px currentColor;
}

.match-time { 
  color: #e5e5e5; 
  font-size: 0.9rem; 
  font-weight: 600;
  position: relative;
  z-index: 2;
}


.vs-divider {
  margin: 0 1.5rem;
  font-size: 1.8rem;
  font-weight: 900;
  color: #f97316;
  background: linear-gradient(135deg, #f97316, #fb923c, #fbbf24);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  text-align: center;
  flex-shrink: 0;
  animation: pulse-vs 3s ease-in-out infinite;
  filter: drop-shadow(0 0 8px rgba(249, 115, 22, 0.5));
  transition: all 0.3s ease;
}

.match-card:hover .vs-divider {
  transform: scale(1.1);
  filter: drop-shadow(0 0 15px rgba(249, 115, 22, 0.8));
}

.team:last-child {
  flex-direction: row-reverse;
  text-align: right;
}

.team:last-child .team-info { 
  text-align: right; 
}

.tournament-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 1rem;
  border-top: 1px solid rgba(255, 255, 255, 0.15);
  position: relative;
  z-index: 2;
}

.tournament-info {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  flex: 1;
  min-width: 0;
}

.tournament-name {
  color: #e5e5e5;
  font-size: 0.9rem;
  font-weight: 600;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.match-id {
  color: #a3a3a3;
  font-size: 0.75rem;
  font-family: 'JetBrains Mono', monospace;
  background: rgba(255, 255, 255, 0.08);
  padding: 0.5rem 0.75rem;
  border-radius: 0.75rem;
  border: 1px solid rgba(255, 255, 255, 0.15);
  backdrop-filter: blur(10px);
  transition: all 0.3s ease;
  font-weight: 600;
}

.match-id:hover {
  background: rgba(249, 115, 22, 0.1);
  border-color: rgba(249, 115, 22, 0.3);
  color: #fb923c;
}

@keyframes pulse-bright { 
  0%, 100% { 
    opacity: 1; 
    box-shadow: 0 0 15px currentColor;
  } 
  50% { 
    opacity: 0.7; 
    box-shadow: 0 0 25px currentColor;
  } 
}

@keyframes pulse-vs {
  0%, 100% {
    transform: scale(1);
    filter: drop-shadow(0 0 8px rgba(249, 115, 22, 0.5));
  }
  50% {
    transform: scale(1.05);
    filter: drop-shadow(0 0 12px rgba(249, 115, 22, 0.7));
  }
}

@keyframes rotate-slow {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

@media (max-width: 768px) {
  .teams-section { 
    flex-direction: column; 
    gap: 1rem; 
  }
  
  .team { 
    width: 100%; 
    justify-content: flex-start; 
  }
  
  .team:last-child { 
    flex-direction: row; 
    text-align: left; 
  }
  
  .team:last-child .team-info { 
    text-align: left; 
  }
  
  .vs-divider { 
    transform: rotate(90deg); 
    font-size: 1.5rem; 
    margin: 0;
  }
  
  .tournament-section { 
    flex-direction: column; 
    gap: 0.75rem; 
    text-align: center; 
  }
  
  .status-bar { 
    flex-direction: column; 
    gap: 0.75rem; 
    text-align: center; 
  }

  .match-card {
    padding: 1.5rem;
  }
}
</style>

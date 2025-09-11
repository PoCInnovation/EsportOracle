<template>
  <div v-if="visible" class="match-details-overlay" @click="closePopup">
    <div class="match-details-popup" @click.stop>
      <!-- Header with close button -->
      <div class="popup-header">
        <h2 class="popup-title">
          <i class="pi pi-info-circle"></i>
          Match Details
        </h2>
        <button @click="closePopup" class="close-button">
          <i class="pi pi-times"></i>
        </button>
      </div>

      <!-- Scrollable content -->
      <div class="popup-content">
        <!-- Match Status and Time -->
      <div class="match-status-section">
        <div class="status-badge-large" :class="getStatusClass(match.status)">
          <span class="status-icon-large"></span>
          {{ getStatusText(match.status) }}
        </div>
        <div class="match-timing">
          <div class="time-info">
            <i class="pi pi-calendar"></i>
            <span>{{ formatDate(match.begin_at || match.scheduled_at) }}</span>
          </div>
          <div v-if="match.status === 'not_started'" class="countdown">
            <i class="pi pi-clock"></i>
            <span>{{ getCountdown(match.begin_at || match.scheduled_at) }}</span>
          </div>
        </div>
      </div>

      <!-- Teams Section -->
      <div class="teams-details-section">
        <div class="team-detail team-left" v-if="match.opponents && match.opponents[0]">
          <div class="team-visual">
            <div class="team-logo-large">
              <img 
                v-if="getTeamImageUrl(match.opponents[0])"
                :src="getTeamImageUrl(match.opponents[0])" 
                :alt="match.opponents[0].opponent.name"
                class="logo-img"
                @error="handleImageError"
              >
              <div v-else class="team-fallback-large">
                {{ getTeamInitials(match.opponents[0].opponent.name) }}
              </div>
            </div>
            <div class="team-details-info">
              <h3 class="team-name-large">{{ match.opponents[0].opponent.name }}</h3>
              <p class="team-tag-large">{{ match.opponents[0].opponent.acronym }}</p>
            </div>
          </div>
          <div class="team-stats" v-if="match.opponents[0].opponent.stats">
            <div class="stat-item">
              <span class="stat-label">Wins:</span>
              <span class="stat-value">{{ match.opponents[0].opponent.stats?.wins || '-' }}</span>
            </div>
            <div class="stat-item">
              <span class="stat-label">Losses:</span>
              <span class="stat-value">{{ match.opponents[0].opponent.stats?.losses || '-' }}</span>
            </div>
          </div>
        </div>

        <div class="vs-divider-large">
          <span class="vs-text-large">VS</span>
          <div class="match-format" v-if="match.match_type">
            {{ match.match_type || 'Best of 3' }}
          </div>
        </div>

        <div class="team-detail team-right" v-if="match.opponents && match.opponents[1]">
          <div class="team-visual">
            <div class="team-logo-large">
              <img 
                v-if="getTeamImageUrl(match.opponents[1])"
                :src="getTeamImageUrl(match.opponents[1])" 
                :alt="match.opponents[1].opponent.name"
                class="logo-img"
                @error="handleImageError"
              >
              <div v-else class="team-fallback-large">
                {{ getTeamInitials(match.opponents[1].opponent.name) }}
              </div>
            </div>
            <div class="team-details-info">
              <h3 class="team-name-large">{{ match.opponents[1].opponent.name }}</h3>
              <p class="team-tag-large">{{ match.opponents[1].opponent.acronym }}</p>
            </div>
          </div>
          <div class="team-stats" v-if="match.opponents[1].opponent.stats">
            <div class="stat-item">
              <span class="stat-label">Wins:</span>
              <span class="stat-value">{{ match.opponents[1].opponent.stats?.wins || '-' }}</span>
            </div>
            <div class="stat-item">
              <span class="stat-label">Losses:</span>
              <span class="stat-value">{{ match.opponents[1].opponent.stats?.losses || '-' }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Tournament and Match Details -->
      <div class="match-info-section">
        <div class="info-grid">
          <div class="info-item">
            <div class="info-header">
              <i class="pi pi-trophy"></i>
              <span>Tournament</span>
            </div>
            <p class="info-content">{{ match.league?.name || match.tournament?.name || 'Tournament TBD' }}</p>
          </div>
          
          <div class="info-item">
            <div class="info-header">
              <i class="pi pi-gamepad-2"></i>
              <span>Game</span>
            </div>
            <p class="info-content">{{ match.videogame?.name || 'CS:GO' }}</p>
          </div>

          <div class="info-item">
            <div class="info-header">
              <i class="pi pi-list"></i>
              <span>Format</span>
            </div>
            <p class="info-content">{{ match.match_type || 'Best of 3' }}</p>
          </div>

          <div class="info-item">
            <div class="info-header">
              <i class="pi pi-hashtag"></i>
              <span>Match ID</span>
            </div>
            <p class="info-content">#{{ match.id }}</p>
          </div>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="action-buttons">
        <button 
          v-if="canShowBetting" 
          @click="PushMatchBets(match.id)"
          class="betting-button"
        >
          <i class="pi pi-money-bill"></i>
          Bet on this match
        </button>
        
        <button @click="shareMatch" class="share-button">
          <i class="pi pi-share-alt"></i>
          Share
        </button>
        
        <button 
          v-if="match.status === 'not_started'"
          @click="addToCalendar" 
          class="calendar-button"
        >
          <i class="pi pi-calendar-plus"></i>
          Add to calendar
        </button>
      </div>

      <!-- Additional Info Section -->
      <div v-if="match.streams || match.official_stream_url" class="streams-section">
        <h4 class="section-title">
          <i class="pi pi-video"></i>
          Streams
        </h4>
        <div class="stream-links">
          <a 
            v-if="match.official_stream_url" 
            :href="match.official_stream_url" 
            target="_blank" 
            class="stream-link official"
          >
            <i class="pi pi-play"></i>
            Official Stream
          </a>
          <div v-if="match.streams" class="other-streams">
            <a 
              v-for="stream in match.streams" 
              :key="stream.id"
              :href="stream.raw_url" 
              target="_blank" 
              class="stream-link"
            >
              <i class="pi pi-video"></i>
              {{ stream.language }}
            </a>
          </div>
        </div>
      </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

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
      image_url?: string | null
      stats?: {
        wins?: number
        losses?: number
      }
    }
  }>
  league?: { name: string }
  tournament?: { name: string }
  streams?: Array<{
    id: number
    language: string
    raw_url: string
  }>
  official_stream_url?: string
}

const PushMatchBets = (matchId: number) => {
  return router.push(`/bets/upcoming/${matchId}`)
}

const props = defineProps<{
  visible: boolean
  match: Match
}>()

const emit = defineEmits<{
  close: []
  openBetting: []
}>()

const failedImages = ref<Set<string>>(new Set())

watch(() => props.visible, (newVisible) => {
  if (newVisible) {
    const scrollY = window.scrollY
    const scrollX = window.scrollX
    
    document.body.style.overflow = 'hidden'
    document.documentElement.style.overflow = 'hidden'
    document.body.style.position = 'fixed'
    document.body.style.top = `-${scrollY}px`
    document.body.style.left = `-${scrollX}px`
    document.body.style.width = '100%'
    
    document.body.setAttribute('data-scroll-y', scrollY.toString())
    document.body.setAttribute('data-scroll-x', scrollX.toString())
  } else {
    const scrollY = parseInt(document.body.getAttribute('data-scroll-y') || '0')
    const scrollX = parseInt(document.body.getAttribute('data-scroll-x') || '0')
    
    document.body.style.overflow = ''
    document.documentElement.style.overflow = ''
    document.body.style.position = ''
    document.body.style.top = ''
    document.body.style.left = ''
    document.body.style.width = ''
    
    window.scrollTo(scrollX, scrollY)
    
    document.body.removeAttribute('data-scroll-y')
    document.body.removeAttribute('data-scroll-x')
  }
})

onUnmounted(() => {
  if (props.visible) {
    const scrollY = parseInt(document.body.getAttribute('data-scroll-y') || '0')
    const scrollX = parseInt(document.body.getAttribute('data-scroll-x') || '0')
    
    document.body.style.overflow = ''
    document.documentElement.style.overflow = ''
    document.body.style.position = ''
    document.body.style.top = ''
    document.body.style.left = ''
    document.body.style.width = ''
    
    window.scrollTo(scrollX, scrollY)
    
    document.body.removeAttribute('data-scroll-y')
    document.body.removeAttribute('data-scroll-x')
  }
})

const canShowBetting = computed(() => {
  return props.match.status === 'not_started' || props.match.status === 'upcoming'
})

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
    weekday: 'long',
    day: '2-digit',
    month: 'long',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const getCountdown = (dateString?: string): string => {
  if (!dateString) return 'Time TBD'
  
  const now = new Date()
  const matchDate = new Date(dateString)
  const diff = matchDate.getTime() - now.getTime()
  
  if (diff <= 0) return 'Match started'
  
  const days = Math.floor(diff / (1000 * 60 * 60 * 24))
  const hours = Math.floor((diff % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60))
  const minutes = Math.floor((diff % (1000 * 60 * 60)) / (1000 * 60))
  
  if (days > 0) return `In ${days}d ${hours}h`
  if (hours > 0) return `In ${hours}h ${minutes}min`
  return `In ${minutes}min`
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

const closePopup = () => {
  emit('close')
}

const openBetting = () => {
  emit('openBetting')
}

const shareMatch = async () => {
  const baseUrl = window.location.origin + window.location.pathname
  const shareUrl = `${baseUrl}#match-${props.match.id}`
  
  const matchTitle = `Match: ${props.match.opponents?.[0]?.opponent.name} vs ${props.match.opponents?.[1]?.opponent.name}`
  const matchDescription = `Watch this match ${props.match.league?.name || props.match.tournament?.name || ''}`
  
  if (navigator.share) {
    try {
      await navigator.share({
        title: matchTitle,
        text: matchDescription,
        url: shareUrl
      })
    } catch (err) {
      console.log('Share cancelled')
      await copyToClipboard(shareUrl)
    }
  } else {
    await copyToClipboard(shareUrl)
  }
}

const copyToClipboard = async (url: string) => {
  try {
    await navigator.clipboard.writeText(url)
    console.log('Link copied to clipboard:', url)
  } catch (err) {
    console.error('Error copying link:', err)
    const textArea = document.createElement('textarea')
    textArea.value = url
    document.body.appendChild(textArea)
    textArea.select()
    document.execCommand('copy')
    document.body.removeChild(textArea)
  }
}

const addToCalendar = () => {
  if (!props.match.begin_at && !props.match.scheduled_at) return
  
  const startDate = new Date(props.match.begin_at || props.match.scheduled_at!)
  const endDate = new Date(startDate.getTime() + 2 * 60 * 60 * 1000) // +2 hours
  
  const title = `${props.match.opponents?.[0]?.opponent.name} vs ${props.match.opponents?.[1]?.opponent.name}`
  const details = `Match ${props.match.league?.name || props.match.tournament?.name || ''}`
  
  const googleCalendarUrl = `https://calendar.google.com/calendar/render?action=TEMPLATE&text=${encodeURIComponent(title)}&dates=${startDate.toISOString().replace(/[-:]/g, '').split('.')[0]}Z/${endDate.toISOString().replace(/[-:]/g, '').split('.')[0]}Z&details=${encodeURIComponent(details)}`
  
  window.open(googleCalendarUrl, '_blank')
}
</script>

<style scoped>
.match-details-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(circle at center, rgba(26, 15, 8, 0.95) 0%, rgba(10, 10, 10, 0.98) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
  backdrop-filter: blur(20px);
  animation: overlayFadeIn 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  overflow: hidden;
}

.match-details-popup {
  background: linear-gradient(135deg, 
    rgba(26, 26, 26, 0.95) 0%, 
    rgba(26, 15, 8, 0.97) 25%,
    rgba(26, 26, 26, 0.95) 50%,
    rgba(26, 15, 8, 0.97) 75%,
    rgba(26, 26, 26, 0.95) 100%
  );
  border-radius: 2rem;
  width: 90%;
  max-width: 900px;
  height: 65vh;
  max-height: 65vh;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  box-shadow: 
    0 30px 80px rgba(0, 0, 0, 0.7),
    0 0 100px rgba(249, 115, 22, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(249, 115, 22, 0.2);
  backdrop-filter: blur(25px);
  animation: popupSlideIn 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
  position: relative;
}

.popup-content {
  flex: 1;
  overflow-y: auto;
  overflow-x: hidden;
}

.popup-content::-webkit-scrollbar {
  width: 8px;
}

.popup-content::-webkit-scrollbar-track {
  background: rgba(26, 15, 8, 0.3);
  border-radius: 4px;
}

.popup-content::-webkit-scrollbar-thumb {
  background: linear-gradient(180deg, #f97316, #fb923c);
  border-radius: 4px;
  transition: all 0.3s ease;
}

.popup-content::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(180deg, #fb923c, #fbbf24);
  box-shadow: 0 0 10px rgba(249, 115, 22, 0.5);
}

/* Pour Firefox */
.popup-content {
  scrollbar-width: thin;
  scrollbar-color: #f97316 rgba(26, 15, 8, 0.3);
}


.match-details-popup::before {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: conic-gradient(
    from 0deg at 50% 50%,
    transparent 0deg,
    rgba(249, 115, 22, 0.03) 60deg,
    transparent 120deg,
    rgba(251, 146, 60, 0.03) 180deg,
    transparent 240deg,
    rgba(249, 115, 22, 0.03) 300deg,
    transparent 360deg
  );
  animation: rotate-slow 25s linear infinite;
  pointer-events: none;
}

.popup-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 2rem;
  border-bottom: 1px solid rgba(249, 115, 22, 0.2);
  background: rgba(26, 15, 8, 0.3);
  position: relative;
  z-index: 2;
  flex-shrink: 0;
}

.popup-header::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  width: 60%;
  height: 1px;
  background: linear-gradient(90deg, transparent, #f97316, transparent);
  transform: translateX(-50%);
}

.popup-title {
  color: #ffffff;
  font-size: 1.75rem;
  font-weight: 700;
  margin: 0;
  display: flex;
  align-items: center;
  gap: 1rem;
}

.popup-title i {
  color: #fb923c;
  filter: drop-shadow(0 0 8px rgba(249, 115, 22, 0.5));
  animation: iconGlow 3s ease-in-out infinite;
}

.close-button {
  background: rgba(249, 115, 22, 0.1);
  border: 1px solid rgba(249, 115, 22, 0.2);
  color: #fb923c;
  font-size: 1.5rem;
  cursor: pointer;
  padding: 0.75rem;
  border-radius: 12px;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  overflow: hidden;
}

.close-button::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(249, 115, 22, 0.3), transparent);
  transition: left 0.6s ease;
}

.close-button:hover {
  background: rgba(249, 115, 22, 0.2);
  border-color: #f97316;
  color: #ffffff;
  transform: scale(1.1) rotate(90deg);
  box-shadow: 0 4px 20px rgba(249, 115, 22, 0.4);
}

.close-button:hover::before {
  left: 100%;
}

.match-status-section {
  padding: 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: linear-gradient(135deg, rgba(26, 15, 8, 0.2), rgba(26, 26, 26, 0.1));
  position: relative;
  z-index: 2;
}

.match-status-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(ellipse at center, rgba(249, 115, 22, 0.05) 0%, transparent 70%);
  pointer-events: none;
}

.status-badge-large {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1rem 2rem;
  border-radius: 50px;
  font-size: 1rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 1px;
  backdrop-filter: blur(15px);
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.3);
}

.status-icon-large {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background: currentColor;
  box-shadow: 0 0 20px currentColor;
}

.status-live { 
  background: linear-gradient(135deg, rgba(249, 115, 22, 0.3), rgba(234, 88, 12, 0.2)); 
  color: #fb923c; 
  border: 1px solid rgba(249, 115, 22, 0.5);
  box-shadow: 0 0 30px rgba(249, 115, 22, 0.3);
}

.status-upcoming { 
  background: linear-gradient(135deg, rgba(251, 146, 60, 0.3), rgba(249, 115, 22, 0.2)); 
  color: #fbbf24; 
  border: 1px solid rgba(251, 146, 60, 0.5);
  box-shadow: 0 0 25px rgba(251, 146, 60, 0.2);
}

.status-finished { 
  background: linear-gradient(135deg, rgba(156, 163, 175, 0.3), rgba(107, 114, 128, 0.2)); 
  color: #d1d5db; 
  border: 1px solid rgba(156, 163, 175, 0.4);
}

.status-live .status-icon-large {
  animation: pulse-bright 1.5s infinite, glow-pulse 2s ease-in-out infinite;
  box-shadow: 0 0 20px currentColor;
}

.match-timing {
  text-align: right;
  color: #e5e5e5;
}

.time-info, .countdown {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
  font-weight: 600;
}

.countdown {
  color: #fb923c;
  font-size: 1.1rem;
  background: rgba(249, 115, 22, 0.1);
  padding: 0.5rem 1rem;
  border-radius: 12px;
  border: 1px solid rgba(249, 115, 22, 0.2);
  animation: countdownPulse 2s ease-in-out infinite;
}

.teams-details-section {
  padding: 2rem;
  display: grid;
  grid-template-columns: 1fr auto 1fr;
  gap: 2rem;
  align-items: center;
  position: relative;
  z-index: 2;
}

.teams-details-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, 
    rgba(249, 115, 22, 0.02) 0%, 
    transparent 25%, 
    transparent 75%, 
    rgba(251, 146, 60, 0.02) 100%
  );
  border-radius: 1rem;
  pointer-events: none;
}

.team-detail {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.team-visual {
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

.team-right .team-visual {
  flex-direction: row-reverse;
}

.team-right .team-details-info {
  text-align: right;
}

.team-logo-large {
  width: 80px;
  height: 80px;
  border-radius: 18px;
  background: linear-gradient(135deg, rgba(26, 15, 8, 0.8), rgba(26, 26, 26, 0.9));
  display: flex;
  align-items: center;
  justify-content: center;
  border: 2px solid rgba(249, 115, 22, 0.4);
  overflow: hidden;
  backdrop-filter: blur(25px);
  box-shadow: 
    0 10px 35px rgba(0, 0, 0, 0.4),
    0 0 20px rgba(249, 115, 22, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
}

.team-logo-large::before {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: conic-gradient(from 0deg, transparent, rgba(249, 115, 22, 0.1), transparent);
  animation: rotate-slow 20s linear infinite;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.team-logo-large:hover {
  transform: scale(1.15) rotate(5deg);
  border-color: #f97316;
  box-shadow: 
    0 15px 45px rgba(0, 0, 0, 0.5),
    0 0 35px rgba(249, 115, 22, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.2);
}

.team-logo-large:hover::before {
  opacity: 1;
}

.logo-img {
  width: 70%;
  height: 70%;
  object-fit: cover;
  border-radius: 8px;
}

.team-fallback-large {
  font-size: 1.5rem;
  font-weight: 800;
  color: #ffffff;
}

.team-details-info {
  flex: 1;
}

.team-name-large {
  font-size: 1.5rem;
  font-weight: 700;
  color: #ffffff;
  margin: 0 0 0.5rem 0;
}

.team-tag-large {
  font-size: 1rem;
  color: #9ca3af;
  font-weight: 600;
  text-transform: uppercase;
}

.team-stats {
  display: flex;
  gap: 2rem;
  padding: 1rem;
  background: rgba(0, 0, 0, 0.2);
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.stat-item {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
  text-align: center;
}

.stat-label {
  color: #9ca3af;
  font-size: 0.875rem;
  font-weight: 500;
}

.stat-value {
  color: #ffffff;
  font-size: 1.25rem;
  font-weight: 700;
}

.vs-divider-large {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
  text-align: center;
}

.vs-text-large {
  font-size: 2.5rem;
  font-weight: 900;
  background: linear-gradient(135deg, #f97316 0%, #fb923c 25%, #fbbf24 50%, #fb923c 75%, #f97316 100%);
  background-size: 200% 200%;
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  filter: drop-shadow(0 0 20px rgba(249, 115, 22, 0.7));
  animation: gradientShift 3s ease-in-out infinite, textGlow 2s ease-in-out infinite;
  position: relative;
}

.vs-text-large::after {
  content: 'VS';
  position: absolute;
  top: 0;
  left: 0;
  background: linear-gradient(135deg, #f97316, #fb923c);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  opacity: 0.3;
  animation: textPulse 1.5s ease-in-out infinite;
  z-index: -1;
}

.match-format {
  color: #fb923c;
  font-size: 0.875rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  padding: 0.5rem 1rem;
  background: linear-gradient(135deg, rgba(26, 15, 8, 0.4), rgba(249, 115, 22, 0.1));
  border: 1px solid rgba(249, 115, 22, 0.2);
  border-radius: 20px;
  backdrop-filter: blur(10px);
  box-shadow: 0 4px 15px rgba(249, 115, 22, 0.1);
  transition: all 0.3s ease;
}

.match-format:hover {
  background: linear-gradient(135deg, rgba(249, 115, 22, 0.2), rgba(251, 146, 60, 0.1));
  border-color: rgba(249, 115, 22, 0.4);
  color: #ffffff;
  transform: scale(1.05);
}

.match-info-section {
  padding: 2rem;
  background: rgba(0, 0, 0, 0.1);
}

.info-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1.5rem;
}

.info-item {
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 12px;
  padding: 1.5rem;
  transition: all 0.3s ease;
}

.info-item:hover {
  border-color: rgba(249, 115, 22, 0.3);
  background: rgba(249, 115, 22, 0.05);
}

.info-header {
  display: flex;
  align-items: center;
  color: #fb923c;
  font-weight: 600;
  margin-bottom: 0.75rem;
  position: relative;
}

.info-header i {
  animation: iconFloat 2s ease-in-out infinite;
  filter: drop-shadow(0 0 8px rgba(249, 115, 22, 0.3));
}

.info-content {
  color: #ffffff;
  font-size: 1rem;
  margin: 0;
  font-weight: 600;
}

.action-buttons {
  padding: 2rem;
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.betting-button, .share-button, .calendar-button {
  display: flex;
  align-items: center;
  padding: 1rem 2rem;
  border: none;
  border-radius: 12px;
  font-weight: 600;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
}

.betting-button, .share-button, .calendar-button {
  position: relative;
  overflow: hidden;
}

.betting-button::before, .share-button::before, .calendar-button::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.6s ease;
}

.betting-button:hover::before, .share-button:hover::before, .calendar-button:hover::before {
  left: 100%;
}

.betting-button {
  background: linear-gradient(135deg, #f97316 0%, #ea580c 100%);
  color: #ffffff;
  box-shadow: 0 4px 20px rgba(249, 115, 22, 0.3);
}

.betting-button:hover {
  transform: translateY(-3px) scale(1.05);
  box-shadow: 0 12px 40px rgba(249, 115, 22, 0.5);
  background: linear-gradient(135deg, #ea580c 0%, #dc2626 100%);
}

.share-button {
  background: linear-gradient(135deg, #fb923c 0%, #f97316 100%);
  color: #ffffff;
  box-shadow: 0 4px 20px rgba(251, 146, 60, 0.3);
}

.share-button:hover {
  transform: translateY(-3px) scale(1.05);
  box-shadow: 0 12px 40px rgba(251, 146, 60, 0.5);
  background: linear-gradient(135deg, #f97316 0%, #ea580c 100%);
}

.calendar-button {
  background: linear-gradient(135deg, #fbbf24 0%, #fb923c 100%);
  color: #ffffff;
  box-shadow: 0 4px 20px rgba(251, 191, 36, 0.3);
}

.calendar-button:hover {
  transform: translateY(-3px) scale(1.05);
  box-shadow: 0 12px 40px rgba(251, 191, 36, 0.5);
  background: linear-gradient(135deg, #fb923c 0%, #f97316 100%);
}

.streams-section {
  padding: 2rem;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.section-title {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  color: #ffffff;
  font-size: 1.25rem;
  font-weight: 600;
  margin: 0 0 1.5rem 0;
}

.section-title i {
  color: #fb923c;
  animation: iconBounce 2s ease-in-out infinite;
  filter: drop-shadow(0 0 10px rgba(249, 115, 22, 0.4));
}

.stream-links {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
}

.stream-link {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1.5rem;
  background: rgba(255, 255, 255, 0.1);
  color: #ffffff;
  text-decoration: none;
  border-radius: 8px;
  transition: all 0.3s ease;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.stream-link:hover {
  background: linear-gradient(135deg, rgba(249, 115, 22, 0.3), rgba(251, 146, 60, 0.2));
  border-color: rgba(249, 115, 22, 0.5);
  transform: translateY(-2px) scale(1.05);
  box-shadow: 0 8px 25px rgba(249, 115, 22, 0.3);
}

.stream-link.official {
  background: linear-gradient(135deg, rgba(249, 115, 22, 0.3), rgba(234, 88, 12, 0.2));
  border-color: rgba(249, 115, 22, 0.5);
  color: #fb923c;
  position: relative;
  overflow: hidden;
}

.stream-link.official::before {
  content: '';
  position: absolute;
  top: 0;
  right: 0;
  width: 4px;
  height: 100%;
  background: linear-gradient(180deg, #f97316, #fb923c);
  animation: officialGlow 2s ease-in-out infinite;
}

.stream-link.official:hover {
  background: linear-gradient(135deg, rgba(249, 115, 22, 0.4), rgba(234, 88, 12, 0.3));
  color: #ffffff;
}

/* Responsive */
@media (max-width: 768px) {
  .match-details-popup {
    width: 95%;
    height: 70vh;
    max-height: 70vh;
    margin: 0;
  }
  
  .popup-header {
    padding: 1.5rem;
  }
  
  .teams-details-section {
    grid-template-columns: 1fr;
    gap: 2rem;
    text-align: center;
  }
  
  .team-visual {
    justify-content: center;
    flex-direction: column;
    text-align: center;
  }
  
  .team-right .team-visual {
    flex-direction: column;
  }
  
  .team-right .team-details-info {
    text-align: center;
  }
  
  .info-grid {
    grid-template-columns: 1fr;
  }
  
  .action-buttons {
    flex-direction: column;
  }
  
  .vs-text-large {
    font-size: 2rem;
  }
}

@media (max-width: 480px) {
  .match-details-popup {
    width: 98%;
    height: 75vh;
    max-height: 75vh;
    border-radius: 1.5rem;
  }
  
  .popup-header {
    padding: 1rem;
  }
  
  .match-status-section,
  .teams-details-section,
  .match-info-section,
  .action-buttons {
    padding: 1.5rem;
  }
}


@keyframes pulse-bright {
  0%, 100% {
    opacity: 1;
    box-shadow: 0 0 20px currentColor;
  }
  50% {
    opacity: 0.7;
    box-shadow: 0 0 30px currentColor;
  }
}

@keyframes glow-pulse {
  0%, 100% {
    box-shadow: 0 0 20px currentColor;
  }
  50% {
    box-shadow: 0 0 35px currentColor, 0 0 50px rgba(249, 115, 22, 0.3);
  }
}

@keyframes iconGlow {
  0%, 100% {
    filter: drop-shadow(0 0 8px rgba(249, 115, 22, 0.5));
    transform: scale(1);
  }
  50% {
    filter: drop-shadow(0 0 15px rgba(249, 115, 22, 0.8));
    transform: scale(1.05);
  }
}

@keyframes countdownPulse {
  0%, 100% {
    background: rgba(249, 115, 22, 0.1);
    border-color: rgba(249, 115, 22, 0.2);
  }
  50% {
    background: rgba(249, 115, 22, 0.2);
    border-color: rgba(249, 115, 22, 0.4);
    box-shadow: 0 0 15px rgba(249, 115, 22, 0.2);
  }
}

@keyframes gradientShift {
  0%, 100% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
}

@keyframes textGlow {
  0%, 100% {
    filter: drop-shadow(0 0 20px rgba(249, 115, 22, 0.7));
  }
  50% {
    filter: drop-shadow(0 0 30px rgba(249, 115, 22, 1));
  }
}

@keyframes textPulse {
  0%, 100% {
    opacity: 0.3;
    transform: scale(1);
  }
  50% {
    opacity: 0.5;
    transform: scale(1.02);
  }
}

@keyframes iconFloat {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-3px);
  }
}

@keyframes iconBounce {
  0%, 100% {
    transform: translateY(0px) scale(1);
  }
  50% {
    transform: translateY(-2px) scale(1.1);
  }
}

@keyframes officialGlow {
  0%, 100% {
    opacity: 0.7;
  }
  50% {
    opacity: 1;
    box-shadow: 0 0 10px rgba(249, 115, 22, 0.5);
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

@keyframes overlayFadeIn {
  from {
    opacity: 0;
    backdrop-filter: blur(0px);
  }
  to {
    opacity: 1;
    backdrop-filter: blur(20px);
  }
}

@keyframes popupSlideIn {
  from {
    opacity: 0;
    transform: translateY(50px) scale(0.9);
    filter: blur(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
    filter: blur(0px);
  }
}
</style>
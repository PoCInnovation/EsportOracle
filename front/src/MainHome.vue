<template>
  <div class="home">
    <div class="home-container">
      <!-- Hero Section -->
      <section class="hero">
        <div class="hero-content">
          <div class="hero-badge">
            <span class="badge-icon">⚡</span>
            #1 Esports Platform
          </div>
          <h1 class="hero-title">
            Experience <span class="gradient-text">Esports</span>
            <br>like never before
          </h1>
          <p class="hero-description">
            Dive into the competitive universe with real-time data, 
            in-depth analysis and an immersive experience for all esports fans.
          </p>
          <div class="hero-actions">
            <router-link to="/matches/current" class="btn btn-primary">
              <span class="btn-icon">🔴</span>
              Live Matches
              <span class="btn-arrow">→</span>
            </router-link>
            <router-link to="/matches/upcoming" class="btn btn-secondary">
              <span class="btn-icon">📅</span>
              Upcoming Matches
            </router-link>
          </div>
          <div class="hero-stats">
            <div class="stat-item">
              <span class="stat-number">250K+</span>
              <span class="stat-label">Tracked matches</span>
            </div>
            <div class="stat-item">
              <span class="stat-number">50+</span>
              <span class="stat-label">Active tournaments</span>
            </div>
            <div class="stat-item">
              <span class="stat-number">24/7</span>
              <span class="stat-label">Real-time data</span>
            </div>
          </div>
        </div>
        <div class="hero-visual">
          <div class="floating-cards">
            <!-- Carte Match Live -->
            <div v-if="liveMatches.length > 0" class="match-card-demo card-1">
              <div class="match-header">
                <span class="live-badge">
                  🔴 LIVE
                  <span class="live-dot"></span>
                </span>
                <span class="match-time">{{ formatTime(liveMatches[0].begin_at) }}</span>
              </div>
              <div class="match-teams">
                <div class="team" v-if="liveMatches[0].opponents?.[0]">
                  <div class="team-logo-container">
                    <img 
                      v-if="getTeamImageUrl(liveMatches[0].opponents[0])"
                      :src="getTeamImageUrl(liveMatches[0].opponents[0])" 
                      :alt="liveMatches[0].opponents[0].opponent.name"
                      class="team-logo"
                      @error="handleImageError"
                      @load="handleImageLoad"
                    />
                    <div v-else class="team-fallback">
                      {{ getTeamInitials(liveMatches[0].opponents[0].opponent.name) }}
                    </div>
                  </div>
                  <div class="team-info">
                    <span class="team-name">{{ liveMatches[0].opponents[0].opponent.acronym || liveMatches[0].opponents[0].opponent.name }}</span>
                    <span class="team-score">--</span>
                  </div>
                </div>
                <div class="vs-separator">VS</div>
                <div class="team" v-if="liveMatches[0].opponents?.[1]">
                  <div class="team-logo-container">
                    <img 
                      v-if="getTeamImageUrl(liveMatches[0].opponents[1])"
                      :src="getTeamImageUrl(liveMatches[0].opponents[1])" 
                      :alt="liveMatches[0].opponents[1].opponent.name"
                      class="team-logo"
                      @error="handleImageError"
                      @load="handleImageLoad"
                    />
                    <div v-else class="team-fallback">
                      {{ getTeamInitials(liveMatches[0].opponents[1].opponent.name) }}
                    </div>
                  </div>
                  <div class="team-info">
                    <span class="team-name">{{ liveMatches[0].opponents[1].opponent.acronym || liveMatches[0].opponents[1].opponent.name }}</span>
                    <span class="team-score">--</span>
                  </div>
                </div>
              </div>
              <div class="match-info">
                <span class="tournament">🏆 {{ liveMatches[0].tournament?.name || liveMatches[0].league?.name || 'Tournament' }}</span>
              </div>
            </div>
            
            <!-- Live Match Loading Card -->
            <div v-else-if="store.loading" class="match-card-demo card-1 loading-card">
              <div class="loading-content">
                <div class="loading-spinner-small"></div>
                <span>Loading live matches...</span>
              </div>
            </div>
            
            <!-- Live Error Card -->
            <div v-else-if="store.error" class="match-card-demo card-1 error-card">
              <div class="error-content">
                <span class="error-icon">⚠️</span>
                <span>Loading error</span>
              </div>
            </div>
            
            <!-- Default card if no live match -->
            <div v-else class="match-card-demo card-1">
              <div class="match-header">
                <span class="live-badge">
                  🔴 LIVE
                  <span class="live-dot"></span>
                </span>
                <span class="match-time">--:--</span>
              </div>
              <div class="match-teams">
                <div class="team">
                  <div class="team-logo-container">
                    <div class="team-fallback">FNC</div>
                  </div>
                  <div class="team-info">
                    <span class="team-name">Aucun match</span>
                    <span class="team-score">--</span>
                  </div>
                </div>
                <div class="vs-separator">VS</div>
                <div class="team">
                  <div class="team-logo-container">
                    <div class="team-fallback">G2</div>
                  </div>
                  <div class="team-info">
                    <span class="team-name">en direct</span>
                    <span class="team-score">--</span>
                  </div>
                </div>
              </div>
              <div class="match-info">
                <span class="tournament">🏆 Tournoi</span>
              </div>
            </div>
            
            <!-- Upcoming Match Card -->
            <div v-if="upcomingMatches.length > 0" class="match-card-demo card-2">
              <div class="match-header">
                <span class="upcoming-badge">📅 UPCOMING</span>
                <span class="match-time">{{ formatTime(upcomingMatches[0].scheduled_at) }}</span>
              </div>
              <div class="match-teams">
                <div class="team" v-if="upcomingMatches[0].opponents?.[0]">
                  <div class="team-logo-container">
                    <img 
                      v-if="getTeamImageUrl(upcomingMatches[0].opponents[0])"
                      :src="getTeamImageUrl(upcomingMatches[0].opponents[0])" 
                      :alt="upcomingMatches[0].opponents[0].opponent.name"
                      class="team-logo"
                      @error="handleImageError"
                      @load="handleImageLoad"
                    />
                    <div v-else class="team-fallback">
                      {{ getTeamInitials(upcomingMatches[0].opponents[0].opponent.name) }}
                    </div>
                  </div>
                  <span class="team-name">{{ upcomingMatches[0].opponents[0].opponent.acronym || upcomingMatches[0].opponents[0].opponent.name }}</span>
                </div>
                <div class="vs-separator">VS</div>
                <div class="team" v-if="upcomingMatches[0].opponents?.[1]">
                  <div class="team-logo-container">
                    <img 
                      v-if="getTeamImageUrl(upcomingMatches[0].opponents[1])"
                      :src="getTeamImageUrl(upcomingMatches[0].opponents[1])" 
                      :alt="upcomingMatches[0].opponents[1].opponent.name"
                      class="team-logo"
                      @error="handleImageError"
                      @load="handleImageLoad"
                    />
                    <div v-else class="team-fallback">
                      {{ getTeamInitials(upcomingMatches[0].opponents[1].opponent.name) }}
                    </div>
                  </div>
                  <span class="team-name">{{ upcomingMatches[0].opponents[1].opponent.acronym || upcomingMatches[0].opponents[1].opponent.name }}</span>
                </div>
              </div>
              <div class="match-info">
                <span class="tournament">🏆 {{ upcomingMatches[0].tournament?.name || upcomingMatches[0].league?.name || 'Tournament' }}</span>
              </div>
            </div>
            
            <!-- Upcoming Loading Card -->
            <div v-else-if="store.loading && store.currentMatchType === 'upcoming'" class="match-card-demo card-2 loading-card">
              <div class="loading-content">
                <div class="loading-spinner-small"></div>
                <span>Loading upcoming matches...</span>
              </div>
            </div>
            
            <!-- Default card if no upcoming match -->
            <div v-else class="match-card-demo card-2">
              <div class="match-header">
                <span class="upcoming-badge">📅 UPCOMING</span>
                <span class="match-time">--:--</span>
              </div>
              <div class="match-teams">
                <div class="team">
                  <div class="team-logo-container">
                    <div class="team-fallback">TL</div>
                  </div>
                  <span class="team-name">No match</span>
                </div>
                <div class="vs-separator">VS</div>
                <div class="team">
                  <div class="team-logo-container">
                    <div class="team-fallback">NV</div>
                  </div>
                  <span class="team-name">scheduled</span>
                </div>
              </div>
              <div class="match-info">
                <span class="tournament">🏆 Tournament</span>
              </div>
            </div>
            
            <div class="floating-elements">
              <div class="floating-icon icon-1">🎮</div>
              <div class="floating-icon icon-2">🏆</div>
              <div class="floating-icon icon-3">⚡</div>
              <div class="floating-icon icon-4">🔥</div>
            </div>
          </div>
        </div>
      </section>

      <!-- Features Section -->
      <section class="features">
        <div class="section-header">
          <h2 class="section-title">
            <span class="title-icon">⚡</span>
            Advanced Features
          </h2>
          <p class="section-description">
            Discover the power of our Web3 oracle for esports
          </p>
        </div>
        <div class="features-grid">
          <div class="feature-card">
            <div class="feature-icon">🌐</div>
            <h3>Multi-Platform</h3>
            <p>Access your esports data from any device or platform</p>
            <div class="feature-highlight">
              <span class="highlight-text">Web, Mobile, API</span>
            </div>
          </div>
          <div class="feature-card">
            <div class="feature-icon">⛓️</div>
            <h3>Web3 OnChain Oracle</h3>
            <p>Verified and secured esports data directly on the blockchain</p>
            <div class="feature-highlight">
              <span class="highlight-text">100% decentralized</span>
            </div>
          </div>
          <div class="feature-card">
            <div class="feature-icon">🖥️</div>
            <h3>Blockchain Integration</h3>
            <p>Connect your esports data to the Web3 and DeFi ecosystem</p>
            <div class="feature-highlight">
              <span class="highlight-text">Smart Contracts</span>
            </div>
          </div>
          <div class="feature-card">
            <div class="feature-icon">🎮</div>
            <h3>Real-Time Data</h3>
            <p>Matches, tournaments and statistics updated live via our oracle</p>
            <div class="feature-highlight">
              <span class="highlight-text">PandaScore API</span>
            </div>
          </div>
        </div>
      </section>

    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, computed, ref } from 'vue'
import { matchStore } from './stores/matchStore'

const store = matchStore()
const failedImages = ref<Set<string>>(new Set())

// Computed pour obtenir les matches live et à venir
const liveMatches = computed(() => store.currentMatches.slice(0, 1))
const upcomingMatches = computed(() => store.upcomingMatches.slice(0, 1))

// Fonction pour formater l'heure
const formatTime = (dateString: string | undefined) => {
  if (!dateString) return '--:--'
  const date = new Date(dateString)
  return date.toLocaleTimeString('fr-FR', { 
    hour: '2-digit', 
    minute: '2-digit' 
  })
}

// Fonction pour obtenir l'URL de l'image d'une équipe depuis PandaScore
const getTeamImageUrl = (opponent: any): string | null => {
  if (!opponent?.opponent?.image_url) return null
  const url = opponent.opponent.image_url.trim()
  if (!url || url === 'null' || failedImages.value.has(url)) return null
  try { 
    new URL(url)
    return url 
  } catch { 
    return null 
  }
}

// Fonction pour obtenir les initiales d'une équipe
const getTeamInitials = (teamName: string): string => {
  if (!teamName) return '?'
  return teamName.split(' ')
    .map(word => word.charAt(0).toUpperCase())
    .join('')
    .substring(0, 3)
}

// Gestion des erreurs d'images
const handleImageError = (event: Event) => {
  const img = event.target as HTMLImageElement
  if (img.src) failedImages.value.add(img.src)
  img.src = ''
}

const handleImageLoad = (event: Event) => {
  (event.target as HTMLImageElement).style.opacity = '1'
}

// Charger les données au montage
onMounted(async () => {
  // Charger les matches en direct
  const currentUrl = store.createUrlMatches('current', undefined)
  await store.fetchMatches(currentUrl, 'current')
  
  // Charger les matches à venir
  const upcomingUrl = store.createUrlMatches('upcoming', undefined)
  await store.fetchMatches(upcomingUrl, 'upcoming')
})
</script>


<style scoped>
.home {
  min-height: calc(100vh - 200px);
}

.home-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
}

/* Hero Section */
.hero {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 4rem;
  align-items: center;
  margin-bottom: 5rem;
  padding: 4rem 0;
  position: relative;
}

.hero::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(ellipse at center, rgba(249, 115, 22, 0.15) 0%, transparent 70%);
  pointer-events: none;
  z-index: -1;
}

.hero::after {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: conic-gradient(from 0deg at 50% 50%, transparent 0deg, rgba(249, 115, 22, 0.1) 45deg, transparent 90deg, rgba(251, 146, 60, 0.1) 135deg, transparent 180deg);
  animation: rotate 20s linear infinite;
  z-index: -2;
}

.hero-title {
  font-size: clamp(2.5rem, 5vw, 4rem);
  font-weight: 800;
  line-height: 1.1;
  margin-bottom: 1.5rem;
  color: #ffffff;
}

.gradient-text {
  background: linear-gradient(135deg, #f97316, #fb923c);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.hero-description {
  font-size: 1.125rem;
  color: #a3a3a3;
  margin-bottom: 2rem;
  line-height: 1.7;
}

.hero-actions {
  display: flex;
  gap: 1rem;
  flex-wrap: wrap;
}

.hero-visual {
  display: flex;
  justify-content: right;
  align-items: center;
}

.floating-card {
  animation: float 6s ease-in-out infinite;
}

.mock-match {
  background: linear-gradient(135deg, #2d2d2d 0%, #2d1308 50%, #2d2d2d 100%);
  border: 1px solid rgba(249, 115, 22, 0.4);
  border-radius: 1.5rem;
  padding: 1.5rem;
  box-shadow: 0 25px 50px -12px rgba(249, 115, 22, 0.15);
  min-width: 300px;
  position: relative;
  overflow: hidden;
}

.mock-match::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(249, 115, 22, 0.15), transparent);
  animation: shimmer 3s infinite;
}

.match-teams {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 1rem;
}

.team {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5rem;
}

.team-logo {
  font-size: 2rem;
  width: 60px;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #1a1a1a;
  border-radius: 0.75rem;
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.vs {
  background: linear-gradient(135deg, #f97316, #fb923c);
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 0.75rem;
  font-weight: 700;
  font-size: 0.875rem;
  box-shadow: 0 4px 12px rgba(249, 115, 22, 0.3);
}

.match-status {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  padding: 0.75rem;
  background: #ef4444;
  color: white;
  border-radius: 0.75rem;
  font-weight: 600;
  font-size: 0.875rem;
}

.live-indicator {
  width: 8px;
  height: 8px;
  background: white;
  border-radius: 50%;
  animation: pulse 2s infinite;
}

/* Features Section */
.features {
  margin-bottom: 5rem;
}

.section-header {
  text-align: center;
  margin-bottom: 3rem;
}

.section-title {
  font-size: 2.5rem;
  font-weight: 800;
  color: #ffffff;
  margin-bottom: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
}

.title-icon {
  font-size: 2rem;
}

.section-description {
  font-size: 1.125rem;
  color: #a3a3a3;
  max-width: 600px;
  margin: 0 auto;
  line-height: 1.6;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.feature-card {
  background: linear-gradient(135deg, #1a1a1a 0%, #1a0f08 50%, #1a1a1a 100%);
  border: 1px solid rgba(249, 115, 22, 0.3);
  border-radius: 1.5rem;
  padding: 2rem;
  text-align: center;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.feature-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(circle at center, rgba(249, 115, 22, 0.08) 0%, transparent 70%);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.feature-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 25px 50px -12px rgba(249, 115, 22, 0.25);
  border-color: #fb923c;
}

.feature-card:hover::before {
  opacity: 1;
}

.feature-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.feature-card h3 {
  color: #ffffff;
  margin-bottom: 0.75rem;
}

.feature-card p {
  color: #a3a3a3;
  margin: 0;
}

/* Stats Section */
.stats {
  margin-bottom: 4rem;
}

.stats-card {
  background: linear-gradient(135deg, #2d2d2d 0%, #2d1810 50%, #2d2d2d 100%);
  border: 1px solid rgba(249, 115, 22, 0.3);
  border-radius: 1.5rem;
  padding: 3rem;
  position: relative;
  overflow: hidden;
}

.stats-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(ellipse at center, rgba(249, 115, 22, 0.1) 0%, transparent 70%);
  z-index: -1;
}

.stats-card {
  text-align: center;
}

.stats-title {
  color: #ffffff;
  margin-bottom: 2rem;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 2rem;
}

.hero-stats {
  margin-top: 30px;
}

.stat-number {
  font-size: 3rem;
  font-weight: 800;
  color: #60a5fa;
  margin-bottom: 0.5rem;
}

.stat-label {
  color: #a3a3a3;
  font-weight: 500;
}

/* Animations */
@keyframes float {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-20px);
  }
}

@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
}

/* Animations */
@keyframes shimmer {
  0% {
    left: -100%;
  }
  100% {
    left: 100%;
  }
}

@keyframes rotate {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

@keyframes float {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-20px);
  }
}

/* Enhanced Animations for Floating Cards */
@keyframes floatCard {
  0%, 100% {
    transform: translateY(0px) rotateZ(0deg);
  }
  25% {
    transform: translateY(-15px) rotateZ(1deg);
  }
  50% {
    transform: translateY(-8px) rotateZ(0deg);
  }
  75% {
    transform: translateY(-20px) rotateZ(-1deg);
  }
}

@keyframes shimmerCard {
  0% {
    left: -100%;
  }
  50% {
    left: 100%;
  }
  100% {
    left: 100%;
  }
}

@keyframes livePulse {
  0% {
    left: -100%;
  }
  100% {
    left: 100%;
  }
}

@keyframes floatIcon {
  0%, 100% {
    transform: translateY(0px) rotateZ(0deg) scale(1);
  }
  25% {
    transform: translateY(-30px) rotateZ(5deg) scale(1.1);
  }
  50% {
    transform: translateY(-15px) rotateZ(0deg) scale(1);
  }
  75% {
    transform: translateY(-25px) rotateZ(-5deg) scale(1.05);
  }
}

@keyframes shimmerVs {
  0% {
    left: -100%;
  }
  100% {
    left: 100%;
  }
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

/* Games Section */
.games {
  margin-bottom: 5rem;
}

.games-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 2rem;
}

.game-card {
  background: linear-gradient(135deg, #1a1a1a 0%, #1a0f08 50%, #1a1a1a 100%);
  border: 1px solid rgba(249, 115, 22, 0.3);
  border-radius: 1.5rem;
  padding: 2rem;
  text-align: center;
  position: relative;
  overflow: hidden;
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
}

.game-card:hover {
  transform: translateY(-8px);
  border-color: rgba(249, 115, 22, 0.6);
  box-shadow: 0 25px 50px -12px rgba(249, 115, 22, 0.25);
}

.game-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(249, 115, 22, 0.1) 0%, transparent 100%);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.game-card:hover::before {
  opacity: 1;
}

.game-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
  display: block;
}

.game-card h3 {
  color: #ffffff;
  font-size: 1.25rem;
  font-weight: 700;
  margin-bottom: 0.75rem;
}

.game-card p {
  color: #a3a3a3;
  font-size: 0.875rem;
  margin-bottom: 1rem;
  line-height: 1.5;
}

.game-stats {
  display: inline-block;
  background: linear-gradient(135deg, #f97316, #fb923c);
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 0.75rem;
  font-size: 0.75rem;
  font-weight: 600;
  box-shadow: 0 4px 12px rgba(249, 115, 22, 0.3);
}

/* Floating Cards */
.floating-cards {
  display: flex;
  flex-direction: column;
  gap: 2rem;
  position: relative;
  perspective: 1000px;
}

.match-card-demo {
  background: linear-gradient(135deg, #1a1a1a 0%, #1a0f08 50%, #1a1a1a 100%);
  border: 1px solid rgba(249, 115, 22, 0.4);
  border-radius: 1.5rem;
  padding: 2.5rem;
  backdrop-filter: blur(15px);
  position: relative;
  overflow: hidden;
  animation: floatCard 8s ease-in-out infinite;
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  box-shadow: 
    0 10px 30px rgba(0, 0, 0, 0.3),
    0 4px 15px rgba(249, 115, 22, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
  transform-style: preserve-3d;
}

.match-card-demo::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    90deg, 
    transparent, 
    rgba(249, 115, 22, 0.3), 
    transparent
  );
  animation: shimmerCard 4s ease-in-out infinite;
  z-index: 1;
}

.match-card-demo::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(
    circle at 50% 50%, 
    rgba(249, 115, 22, 0.15) 0%, 
    transparent 60%
  );
  opacity: 0;
  transition: opacity 0.4s ease;
  z-index: -1;
}

.match-card-demo:hover {
  transform: translateY(-12px) rotateX(5deg) rotateY(5deg);
  border-color: rgba(249, 115, 22, 0.8);
  box-shadow: 
    0 25px 60px rgba(0, 0, 0, 0.4),
    0 8px 30px rgba(249, 115, 22, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.2);
}

.match-card-demo:hover::after {
  opacity: 1;
}

.card-1 {
  animation-delay: 0s;
}

.card-2 {
  animation-delay: 3s;
}

.match-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  position: relative;
  z-index: 2;
}

.live-badge {
  background: linear-gradient(135deg, #ef4444, #dc2626);
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 0.75rem;
  font-size: 0.75rem;
  font-weight: 700;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  box-shadow: 0 4px 15px rgba(239, 68, 68, 0.4);
  position: relative;
  overflow: hidden;
}

.live-badge::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  animation: livePulse 2s ease-in-out infinite;
}

.live-badge .live-dot {
  width: 8px;
  height: 8px;
  background: white;
  border-radius: 50%;
  animation: pulse 1.5s ease-in-out infinite;
}

.match-time {
  color: #f97316;
  font-size: 0.875rem;
  font-weight: 600;
  background: rgba(249, 115, 22, 0.1);
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  border: 1px solid rgba(249, 115, 22, 0.3);
  font-family: 'Monaco', 'Menlo', monospace;
}

/* Match Card Content */
.match-teams {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 2rem;
  position: relative;
  z-index: 2;
}

.team {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
  flex: 1;
}

.team-logo {
  width: 80%;
  height: 80%;
  border-radius: 1rem;
  border: 2px solid rgba(249, 115, 22, 0.3);
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.3);
  transition: all 0.3s ease;
  object-fit: cover;
  opacity: 0;
}

.team-logo-container {
  width: 80px;
  height: 80px;
  border-radius: 1rem;
  background: linear-gradient(135deg, #1a1a1a, #2d2d2d);
  display: flex;
  align-items: center;
  justify-content: center;
  border: 2px solid rgba(249, 115, 22, 0.3);
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.3);
  transition: all 0.3s ease;
  overflow: hidden;
  flex-shrink: 0;
  backdrop-filter: blur(10px);
}

.team-logo-container:hover {
  transform: scale(1.1);
  border-color: rgba(249, 115, 22, 0.6);
  box-shadow: 0 12px 30px rgba(249, 115, 22, 0.2);
}

.team-fallback {
  font-size: 1rem;
  font-weight: 700;
  color: #ffffff;
  text-align: center;
}

.team-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5rem;
}

.team-name {
  color: #ffffff;
  font-weight: 700;
  font-size: 1rem;
  text-align: center;
}

.team-score {
  color: #f97316;
  font-weight: 800;
  font-size: 1.25rem;
  text-shadow: 0 2px 4px rgba(249, 115, 22, 0.3);
}

.vs-separator {
  background: linear-gradient(135deg, #f97316, #fb923c);
  color: white;
  padding: 1rem 1.5rem;
  border-radius: 1rem;
  font-weight: 800;
  font-size: 1rem;
  box-shadow: 0 6px 20px rgba(249, 115, 22, 0.4);
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  position: relative;
  overflow: hidden;
  margin: 0 1.5rem;
}

.vs-separator::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  animation: shimmerVs 3s ease-in-out infinite;
}

.match-info {
  display: flex;
  justify-content: center;
  position: relative;
  z-index: 2;
}

.tournament {
  color: #a3a3a3;
  font-size: 0.75rem;
  font-weight: 600;
  background: rgba(249, 115, 22, 0.1);
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  border: 1px solid rgba(249, 115, 22, 0.2);
}

.upcoming-badge {
  background: linear-gradient(135deg, #0ea5e9, #0284c7);
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 0.75rem;
  font-size: 0.75rem;
  font-weight: 700;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  box-shadow: 0 4px 15px rgba(14, 165, 233, 0.4);
}

/* Loading States */
.loading-card {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 200px;
}

.loading-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
  color: #a3a3a3;
  font-size: 0.875rem;
}

.loading-spinner-small {
  width: 32px;
  height: 32px;
  border: 3px solid rgba(249, 115, 22, 0.1);
  border-left: 3px solid #f97316;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

/* Error States */
.error-card {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 200px;
}

.error-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5rem;
  color: #ef4444;
  font-size: 0.875rem;
}

.error-icon {
  font-size: 2rem;
}

.floating-elements {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;
  z-index: 0;
  overflow: hidden;
}

.floating-icon {
  position: absolute;
  font-size: 2.5rem;
  opacity: 0.2;
  animation: floatIcon 12s ease-in-out infinite;
  filter: drop-shadow(0 0 10px rgba(249, 115, 22, 0.3));
  transition: all 0.3s ease;
}

.floating-icon:hover {
  opacity: 0.6;
  transform: scale(1.2);
}

.icon-1 {
  top: 15%;
  right: 10%;
  animation-delay: 0s;
  color: #f97316;
}

.icon-2 {
  top: 65%;
  right: 25%;
  animation-delay: 3s;
  color: #fb923c;
}

.icon-3 {
  bottom: 25%;
  right: 15%;
  animation-delay: 6s;
  color: #ea580c;
}

.icon-4 {
  top: 35%;
  right: 5%;
  animation-delay: 9s;
  color: #fdba74;
}

/* Responsive */
@media (max-width: 768px) {
  .home-container {
    padding: 0 1rem;
  }
  
  .hero {
    grid-template-columns: 1fr;
    gap: 2rem;
    text-align: center;
  }

  .hero-actions {
    justify-content: center;
  }

  .features-grid {
    grid-template-columns: 1fr;
  }

  .stats-grid {
    grid-template-columns: 1fr;
  }

  .mock-match {
    min-width: 280px;
  }

  .hero {
    padding: 2rem 0;
  }

  .floating-cards {
    gap: 1.5rem;
  }

  .match-card-demo {
    padding: 2rem;
  }

  .team-logo {
    width: 80%;
    height: 80%;
  }

  .team-logo-container {
    width: 60px;
    height: 60px;
  }

  .team-name {
    font-size: 0.75rem;
  }

  .vs-separator {
    padding: 0.5rem 1rem;
    font-size: 0.75rem;
  }
}

@media (max-width: 480px) {
  .home-container {
    padding: 0 0.5rem;
  }

  .floating-cards {
    gap: 1rem;
  }

  .match-card-demo {
    padding: 1.5rem;
  }

  .team-logo {
    width: 80%;
    height: 80%;
  }

  .team-logo-container {
    width: 50px;
    height: 50px;
  }

  .match-teams {
    margin-bottom: 1rem;
  }

  .vs-separator {
    padding: 0.75rem 1rem;
    font-size: 0.875rem;
    margin: 0 1rem;
  }
}
</style>

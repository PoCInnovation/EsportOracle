<template>
  <div class="matches-selector">
    <div class="container">
      <!-- Hero Section -->
      <section class="hero">
        <h1 class="hero-title">
          <span class="title-icon">🎮</span>
          Matches Esport
        </h1>
        <p class="hero-description">
          Choisissez le type de matches que vous souhaitez consulter
        </p>
      </section>

      <!-- Categories Grid -->
      <section class="categories">
        <div class="categories-grid">
          <!-- Live Matches -->
          <router-link to="/matches/current" class="category-card live">
            <div class="card-header">
              <div class="card-icon">🔴</div>
              <div class="live-indicator">
                <span class="pulse-dot"></span>
                <span class="live-text">EN DIRECT</span>
              </div>
            </div>
            <div class="card-content">
              <h2 class="card-title">Matches en Direct</h2>
              <p class="card-description">
                Suivez les matches en cours avec des mises à jour en temps réel
              </p>
            </div>
            <div class="card-footer">
              <span class="card-action">Voir maintenant →</span>
            </div>
          </router-link>

          <!-- Upcoming Matches -->
          <router-link to="/matches/upcoming" class="category-card upcoming">
            <div class="card-header">
              <div class="card-icon">⏰</div>
              <div class="status-badge upcoming-badge">À VENIR</div>
            </div>
            <div class="card-content">
              <h2 class="card-title">Matches à Venir</h2>
              <p class="card-description">
                Planifiez votre programme avec les prochains matches programmés
              </p>
            </div>
            <div class="card-footer">
              <span class="card-action">Découvrir →</span>
            </div>
          </router-link>

          <!-- Past Matches -->
          <router-link to="/matches/past" class="category-card past">
            <div class="card-header">
              <div class="card-icon">📚</div>
              <div class="status-badge past-badge">TERMINÉS</div>
            </div>
            <div class="card-content">
              <h2 class="card-title">Matches Passés</h2>
              <p class="card-description">
                Consultez les résultats et statistiques des matches terminés
              </p>
            </div>
            <div class="card-footer">
              <span class="card-action">Explorer →</span>
            </div>
          </router-link>
        </div>
      </section>

      <!-- Quick Stats -->
      <section class="quick-stats">
        <div class="stats-card">
          <h3 class="stats-title">Statistiques Rapides</h3>
          <div class="stats-grid">
            <div class="stat">
              <div class="stat-icon">🎯</div>
              <div class="stat-content">
                <div class="stat-number">250+</div>
                <div class="stat-label">Matches / semaine</div>
              </div>
            </div>
            <div class="stat">
              <div class="stat-icon">🏆</div>
              <div class="stat-content">
                <div class="stat-number">15</div>
                <div class="stat-label">Jeux supportés</div>
              </div>
            </div>
            <div class="stat">
              <div class="stat-icon">⚡</div>
              <div class="stat-content">
                <div class="stat-number">24/7</div>
                <div class="stat-label">Mises à jour</div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<style scoped>
.matches-selector {
  min-height: calc(100vh - 200px);
  padding: 2rem 0;
}

/* Hero Section */
.hero {
  text-align: center;
  margin-bottom: 4rem;
}

.hero-title {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 1rem;
  font-size: clamp(2.5rem, 5vw, 4rem);
  font-weight: 800;
  margin-bottom: 1.5rem;
  background: linear-gradient(135deg, #f97316, #fb923c);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.title-icon {
  font-size: 3rem;
}

.hero-description {
  font-size: 1.125rem;
  color: #a3a3a3;
  max-width: 600px;
  margin: 0 auto;
  line-height: 1.7;
}

/* Categories Grid */
.categories {
  margin-bottom: 4rem;
}

.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
  gap: 2rem;
}

.category-card {
  background: #1a1a1a;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 1.5rem;
  padding: 2rem;
  text-decoration: none;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  min-height: 280px;
}

.category-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  transition: all 0.3s ease;
}

.category-card.live::before {
  background: linear-gradient(90deg, #ef4444, #f87171);
}

.category-card.upcoming::before {
  background: linear-gradient(90deg, #3b82f6, #60a5fa);
}

.category-card.past::before {
  background: linear-gradient(90deg, #737373, #9ca3af);
}

.category-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
  border-color: #3b82f6;
}

.category-card:hover::before {
  height: 6px;
}

/* Card Header */
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 1.5rem;
}

.card-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

/* Live Indicator */
.live-indicator {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background: rgba(239, 68, 68, 0.1);
  border: 1px solid rgba(239, 68, 68, 0.2);
  padding: 0.5rem 0.75rem;
  border-radius: 0.75rem;
}

.pulse-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #ef4444;
  animation: pulse 2s infinite;
}

.live-text {
  font-size: 0.75rem;
  font-weight: 700;
  color: #ef4444;
  letter-spacing: 0.05em;
}

/* Status Badges */
.status-badge {
  padding: 0.5rem 0.75rem;
  border-radius: 0.75rem;
  font-size: 0.75rem;
  font-weight: 700;
  letter-spacing: 0.05em;
}

.upcoming-badge {
  background: rgba(59, 130, 246, 0.1);
  border: 1px solid rgba(59, 130, 246, 0.2);
  color: #3b82f6;
}

.past-badge {
  background: rgba(107, 114, 128, 0.1);
  border: 1px solid rgba(107, 114, 128, 0.2);
  color: #737373;
}

/* Card Content */
.card-content {
  flex: 1;
}

.card-title {
  font-size: 1.5rem;
  font-weight: 700;
  color: #ffffff;
  margin-bottom: 1rem;
}

.card-description {
  color: #a3a3a3;
  line-height: 1.6;
  margin: 0;
}

/* Card Footer */
.card-footer {
  margin-top: 1.5rem;
  padding-top: 1rem;
  border-top: 1px solid rgba(255, 255, 255, 0.2);
}

.card-action {
  color: #60a5fa;
  font-weight: 600;
  font-size: 0.875rem;
  transition: color 0.2s ease;
}

.category-card:hover .card-action {
  color: #93c5fd;
}

/* Quick Stats */
.quick-stats {
  margin-bottom: 3rem;
}

.stats-card {
  background: linear-gradient(135deg, #2d2d2d, #1a1a1a);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 1.5rem;
  padding: 2.5rem;
  text-align: center;
}

.stats-title {
  color: #ffffff;
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 2rem;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 2rem;
}

.stat {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1.5rem;
  background: #1a1a1a;
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 0.75rem;
  transition: all 0.3s ease;
}

.stat:hover {
  transform: translateY(-2px);
  border-color: #3b82f6;
  box-shadow: 0 10px 25px -3px rgba(0, 0, 0, 0.1);
}

.stat-icon {
  font-size: 2rem;
  width: 60px;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #3b82f6, #34d399);
  border-radius: 0.75rem;
  flex-shrink: 0;
}

.stat-content {
  text-align: left;
}

.stat-number {
  font-size: 1.5rem;
  font-weight: 800;
  color: #ffffff;
  margin-bottom: 0.25rem;
}

.stat-label {
  color: #a3a3a3;
  font-size: 0.875rem;
  font-weight: 500;
}

/* Animations */
@keyframes pulse {
  0%, 100% {
    opacity: 1;
    transform: scale(1);
  }
  50% {
    opacity: 0.7;
    transform: scale(1.1);
  }
}

/* Responsive Design */
@media (max-width: 768px) {
  .hero-title {
    flex-direction: column;
    gap: 0.5rem;
  }

  .title-icon {
    font-size: 2.5rem;
  }

  .categories-grid {
    grid-template-columns: 1fr;
    gap: 1.5rem;
  }

  .stats-grid {
    grid-template-columns: 1fr;
    gap: 1rem;
  }

  .stat {
    flex-direction: column;
    text-align: center;
  }

  .stat-content {
    text-align: center;
  }

  .card-header {
    flex-direction: column;
    align-items: center;
    text-align: center;
    gap: 1rem;
  }
}

@media (max-width: 480px) {
  .matches-selector {
    padding: 1rem 0;
  }

  .category-card {
    padding: 1.5rem;
    min-height: 240px;
  }

  .stats-card {
    padding: 1.5rem;
  }
}
</style>
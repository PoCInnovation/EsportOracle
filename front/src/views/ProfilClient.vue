<template>
  <div class="profil-container">
    <Card class="profil-card">
      <template #header>
        <div class="card-header">
          <div class="user-info">
            <h2 class="username">{{ userPseudo || 'Utilisateur Anonyme' }}</h2>
            <p class="wallet-address">{{ formatWalletAddress(walletStore.connectedAccount) }}</p>
          </div>
        </div>
      </template>

      <template #content>
        <div class="profile-content">
          
          <!-- Identité -->
          <div class="section">
            <h3 class="section-title">
              <i class="pi pi-user"></i>
              Informations du profil
            </h3>
            <div class="form-group-inline">
              <InputText
                id="pseudo"
                v-model="userPseudo" 
                placeholder="Entrez votre pseudo"
                class="profile-input"
                @blur="saveNameUser"
              />
              <Button label="Ok" @click="saveNameUser" />
            </div>
          </div>

          <!-- Équipe Favorite -->
          <div class="section">
            <h3 class="section-title">
              <i class="pi pi-trophy"></i>
              Équipe favorite
            </h3>
            <div class="form-group">
              <label for="team">Équipe CS:GO favorite</label>
              <Dropdown
                v-model="selectedTeam" 
                :options="csgoTeams" 
                optionLabel="name"
                optionValue="value"
                placeholder="Choisissez votre équipe favorite"
                class="profile-dropdown"
                @change="saveTeamPreference"
              >
                <template #value="{ value }">
                  <div v-if="value" class="team-option">
                    <!-- Img pour l'image de la team-->
                    <span>{{ getTeamName(value) }}</span>
                  </div>
                  <span v-else>{{ "Choisissez votre équipe favorite" }}</span>
                </template>
                <template #option="{ option }">
                  <div class="team-option">
                    <span>{{ option.name }}</span>
                  </div>
                </template>
              </Dropdown>
            </div>
            
            <div v-if="selectedTeam" class="favorite-team-display">
              <div class="team-card">
                <!-- Img pour l'image de la team-->
                <div class="team-info">
                  <h4>{{ getTeamName(selectedTeam) }}</h4>
                  <p>Votre équipe favorite</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Section Investissements -->
          <div class="section">
            <h3 class="section-title">
              <i class="pi pi-wallet"></i>
              Portefeuille & Investissements
            </h3>
            
            <div class="wallet-stats">
              <div class="stat-card">
                <div class="stat-icon">
                  <i class="pi pi-money-bill"></i>
                </div>
                <div class="stat-content">
                  <h4>Balance</h4>
                  <p class="stat-value">{{ walletStore.balance }} ETH</p>
                </div>
              </div>
              
              <div class="stat-card">
                <div class="stat-icon">
                  <i class="pi pi-chart-line"></i>
                </div>
                <div class="stat-content">
                  <h4>Paris totaux</h4>
                  <p class="stat-value">{{ userStats.totalBets }}</p>
                </div>
              </div>
              
              <div class="stat-card">
                <div class="stat-icon">
                  <i class="pi pi-percentage"></i>
                </div>
                <div class="stat-content">
                  <h4>Taux de réussite</h4>
                  <p class="stat-value">{{ userStats.winRate }}%</p>
                </div>
              </div>
              
              <div class="stat-card">
                <div class="stat-icon">
                  <i class="pi pi-star-fill"></i>
                </div>
                <div class="stat-content">
                  <h4>Gains totaux</h4>
                  <p class="stat-value positive">+{{ userStats.totalWinnings }} ETH</p>
                </div>
              </div>
            </div>
          </div>

        </div>
      </template>
    </Card>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useWalletStore } from '@/stores/useWalletStore'
import Card from 'primevue/card'
import InputText from 'primevue/inputtext'
import Dropdown from 'primevue/dropdown'
import { Button } from 'primevue'
import { updateUserPseudo } from '@/db/queries'

const walletStore = useWalletStore()

const userPseudo = ref('')
const selectedTeam = ref('')

// Use with Supabase API.
const userStats = ref({
  totalBets: 0,
  winRate: 0,
  totalWinnings: 0
})

// recents Bets user with Supabase.
const recentBets = ref([
  //
])

// CSGO teams
const csgoTeams = ref([
  { name: 'Astralis', value: 'astralis'},
  { name: 'FaZe Clan', value: 'faze'},
  { name: 'G2 Esports', value: 'g2'},
  { name: 'NAVI', value: 'navi'},
  { name: 'Team Liquid', value: 'liquid'},
  { name: 'Vitality', value: 'vitality'},
  { name: 'MOUZ', value: 'mouz'},
  { name: 'Spirit', value: 'spirit'}
])

// Computed
const avatarLabel = computed(() => {
  if (userPseudo.value) {
    return userPseudo.value.substring(0, 2).toUpperCase()
  }
  return walletStore.connectedAccount ? walletStore.connectedAccount.substring(2, 4).toUpperCase() : 'U'
})

const formatWalletAddress = (address: string) => {
  if (!address) return ''
  return `${address.substring(0, 6)}...${address.substring(address.length - 4)}`
}

const getTeamName = (value: string) => {
  return csgoTeams.value.find(team => team.value === value)?.name || ''
}

const saveNameUser = async () => {
  if (userPseudo.value.trim()) {
    try {
      console.log('Save nameUser:', userPseudo.value)
      await updateUserPseudo(walletStore.connectedAccount, userPseudo.value)
    } catch (error) {
      console.error('Error to save nameUser:', error)
    }
  }
}

const saveTeamPreference = async () => {
  try {
    console.log('Sauvegarde de l\'équipe favorite:', selectedTeam.value)
  } catch (error) {
    console.error('Erreur lors de la sauvegarde de l\'équipe favorite:', error)
  }
}

// Chargement des données utilisateur
const loadUserData = async () => {
  try {
    //Charger les données depuis Supabase
    console.log('Chargement des données utilisateur pour:', walletStore.connectedAccount)
    
    // Exemple de données mockées
    userStats.value = {
      totalBets: 0,
      winRate: 0,
      totalWinnings: 0
    }
    //Donner à charger dès qu'il y a un total, winrate etc.
  } catch (error) {
    console.error('Error loading data:', error)
  }
}

// Lifecycle
onMounted(() => {
  if (walletStore.isConnected) {
    loadUserData()
  }
})
</script>

<style scoped>

.form-group-inline {
  display: flex;
  align-items: center; /* aligne verticalement le bouton et l’input */
  gap: 0.5rem; /* espace entre les deux */
}

.form-group-inline :deep(.profile-input) {
  flex: 1; /* pour que le champ prenne toute la largeur dispo */
}

.profil-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem;
}

.profil-card {
  background: rgba(15, 23, 42, 0.95) !important;
  border: 1px solid rgba(234, 88, 12, 0.3);
  box-shadow: 0 8px 32px rgba(234, 88, 12, 0.1);
}

.card-header {
  display: flex;
  align-items: center;
  gap: 1.5rem;
  padding: 2rem;
  background: linear-gradient(135deg, rgba(234, 88, 12, 0.1), rgba(249, 115, 22, 0.05));
  border-bottom: 1px solid rgba(234, 88, 12, 0.2);
}

.user-avatar {
  background: linear-gradient(135deg, #ea580c, #f97316) !important;
  color: white !important;
  font-weight: 700;
}

.user-info h2 {
  color: white;
  margin: 0;
  font-size: 1.8rem;
  font-weight: 700;
}

.wallet-address {
  color: #94a3b8;
  font-family: monospace;
  font-size: 0.9rem;
  margin: 0.5rem 0 0 0;
}

.profile-content {
  padding: 0;
}

.section {
  padding: 2rem;
  border-bottom: 1px solid rgba(71, 85, 105, 0.3);
}

.section:last-child {
  border-bottom: none;
}

.section-title {
  color: #ea580c;
  font-size: 1.3rem;
  font-weight: 600;
  margin-bottom: 1.5rem;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group label {
  display: block;
  color: #e2e8f0;
  font-weight: 500;
  margin-bottom: 0.5rem;
}

.profile-input, .profile-dropdown {
  width: 100%;
}

:deep(.profile-input) {
  background: rgba(30, 41, 59, 0.8) !important;
  border: 1px solid rgba(71, 85, 105, 0.5) !important;
  color: white !important;
}

:deep(.profile-input:focus) {
  border-color: #ea580c !important;
  box-shadow: 0 0 0 1px #ea580c !important;
}

:deep(.profile-dropdown) {
  background: rgba(30, 41, 59, 0.8) !important;
  border: 1px solid rgba(71, 85, 105, 0.5) !important;
  color: white !important;
}

:deep(.profile-dropdown:focus) {
  border-color: #ea580c !important;
  box-shadow: 0 0 0 1px #ea580c !important;
}

.team-option {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.team-logo-small {
  width: 24px;
  height: 24px;
  object-fit: contain;
}

.team-logo-large {
  width: 64px;
  height: 64px;
  object-fit: contain;
}

.favorite-team-display {
  margin-top: 1.5rem;
}

.team-card {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1.5rem;
  background: linear-gradient(135deg, rgba(234, 88, 12, 0.1), rgba(249, 115, 22, 0.05));
  border: 1px solid rgba(234, 88, 12, 0.3);
  border-radius: 0.75rem;
}

.team-info h4 {
  color: white;
  margin: 0;
  font-size: 1.2rem;
  font-weight: 600;
}

.team-info p {
  color: #94a3b8;
  margin: 0.25rem 0 0 0;
}

.wallet-stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
  margin-bottom: 2rem;
}

.stat-card {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1.5rem;
  background: rgba(30, 41, 59, 0.6);
  border: 1px solid rgba(71, 85, 105, 0.3);
  border-radius: 0.75rem;
  transition: all 0.3s ease;
}

.stat-card:hover {
  border-color: rgba(234, 88, 12, 0.5);
  transform: translateY(-2px);
}

.stat-icon {
  width: 48px;
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #ea580c, #f97316);
  border-radius: 50%;
  color: white;
  font-size: 1.2rem;
}

.stat-content h4 {
  color: #e2e8f0;
  margin: 0;
  font-size: 0.9rem;
  font-weight: 500;
}

.stat-value {
  color: white;
  margin: 0.25rem 0 0 0;
  font-size: 1.4rem;
  font-weight: 700;
}

.stat-value.positive {
  color: #10b981;
}

.recent-activity h4 {
  color: white;
  margin-bottom: 1rem;
  font-size: 1.1rem;
  font-weight: 600;
}

.empty-state {
  text-align: center;
  padding: 2rem;
  color: #94a3b8;
}

.empty-state i {
  font-size: 3rem;
  margin-bottom: 1rem;
  color: rgba(234, 88, 12, 0.5);
}

.activity-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.activity-item {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1rem;
  background: rgba(30, 41, 59, 0.4);
  border: 1px solid rgba(71, 85, 105, 0.2);
  border-radius: 0.5rem;
}

.activity-icon {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: rgba(234, 88, 12, 0.2);
  color: #ea580c;
}

.activity-content {
  flex: 1;
}

.activity-content h5 {
  color: white;
  margin: 0;
  font-size: 1rem;
  font-weight: 600;
}

.activity-content p {
  color: #94a3b8;
  margin: 0.25rem 0 0 0;
  font-size: 0.9rem;
}
</style>
<template>
<div class="matches-container">
    <div class="matches-header">
      <h1 class="page-title">
        <i class="pi pi-wallet"></i>
        Bets
      </h1>
      <router-link to="#" @click.prevent="goBack" class="page-Backtitle">
        <h2>
            <i class="pi pi-arrow-left"></i>
            Back to Matches
        </h2>
      </router-link>
    </div>
    <div>
        <div class="matchesBets-container">
          <!-- Team 1 -->
          <div v-if="matches && firstOpponent" class="team-centered">
            <div class="team-logo-container">
              <img
              v-if="matchesStore.getTeamImageUrl(firstOpponent)"
              :src="matchesStore.getTeamImageUrl(firstOpponent)"
              :alt="firstOpponent.name"
              class="team-logo"
              @error="handleImageError"
              @load="handleImageLoad"
              >
              <div v-else class="team-fallback">
                {{ matchesStore.getTeamInitials(firstOpponent.name) }}
              </div>
              </div>
              <div class="team-info-centered">
                <h3 class="team-name-centered">{{ firstOpponent.name }}</h3>
              </div>
          </div>

          <!-- Score Match-->
          <div>
            {{  }}
          </div>

          <!-- Team 2 -->
          <div v-if="matches && firstOpponent" class="team-centered">
            <div v-if="matches && secondOpponent" class="team-logo-container">
              <img
              v-if="matchesStore.getTeamImageUrl(secondOpponent)"
              :src="matchesStore.getTeamImageUrl(secondOpponent)"
              :alt="secondOpponent.name"
              class="team-logo"
              @error="handleImageError"
              @load="handleImageLoad"
              >
              <div v-else class="team-fallback">
                {{ matchesStore.getTeamInitials(secondOpponent.name) }}
              </div>
              </div>
              <div  v-if="secondOpponent?.name" class="team-info-centered">
                <h3 class="team-name-centered">{{ secondOpponent.name }}</h3>
              </div>
            </div>
        </div>

        <div class="matchesBets-container">
          <div class="button-container">
            <button class="button-bets">Parier</button>
          </div>
        </div>
    </div>
</div>
</template>


<script setup lang="ts">
import { matchStore } from '@/stores/matchStore'
import { computed, ref } from 'vue'
import { storeToRefs } from 'pinia'
import { useRoute, useRouter } from 'vue-router'

import 'primeicons/primeicons.css'

const route = useRoute()
const router = useRouter()

const MatchId = ref(Number(route.params.id))
const matchesStore = matchStore()
const failedImages = ref<Set<string>>(new Set())

const { matches } = storeToRefs(matchesStore)

const specificMatch = computed(() => {
  return matches.value.find(match => match.id === MatchId.value)
})
const firstOpponent = computed(() => specificMatch.value?.opponents?.[0]?.opponent || null)
const secondOpponent = computed(() => specificMatch.value?.opponents?.[1]?.opponent || null)


const handleImageError = (event: Event) => {
  const img = event.target as HTMLImageElement
  if (img.src) failedImages.value.add(img.src)
  img.src = ''
}

const handleImageLoad = (event: Event) => {
  (event.target as HTMLImageElement).style.opacity = '1'
}

const goBack = () => {
  router.go(-1)
}

</script>


<style scoped>

@import "../../components/matches.css";

.page-Backtitle {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: rgba(249, 115, 22, 0.3);
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: medium;
  margin: 1rem 0;
  text-decoration: none;
}

.page-Backtitle:hover {
  color: #fb923c64;
  transform: translateX(-2px);
}

.page-Backtitle i {
  font-size: 1.2rem;
}

.matchesBets-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 3rem;
  padding: 2.5rem;
  background: rgba(26, 26, 26, 0.85);
  border: 1px solid rgba(249, 115, 22, 0.3);
  border-radius: 2rem;
  backdrop-filter: blur(25px);
  box-shadow: 
    0 8px 32px rgba(249, 115, 22, 0.15),
    0 0 0 1px rgba(255, 255, 255, 0.05),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
  position: relative;
  overflow: hidden;
  transition: all 0.4s ease;
  z-index: 10;
  gap: 2rem;
}

.matchesBets-container:hover {
  transform: translateY(-2px);
  box-shadow: 
    0 12px 40px rgba(249, 115, 22, 0.2),
    0 0 0 1px rgba(249, 115, 22, 0.2),
    inset 0 1px 0 rgba(255, 255, 255, 0.15);
}

.team-logo-container {
  width: 3.5rem;
  height: 3.5rem;
  border-radius: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
}

.team-logo {
  width: 80%;
  height: 80%;
  object-fit: cover;
  opacity: 0;
  transition: all 0.4s ease;
  filter: brightness(1.1) contrast(1.05);
  border-radius: 0.5rem;
}

.team-centered {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
  flex: 1;
  min-width: 0;
  transition: all 0.3s ease;
}

.team-centered:hover {
  transform: scale(1.02);
}

.team-info-centered {
  text-align: center;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.team-fallback {
  font-size: 1rem;
  font-weight: 800;
  color: #ffffff;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, rgba(249, 115, 22, 0.3), rgba(251, 146, 60, 0.2));
  border-radius: 0.5rem;
}

.team-name-centered {
  font-size: 1.1rem;
  font-weight: 700;
  color: #ffffff;
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  transition: all 0.3s ease;
  text-overflow: unset;
  white-space: nowrap;
  overflow: visible;
  text-overflow: ellipsis;
}

.team-centered:hover .team-name-centered {
  color: #fb923c;
}

.button-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%; 
}

.button-bets {
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

.button-bets {
  position: relative;
  overflow: hidden;
}

.button-bets::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.6s ease;
}

.button-bets:hover::before{
  left: 100%;
}

.button-bets {
  background: linear-gradient(135deg, #f97316 0%, #ea580c 100%);
  color: #ffffff;
  box-shadow: 0 4px 20px rgba(249, 115, 22, 0.3);
}

.button-bets:hover {
  transform: translateY(-3px) scale(1.05);
  box-shadow: 0 12px 40px rgba(249, 115, 22, 0.5);
  background: linear-gradient(135deg, #ea580c 0%, #dc2626 100%);
}


</style>
<template>
    <div v-if="firstOpponent && secondOpponent" @click="openDetailsPopup">
      <h1 class="page-title-card">Match of the Day</h1>
    <Card class="tournament-card">
        <template #title>
            <div class="tournament-title">
                <span class="league-name"> {{ props.match?.league?.name }}</span>
                <span class="dot">•</span>
                <span class="round-label">{{ props.match?.tournament?.name }}</span>
            </div>
        </template>

        <template #content>
            <Card class="motd-card">
                <template #content>
                  <div class="teams-sectionOfTheDay">
                        <!-- Équipe 1 -->
                    <div class="team-centered">
                            <div class="team-logo-container" v-if="firstOpponent">
                              <img
                                v-if="matchesStore.getTeamImageUrl(firstOpponent)"
                                :src="matchesStore.getTeamImageUrl(firstOpponent)"
                                :alt="firstOpponent.name"
                                class="team-logo"
                                @error="handleImageError"
                                @load="handleImageLoad"
                              />
                              <div v-else class="team-fallback">
                                {{ matchesStore.getTeamInitials(firstOpponent.name) }}
                              </div>
                            </div>
                            <h3 v-if="firstOpponent" class="team-name-centered">{{ firstOpponent.name }}</h3>
                          </div>

                          <!-- VS -->
                          <div class="vs-text">VS</div>

                          <!-- Équipe 2 -->
                          <div class="team-centered" v-if="secondOpponent">
                            <div class="team-logo-container">
                              <img
                                v-if="matchesStore.getTeamImageUrl(secondOpponent)"
                                :src="matchesStore.getTeamImageUrl(secondOpponent)"
                                :alt="secondOpponent.name"
                                class="team-logo"
                                @error="handleImageError"
                                @load="handleImageLoad"
                              />
                              <div v-else class="team-fallback">
                                {{ matchesStore.getTeamInitials(secondOpponent.name) }}
                              </div>
                            </div>
                            <h3 class="team-name-centered">{{ secondOpponent.name }}</h3>
                          </div>
                        </div>

                    <div class="center-col">
                        <div class="score">
                            <span class="score-name">{{ "Score" }}</span>
                            <span class="sep">-</span>
                            <span class="score-name">{{ "Score" }}</span>
                        </div>
                    </div>
                  <div class="odds">
                    <button class="odd-pill" >
                      <div class="odd-title">{{ firstOpponent?.name }}</div>
                      <div class="odd-value">{{ "Cote 1" }}</div>
                    </button>

                    <button class="odd-pill warning" >
                      <span class="corner-badge">{{ "!!" }}</span>
                      <div class="odd-title">{{ secondOpponent?.name }}</div>
                      <div class="odd-value">{{ "Cote 2" }}</div>
                    </button>
                  </div>
                      </template>
                  </Card>
              </template>
          </Card>
          <Divider/>
          </div>

          <!-- Match Details Popup -->
  <MatchDetailsPopup
    v-if="props.match"
    :visible="showDetailsPopup" 
    :match="props.match" 
    @close="closeDetailsPopup"
    @openBetting="openBettingFromDetails"
  />
</template>

<script setup lang="ts">
import { matchStore } from '@/stores/matchStore';
import { storeToRefs } from 'pinia';

import Card from 'primevue/card';
import Divider from 'primevue/divider';
import 'primeicons/primeicons.css';

const showDetailsPopup = ref(false)

import { computed,ref, type UnwrapRef } from 'vue';
import MatchDetailsPopup from './MatchDetailsPopup.vue';
import { useRouter } from 'vue-router'

const router = useRouter()

const failedImages = ref<Set<string>>(new Set());

const matchesStore = matchStore();

let Url = matchesStore.createUrlMatches("current", "");

const { matches } = storeToRefs(matchesStore)

type MatchesArray = UnwrapRef<typeof matchesStore.matches>
type ArrayElement<T> = T extends (infer U)[] ? U : never
type MatchFromStore = ArrayElement<MatchesArray>  

interface Props {
  match?: MatchFromStore | null
}

const props = defineProps<Props>();


const firstOpponent = computed(() => props.match?.opponents?.[0]?.opponent || null)
const secondOpponent = computed(() => props.match?.opponents?.[1]?.opponent || null)

const emit = defineEmits<{
  openBetting: [match: MatchFromStore]
}>()

const handleImageError = (event: Event) => {
  const img = event.target as HTMLImageElement
  if (img.src) failedImages.value.add(img.src)
  img.src = ''
}

const handleImageLoad = (event: Event) => {
  (event.target as HTMLImageElement).style.opacity = '1'
}

const closeDetailsPopup = () => {
  showDetailsPopup.value = false
  router.replace({ hash: '' })
}

const openDetailsPopup = () => {
  showDetailsPopup.value = true
  if (props.match) {
    router.replace({ hash: `#match-${props.match.id}` })
  }
}

const openBettingFromDetails = () => {
  closeDetailsPopup()
  if (props.match) {
    emit('openBetting', props.match)
  }
}

const openPopupFromHash = () => {
  showDetailsPopup.value = true
}

defineExpose({
  openPopupFromHash
})

</script>

<style scoped>

@import "../styles/matches.css";
@import "../styles/team.css";


/* ----- Card externe (Tournoi) ----- */
.tournament-card {
  background: rgba(26, 26, 26, 0.85);
  border: 1px solid rgba(249, 115, 22, 0.3);
  border-radius: 16px;
  padding: 12px;
  transition: all 0.4s ease;
  z-index: 10;
}

.tournament-card:hover {
  transform: translateY(-2px);
  box-shadow: 
    0 12px 40px rgba(249, 115, 22, 0.2),
    0 0 0 1px rgba(249, 115, 22, 0.2),
    inset 0 1px 0 rgba(255, 255, 255, 0.15);
}

.tournament-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;
}

.page-title-card {
  display: flex;
  align-items: center;
  gap: 1rem;
  font-size: 1.5rem;
  font-weight: 800;
  margin: 1rem 0;
  background: linear-gradient(135deg, #f97316, #fb923c, #fbbf24);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  position: relative;
  z-index: 2;
  animation: glow-text 3s ease-in-out infinite alternate;
  justify-content: center;
}

.tournament-card::after {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  animation: rotate-slow 20s linear infinite;
  pointer-events: none;
  opacity: 0.5;
}

.tournament-title {
  display: flex; align-items: center; gap: 8px;
  font-weight: 700;
}
.dot { 
    opacity: .5; 
}
.round-label { 
    opacity: .85; 
}

.vs-text {
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

/* ----- Card interne (Match) ----- */
.motd-card {
  background: rgba(26, 26, 26, 0.85);
  border-radius: 16px;
  border: 1px solid rgba(255,255,255,.06);
  box-shadow: 
  0 8px 32px rgba(0,0,0,.35), 
  inset 0 1px 0 rgba(255,255,255,.06);

  padding: 12px;
  margin-top: 8px;
}

/* Header compact du match */
.motd-top {
  display: flex; align-items: center; 
  justify-content: space-between;
  padding: 6px 6px 10px;
  border-bottom: 1px solid rgba(255,255,255,.06);
  margin-bottom: 10px;
}

.motd-top {
  display: flex;
  justify-content: center; 
  align-items: center;     
  margin-bottom: 1rem;
  gap: 5rem;               
}

.motd-top-left,
.motd-top-right {
  display: flex;
  flex-direction: column;  
  align-items: center;
  text-align: center;
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

.motd-top-left { 
    display: inline-flex; 
    align-items: center; 
    gap: 8px;
    width: 3.5rem;
    height: 3.5rem;
    border-radius: 1rem;
    background: rgba(45, 45, 45, 0.9);
    justify-content: center;
    border: 1px solid rgba(249, 115, 22, 0.3);
    overflow: hidden;
    flex-shrink: 0;
    backdrop-filter: blur(15px);
    transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
    box-shadow: 
        0 4px 15px rgba(0, 0, 0, 0.2),
        inset 0 1px 0 rgba(255, 255, 255, 0.1);
    position: relative;
}

.motd-top-left::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(249, 115, 22, 0.1), transparent);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.motd-top-left:hover {
  border-color: #f97316;
  transform: scale(1.1) rotate(2deg);
  box-shadow: 
    0 8px 25px rgba(249, 115, 22, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.2);
}

.motd-top-left:hover::before {
  opacity: 1;
}

.motd-top-right::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(249, 115, 22, 0.1), transparent);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.motd-top-right:hover {
  border-color: #f97316;
  transform: scale(1.1) rotate(2deg);
  box-shadow: 
    0 8px 25px rgba(249, 115, 22, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.2);
}

.motd-top-right:hover::before {
  opacity: 1;
}

.league-flag { width: 18px; height: 18px; border-radius: 50%; object-fit: cover; }
.league-name { font-weight: 600; }
.motd-top-right {
    display: inline-flex; 
    align-items: center; 
    gap: 6px;
    width: 3.5rem;
  height: 3.5rem;
  border-radius: 1rem;
  background: rgba(45, 45, 45, 0.9);
  justify-content: center;
  border: 1px solid rgba(249, 115, 22, 0.3);
  overflow: hidden;
  flex-shrink: 0;
  backdrop-filter: blur(15px);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 
    0 4px 15px rgba(0, 0, 0, 0.2),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
  position: relative;
}
.pill {
  font-size: .75rem; padding: 4px 8px; border-radius: 999px;
  border: 1px solid rgba(255,255,255,.1); background: rgba(255,255,255,.04);
}
.pill.danger { border-color: rgba(239,68,68,.35); }

/* Corps 3 colonnes */
.motd-body {
  display: grid; 
  grid-template-columns: 1fr auto 1fr;
  align-items: center; 
  gap: 12px; 
  padding: 6px 4px 14px;
}

.team-col { 
    display: grid; 
    justify-items: center; 
    gap: 6px; 
}

.score .score-name { 
  font-size: 1.25rem; 
  font-weight: 800; 
}

.score .sep { 
  margin: 0 6px; 
  opacity: .7; 
  font-weight: 800; 
}

.chrono { 
  display: inline-flex; 
  align-items: center; 
  gap: 6px; 
  font-size: .8rem; opacity: .8; 
}

.center-col { 
  display: grid; 
  justify-items: center; 
  gap: 6px; 
  min-width: 120px; 
}

.chrono.live { 
  color: #ef4444; 
  opacity: 1; 
}

/* Cotes */
.odds { 
  display: grid; 
  grid-template-columns: repeat(2, 1fr); 
  gap: 8px; 
  padding-top: 8px; 
}

.odd-pill {
  position: relative;
  border: 0;
  border-radius: 14px; 
  padding: 10px 12px;
  background: #0e1013;
  outline: 1px solid rgba(255,255,255,.06);
  box-shadow: 0 4px 12px rgba(0,0,0,.35), inset 0 1px 0 rgba(255,255,255,.06);
  text-align: center;
  cursor: pointer;
  transition: transform .12s ease,
  outline-color .12s ease;
}


.odd-pill:hover { 
  transform: translateY(-3px) scale(1.05);
  box-shadow: 0 12px 40px rgba(249, 115, 22, 0.5);
  background: linear-gradient(135deg, #ea580c 0%, #dc2626 100%);
}
.odd-pill .odd-title { 
    font-size: .8rem; 
    opacity: .8; 
}
.odd-pill .odd-value { 
    font-size: 1.1rem; 
    font-weight: 800; 
}

.odd-pill.warning { 
    outline-color: rgba(239,68,68,.35); 
}

.corner-badge {
  position: absolute; 
  top: -6px; 
  right: -6px;
  background: #ef4444; 
  color: #fff; 
  font-size: .7rem; 
  font-weight: 800;
  padding: 2px 6px; 
  border-radius: 8px; 
  box-shadow: 0 2px 6px rgba(239,68,68,.4);
}

/* Responsive */
@media (max-width: 768px) {
  .motd-body { 
    grid-template-columns: 1fr; 
  }


  .odds { 
    grid-template-columns: 1fr; 
  }
}
</style>
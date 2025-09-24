<template>
  <div>
    <Card class="bet-card">
      <template #title>
        <div class="betHistory-header">
          <div class="betHistory-league"> {{ "ESL Pro Division" }}</div>
          <div class="betHistory-meta" 
          :class="{ open: isOpen, closed: !isOpen && !props.bet.resolved, finish: !isOpen && props.bet.resolved }">
            {{ GetStatusHistory() }}
          </div>
        </div>
      </template>

      <template #subtitle>
      </template>

      <template #content>
        <Card class="betHistory-card">

          <template #content>
            <div class="betHistory-section">
                <div class="betHistory-team-container">
                  <div class="betHistory-team" v-if="props.bet.team1Name">
                    <img
                      v-if="props.bet.team1Logo"
                      :src="props.bet.team1Logo"
                      :alt="props.bet.team1Name"
                      class="team-logo"
                      @error="handleImageError"
                      @load="handleImageLoad">
                  </div>
                  <h3 class="betHistory-team-name">{{ props.bet.team1Name }}</h3>
                </div>

                <div class="betHistory-deadline" :class="{ open: isOpen, closed: !isOpen }">
                  <template v-if="isOpen">
                    <div class="date">{{ datePart }}</div>
                    <div class="hour">{{ timePart }}</div>
                  </template>
                  <template v-else>
                    <div class="finish"> {{ "Terminé" }}</div>
                  </template>
                </div>

                <div class="betHistory-team-container">
                  <div class="betHistory-team" v-if="props.bet.team2Name">
                    <img
                      v-if="props.bet.team2Logo"
                      :src="props.bet.team2Logo"
                      :alt="props.bet.team2Name"
                      class="team-logo"
                      @error="handleImageError"
                      @load="handleImageLoad">
                  </div>
                  <h3 class="betHistory-team-name">{{ props.bet.team2Name }}</h3>
                </div>
            </div>
            <div class="betHistory-pools">
              <Button class="button" :label="t1EthLabel" />
              <!--Barre de progression-->
              <div class="bar">
                <div class="fill" :style="{ width: pcts.t1 + '%' }"></div>
              </div>
              <Button class="button" :label="t2EthLabel" />
            </div>

          </template>
        </Card>
        <div class="betHistory-description">
          {{ props.bet.description }}
        </div>
        <div class="betHistory-creator">
           {{ truncateAddress(props.bet.creator) }}
        </div>
      </template>
    </Card>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue'
import Card from 'primevue/card'
import { Button } from 'primevue'
import type { BetHistoryVM } from '@/views/HistoryBets.vue'

const props = defineProps<{
    bet: BetHistoryVM
}>()

const failedImages = ref<Set<string>>(new Set());

const isOpen = computed(() => Date.now() < props.bet.deadline.getTime())
const datePart = computed(() =>
  new Intl.DateTimeFormat('fr-FR', {month: 'long',day: 'numeric'}).format(props.bet.deadline)
)
const timePart = computed(() =>
  new Intl.DateTimeFormat('fr-FR', { timeStyle: 'short' }).format(props.bet.deadline)
)

function truncateAddress(creator: string): string {
  return creator.slice(0, 6) + '...' + creator.slice(-4)
}

function GetStatusHistory(): string {
  if (isOpen.value) {
    return "Open"
  }
  else if (!isOpen.value && !props.bet.resolved) {
    return "Fermé"
  } else {
    return "Terminé"
  }
}

const Pool1ETH = computed(() => props.bet.team1PoolEth.toString())
const Pool2ETH = computed(() => props.bet.team2PoolEth.toString())

const NumberFormat = new Intl.NumberFormat('fr-FR', { maximumFractionDigits: 3})
const t1EthLabel = computed(() => NumberFormat.format(Number(Pool1ETH.value)) + " ETH")
const t2EthLabel = computed(() => NumberFormat.format(Number(Pool2ETH.value)) + " ETH")

const pcts = computed(() => {
  const t1 = BigInt(props.bet.team1PoolEth as string || '0')
  const t2 = BigInt(props.bet.team2PoolEth as string || '0')
  const tot = t1 + t2

  if (tot === 0n)  {
    return { t1: 50, t2: 50 }
  }

  const p1 = Number((t1 * 100n) / tot)
  return { t1: p1, t2: 100 - p1 }
})


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

@import "../styles/team.css";

.bet-card {
  max-width: 640px;
  max-height: 400px;
  margin: 13px auto;
  padding: 12px;
  background: rgba(26, 26, 26, 0.85);
  border: 1px solid rgba(249, 115, 22, 0.3);
  border-radius: 16px;
  padding: 12px;
  transition: all 0.4s ease;
  z-index: 10;
  overflow: hidden;
  box-sizing: border-box;
  box-shadow: 0 2px 8px rgba(0,0,0,.08);
  position: relative;
}

.bet-card:hover {
  transform: translateY(-12px) scale(1.03);
  box-shadow: 
    0 10px 20px rgba(249, 115, 22, 0.1),
    0 0 0 1px rgba(249, 115, 22, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
  border-color: #fb923c64;
}

.bet-card::before {
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

.bet-card::after {
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

.bet-card:hover::before {
  opacity: 1;
}

.bet-card:hover::after {
  opacity: 0.3;
}

.betHistory-league {
  font-weight: 700; 
  font-size: 17px; 
  line-height: 18px; 
  white-space: nowrap; 
  overflow: hidden; 
  text-overflow: ellipsis; 
  max-width: 65%;
}

.betHistory-header {
  display:flex; 
  align-items:center; 
  justify-content:space-between; 
  gap:8px;
  margin-bottom: 8px;

}

.betHistory-meta {
  font-size: 16px;
  background: linear-gradient(90deg, #6b7280, #9ca3af, #6b7280);
  background-size: 200% auto;
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent; /** Effet gradient sur le texte: webkit/background-clip: text, webkit-text: transparent */
  animation: pulse-vs 3s linear infinite;
}

.betHistory-meta.open {
  background: linear-gradient(135deg, #34d399, #10b981, #059669);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.betHistory-meta.closed {
  background: linear-gradient(135deg, #fbbf24, #f59e0b, #d97706);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.betHistory-meta.finish {
  background: linear-gradient(135deg, #60a5fa, #3b82f6, #1e40af);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.betHistory-card {
  background: rgba(26, 26, 26, 0.85);
  border-radius: 16px;
  border: 1px solid rgba(255,255,255,.06);
  box-shadow: 
  0 8px 32px rgba(0,0,0,.35), 
  inset 0 1px 0 rgba(255,255,255,.06);
  padding: 12px;
  margin-top: 8px;
  margin-bottom: 1rem;
}

.betHistory-section {
  display: grid;
  grid-template-columns: 1fr auto 1fr;
  align-items: center;
  text-align: center;
  /**Gap: 1rem */
  margin-bottom: 1.5rem;
  flex-direction: row;
  width: 100%;
  align-items: center;
  justify-content: center;
  position: relative;
  z-index: 2;
}

.betHistory-team-container.left  { justify-self: end; }   /* pousse vers le centre */
.betHistory-team-container.right { justify-self: start; }

.betHistory-team-container {
  display: flex;
  flex-direction: column;
  align-items: center; 
  gap: 0.6rem;         
}


.betHistory-team {
  width: 3.5rem;
  height: 3.5rem;
  border-radius: 1rem;
  background: rgba(45, 45, 45, 0.9);
  align-items: center;
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
  display: flex;
    gap: 8px; 
}

.betHistory-team-name {
  text-align: center;

}

.betHistory-team::before {
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

.betHistory-team:hover {
  border-color: #f97316;
  transform: scale(1.1) rotate(2deg);
  box-shadow: 
    0 8px 25px rgba(249, 115, 22, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.2);
}

.betHistory-team:hover::before {
  opacity: 1;
}

.betHistory-pools { 
  display: grid; 
  grid-template-columns: 1fr auto 1fr; /** Three columns */
  gap: .5rem; 
  align-items: center; 
}

.betHistory-pools .button {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0.6rem 1.2rem;
  border: none;
  border-radius: 12px;
  font-weight: 600;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);

  background: linear-gradient(135deg, #f97316 0%, #ea580c 100%);
  color: #ffffff;
  box-shadow: 0 4px 20px rgba(249, 115, 22, 0.3);

  position: relative;
  overflow: hidden;
}

.betHistory-pools .button::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(255, 255, 255, 0.2),
    transparent
  );
  transition: left 0.6s ease;
}

.betHistory-pools .button:hover::before {
  left: 100%;
}

.betHistory-pools .button:hover {
  transform: translateY(-3px) scale(1.05);
  box-shadow: 0 12px 40px rgba(249, 115, 22, 0.5);
  background: linear-gradient(135deg, #ea580c 0%, #dc2626 100%);
}

.betHistory-description,
.betHistory-creator {
  font-size: 0.85rem;
  font-weight: 500;
  text-align: center;
  margin-top: 0.5rem;

  /* Shine */
  background: linear-gradient(
    90deg,
    #9ca3af 0%,
    #e5e7eb 50%,
    #9ca3af 100%
  );
  background-size: 200% auto;
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;

  animation: shine-move 4s linear infinite;
}

.betHistory-creator {
  animation: none !important;
  -webkit-text-fill-color: #9ca3af; 
  font-size: 0.80rem;
}

.betHistory-deadline {
  margin: 0 1.5rem;
  text-align: center;
  flex-shrink: 0;
}

.betHistory-deadline.open {
  filter: drop-shadow(0 0 8px rgba(249,115,22,.5));
  animation: pulse-vs 3s ease-in-out infinite;
}

.betHistory-deadline.open .date,
.betHistory-deadline.open .hour {
  font-weight: 900;
  background: linear-gradient(135deg,#f97316,#fb923c,#fbbf24);
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
}

.betHistory-deadline.open .date {
  font-weight: 700;
  font-size: 1.4rem; 
  line-height: 1.2;
  color: #f97316;
}

.betHistory-deadline.open .hour { 
  font-weight: 600;
  font-size: 1.2rem; 
  line-height: 1.2; 
  opacity: .95;
}

.betHistory-deadline.closed{
  filter: none;
  animation: none;
}
.betHistory-deadline.closed .finish{
  font-size: 1.4rem;
  font-weight: 700;
  color: #6b7280;
  animation: blink 1.5s ease-in-out infinite;
}

@keyframes blink {
  0%, 100% { opacity: 1; transform: scale(1); }
  50%      { opacity: 0.4; transform: scale(1.05); }
}

@keyframes pulse-vs {
  0%,100%{ 
    transform: translateZ(0) scale(1); 
  }
  50%{ transform: translateZ(0) scale(1.02); }
}


@keyframes shine-move {
  0% { background-position: -100% 0; }
  100% { background-position: 200% 0; }
}

</style>

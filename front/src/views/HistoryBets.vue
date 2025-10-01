<template>
    <div>
        <nav class="nav-desktop" :class="{ 'history-nav': isHistoryPage, 'historyUser-nav': !isHistoryPage }">
            <RouterLink class="page-title" :to="{ name: 'BetsHistory' }">All</RouterLink>
            <RouterLink
                class="page-title"
                :key="userLinkKey"
                :to="userLink"
                @click="handleUserClick"
            >
                User
            </RouterLink>
        </nav>

        <div v-if="betsVM.length === 0 && !loading" class="no-bets-message">
            <span class="no-bets-icon">📭</span>
            <h3>No Bets</h3>
            <p>{{ isHistoryPage ? 'No bets have been placed yet.' : 'You have not placed any bets yet.' }}</p>
        </div>

        <div v-if="route.query.error === 'wallet_required'" class="connect-warning">
            <span class="warning-icon">⚠️</span>
            <p>You must connect your wallet to access your bets.</p>
        </div>

        <div v-else class="betHistory-grid">
            <BetHistoryCard
                v-for="b in betsVM"
                :key="b.id"
                :bet="b"
            />
        </div>
    </div>
</template>

<script setup lang="ts">

import BetHistoryCard from '@/components/BetHistoryCard.vue'
import { formatEther } from 'viem'
import { computed, onUnmounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const router = useRouter()
const route = useRoute()
const isHistoryPage = route.name === 'BetsHistory'

const props = withDefaults(defineProps<{
  ethAddress?: string
  ethIsConnected?: boolean
}>(), {
  ethAddress: '',
  ethIsConnected: false
})

const accountWallet = computed(() => props.ethAddress)

console.log("Pour voir === ", props.ethIsConnected)


const userClicked = ref(false)

const userLink = computed(() => {
  return accountWallet.value
    ? { name: 'BetsHistoryUserId', params: { id: accountWallet.value } }
    : { name: 'BetsHistory' }
})

const userLinkKey = computed(() =>
  accountWallet.value ? `user-${accountWallet.value}` : 'all'
)

const handleUserClick = (event: Event) => {
  
  
  event.preventDefault()

  if (!accountWallet.value || !props.ethIsConnected) {
    router.replace({
      name: 'BetsHistory',
      query: { error: 'wallet_required' }
    })
    return
  }

  userClicked.value = true
  router.replace({ 
    name: 'BetsHistoryUserId', 
    params: { id: accountWallet.value } 
  })
}

console.log(`ACCOUNT ===${accountWallet.value}`)

console.log(`USERLINK ===${userLink.value.params}`)

function asString(x: string | string[] | null | undefined): string {
  return Array.isArray(x) ? (x[0] ?? '') : (x ?? '')
}

const userId = computed<string>(() => asString(route.params.id))

const localhost = 'http://localhost:8080/bets/history'

const loading = ref(false)
const error = ref<string>('')
const HistoryType = ref<'all' | 'user'>('all');

let autoRefreshInterval: NodeJS.Timeout | null = null
let currentAbortController: AbortController | null = null

const routeUser = computed(() => `/bets/history/${accountWallet.value}`)

console.log(`La routeUser ===${routeUser.value}`)
const bets = ref<BetHistory[]>([])

const betsVM = computed<BetHistoryVM[]>(() =>
    (bets.value ?? []).map(mapToBetVM)
)

export interface BetHistory {
    betId: number
    bet: {
        description: string,
        team1Id: number,
        team2Id: number,
        deadline: number,
        team1Pool: string,
        team2Pool: string,
        winningTeam: '0' | '1' | '2',
        resolved: false | true,
        creator: string,
        matchId: string

        team1Name?: string
        team2Name?: string
        team1Logo?: string
        team2Logo?: string
    }
    userBet?: {
        amount: string,
        teamChosen: string,
        claimed: false | true 
    }
}

export interface BetHistoryVM {
    id: number
    description: string
    team1Id: number
    team2Id: number
    deadline: Date                 // Date ready to display
    team1PoolEth: string           // ETH formated
    team2PoolEth: string
    winningTeam: 0 | 1 | 2
    resolved: boolean
    creator: string
    matchId: number

    team1Name?: string
    team2Name?: string
    team1Logo?: string
    team2Logo?: string

    user?: {
        amountEth: string
        teamChosen: 1 | 2
        claimed: boolean
    }
}

function mapToBetVM(bets: BetHistory): BetHistoryVM {
    return {
        id: bets.betId,
        description: bets.bet.description,
        team1Id: bets.bet.team1Id,
        team2Id: bets.bet.team2Id,
        deadline: new Date(bets.bet.deadline * 1000),
        team1PoolEth: formatEther(BigInt(bets.bet.team1Pool)),
        team2PoolEth: formatEther(BigInt(bets.bet.team2Pool)),
        winningTeam: parseInt(bets.bet.winningTeam, 10) as 0 | 1 | 2,
        resolved: bets.bet.resolved,
        creator: bets.bet.creator,
        matchId: parseInt(bets.bet.matchId, 10),

        team1Name: bets.bet.team1Name ?? `Team #${bets.bet.team1Id}`,
        team2Name: bets.bet.team2Name ?? `Team #${bets.bet.team2Id}`,
        team1Logo: bets.bet.team1Logo ?? '',
        team2Logo: bets.bet.team2Logo ?? '',

        user: bets.userBet
        ? {
            amountEth: formatEther(BigInt(bets.userBet?.amount)),
            teamChosen: parseInt(bets.userBet.teamChosen, 10) as 1 | 2,
            claimed: bets.userBet.claimed,
        }
        : undefined,
    }
}


const UrlHistory = computed(() => {
    if (userId.value) {
        return `${localhost}/${userId.value}`
    } else {
        return localhost
    }
})

const fetchHistory = async () => {
    try {
        if (currentAbortController) {
            currentAbortController.abort()
        }

        currentAbortController = new AbortController()
        loading.value = true
        error.value = ''

        console.log(`URL History = ${UrlHistory.value}`);
        
        const response = await fetch(UrlHistory.value, {
            method: 'GET',
            signal: currentAbortController.signal,
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            }
        })
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`)
        }

        bets.value = await response.json()
        console.log(`Data received === ${JSON.stringify(bets.value, null, 2)}`)
        console.log("Successfully loaded user history")
        
    } catch (err) {
        console.error('Error fetching matches:', err)
        error.value = err instanceof Error ? err.message : 'Une erreur inconnue est survenue'
    } finally {
        loading.value = false
        currentAbortController = null
    }
}

const startAutoRefresh = () => {
  if (autoRefreshInterval) {
    clearInterval(autoRefreshInterval)
  }
  
  autoRefreshInterval = setInterval(async () => {
    await fetchHistory()
  }, 30000)
}

const stopAutoRefresh = () => {
  if (autoRefreshInterval) {
    clearInterval(autoRefreshInterval)
    autoRefreshInterval = null
  }
}

/**
 * Component lifecycle - fetch matches on mount
 */

onUnmounted(() => {
    stopAutoRefresh()
})

// Watch pour fetch les données quand la route change
watch(userId, async (id) => {
    if (id) {
        HistoryType.value = "user"
        await fetchHistory()
    } else {
        HistoryType.value = "all"
        await fetchHistory()
    }

    // Set up auto-refresh every 30 seconds for live updates
    startAutoRefresh()
}, {immediate: true})

</script>

<style lang="css" scoped>

@import "../styles/navbar.css";

.betHistory-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill);
    gap: 12px;
    align-items: start;
}

.connect-warning {
    display: flex;
    align-items: center;
    gap: 1rem;
    padding: 1rem 1.5rem;
    margin: 1rem 0;
    max-width: 1000px;
    margin: 0 auto 1.5rem;
    background: rgba(249, 115, 22, 0.1);
    border: 1px solid rgba(249, 115, 22, 0.3);
    border-radius: 12px;
    color: #fb923c;
}

.warning-icon {
    font-size: 1.5rem;
}

.connect-warning p {
    margin: 0;
    font-weight: 500;
}

.no-bets-message {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 3rem 2rem;
  margin: 2rem auto;
  max-width: 500px;
  background: rgba(26, 26, 26, 0.6);
  border: 1px solid rgba(249, 115, 22, 0.2);
  border-radius: 16px;
  text-align: center;
  position: relative;
}

.no-bets-message:hover {
  transform: translateY(-12px) scale(1.03);
  box-shadow: 
    0 10px 20px rgba(249, 115, 22, 0.1),
    0 0 0 1px rgba(249, 115, 22, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
  border-color: #fb923c64;
}

.no-bets-message::before {
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

.no-bets-message::after {
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

.no-bets-message:hover::before {
  opacity: 1;
}

.no-bets-message:hover::after {
  opacity: 0.3;
}

.no-bets-icon {
  font-size: 4rem;
  margin-bottom: 1rem;
  opacity: 0.8;
}

.no-bets-message h3 {
  font-size: 1.5rem;
  font-weight: 700;
  color: #f97316;
  margin: 0 0 0.5rem 0;
}

.no-bets-message p {
  color: #9ca3af;
  font-size: 1rem;
  margin: 0;
}

.loading-message {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 3rem 2rem;
  text-align: center;
}

.loading-spinner {
  font-size: 3rem;
  animation: spin 2s linear infinite;
}

.loading-message p {
  color: #9ca3af;
  font-size: 1rem;
  margin-top: 1rem;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

</style>
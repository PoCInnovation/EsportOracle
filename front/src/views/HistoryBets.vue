
<template>
    <div>
        <nav class="nav-desktop" :class="{ 'history-nav': isHistoryPage }">
            <router-link
            v-for="item in navigation"
            :key="item.name"
            :to="item.to">
            <span class="page-title">{{ item.name }}</span>
            </router-link>
        </nav>
        <div class="betHistory-grid">
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
import { useRoute } from 'vue-router'

const route = useRoute()
const isHistoryPage = route.name === 'BetsHistory'

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

const navigation = [
  { name: 'All', to: '/bets/history' },
  { name: 'User',  to: '/bets/history' }
]


const bets = ref<BetHistory[]>([])
console.log('history is', bets.value, Array.isArray(bets.value))
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

/*watch(
  () => route.fullPath,
   async (newPath) => {
    if (newPath) {
        HistoryType.value = "user"
        await fetchHistory()
    } else {
        HistoryType.value = "all"
        await fetchHistory()
    }
    startAutoRefresh()
  },
  { immediate: true }
)*/

</script>

<style lang="css" scoped>

@import "../styles/navbar.css";

.betHistory-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill);
    gap: 12px;
    align-items: start;
}

</style>
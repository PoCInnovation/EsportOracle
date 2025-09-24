import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

const API_BASE_URL = import.meta.env.VITE_API_URL ?? 'http://localhost:8080'
const BETS_ENDPOINT = `${API_BASE_URL}/bets`

export interface ContractBetPayload {
  betId: string
  bet: {
    description: string
    team1Id: string
    team2Id: string
    deadline: string
    team1Pool: string
    team2Pool: string
    winningTeam: number
    resolved: boolean
    creator: string
    matchId: string
  }
  userBet?: {
    amount: string
    teamChosen: number
    claimed: boolean
  }
}

const toBigInt = (value?: string | null) => {
  try {
    if (!value) return 0n
    return BigInt(value)
  } catch {
    return 0n
  }
}

const formatDeadline = (deadline: string) => {
  const seconds = Number(deadline)
  if (!Number.isFinite(seconds) || seconds <= 0) return null
  const date = new Date(seconds * 1000)
  return Number.isNaN(date.getTime()) ? null : date
}

export const betStore = defineStore('bet', () => {
  const bets = ref<ContractBetPayload[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)
  const lastUpdated = ref<Date | null>(null)

  const fetchBets = async () => {
    try {
      loading.value = true
      error.value = null

      const response = await fetch(BETS_ENDPOINT)
      if (!response.ok) {
        throw new Error(`HTTP ${response.status}`)
      }

      const data: ContractBetPayload[] = await response.json()
      bets.value = Array.isArray(data) ? data : []
      lastUpdated.value = new Date()
    } catch (err) {
      console.error('[betStore] fetchBets failed', err)
      error.value = err instanceof Error ? err.message : 'Unknown error'
    } finally {
      loading.value = false
    }
  }

  const refresh = async () => {
    await fetchBets()
  }

  const openBets = computed(() => bets.value.filter(bet => !bet.bet.resolved))
  const resolvedBets = computed(() => bets.value.filter(bet => bet.bet.resolved))

  const getBetByMatchId = (matchId: number | string) => {
    return bets.value.find(bet => Number(bet.bet.matchId) === Number(matchId)) ?? null
  }

  const getTotalPool = (bet: ContractBetPayload) => {
    const team1 = toBigInt(bet.bet.team1Pool)
    const team2 = toBigInt(bet.bet.team2Pool)
    return team1 + team2
  }

  const getDeadline = (bet: ContractBetPayload) => formatDeadline(bet.bet.deadline)

  return {
    bets,
    loading,
    error,
    lastUpdated,
    fetchBets,
    refresh,
    openBets,
    resolvedBets,
    getBetByMatchId,
    getTotalPool,
    getDeadline,
  }
})

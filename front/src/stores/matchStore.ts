import { defineStore } from "pinia";
import { computed, ref } from "vue";

const localhost = 'http://localhost:8080/matches/'

export const matchStore = defineStore('match', () => {
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
            image_url?: string
            }
        }>
        league?: {
            name: string
        }
        tournament?: {
            name: string
        }
    }

    // État séparé pour chaque type de match
    const upcomingMatches = ref<Match[]>([])
    const currentMatches = ref<Match[]>([])
    const pastMatches = ref<Match[]>([])
    
    const loading = ref(false)
    const error = ref<string>('')
    const lastUpdate = ref<Date | null>(null)

    let currentAbortController: AbortController | null = null;
    const currentMatchType = ref<'upcoming' | 'current' | 'past'>('upcoming');

    const matches = computed(() => {
        switch (currentMatchType.value) {
            case 'upcoming':
                return upcomingMatches.value
            case 'current':
                return currentMatches.value
            case 'past':
                return pastMatches.value
            default:
                return []
        }
    })

   const lastUpdated = computed(() => {
    if (!lastUpdate.value) return 'Jamais'
    
    return lastUpdate.value.toLocaleString('fr-FR', {
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit'
    })
   })

   const getMatchById = computed(() => {
    return (id: number) => matches.value.find(match => match.id === id)
  })

   const createUrlMatches = (status: 'upcoming' | 'current' | 'past', teamID: string | undefined): string => {
    let Url: string;

    if (status === "upcoming") {
        if (teamID) {
            Url = `${localhost}upcoming/${teamID}`
        } else {
            Url = `${localhost}upcoming`
        }
    } else if (status === "current") {
        if (teamID) {
            Url = `${localhost}current/${teamID}`
        } else {
            Url = `${localhost}current`
        }
    } else {
        if (teamID) {
            Url = `${localhost}past/${teamID}`
        } else {
            Url = `${localhost}past`
        }
    }
    return Url;
   }

   const fetchMatches = async (Url: string, matchType: 'upcoming' | 'current' | 'past'): Promise<void> => {
    try {
        if (currentAbortController) {
            currentAbortController.abort()
        }

        currentAbortController = new AbortController()
        loading.value = true
        error.value = ''
        currentMatchType.value = matchType

        console.log(`URL = ${Url}`);

        const response = await fetch(Url, {
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

        const data: Match[] = await response.json()

      // Sort matches by status priority (live first, then upcoming, then finished)
      const sortedMatches = data.sort((a, b) => {
      const statusPriority = { 'running': 0, 'live': 0, 'not_started': 1, 'upcoming': 1, 'finished': 2 }
      const aPriority = statusPriority[a.status?.toLowerCase() as keyof typeof statusPriority] ?? 3
      const bPriority = statusPriority[b.status?.toLowerCase() as keyof typeof statusPriority] ?? 3
      
      return aPriority - bPriority
    })

    switch (matchType) {
        case 'upcoming':
                upcomingMatches.value = sortedMatches
                break
        case 'current':
                currentMatches.value = sortedMatches
                break
        case 'past':
                pastMatches.value = sortedMatches
                break
    }
    lastUpdate.value = new Date()
    
    console.log(`Successfully loaded ${data.length} matches`)

        //console.log("Matches:", JSON.stringify(data, null, 2))
    } catch (err) {
            console.error('Error fetching matches:', err)
            error.value = err instanceof Error ? err.message : 'Une erreur inconnue est survenue'
    } finally {
            loading.value = false
            currentAbortController = null
    }
}
   return {
     // État
        matches,
        upcomingMatches,
        currentMatches,
        pastMatches,
        loading,
        error,
        lastUpdate,
        currentMatchType,
        // Actions
        getMatchById,
        createUrlMatches,
        fetchMatches,
        lastUpdated,
   }
})
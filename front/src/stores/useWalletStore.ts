import { defineStore } from "pinia";
import { ref } from "vue";

export const useWalletStore = defineStore('wallet', () => {
  const isConnected = ref(false)
  const connectedAccount = ref('')
  const chainId = ref(null)
  const balance = ref('')

  const setConnectionStatus = (status) => {
    isConnected.value = status
  }
  
  const setWalletData = (data) => {
    isConnected.value = data.isConnected
    connectedAccount.value = data.account || ''
    chainId.value = data.chainId || null
    balance.value = data.balance || ''
  }
  
  const resetWallet = () => {
    isConnected.value = false
    connectedAccount.value = ''
    chainId.value = null
    balance.value = ''
    localStorage.removeItem("wallet-store");
  }
  
  return {
    isConnected,
    connectedAccount,
    chainId,
    balance,
    setConnectionStatus,
    setWalletData,
    resetWallet
  }
}, {
  persist: {
    key: 'wallet-store',
    storage: localStorage
  }
})
<template>
  <div class="walletSection">
    <button 
      @click="handleClick" 
      class="connectButton"
    >
      {{ buttonText }}
    </button>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'

const isConnected = ref(false)
const connectedAccount = ref('')

const buttonText = computed(() => {
  if (isConnected.value) {
    return shortenAddress(connectedAccount.value)
  }
  return typeof (window as any)?.ethereum !== 'undefined' ? 'Connect Wallet' : 'Install MetaMask'
})

const handleClick = async () => {
  if (isConnected.value) {
    const confirmDisconnect = confirm('Are you sure you want to disconnect your wallet?')
    
    if (confirmDisconnect) {
      isConnected.value = false
      connectedAccount.value = ''
      console.log('Wallet disconnected')
    }
    return
  }
  
  const ethereum = (window as any).ethereum

  if (!ethereum) {
    window.open('https://metamask.io/download/', '_blank')
    return
  }

  try {
    if (ethereum.providers && ethereum.providers.length > 0) {
      const metamaskProvider = ethereum.providers.find((provider: any) => provider.isMetaMask)
      if (metamaskProvider) {
        const accounts = await metamaskProvider.request({ 
          method: 'eth_requestAccounts' 
        })
        
        if (accounts && accounts.length > 0) {
          connectedAccount.value = accounts[0]
          isConnected.value = true
          console.log('Wallet connected:', accounts[0])
          return
        }
      }
    }
    
    const accounts = await ethereum.request({ 
      method: 'eth_requestAccounts' 
    })
    
    if (accounts && accounts.length > 0) {
      connectedAccount.value = accounts[0]
      isConnected.value = true
      console.log('Wallet connected:', accounts[0])
    }
    
  } catch (err: any) {
    console.error('Erreur MetaMask:', err)
  }
}

const shortenAddress = (address: string) => {
  if (!address) return ''
  return `${address.slice(0, 6)}...${address.slice(-4)}`
}

onMounted(() => {
  setTimeout(() => {
    const ethereum = (window as any)?.ethereum
    console.log('MetaMask provider:', ethereum);
    console.log('Is MetaMask:', ethereum.isMetaMask);
    console.log('available method:', Object.keys(ethereum));
  }, 1000)
})
</script>

<style scoped>
.walletSection {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.connectButton {
  background-color: #ffffff;
  color: #111317;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 6px;
  font-family: 'DM Sans', sans-serif;
  font-weight: 500;
  font-size: 0.875rem;
  cursor: pointer;
  transition: all 0.2s ease;
}

.connectButton:hover {
  background-color: #f0f0f0;
  transform: translateY(-1px);
}

.connectButton:disabled {
  background-color: #cccccc;
  color: #666666;
  cursor: not-allowed;
  transform: none;
}

.walletInfo {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background-color: #ffffff;
  padding: 0.5rem 0.75rem;
  border-radius: 6px;
}

.addressDisplay {
  color: #111317;
  font-family: 'DM Sans', sans-serif;
  font-weight: 500;
  font-size: 0.875rem;
}

.disconnectButton {
  background: none;
  border: none;
  color: #666666;
  cursor: pointer;
  font-size: 1.2rem;
  font-weight: bold;
  padding: 0;
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  transition: all 0.2s ease;
}

.disconnectButton:hover {
  background-color: #f0f0f0;
  color: #333333;
}
</style>

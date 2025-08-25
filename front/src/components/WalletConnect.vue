<template>
  <div class="walletSection">
    <!-- Connected State -->
    <div v-if="isConnected" class="walletInfo">
      <div class="walletIndicator">
        <div class="connectedDot"></div>
        <span class="addressDisplay">{{ shortenAddress(connectedAccount) }}</span>
        <span class="chainBadge">{{ getChainName(chainId) }}</span>
      </div>
      <button @click="confirmAndDisconnect" class="disconnectButton" title="Disconnect wallet">
        <i class="pi pi-sign-out"></i>
      </button>
    </div>

    <!-- Disconnected State -->
    <div v-else class="connectSection">
      <button
        @click="() => connect()" 
        class="connectButton"
        :disabled="isConnecting"
      >
        <i class="pi pi-wallet"></i>
        {{ isConnecting ? 'Connecting...' : buttonText }}
      </button>
      
      <!-- Wallet Selection Button -->
      <button 
        v-if="hasWalletHistory || hasMultipleWallets"
        @click="toggleWalletSelector" 
        class="selectWalletButton"
        :disabled="isConnecting"
        title="Choose a different wallet"
      >
        <i class="pi pi-th-large"></i>
      </button>
    </div>

    <!-- Wallet Selector Popup -->
    <div v-if="showWalletSelector" class="walletSelectorOverlay" @click="closeWalletSelector">
      <div class="walletSelectorModal" @click.stop>
        <div class="modalHeader">
          <h3>Connect Wallet</h3>
          <button @click="closeWalletSelector" class="closeButton">
            <i class="pi pi-times"></i>
          </button>
        </div>
        
        <div class="walletOptions">
          <button 
            v-for="wallet in availableWallets" 
            :key="wallet.id"
            @click="connectWithWallet(wallet)"
            class="walletOption"
            :disabled="isConnecting"
          >
            <div class="walletIcon">
              <i :class="wallet.icon"></i>
            </div>
            <div class="walletDetails">
              <span class="walletName">{{ wallet.name }}</span>
              <span class="walletStatus">{{ wallet.status }}</span>
            </div>
            <div class="walletAction">
              <i class="pi pi-chevron-right"></i>
            </div>
          </button>
        </div>
        
        <div v-if="error" class="errorMessage">
          <i class="pi pi-exclamation-triangle"></i>
          {{ error }}
        </div>
      </div>
    </div>

    <!-- Global Error Display -->
    <div v-if="error && !showWalletSelector" class="errorMessage">
      <i class="pi pi-exclamation-triangle"></i>
      {{ error }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { createPublicClient, createWalletClient, custom, type WalletClient, type PublicClient, type Address, formatEther } from 'viem'
import { mainnet, sepolia, localhost } from 'viem/chains'
import { useWalletStore } from '@/stores/useWalletStore'
import { AddOrCreateNewUser } from '@/db/queries'

// Store
const walletStore = useWalletStore()

// Component reactive state
const isConnected = ref(false)
const isConnecting = ref(false)
const connectedAccount = ref<Address | ''>('')
const walletClient = ref<WalletClient | null>(null)
const publicClient = ref<PublicClient | null>(null)
const chainId = ref<number | null>(null)
const balance = ref<string>('')
const error = ref<string>('')
const hasWalletHistory = ref(false)
const showWalletSelector = ref(false)

// Props
interface Props {
  buttonText?: string
}
const props = withDefaults(defineProps<Props>(), {
  buttonText: 'Connect Wallet'
})

// Ethereum provider interface
interface EthereumProvider {
  request(args: { method: string; params?: any[] }): Promise<any>
  on(event: string, handler: (...args: any[]) => void): void
  removeListener(event: string, handler: (...args: any[]) => void): void
  isMetaMask?: boolean
}

declare global {
  interface Window {
    ethereum?: EthereumProvider
  }
}

// Available wallets configuration
const availableWallets = computed(() => [
  {
    id: 'metamask',
    name: 'MetaMask',
    icon: 'pi pi-wallet',
    status: 'Browser Extension'
  },
  {
    id: 'walletconnect',
    name: 'WalletConnect',
    icon: 'pi pi-mobile',
    status: 'Mobile & Desktop'
  },
  {
    id: 'coinbase',
    name: 'Coinbase Wallet',
    icon: 'pi pi-credit-card',
    status: 'Browser & Mobile'
  }
])

// Check if multiple wallets are available
const hasMultipleWallets = computed(() => {
  return availableWallets.value.length > 1
})

// Chain configuration
const getCurrentChain = (chainId: number) => {
  switch (chainId) {
    case 1: return mainnet
    case 11155111: return sepolia
    case 31337: return localhost
    default: return mainnet
  }
}

// Provider utilities
const getEthereumProvider = (): EthereumProvider | null => {
  if (typeof window === 'undefined') return null
  
  const ethereum = (window as any).ethereum
  if (!ethereum) return null

  if (ethereum.providers && Array.isArray(ethereum.providers) && ethereum.providers.length > 0) {
    const metamaskProvider = ethereum.providers.find((provider: any) => {
      return provider.isMetaMask && !provider.isBraveWallet && !provider.isCoinbaseWallet
    })
    
    if (metamaskProvider) return metamaskProvider
    
    const stableProvider = ethereum.providers.find((provider: any) => {
      return provider.request && typeof provider.request === 'function'
    })
    
    return stableProvider || ethereum.providers[0]
  }
  
  if (ethereum.request && typeof ethereum.request === 'function') {
    return ethereum
  }
  
  return null
}

// Main wallet connection function
async function connect(forceSelection = false) {
  if (!forceSelection && (hasMultipleWallets.value || hasWalletHistory.value)) {
    showWalletSelector.value = true
    return
  }

  error.value = ''
  const provider = getEthereumProvider()
  
  if (!provider) {
    window.open('https://metamask.io/download/', '_blank')
    return
  }

  try {
    isConnecting.value = true
    
    await new Promise(resolve => setTimeout(resolve, 100))
    
    let accounts: Address[]
    
    if (forceSelection) {
      try {
        await provider.request({
          method: 'wallet_revokePermissions',
          params: [{ eth_accounts: {} }]
        }).catch(() => {})
      } catch {}
      
      accounts = await provider.request({
        method: 'eth_requestAccounts'
      }) as Address[]
    } else {
      const accountsPromise = provider.request({
        method: 'eth_requestAccounts'
      })
      
      const timeoutPromise = new Promise((_, reject) => {
        setTimeout(() => reject(new Error('Connection timeout - please try again')), 30000)
      })
      
      accounts = await Promise.race([accountsPromise, timeoutPromise]) as Address[]
    }
    
    if (!accounts || accounts.length === 0) {
      throw new Error('No accounts returned from wallet')
    }

    let chainIdHex: string
    let retryCount = 0
    const maxRetries = 3
    
    while (retryCount < maxRetries) {
      try {
        chainIdHex = await provider.request({
          method: 'eth_chainId'
        }) as string
        break
      } catch (chainError) {
        retryCount++
        if (retryCount === maxRetries) {
          throw new Error('Failed to get chain ID after multiple attempts')
        }
        await new Promise(resolve => setTimeout(resolve, 500))
      }
    }
    
    const currentChainId = parseInt(chainIdHex!, 16)
    const currentChain = getCurrentChain(currentChainId)
    
    const pubClient = createPublicClient({
      chain: currentChain,
      transport: custom(provider)
    })
    const client = createWalletClient({
      chain: currentChain,
      transport: custom(provider)
    })
    
    walletClient.value = client
    publicClient.value = pubClient
    connectedAccount.value = accounts[0]
    chainId.value = currentChainId
    isConnected.value = true
    
    try {
      await updateBalance()
    } catch (balanceError) {
      console.warn('Failed to fetch balance, but connection successful:', balanceError)
    }
    
    localStorage.setItem('walletConnected', 'true')
    hasWalletHistory.value = true

    // Update Pinia store
    walletStore.setWalletData({
      isConnected: isConnected.value,
      account: connectedAccount.value as string,
      chainId: chainId.value,
      balance: balance.value
    })

    await AddOrCreateNewUser(connectedAccount.value)
    
  } catch (err: any) {
    console.error('Connection error:', err)
    
    if (err.message.includes('User rejected')) {
      error.value = 'Connection cancelled by user'
    } else if (err.message.includes('timeout')) {
      error.value = 'Connection timeout - please try again'
    } else {
      error.value = err.message || 'Failed to connect wallet'
    }
  } finally {
    isConnecting.value = false
  }
}

async function connectWithWallet(wallet: any) {
  try {
    await connect(true) // Force connection
    closeWalletSelector()
  } catch (err) {
    console.error('Failed to connect with wallet:', err)
  }
}

async function disconnect() {
  const provider = getEthereumProvider()

  try {
    await provider?.request({
      method: 'wallet_revokePermissions',
      params: [{ eth_accounts: {} }]
    }).catch(() => {})
  } catch {}

  isConnected.value = false
  connectedAccount.value = ''
  walletClient.value = null
  publicClient.value = null
  chainId.value = null
  balance.value = ''
  error.value = ''
  
  localStorage.removeItem('walletConnected')
  walletStore.resetWallet()
}

async function confirmAndDisconnect() {
  if (confirm('Are you sure you want to disconnect your wallet?')) {
    await disconnect()
  }
}

// Wallet selector methods
function toggleWalletSelector() {
  showWalletSelector.value = !showWalletSelector.value
}

function closeWalletSelector() {
  showWalletSelector.value = false
}

async function forceWalletSelection() {
  showWalletSelector.value = true
}

// Balance update
const updateBalance = async () => {
  if (!publicClient.value || !connectedAccount.value) return
  
  try {
    const balanceWei = await publicClient.value.getBalance({
      address: connectedAccount.value
    })
    balance.value = formatEther(balanceWei)
  } catch (err) {
    console.error('Failed to get balance:', err)
  }
}

// Event handlers
const handleAccountsChanged = (accounts: string[]) => {
  if (accounts.length === 0) {
    disconnect()
  } else if (accounts[0] !== connectedAccount.value) {
    connectedAccount.value = accounts[0] as Address
    updateBalance()
  }
}

const handleChainChanged = (chainIdHex: string) => {
  const newChainId = parseInt(chainIdHex, 16)
  chainId.value = newChainId
  
  if (walletClient.value && connectedAccount.value) {
    const provider = getEthereumProvider()
    if (provider) {
      try {
        const currentChain = getCurrentChain(newChainId)
        
        const pubClient = createPublicClient({
          chain: currentChain,
          transport: custom(provider)
        })
        
        const client = createWalletClient({
          chain: currentChain,
          transport: custom(provider)
        })
        
        publicClient.value = pubClient
        walletClient.value = client
        updateBalance()
      } catch (err) {
        console.error('Failed to recreate clients after chain change:', err)
        error.value = 'Failed to switch network - please reconnect your wallet'
      }
    }
  }
}

// Event listener setup
const setupEventListeners = () => {
  const provider = getEthereumProvider()
  if (!provider) return
  
  provider.on('accountsChanged', handleAccountsChanged)
  provider.on('chainChanged', handleChainChanged)
  provider.on('disconnect', () => {
    disconnect()
  })
}

// Auto-connect check
const checkConnection = async () => {
  const provider = getEthereumProvider()
  if (!provider) return

  const walletConnected = localStorage.getItem('walletConnected')
  if (walletConnected === 'true') {
    hasWalletHistory.value = true
    try {
      const accounts = await provider.request({ method: 'eth_accounts' }) as Address[]
      if (accounts && accounts.length > 0) {
        connectedAccount.value = accounts[0]
        await connect(true)
      } else {
        localStorage.removeItem('walletConnected')
      }
    } catch (err) {
      console.error('Failed to check existing connection:', err)
    }
  } else {
    try {
      const accounts = await provider.request({ method: 'eth_accounts' }) as Address[]
      if (accounts && accounts.length > 0) {
        hasWalletHistory.value = true
      }
    } catch (err) {
      console.error('Failed to check existing connection:', err)
    }
  }
}

// Utility methods
function shortenAddress(address: string): string {
  if (!address) return ''
  return `${address.slice(0, 6)}...${address.slice(-4)}`
}

function getChainName(chainId?: number): string {
  const chainNames: Record<number, string> = {
    1: 'Ethereum',
    137: 'Polygon',
    42161: 'Arbitrum',
    10: 'Optimism',
    1337: 'Local'
  }
  return chainNames[chainId || 1] || 'Unknown'
}

// Lifecycle
onMounted(async () => {
  setupEventListeners()
  await checkConnection()
})

// Watchers
watch(isConnected, (connected) => {
  if (connected) {
    updateBalance()
  }
})

// Expose for parent components
defineExpose({
  isConnected,
  connectedAccount,
  walletClient,
  chainId,
  balance,
  connect,
  disconnect
})
</script>

<style scoped>
@import "./button.css";

.walletSection {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

/* Connected State */
.walletInfo {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.5rem 1rem;
  background: rgba(255, 140, 0, 0.1);
  border: 1px solid rgba(255, 140, 0, 0.3);
  border-radius: 12px;
  backdrop-filter: blur(10px);
}

.walletIndicator {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.connectedDot {
  width: 8px;
  height: 8px;
  background: #00ff88;
  border-radius: 50%;
  box-shadow: 0 0 8px rgba(0, 255, 136, 0.6);
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.6; }
}

.addressDisplay {
  font-family: 'Courier New', monospace;
  font-size: 0.875rem;
  color: #fff;
  font-weight: 500;
}

.chainBadge {
  padding: 0.25rem 0.5rem;
  background: rgba(255, 140, 0, 0.2);
  border: 1px solid rgba(255, 140, 0, 0.4);
  border-radius: 6px;
  font-size: 0.75rem;
  color: #ff8c00;
  font-weight: 500;
}

.disconnectButton {
  padding: 0.5rem;
  background: rgba(255, 69, 69, 0.1);
  border: 1px solid rgba(255, 69, 69, 0.3);
  border-radius: 8px;
  color: #ff4545;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.disconnectButton:hover {
  background: rgba(255, 69, 69, 0.2);
  border-color: rgba(255, 69, 69, 0.5);
  transform: translateY(-1px);
}

/* Disconnected State */
.connectSection {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.connectButton {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1.5rem;
  background: linear-gradient(135deg, #ff8c00 0%, #ff6b35 100%);
  border: none;
  border-radius: 12px;
  color: white;
  font-weight: 600;
  font-size: 0.875rem;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 15px rgba(255, 140, 0, 0.3);
}

.connectButton:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(255, 140, 0, 0.4);
  background: linear-gradient(135deg, #ff9500 0%, #ff7043 100%);
}

.connectButton:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}

.selectWalletButton {
  padding: 0.75rem;
  background: rgba(255, 140, 0, 0.1);
  border: 1px solid rgba(255, 140, 0, 0.3);
  border-radius: 10px;
  color: #ff8c00;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.selectWalletButton:hover {
  background: rgba(255, 140, 0, 0.2);
  border-color: rgba(255, 140, 0, 0.5);
  transform: translateY(-1px);
}

.selectWalletButton:disabled {
  opacity: 0.5;
  cursor: not-allowed;
  transform: none;
}

/* Wallet Selector Popup */
.walletSelectorOverlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.7);
  backdrop-filter: blur(5px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  animation: fadeIn 0.3s ease;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

.walletSelectorModal {
  background: linear-gradient(145deg, 
    rgba(30, 30, 30, 0.95) 0%, 
    rgba(45, 45, 45, 0.95) 100%
  );
  border: 1px solid rgba(255, 140, 0, 0.3);
  border-radius: 20px;
  padding: 2rem;
  min-width: 400px;
  max-width: 500px;
  backdrop-filter: blur(20px);
  box-shadow: 
    0 20px 40px rgba(0, 0, 0, 0.5),
    0 0 50px rgba(255, 140, 0, 0.1);
  animation: slideIn 0.3s ease;
}

@keyframes slideIn {
  from { 
    opacity: 0;
    transform: translateY(-20px) scale(0.95);
  }
  to { 
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.modalHeader {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 1.5rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid rgba(255, 140, 0, 0.2);
}

.modalHeader h3 {
  margin: 0;
  color: #fff;
  font-size: 1.5rem;
  font-weight: 600;
  background: linear-gradient(135deg, #ff8c00 0%, #ff6b35 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.closeButton {
  padding: 0.5rem;
  background: rgba(255, 140, 0, 0.1);
  border: 1px solid rgba(255, 140, 0, 0.3);
  border-radius: 8px;
  color: #ff8c00;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.closeButton:hover {
  background: rgba(255, 140, 0, 0.2);
  border-color: rgba(255, 140, 0, 0.5);
  transform: scale(1.05);
}

/* Wallet Options */
.walletOptions {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.walletOption {
  display: flex;
  align-items: center;
  padding: 1rem 1.25rem;
  background: rgba(255, 140, 0, 0.05);
  border: 1px solid rgba(255, 140, 0, 0.2);
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  width: 100%;
  text-align: left;
}

.walletOption:hover {
  background: rgba(255, 140, 0, 0.1);
  border-color: rgba(255, 140, 0, 0.4);
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(255, 140, 0, 0.2);
}

.walletOption:disabled {
  opacity: 0.5;
  cursor: not-allowed;
  transform: none;
}

.walletIcon {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(255, 140, 0, 0.1);
  border: 1px solid rgba(255, 140, 0, 0.3);
  border-radius: 10px;
  margin-right: 1rem;
  color: #ff8c00;
  font-size: 1.25rem;
}

.walletDetails {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.walletName {
  color: #fff;
  font-weight: 600;
  font-size: 1rem;
}

.walletStatus {
  color: rgba(255, 255, 255, 0.7);
  font-size: 0.875rem;
}

.walletAction {
  color: #ff8c00;
  font-size: 0.875rem;
  opacity: 0.7;
  transition: all 0.3s ease;
}

.walletOption:hover .walletAction {
  opacity: 1;
  transform: translateX(2px);
}

/* Error Messages */
.errorMessage {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1rem;
  background: rgba(255, 69, 69, 0.1);
  border: 1px solid rgba(255, 69, 69, 0.3);
  border-radius: 8px;
  color: #ff4545;
  font-size: 0.875rem;
  margin-top: 1rem;
  animation: shake 0.5s ease;
}

@keyframes shake {
  0%, 100% { transform: translateX(0); }
  25% { transform: translateX(-2px); }
  75% { transform: translateX(2px); }
}

/* Mobile responsiveness */
@media (max-width: 640px) {
  .walletSelectorModal {
    margin: 1rem;
    min-width: unset;
    max-width: unset;
    width: calc(100% - 2rem);
  }
  
  .connectSection {
    flex-direction: column;
    gap: 0.5rem;
    width: 100%;
  }
  
  .connectButton, .selectWalletButton {
    width: 100%;
    justify-content: center;
  }
}
</style>
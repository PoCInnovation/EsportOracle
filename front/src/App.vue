<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { createWalletClient, createPublicClient, custom, formatEther } from 'viem'
import { mainnet, sepolia } from 'viem/chains'

// Types TypeScript
interface Wallet {
  name: string
  icon: string
  provider: any
  detected: boolean
}

type WalletName = 'MetaMask' | 'Coinbase Wallet' | 'Trust Wallet' | 'Brave Wallet' | 'WalletConnect'

const navigation = [
  { name: 'Live Matches', href: '/matches/current', icon: '🔴' },
  { name: 'Upcoming Matches', href: '/matches/upcoming', icon: '⏰' },
  { name: 'Past Matches', href: '/matches/past', icon: '📚' }
]

const mobileMenuOpen = ref(false)
const showWalletModal = ref(false)
const isConnecting = ref(false)
const connectionError = ref('')
const isConnected = ref(false)
const connectedAccount = ref('')
const walletBalance = ref('')
const availableWallets = ref<Wallet[]>([])

// Extended Window interface to include ethereum with providers
declare global {
  interface Window {
    ethereum?: any
  }
}

// Functions for automatic save and restore
const WALLET_STORAGE_KEY = 'esport-oracle-wallet'

const saveWalletConnection = (account: string, walletName: string) => {
  const walletData = {
    account,
    walletName,
    timestamp: Date.now()
  }
  localStorage.setItem(WALLET_STORAGE_KEY, JSON.stringify(walletData))
}

const getSavedWalletConnection = () => {
  try {
    const saved = localStorage.getItem(WALLET_STORAGE_KEY)
    if (saved) {
      const walletData = JSON.parse(saved)
      // Check if connection is not too old (7 days)
      const sevenDays = 7 * 24 * 60 * 60 * 1000
      if (Date.now() - walletData.timestamp < sevenDays) {
        return walletData
      }
    }
  } catch (error) {
    console.error('Error reading saved data:', error)
  }
  return null
}

const clearSavedWalletConnection = () => {
  localStorage.removeItem(WALLET_STORAGE_KEY)
}

// Detect all available wallets
const detectWallets = () => {
  const wallets: Wallet[] = []
  
  if (typeof window.ethereum !== 'undefined') {
    // Single wallet
    if (!window.ethereum.providers) {
      const wallet = {
        name: getWalletName(window.ethereum),
        icon: getWalletIcon(window.ethereum),
        provider: window.ethereum,
        detected: true
      }
      wallets.push(wallet)
    }
    // Multiple wallets detected
    else if (Array.isArray(window.ethereum.providers)) {
      window.ethereum.providers.forEach((provider: any, index: number) => {
        const wallet = {
          name: getWalletName(provider),
          icon: getWalletIcon(provider),
          provider,
          detected: true
        }
        wallets.push(wallet)
      })
    }
  }
  
  // Add common wallets not detected
  const commonWallets = [
    { name: 'MetaMask', icon: '🦊' },
    { name: 'Coinbase Wallet', icon: '🔵' },
    { name: 'Trust Wallet', icon: '�️' },
    { name: 'Brave Wallet', icon: '🦁' },
    { name: 'WalletConnect', icon: '🔗' }
  ]
  
  commonWallets.forEach(commonWallet => {
    if (!wallets.find(w => w.name === commonWallet.name)) {
      wallets.push({
        ...commonWallet,
        provider: null,
        detected: false
      })
    }
  })
  
  availableWallets.value = wallets
  console.log('Detected wallets:', wallets)
}

// Identify wallet name
const getWalletName = (provider: any): string => {
  if (provider.isMetaMask && !provider.isBraveWallet) return 'MetaMask'
  if (provider.isBraveWallet) return 'Brave Wallet'
  if (provider.isCoinbaseWallet) return 'Coinbase Wallet'
  if (provider.isTrustWallet) return 'Trust Wallet'
  if (provider.isExodus) return 'Exodus'
  if (provider.isRabby) return 'Rabby'
  if (provider.isOkxWallet) return 'OKX Wallet'
  if (provider.isTalisman) return 'Talisman'
  if (provider.isPhantom) return 'Phantom'
  return 'Unknown Wallet'
}

// Get wallet icon
const getWalletIcon = (provider: any): string => {
  if (provider.isMetaMask && !provider.isBraveWallet) return '🦊'
  if (provider.isBraveWallet) return '🦁'
  if (provider.isCoinbaseWallet) return '🏦'
  if (provider.isTrustWallet) return '🛡️'
  if (provider.isExodus) return '🚀'
  if (provider.isRabby) return '🐰'
  if (provider.isOkxWallet) return '⭕'
  if (provider.isTalisman) return '💎'
  if (provider.isPhantom) return '👻'
  return '💳'
}

const connectWallet = () => {
  showWalletModal.value = true
  connectionError.value = ''
  detectWallets()
}

const closeWalletModal = () => {
  showWalletModal.value = false
}

const disconnectWallet = () => {
  isConnected.value = false
  connectedAccount.value = ''
  walletBalance.value = ''
  clearSavedWalletConnection()
  console.log('Wallet disconnected')
}

const connectWithProvider = async (wallet: Wallet) => {
  try {
    isConnecting.value = true
    connectionError.value = ''
    
    // If wallet is not detected, redirect to installation
    if (!wallet.detected || !wallet.provider) {
      const installUrls: Record<string, string> = {
        'MetaMask': 'https://metamask.io/download/',
        'Coinbase Wallet': 'https://www.coinbase.com/wallet/downloads',
        'Trust Wallet': 'https://trustwallet.com/download',
        'Brave Wallet': 'https://brave.com/wallet/',
        'WalletConnect': 'https://walletconnect.com/registry'
      }
      
      connectionError.value = `${wallet.name} is not installed`
      if (installUrls[wallet.name]) {
        window.open(installUrls[wallet.name], '_blank')
      }
      return
    }

    console.log(`Connecting with ${wallet.name} using Viem...`)
    
    // Create wallet client with Viem and specific provider
    const walletClient = createWalletClient({
      chain: mainnet,
      transport: custom(wallet.provider)
    })
    
    // Create public client to read data
    const publicClient = createPublicClient({
      chain: mainnet,
      transport: custom(wallet.provider)
    })
    
    // Request accounts
    const accounts = await walletClient.requestAddresses()
    
    if (accounts && accounts.length > 0) {
      const account = accounts[0]
      connectedAccount.value = account
      
      console.log(`Connected account via ${wallet.name}:`, account)
      
      // Get balance with Viem
      try {
        const balance = await publicClient.getBalance({
          address: account
        })
        walletBalance.value = formatEther(balance)
        console.log('Account balance:', walletBalance.value, 'ETH')
      } catch (balanceError) {
        console.warn('Could not fetch balance:', balanceError)
        walletBalance.value = '0.0'
      }
      
      // Get chain information
      const chainId = await walletClient.getChainId()
      console.log('Connected to chain ID:', chainId)
      
      // Mark as connected
      isConnected.value = true
      
      // Save connection for auto-reconnect
      saveWalletConnection(account, wallet.name)
      
      // Close modal
      closeWalletModal()
      console.log(`✅ ${wallet.name} connected successfully`)
      
    } else {
      connectionError.value = `No account found in ${wallet.name}`
    }
    
  } catch (error: any) {
    console.error(`Error connecting ${wallet.name}:`, error)
    
    // Specific error handling
    if (error?.code === 4001) {
      connectionError.value = 'Connection rejected by user'
    } else if (error?.code === -32002) {
      connectionError.value = 'Connection request already pending'
    } else if (error?.message?.includes('User rejected')) {
      connectionError.value = 'Connection cancelled by user'
    } else if (error?.message?.includes('transport')) {
      connectionError.value = `Transport error with ${wallet.name}`
    } else {
      connectionError.value = `Error ${wallet.name}: ${error?.message || error?.shortMessage || 'Unknown error'}`
    }
  } finally {
    isConnecting.value = false
  }
}

const connectWalletConnect = () => {
  console.log('WalletConnect integration coming soon!')
  connectionError.value = 'WalletConnect integration coming soon!'
  closeWalletModal()
}

const connectCoinbase = () => {
  console.log('Coinbase Wallet integration coming soon!')
  connectionError.value = 'Coinbase Wallet integration coming soon!'
  closeWalletModal()
}

// Auto-reconnect function
const tryAutoReconnect = async () => {
  const savedWallet = getSavedWalletConnection()
  if (!savedWallet) return

  console.log('Attempting auto-reconnect for:', savedWallet.walletName)
  
  try {
    // Detect available wallets
    detectWallets()
    
    // Find saved wallet in detected wallets list
    const wallet = availableWallets.value.find(w => 
      w.name === savedWallet.walletName && w.detected
    )
    
    if (!wallet || !wallet.provider) {
      console.log('Saved wallet not available for reconnection')
      clearSavedWalletConnection()
      return
    }

    // Create wallet client with Viem
    const walletClient = createWalletClient({
      chain: mainnet,
      transport: custom(wallet.provider)
    })
    
    // Create public client to read data
    const publicClient = createPublicClient({
      chain: mainnet,
      transport: custom(wallet.provider)
    })
    
    // Check if saved account is still accessible
    const accounts = await walletClient.getAddresses()
    
    if (accounts && accounts.includes(savedWallet.account)) {
      // Successful reconnection
      connectedAccount.value = savedWallet.account
      isConnected.value = true
      
      // Get balance
      try {
        const balance = await publicClient.getBalance({
          address: savedWallet.account
        })
        walletBalance.value = formatEther(balance)
      } catch (error) {
        console.warn('Unable to retrieve balance:', error)
        walletBalance.value = '0.0'
      }
      
      console.log(`✅ Auto-reconnection successful with ${savedWallet.walletName}`)
    } else {
      console.log('Saved account not accessible, removing save')
      clearSavedWalletConnection()
    }
    
  } catch (error) {
    console.log('Error during auto-reconnection:', error)
    clearSavedWalletConnection()
  }
}

// Component initialization on mount
onMounted(async () => {
  await tryAutoReconnect()
})
</script>

<template>
  <div class="app">
    <!-- Header -->
    <header class="header">
      <div class="container">
        <div class="header-content">
          <!-- Logo -->
          <div class="logo">
            <router-link to="/" class="logo-link">
              <div class="logo-icon">⚡</div>
              <span class="logo-text">EsportOracle</span>
            </router-link>
          </div>

          <!-- Desktop Navigation -->
          <nav class="nav-desktop">
            <router-link
              v-for="item in navigation"
              :key="item.name"
              :to="item.href"
              class="nav-link"
              active-class="nav-link-active"
            >
              <span class="nav-icon">{{ item.icon }}</span>
              {{ item.name }}
            </router-link>
          </nav>

          <!-- Wallet Section -->
          <div class="wallet-section">
            <button v-if="!isConnected" class="wallet-btn" @click="connectWallet">
              <span class="wallet-icon">💳</span>
              <span class="wallet-text">Connect Wallet</span>
            </button>
            
            <!-- Connected State -->
            <div v-else class="wallet-connected">
              <span class="wallet-address">{{ connectedAccount.slice(0, 6) }}...{{ connectedAccount.slice(-4) }}</span>
              <button class="wallet-disconnect" @click="disconnectWallet">
                <span>✕</span>
              </button>
            </div>
          </div>

          <!-- Mobile menu button -->
          <button 
            class="mobile-menu-btn"
            @click="mobileMenuOpen = !mobileMenuOpen"
          >
            <span class="sr-only">Ouvrir le menu</span>
            <div class="hamburger" :class="{ active: mobileMenuOpen }">
              <span></span>
              <span></span>
              <span></span>
            </div>
          </button>
        </div>
      </div>

      <!-- Mobile Navigation -->
      <Transition name="mobile-menu">
        <nav v-if="mobileMenuOpen" class="nav-mobile">
          <div class="nav-mobile-content">
            <router-link
              v-for="item in navigation"
              :key="item.name"
              :to="item.href"
              class="nav-mobile-link"
              active-class="nav-mobile-link-active"
              @click="mobileMenuOpen = false"
            >
              <span class="nav-icon">{{ item.icon }}</span>
              {{ item.name }}
            </router-link>
            
            <!-- Mobile Wallet Button -->
            <button class="wallet-btn-mobile" @click="connectWallet(); mobileMenuOpen = false">
              <span class="wallet-icon">💳</span>
              Connect Wallet
            </button>
          </div>
        </nav>
      </Transition>
    </header>

    <!-- Main Content -->
    <main class="main">
      <RouterView v-slot="{ Component }">
        <Transition name="page" mode="out-in">
          <component :is="Component"/>
        </Transition>
      </RouterView>
    </main>

    <!-- Wallet Modal -->
    <Transition name="modal">
      <div v-if="showWalletModal" class="wallet-modal-overlay" @click="closeWalletModal">
        <div class="wallet-modal" @click.stop>
          <div class="wallet-modal-header">
            <h3>Connect Your Wallet</h3>
            <button @click="closeWalletModal" class="wallet-modal-close">
              <span>✕</span>
            </button>
          </div>
          
          <div class="wallet-modal-content">
            <div class="wallet-options">
              <p class="wallet-description">
                Choose your preferred wallet to connect to EsportOracle
              </p>
              
              <!-- Error display -->
              <div v-if="connectionError" class="wallet-error">
                <span class="error-icon">⚠️</span>
                {{ connectionError }}
              </div>
              
              <!-- Dynamic wallet list -->
              <button 
                v-for="wallet in availableWallets" 
                :key="wallet.name"
                class="wallet-option" 
                @click="() => connectWithProvider(wallet)"
                :disabled="isConnecting"
                :class="{ 
                  'connecting': isConnecting,
                  'not-detected': !wallet.detected 
                }"
              >
                <div class="wallet-option-icon">{{ wallet.icon }}</div>
                <div class="wallet-option-info">
                  <div class="wallet-option-name">
                    {{ wallet.name }}
                    <span v-if="!wallet.detected" class="not-installed-badge">Not installed</span>
                  </div>
                  <div class="wallet-option-desc">
                    {{ isConnecting ? 'Connecting...' : 
                       wallet.detected ? 'Extension detected' : 'Click to install' }}
                  </div>
                </div>
                <div class="wallet-option-arrow">
                  {{ isConnecting ? '⏳' : wallet.detected ? '→' : '⬇️' }}
                </div>
              </button>
              
              <button class="wallet-option" @click="connectCoinbase" :disabled="isConnecting">
                <div class="wallet-option-icon">🏦</div>
                <div class="wallet-option-info">
                  <div class="wallet-option-name">Coinbase Wallet</div>
                  <div class="wallet-option-desc">Browser & Mobile</div>
                </div>
                <div class="wallet-option-arrow">→</div>
              </button>
            </div>
          </div>
        </div>
      </div>
    </Transition>

    <!-- Footer -->
    <footer class="footer">
      <div class="container">
        <div class="footer-content">
          <p class="footer-text">
            © 2025 EsportOracle. Data provided by PandaScore API.
          </p>
        </div>
      </div>
    </footer>
  </div>
</template>


<style scoped>
.app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background: linear-gradient(135deg, #0a0a0a 0%, #1a0f08 50%, #0a0a0a 100%);
  color: #ffffff;
  position: relative;
  overflow-x: hidden;
}

.app::before {
  content: '';
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-image: 
    radial-gradient(circle at 20% 80%, rgba(249, 115, 22, 0.1) 0%, transparent 50%),
    radial-gradient(circle at 80% 20%, rgba(251, 146, 60, 0.1) 0%, transparent 50%),
    radial-gradient(circle at 40% 40%, rgba(253, 124, 43, 0.05) 0%, transparent 50%);
  pointer-events: none;
  z-index: -1;
  animation: float 8s ease-in-out infinite;
}

/* Header */
.header {
  position: sticky;
  top: 0;
  z-index: 50;
  background: rgba(10, 10, 10, 0.95);
  backdrop-filter: blur(20px);
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 4rem;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1.5rem;
}

.logo-link {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  text-decoration: none;
  transition: opacity 0.2s ease;
}

.logo-link:hover {
  opacity: 0.8;
}

.logo-icon {
  font-size: 2rem;
  background: linear-gradient(135deg, #f97316, #fb923c);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.logo-text {
  font-size: 1.5rem;
  font-weight: 800;
  color: #ffffff;
}

/* Desktop Navigation */
.nav-desktop {
  display: flex;
  gap: 0.5rem;
}

.nav-link {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1rem;
  color: #a3a3a3;
  text-decoration: none;
  border-radius: 0.75rem;
  font-weight: 500;
  transition: all 0.2s ease;
}

.nav-link:hover {
  color: #ffffff;
  background: #141414;
}

.nav-link-active {
  color: #f97316;
  background: rgba(249, 115, 22, 0.1);
  border: 1px solid rgba(249, 115, 22, 0.2);
}

.nav-icon {
  font-size: 1rem;
}

/* Wallet Section */
.wallet-section {
  display: flex;
  align-items: center;
}

.wallet-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1.5rem;
  background: linear-gradient(135deg, #f97316 0%, #ea580c 100%);
  border: none;
  border-radius: 12px;
  color: white;
  font-weight: 600;
  font-size: 0.875rem;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 15px rgba(249, 115, 22, 0.3);
  font-family: inherit;
}

.wallet-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(249, 115, 22, 0.4);
  background: linear-gradient(135deg, #ea580c 0%, #dc2626 100%);
}

.wallet-icon {
  font-size: 1.1rem;
}

.wallet-text {
  font-weight: 600;
}

/* Connected Wallet State */
.wallet-connected {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.5rem 1rem;
  background: rgba(34, 197, 94, 0.1);
  border: 1px solid rgba(34, 197, 94, 0.3);
  border-radius: 12px;
  font-family: inherit;
}

.wallet-address {
  font-family: 'Courier New', monospace;
  font-size: 0.875rem;
  color: #22c55e;
  font-weight: 600;
  letter-spacing: 0.5px;
}

.wallet-disconnect {
  background: rgba(239, 68, 68, 0.1);
  border: 1px solid rgba(239, 68, 68, 0.3);
  border-radius: 50%;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 0.75rem;
  color: #ef4444;
  font-weight: bold;
}

.wallet-disconnect:hover {
  background: rgba(239, 68, 68, 0.2);
  border-color: rgba(239, 68, 68, 0.5);
  transform: scale(1.05);
}

/* Mobile menu button */
.mobile-menu-btn {
  display: none;
  background: none;
  border: none;
  padding: 0.5rem;
}

.hamburger {
  position: relative;
  width: 24px;
  height: 18px;
}

.hamburger span {
  display: block;
  position: absolute;
  height: 2px;
  width: 100%;
  background: #ffffff;
  border-radius: 1px;
  transition: all 0.3s ease;
}

.hamburger span:nth-child(1) { top: 0; }
.hamburger span:nth-child(2) { top: 8px; }
.hamburger span:nth-child(3) { top: 16px; }

.hamburger.active span:nth-child(1) {
  transform: rotate(45deg);
  top: 8px;
}

.hamburger.active span:nth-child(2) {
  opacity: 0;
}

.hamburger.active span:nth-child(3) {
  transform: rotate(-45deg);
  top: 8px;
}

/* Mobile Navigation */
.nav-mobile {
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  background: #0a0a0a;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.nav-mobile-content {
  padding: 1rem;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.nav-mobile-link {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 1rem;
  color: #a3a3a3;
  text-decoration: none;
  border-radius: 0.75rem;
  font-weight: 500;
  transition: all 0.2s ease;
}

.nav-mobile-link:hover {
  color: #ffffff;
  background: #141414;
}

.nav-mobile-link-active {
  color: #f97316;
  background: rgba(249, 115, 22, 0.1);
  border: 1px solid rgba(249, 115, 22, 0.2);
}

.wallet-btn-mobile {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 1rem;
  background: linear-gradient(135deg, #f97316 0%, #ea580c 100%);
  border: none;
  border-radius: 12px;
  color: white;
  font-weight: 600;
  font-size: 0.875rem;
  cursor: pointer;
  transition: all 0.3s ease;
  margin-top: 0.5rem;
  font-family: inherit;
}

.wallet-btn-mobile:hover {
  background: linear-gradient(135deg, #ea580c 0%, #dc2626 100%);
}

/* Main */
.main {
  flex: 1;
  padding: 2rem 0;
}

/* Footer */
.footer {
  background: #141414;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  padding: 1.5rem 0;
}

.footer-content {
  text-align: center;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1.5rem;
}

.footer-text {
  color: #737373;
  font-size: 0.875rem;
  margin: 0;
}

/* Transitions */
.page-enter-active,
.page-leave-active {
  transition: all 0.3s ease;
}

.page-enter-from {
  opacity: 0;
  transform: translateY(20px);
}

.page-leave-to {
  opacity: 0;
  transform: translateY(-20px);
}

.mobile-menu-enter-active,
.mobile-menu-leave-active {
  transition: all 0.3s ease;
}

.mobile-menu-enter-from,
.mobile-menu-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

.modal-enter-active,
.modal-leave-active {
  transition: all 0.3s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}

/* Wallet Modal */
.wallet-modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.8);
  backdrop-filter: blur(8px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 100;
  padding: 1rem;
}

.wallet-modal {
  background: linear-gradient(145deg, 
    rgba(20, 20, 20, 0.95) 0%, 
    rgba(30, 20, 15, 0.95) 100%
  );
  border: 1px solid rgba(249, 115, 22, 0.3);
  border-radius: 20px;
  padding: 0;
  min-width: 400px;
  max-width: 500px;
  width: 100%;
  backdrop-filter: blur(20px);
  box-shadow: 
    0 20px 40px rgba(0, 0, 0, 0.5),
    0 0 50px rgba(249, 115, 22, 0.1);
  animation: modalSlideIn 0.3s ease;
  overflow: hidden;
}

@keyframes modalSlideIn {
  from { 
    opacity: 0;
    transform: translateY(-20px) scale(0.95);
  }
  to { 
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.wallet-modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1.5rem 2rem;
  border-bottom: 1px solid rgba(249, 115, 22, 0.2);
  background: rgba(249, 115, 22, 0.05);
}

.wallet-modal-header h3 {
  margin: 0;
  color: #fff;
  font-size: 1.25rem;
  font-weight: 600;
  background: linear-gradient(135deg, #f97316, #fb923c);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.wallet-modal-close {
  background: rgba(249, 115, 22, 0.1);
  border: 1px solid rgba(249, 115, 22, 0.3);
  border-radius: 8px;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s ease;
  color: #f97316;
  font-size: 1rem;
  font-weight: bold;
}

.wallet-modal-close:hover {
  background: rgba(249, 115, 22, 0.2);
  border-color: rgba(249, 115, 22, 0.5);
  transform: scale(1.05);
}

.wallet-modal-content {
  padding: 2rem;
}

.wallet-description {
  color: rgba(255, 255, 255, 0.8);
  text-align: center;
  margin: 0 0 1.5rem 0;
  font-size: 0.95rem;
  line-height: 1.5;
}

.wallet-options {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.wallet-error {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1rem;
  background: rgba(239, 68, 68, 0.1);
  border: 1px solid rgba(239, 68, 68, 0.3);
  border-radius: 8px;
  color: #fca5a5;
  font-size: 0.875rem;
  margin-bottom: 0.5rem;
}

.error-icon {
  font-size: 1rem;
  flex-shrink: 0;
}

.wallet-option {
  display: flex;
  align-items: center;
  padding: 1rem 1.25rem;
  background: rgba(249, 115, 22, 0.05);
  border: 1px solid rgba(249, 115, 22, 0.2);
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  width: 100%;
  text-align: left;
  font-family: inherit;
}

.wallet-option:hover {
  background: rgba(249, 115, 22, 0.1);
  border-color: rgba(249, 115, 22, 0.4);
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(249, 115, 22, 0.2);
}

.wallet-option-icon {
  font-size: 2rem;
  margin-right: 1rem;
  flex-shrink: 0;
}

.wallet-option-info {
  flex: 1;
}

.wallet-option-name {
  color: #fff;
  font-weight: 600;
  font-size: 1rem;
  margin-bottom: 0.25rem;
}

.wallet-option-desc {
  color: rgba(255, 255, 255, 0.6);
  font-size: 0.875rem;
}

.wallet-option-arrow {
  color: #f97316;
  font-size: 1.25rem;
  font-weight: bold;
  transition: transform 0.3s ease;
}

.wallet-option:hover .wallet-option-arrow {
  transform: translateX(4px);
}

.wallet-option:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none !important;
}

.wallet-option.connecting {
  background: rgba(249, 115, 22, 0.15);
  border-color: rgba(249, 115, 22, 0.4);
}

/* Nouveaux styles pour les wallets non détectés */
.wallet-option.not-detected {
  opacity: 0.7;
  border-color: rgba(156, 163, 175, 0.3);
}

.wallet-option.not-detected:hover {
  background: rgba(156, 163, 175, 0.05);
  border-color: rgba(156, 163, 175, 0.4);
}

.not-installed-badge {
  color: #ef4444;
  font-size: 0.75rem;
  font-weight: 500;
  margin-left: 0.5rem;
  padding: 0.125rem 0.375rem;
  background: rgba(239, 68, 68, 0.1);
  border-radius: 0.25rem;
  border: 1px solid rgba(239, 68, 68, 0.2);
}

/* Responsive */
@media (max-width: 768px) {
  .nav-desktop {
    display: none;
  }

  .wallet-section {
    display: none;
  }

  .wallet-connected {
    display: none;
  }

  .mobile-menu-btn {
    display: block;
  }

  .logo-text {
    font-size: 1.25rem;
  }

  .header-content {
    padding: 0 1rem;
  }

  .main {
    padding: 1.5rem 0;
  }

  .wallet-modal {
    min-width: unset;
    margin: 0.5rem;
    width: calc(100% - 1rem);
  }

  .wallet-modal-header {
    padding: 1rem 1.5rem;
  }

  .wallet-modal-content {
    padding: 1.5rem;
  }
}

/* Animations */
@keyframes float {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
  }
  33% {
    transform: translateY(-10px) rotate(1deg);
  }
  66% {
    transform: translateY(-5px) rotate(-1deg);
  }
}
</style>
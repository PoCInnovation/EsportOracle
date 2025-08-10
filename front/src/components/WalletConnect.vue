<template>
  <div class="walletSection">
    <div v-if="isConnected" class="walletInfo">
      <span class="addressDisplay">{{ shortenAddress(connectedAccount) }}</span>
      <button @click="disconnect" class="disconnectButton" title="Disconnect wallet">
        ×
      </button>
    </div>
    <button 
      v-else
      @click="connect" 
      class="connectButton"
      :disabled="isConnecting"
    >
      {{ isConnecting ? 'Connecting...' : buttonText }}
    </button>
    <div v-if="error" class="errorMessage">
      {{ error }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue'
import { createPublicClient, createWalletClient, custom, http, type WalletClient, type PublicClient, type Address, formatEther } from 'viem'
import { mainnet, sepolia, localhost } from 'viem/chains'

/**
 * Interface for Ethereum provider (MetaMask, WalletConnect, etc.)
 * Defines the standard methods and events that wallet providers should implement
 */
interface EthereumProvider {
  request(args: { method: string; params?: any[] }): Promise<any>
  on(event: string, handler: (...args: any[]) => void): void
  removeListener(event: string, handler: (...args: any[]) => void): void
  isMetaMask?: boolean
}

/**
 * Extend the global Window interface to include the ethereum property
 * This allows TypeScript to recognize window.ethereum without type errors
 */
declare global {
  interface Window {
    ethereum?: EthereumProvider
  }
}

/**
 * Reactive state management for wallet connection
 * These refs track the current state of the wallet connection and user interaction
 */
const isConnected = ref(false)          // Whether a wallet is currently connected
const isConnecting = ref(false)         // Whether a connection attempt is in progress
const connectedAccount = ref<Address | ''>('')  // The currently connected wallet address
const walletClient = ref<WalletClient | null>(null)    // Viem wallet client for transactions
const publicClient = ref<PublicClient | null>(null)    // Viem public client for reading blockchain data
const chainId = ref<number | null>(null)        // Current blockchain network ID
const balance = ref<string>('')         // Current ETH balance in human-readable format
const error = ref<string>('')           // Current error message to display to user

/**
 * Maps chain IDs to their corresponding Viem chain configurations
 * Supports mainnet (1), Sepolia testnet (11155111), and local development (31337)
 * @param chainId - The numeric chain ID from the wallet
 * @returns The corresponding Viem chain configuration
 */
const getCurrentChain = (chainId: number) => {
  switch (chainId) {
    case 1: return mainnet
    case 11155111: return sepolia
    case 31337: return localhost
    default: return mainnet
  }
}

/**
 * Computed property that determines the text to display on the connect button
 * Changes based on whether MetaMask is installed or not
 * @returns String indicating either to install MetaMask or connect wallet
 */
const buttonText = computed(() => {
  if (!window.ethereum) {
    return 'Install MetaMask'
  }
  return 'Connect Wallet'
})

/**
 * Safely retrieves the Ethereum provider from the window object
 * Handles cases where multiple wallet providers are installed (e.g., MetaMask + WalletConnect)
 * Prioritizes MetaMask if available among multiple providers
 * @returns The Ethereum provider instance or null if not available
 */
const getEthereumProvider = (): EthereumProvider | null => {
  if (typeof window === 'undefined') return null
  
  const ethereum = (window as any).ethereum
  if (!ethereum) return null

  // Handle multiple providers (e.g., MetaMask + other wallets)
  if (ethereum.providers && ethereum.providers.length > 0) {
    return ethereum.providers.find((provider: any) => provider.isMetaMask) || ethereum.providers[0]
  }
  
  return ethereum
}

/**
 * Main function to establish a connection with the user's wallet
 * Handles the complete flow from requesting accounts to setting up Viem clients
 * Updates all relevant state variables upon successful connection
 * @throws Will set error state if connection fails at any step
 */
const connect = async () => {
  error.value = ''
  const provider = getEthereumProvider()
  
  if (!provider) {
    window.open('https://metamask.io/download/', '_blank')
    return
  }

  try {
    isConnecting.value = true
    
    // Request account access from the wallet
    const accounts = await provider.request({
      method: 'eth_requestAccounts'
    }) as Address[]
    
    if (!accounts || accounts.length === 0) {
      throw new Error('No accounts returned from wallet')
    }

    // Get the current chain ID from the wallet
    const chainIdHex = await provider.request({
      method: 'eth_chainId'
    }) as string
    
    const currentChainId = parseInt(chainIdHex, 16)
    const currentChain = getCurrentChain(currentChainId)
    
    // Create public client for reading blockchain data (balances, contract state, etc.)
    const pubClient = createPublicClient({
      chain: currentChain,
      transport: custom(provider)
    })
    
    // Create wallet client for signing transactions (without pre-specifying account)
    const client = createWalletClient({
      chain: currentChain,
      transport: custom(provider)
    })
    
    // Update component state with connection details
    walletClient.value = client
    publicClient.value = pubClient
    connectedAccount.value = accounts[0]
    chainId.value = currentChainId
    isConnected.value = true
    
    // Fetch and display the user's ETH balance
    await updateBalance()
    
    console.log('Wallet connected successfully:', {
      account: accounts[0],
      chainId: currentChainId,
      balance: balance.value
    })
    
  } catch (err: any) {
    console.error('Connection error:', err)
    error.value = err.message || 'Failed to connect wallet'
  } finally {
    isConnecting.value = false
  }
}

/**
 * Safely disconnects the wallet after user confirmation
 * Clears all wallet-related state and logs the disconnection
 * Shows a confirmation dialog to prevent accidental disconnections
 */
const disconnect = async () => {
  const confirmDisconnect = confirm('Are you sure you want to disconnect your wallet?')
  
  if (!confirmDisconnect) return
  
  // Clear all wallet-related state
  isConnected.value = false
  connectedAccount.value = ''
  walletClient.value = null
  publicClient.value = null
  chainId.value = null
  balance.value = ''
  error.value = ''
  
  console.log('Wallet disconnected')
}

/**
 * Fetches and updates the ETH balance for the connected account
 * Uses the public client to query the blockchain for the current balance
 * Converts the balance from Wei to Ether for human-readable display
 * @throws Logs error if balance retrieval fails but doesn't break the UI
 */
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

/**
 * Handles wallet account changes triggered by the user switching accounts in MetaMask
 * Automatically disconnects if no accounts are available, or updates to the new account
 * Updates the balance display when switching to a different account
 * @param accounts - Array of account addresses from the wallet provider
 */
const handleAccountsChanged = (accounts: string[]) => {
  console.log('Accounts changed:', accounts)
  
  if (accounts.length === 0) {
    disconnect()
  } else if (accounts[0] !== connectedAccount.value) {
    connectedAccount.value = accounts[0] as Address
    updateBalance()
  }
}

/**
 * Handles blockchain network changes triggered by the user switching networks in MetaMask
 * Recreates the Viem clients with the new chain configuration
 * Updates the balance display for the new network
 * @param chainIdHex - Hexadecimal string representation of the new chain ID
 */
const handleChainChanged = (chainIdHex: string) => {
  const newChainId = parseInt(chainIdHex, 16)
  console.log('Chain changed:', newChainId)
  
  chainId.value = newChainId
  
  if (walletClient.value && connectedAccount.value) {
    // Recreate wallet client with new chain
    const provider = getEthereumProvider()
    if (provider) {
      const client = createWalletClient({
        chain: newChainId === 1 ? mainnet : sepolia,
        transport: custom(provider)
      })
      walletClient.value = client
      updateBalance()
    }
  }
}

/**
 * Sets up event listeners for wallet provider events
 * Listens for account changes, network changes, and disconnection events
 * Must be called during component initialization to ensure proper wallet state synchronization
 */
const setupEventListeners = () => {
  const provider = getEthereumProvider()
  if (!provider) return
  
  provider.on('accountsChanged', handleAccountsChanged)
  provider.on('chainChanged', handleChainChanged)
  
  // Handle wallet disconnection initiated from the wallet itself
  provider.on('disconnect', () => {
    console.log('Provider disconnected')
    disconnect()
  })
}

/**
 * Removes all wallet provider event listeners to prevent memory leaks
 * Should be called during component cleanup or before setting up new listeners
 */
const cleanupEventListeners = () => {
  const provider = getEthereumProvider()
  if (!provider) return
  
  provider.removeListener('accountsChanged', handleAccountsChanged)
  provider.removeListener('chainChanged', handleChainChanged)
}

/**
 * Checks if the wallet was previously connected and auto-reconnects if so
 * This provides a seamless user experience by maintaining wallet connection across page reloads
 * Uses eth_accounts method which returns accounts only if previously authorized
 */
const checkConnection = async () => {
  const provider = getEthereumProvider()
  if (!provider) return
  
  try {
    const accounts = await provider.request({
      method: 'eth_accounts'
    }) as Address[]
    
    if (accounts && accounts.length > 0) {
      // Auto-connect if previously connected
      await connect()
    }
  } catch (err) {
    console.error('Failed to check existing connection:', err)
  }
}

/**
 * Utility function to create a user-friendly shortened version of an Ethereum address
 * Displays the first 6 and last 4 characters with ellipsis in between
 * @param address - The full Ethereum address to shorten
 * @returns Shortened address in format "0x1234...abcd" or empty string if no address
 */
const shortenAddress = (address: string) => {
  if (!address) return ''
  return `${address.slice(0, 6)}...${address.slice(-4)}`
}

// Lifecycle hooks
onMounted(async () => {
  setupEventListeners()
  await checkConnection()
})

// Watch for connection changes to update balance
watch(isConnected, (connected) => {
  if (connected) {
    updateBalance()
  }
})

// Expose reactive state for parent components
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
.walletSection {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 0.5rem;
}

.connectButton {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #ffffff;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  font-family: 'DM Sans', sans-serif;
  font-weight: 600;
  font-size: 0.875rem;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.connectButton:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

.connectButton:active:not(:disabled) {
  transform: translateY(0);
}

.connectButton:disabled {
  background: linear-gradient(135deg, #a0a0a0 0%, #808080 100%);
  cursor: not-allowed;
  transform: none;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.walletInfo {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  padding: 0.75rem 1rem;
  border-radius: 8px;
  border: 1px solid #e2e8f0;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.addressDisplay {
  color: #1a202c;
  font-family: 'DM Sans', sans-serif;
  font-weight: 600;
  font-size: 0.875rem;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
}

.disconnectButton {
  background: #ef4444;
  border: none;
  color: #ffffff;
  cursor: pointer;
  font-size: 0.875rem;
  font-weight: 600;
  padding: 0.25rem 0.5rem;
  width: auto;
  height: auto;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 4px;
  transition: all 0.2s ease;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.disconnectButton:hover {
  background: #dc2626;
  transform: translateY(-1px);
  box-shadow: 0 2px 4px rgba(239, 68, 68, 0.3);
}

.errorMessage {
  background: #fef2f2;
  color: #dc2626;
  padding: 0.5rem 0.75rem;
  border-radius: 6px;
  border: 1px solid #fecaca;
  font-size: 0.75rem;
  font-weight: 500;
  max-width: 300px;
}

/* Responsive design */
@media (max-width: 640px) {
  .walletSection {
    width: 100%;
  }
  
  .connectButton,
  .walletInfo {
    width: 100%;
    justify-content: center;
  }
  
  .addressDisplay {
    font-size: 0.75rem;
  }
}

/* Animation for connection state */
.walletInfo {
  animation: slideIn 0.3s ease-out;
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>

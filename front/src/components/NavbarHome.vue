<template>
  <nav class="navbar">
    <div class="navContainer">
      <!-- Logo/Brand -->
      <div class="navBrand">
        <RouterLink to="/" class="brandLink">
          <h2>EsportOracleeee</h2>
        </RouterLink>
      </div>
      
      <!-- Navigation Links -->
      <ul class="navMenu" :class="{ active: isMenuOpen }">
        <li class="navItem">
          <RouterLink to="/" class="navLink">
            <i class="pi pi-home"></i>
            Home
          </RouterLink>
        </li>
        <li class="navItem">
          <RouterLink to="/matches" class="navLink">
            <i class="pi pi-calendar"></i>
            Matches
          </RouterLink>
        </li>
        <li class="navItem">
          <RouterLink to="/profil" class="navLink">
            <i class="pi pi-user"></i>
            Profil
          </RouterLink>
        </li>
      </ul>

      <!-- Wallet Connection Section -->
      <div class="navActions">
        <div class="wallet-container">
          <!-- Simple wallet button for now -->
          <button class="wallet-btn" @click="connectWallet">
            <i class="pi pi-wallet"></i>
            <span>Connect Wallet</span>
          </button>
        </div>
        <button class="menuToggle" @click="toggleMenu" :class="{ active: isMenuOpen }">
          <span class="bar"></span>
          <span class="bar"></span>
          <span class="bar"></span>
        </button>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const isMenuOpen = ref(false)

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}

const connectWallet = () => {
  console.log('Connecting wallet...')
  alert('Wallet connection feature - to be implemented!')
}
</script>

<style scoped>
.navbar {
  background: rgba(17, 19, 23, 0.95);
  backdrop-filter: blur(20px);
  border-bottom: 1px solid rgba(249, 115, 22, 0.2);
  position: sticky;
  top: 0;
  z-index: 1000;
  margin: 0;
  padding: 0;
  width: 100%;
  box-shadow: 
    0 4px 20px rgba(0, 0, 0, 0.3),
    0 0 40px rgba(249, 115, 22, 0.1);
  transition: all 0.3s ease;
}

.navbar:hover {
  box-shadow: 
    0 4px 25px rgba(0, 0, 0, 0.4),
    0 0 50px rgba(249, 115, 22, 0.15);
}

.navContainer {
  max-width: 1400px;
  margin: 0 auto;
  padding: 0 1.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 70px;
  position: relative;
}

.navContainer::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(90deg, 
    transparent 0%, 
    rgba(249, 115, 22, 0.05) 50%, 
    transparent 100%
  );
  pointer-events: none;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.navbar:hover .navContainer::before {
  opacity: 1;
}

/* Brand Section */
.navBrand {
  flex-shrink: 0;
}

.brandLink {
  text-decoration: none;
  transition: all 0.3s ease;
  display: block;
}

.brandLink:hover {
  transform: translateY(-1px);
}

.navBrand h2 {
  background: linear-gradient(135deg, #f97316, #fb923c, #fbbf24);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  margin: 0;
  font-size: 1.75rem;
  font-weight: 800;
  cursor: pointer;
  font-family: 'DM Sans', sans-serif;
  transition: all 0.3s ease;
  position: relative;
  text-shadow: 0 0 20px rgba(249, 115, 22, 0.3);
}

.navBrand h2:hover {
  filter: drop-shadow(0 0 10px rgba(249, 115, 22, 0.5));
  transform: scale(1.02);
}

/* Navigation Menu */
.navMenu {
  display: flex;
  list-style: none;
  margin: 0;
  padding: 0;
  gap: 0.5rem;
  align-items: center;
}

.navItem {
  margin: 0;
}

.navLink {
  color: rgba(255, 255, 255, 0.8);
  font-family: 'DM Sans', sans-serif;
  font-size: 0.95rem;
  text-decoration: none;
  font-weight: 500;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  padding: 0.75rem 1.25rem;
  border-radius: 12px;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  position: relative;
  overflow: hidden;
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid transparent;
}

.navLink::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, 
    transparent, 
    rgba(249, 115, 22, 0.1), 
    transparent
  );
  transition: left 0.5s ease;
}

.navLink:hover {
  color: #ffffff;
  background: rgba(249, 115, 22, 0.1);
  border-color: rgba(249, 115, 22, 0.3);
  transform: translateY(-1px);
  box-shadow: 0 4px 15px rgba(249, 115, 22, 0.2);
}

.navLink:hover::before {
  left: 100%;
}

.navLink.router-link-active {
  color: #ffffff;
  background: rgba(249, 115, 22, 0.15);
  border-color: rgba(249, 115, 22, 0.4);
  box-shadow: 0 2px 10px rgba(249, 115, 22, 0.3);
}

.navLink i {
  font-size: 1rem;
  opacity: 0.8;
  transition: all 0.3s ease;
}

.navLink:hover i {
  opacity: 1;
  transform: scale(1.1);
}

/* Navigation Actions */
.navActions {
  display: flex;
  align-items: center;
  gap: 1rem;
  flex-shrink: 0;
}

.wallet-container {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  min-height: 40px;
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
  font-family: 'DM Sans', sans-serif;
}

.wallet-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(249, 115, 22, 0.4);
  background: linear-gradient(135deg, #ea580c 0%, #dc2626 100%);
}

.wallet-btn i {
  font-size: 1rem;
}

/* Mobile Menu Toggle */
.menuToggle {
  display: none;
  flex-direction: column;
  cursor: pointer;
  background: rgba(249, 115, 22, 0.1);
  border: 1px solid rgba(249, 115, 22, 0.3);
  border-radius: 8px;
  padding: 0.5rem;
  transition: all 0.3s ease;
}

.menuToggle:hover {
  background: rgba(249, 115, 22, 0.2);
  border-color: rgba(249, 115, 22, 0.5);
}

.bar {
  width: 20px;
  height: 2px;
  background: #f97316;
  margin: 2px 0;
  transition: all 0.3s ease;
  border-radius: 1px;
}

.menuToggle.active .bar:nth-child(1) {
  transform: rotate(-45deg) translate(-4px, 4px);
}

.menuToggle.active .bar:nth-child(2) {
  opacity: 0;
}

.menuToggle.active .bar:nth-child(3) {
  transform: rotate(45deg) translate(-4px, -4px);
}

/* Mobile Responsive */
@media (max-width: 968px) {
  .navContainer {
    padding: 0 1rem;
    height: 65px;
  }

  .navBrand h2 {
    font-size: 1.5rem;
  }

  .navMenu {
    position: absolute;
    top: 100%;
    left: 0;
    right: 0;
    background: rgba(17, 19, 23, 0.98);
    backdrop-filter: blur(25px);
    border-top: 1px solid rgba(249, 115, 22, 0.2);
    border-bottom: 1px solid rgba(249, 115, 22, 0.2);
    flex-direction: column;
    padding: 1.5rem;
    gap: 0.5rem;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
    transform: translateY(-10px);
    opacity: 0;
    visibility: hidden;
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .navMenu.active {
    transform: translateY(0);
    opacity: 1;
    visibility: visible;
  }

  .navItem {
    width: 100%;
  }

  .navLink {
    width: 100%;
    justify-content: flex-start;
    padding: 1rem 1.25rem;
    font-size: 1rem;
  }

  .menuToggle {
    display: flex;
  }

  .navActions {
    gap: 0.75rem;
  }

  .wallet-btn {
    padding: 0.5rem 1rem;
    font-size: 0.8rem;
  }
}

@media (max-width: 640px) {
  .navContainer {
    padding: 0 0.75rem;
    height: 60px;
  }

  .navBrand h2 {
    font-size: 1.3rem;
  }

  .navActions {
    gap: 0.5rem;
  }

  .wallet-btn {
    padding: 0.5rem 0.75rem;
    font-size: 0.75rem;
  }

  .wallet-btn span {
    display: none;
  }
}

/* Enhanced hover effects for desktop */
@media (min-width: 969px) {
  .navMenu:hover .navLink:not(:hover) {
    opacity: 0.6;
  }

  .navMenu .navLink:hover {
    opacity: 1;
  }
}

/* Accessibility improvements */
.navLink:focus-visible {
  outline: 2px solid #f97316;
  outline-offset: 2px;
}

.menuToggle:focus-visible {
  outline: 2px solid #f97316;
  outline-offset: 2px;
}

/* Animation for brand logo */
@keyframes pulse-brand {
  0%, 100% {
    filter: drop-shadow(0 0 5px rgba(249, 115, 22, 0.3));
  }
  50% {
    filter: drop-shadow(0 0 15px rgba(249, 115, 22, 0.5));
  }
}

.navBrand h2 {
  animation: pulse-brand 4s ease-in-out infinite;
}

/* Smooth scroll behavior for navigation */
html {
  scroll-behavior: smooth;
}
</style>

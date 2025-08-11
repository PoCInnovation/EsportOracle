<template>
  <nav class="navbar">
    <div class="navContainer">
      <RouterLink to="/" class="navLink">Home</RouterLink>
      <!-- Logo/Brand -->
      <div class="navBrand">
        <h2>EsportOracle</h2>
      </div>
      
      <!-- Navigation Links -->
      <ul class="navMenu" :class="{ active: isMenuOpen }">
        <li class="navItem">
          <RouterLink to="/matches" class="navLink">Matches</RouterLink>
        </li>
        <li class="navItem">
          <a href="#" class="navLink">Option 2</a>
        </li>
      </ul>

      <!-- MetaMask integration -->
      <WalletConnect />
      <RouterLink to="/profil" class="navLink">Profil</RouterLink>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import WalletConnect from './WalletConnect.vue'

const isMenuOpen = ref(false)

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}
</script>

<style scoped>
.navbar {
  background-color: #111317;
  position: sticky;
  top: 0;
  z-index: 1000;
  margin: 0;
  padding: 0;
  width: 100%;
}

.navContainer {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 60px;
}

.navBrand h2 {
  color: #db6715;
  margin: 0;
  font-size: 1.5rem;
  font-weight: bold;
  cursor: pointer;
  font-family: 'DM Sans', sans-serif;
}

.navMenu {
  display: flex;
  list-style: none;
  margin: 0;
  padding: 0;
  gap: 2rem;
}

.navItem {
  margin: 0;
  color: #e48d31;
}

.navLink {
  color: #d0d0d0;
  font-family: 'DM Sans', sans-serif;
  font-size: 13px;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.3s ease;
  padding: 0.5rem 0;
}

.navLink:hover {
  color: #ffffff;
}

.nav-toggle {
  display: none;
  flex-direction: column;
  cursor: pointer;
}

.bar {
  width: 25px;
  height: 3px;
  background-color: #ffffff;
  margin: 3px 0;
  transition: 0.3s;
}
</style>

<script setup>
import { ref } from 'vue';
const navOpen = ref(false)

</script>

<template>
    <div>
        <div>
            <div>
                <button @click="navOpen = !navOpen">
                    <svg :class="navOpen ? '' : 'hidden'"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    class="icon"
                    >
                    <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        d="M8.25 6.75h12M8.25 12h12m-12 5.25h12M3.75 6.75h.007v.008H3.75V6.75Zm.375 0a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0ZM3.75 12h.007v.008H3.75V12Zm.375 0a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Zm-.375 5.25h.007v.008H3.75v-.008Zm.375 0a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Z"
                    />
                    </svg>
                    <svg :class="navOpen ? '' : 'hidden'"
                        xmlns="http://www.w3.org/2000/svg" 
                        fill="none"
                        viewBox="0 0 24 24" 
                        stroke-width="1.5"
                        class="icon">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M6 18 18 6M6 6l12 12" />
                    </svg>
                </button>
            </div>
            <nav>

            </nav>
        </div>

        {{ navOpen }}
    </div>
</template>

<style scoped>
.icon {
  width: 24px;
  height: 24px;
  stroke: black;
}
</style>

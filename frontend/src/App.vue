<script setup lang="ts">
import { RouterView } from 'vue-router'

import { Icon } from '@iconify/vue'
import { userSession, toggleTheme } from './services/sessionService';
import { useTheme } from 'vuetify/lib/framework.mjs';

const theme = useTheme()

const switchTheme = () => {
  toggleTheme()
  theme.global.name.value = userSession.currentTheme.value!
}
</script>

<template>
  <main class="flex flex-col w-full gap-4 p-2 transition-all sm:h-auto md:h-screen font-roboto" :class="{
    'bg-gray-200': userSession.currentTheme.value === 'light',
    'bg-stone-800': userSession.currentTheme.value === 'dark'
  }">
    <header class="w-full h-[65px] px-4 py-2 flex justify-between rounded-md shadow-md transition-all" :class="{
      'bg-primary-500': userSession.currentTheme.value === 'light',
      'bg-primary-700': userSession.currentTheme.value === 'dark'
    }">
      <section class="flex w-auto h-auto align-middle text-light-inv_contrast">
        <Icon icon="mdi:world" width="48" class="mr-3"></Icon>
        <span class="my-auto text-2xl font-normal">Mail</span>
        <span class="my-auto text-2xl font-normal">-</span>
        <span class="my-auto text-2xl font-bold uppercase">Lens</span>
      </section>

      <section class="flex items-center justify-center w-auto h-auto">
        <Icon @click="switchTheme" class="transition-all hover:opacity-80 hover:cursor-pointer text-light-inv_contrast"
          :icon="userSession.currentTheme.value === 'light' ? 'solar:moon-bold' : 'solar:sun-bold'" width="32"></Icon>
      </section>
    </header>

    <RouterView />
  </main>
</template>
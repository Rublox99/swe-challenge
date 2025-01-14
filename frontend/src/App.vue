<script setup lang="ts">
import { RouterView } from 'vue-router'
import { Icon } from '@iconify/vue'
</script>

<template>
  <main class="flex flex-col items-center w-screen h-screen gap-10 p-2 font-body">
    <header class="w-full h-[65px] px-4 py-2 flex justify-between rounded-md shadow-md bg-slate-600">
      <section class="flex w-auto h-auto gap-2 align-middle">
        <Icon icon="mdi:world" width="48" color="var(--inverted-contrast)"></Icon>
        <span class="my-auto text-2xl font-normal text-white">Mail</span>
        <span class="my-auto text-2xl font-bold text-white uppercase">Lens</span>
      </section>

      <section class="flex items-center justify-center w-auto h-auto">
        <Icon icon="solar:moon-bold" width="32" color="var(--inverted-contrast)"></Icon>
      </section>
    </header>

    <RouterView />
  </main>
</template>
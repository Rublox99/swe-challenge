<script setup>
import { userSession } from '@/services/sessionService';
import { Icon } from '@iconify/vue';
import { ref } from 'vue';
import { DateTime } from 'luxon'

const startDate = ref()
const endDate = ref()

const startMenu = ref(false)
const endMenu = ref(false)

const formatDate = (date) => {
  const d = DateTime.fromJSDate(new Date(date))

  return d.toFormat('yyyy/MM/dd')
}
</script>

<template>
  <main class="flex flex-col gap-[15px] md:justify-between w-full rounded-md md:gap-0 md:flex-row lg:flex-row">
    <!-- Search Section -->
    <section class="flex items-center gap-[10px] w-full lg:w-[400px] md:w-[300px]">
      <Icon icon="mdi:magnify" width="30"
        :class="userSession.currentTheme.value === 'light' ? 'text-light-contrast' : 'text-dark-contrast'"></Icon>
      <v-text-field bg-color="transparent" hide-details="auto" clearable label="Search"
        variant="outlined"></v-text-field>
    </section>

    <!-- Dates Section -->
    <section class="flex items-center gap-[10px] w-full lg:w-[600px] md:w-[400px] sm:w-[200px]">
      <v-menu v-model="startMenu" transition="slide-y-transition" min-width="auto" :close-on-content-click="true">
        <template v-slot:activator="{ props }">
          <Icon icon="mdi:date-range" width="30" class="my-auto"
            :class="userSession.currentTheme.value === 'light' ? 'text-gray-800' : 'text-gray-600'"></Icon>
          <v-text-field label="Start Date" placeholder="YYYY-mm-dd" hide-details="auto" variant="outlined"
            v-bind="props" v-model="startDate" :readonly="true" />
        </template>

        <v-date-picker show-adjacent-months width="350"
          @update:model-value="startDate = formatDate($event)"
          :color="userSession.currentTheme.value === 'light' ? '#185E5E' : '#0E3838'"></v-date-picker>
      </v-menu>

      <Icon icon="mdi:arrow-left-right" width="30"
        :class="userSession.currentTheme.value === 'light' ? 'text-gray-800' : 'text-gray-600'"></Icon>

      <v-menu v-model="endMenu" transition="slide-y-transition" min-width="auto" :close-on-content-click="true">
        <template v-slot:activator="{ props }">
          <v-text-field label="End Date" placeholder="YYYY-mm-dd" hide-details="auto" variant="outlined" v-bind="props"
            v-model="endDate" :readonly="true" />
          <Icon icon="mdi:date-range" width="30" class="my-auto"
            :class="userSession.currentTheme.value === 'light' ? 'text-gray-800' : 'text-gray-600'"></Icon>
        </template>

        <v-date-picker show-adjacent-months width="350"
          @update:model-value="endDate = formatDate($event)"
          :color="userSession.currentTheme.value === 'light' ? '#185E5E' : '#0E3838'"></v-date-picker>
      </v-menu>
    </section>
  </main>
</template>

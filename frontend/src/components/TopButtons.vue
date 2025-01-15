<script setup lang="ts">
import { userSession } from '@/services/sessionService';
import { Icon } from '@iconify/vue';
import { ref, watch } from 'vue';
import { debounce } from '@/services/sessionService';
import { DateTime } from 'luxon';

// Define the emit event
const emit = defineEmits<{
  (e: 'onSearchEmit', term: string, start: string, end: string): void; // 'search' event with text payload
}>();

const term = ref('');

const startDate = ref();
const endDate = ref();

const startMenu = ref(false);
const endMenu = ref(false);

const formatDate = (date: any) => {
  const d = DateTime.fromJSDate(new Date(date));
  return d.toFormat('yyyy/MM/dd');
};

const onSearchApply = debounce((text: string) => {
  emit('onSearchEmit', text === null ? '' : text, startDate.value, endDate.value);
}, 300);

watch([startDate, endDate], () => {
  if (startDate.value && endDate.value)
    onSearchApply(term.value);
})

</script>

<template>
  <main class="flex flex-col gap-[15px] md:justify-between w-full rounded-md md:gap-0 md:flex-row lg:flex-row">
    <!-- Search Section -->
    <section class="flex items-center gap-[10px] w-full lg:w-[425px] md:w-[325px]">
      <Icon icon="mdi:magnify" width="30"
        :class="userSession.currentTheme.value === 'light' ? 'text-light-contrast' : 'text-dark-contrast'" />
      <v-text-field bg-color="transparent" hide-details="auto" clearable label="Search"
        placeholder="Search on 'Subject', 'From', 'To' and 'Body' fields" variant="outlined" v-model="term"
        v-on:update:model-value="onSearchApply" />
    </section>

    <!-- Dates Section -->
    <section class="flex items-center gap-[10px] w-full lg:w-[600px] md:w-[400px] sm:w-full">
      <v-menu v-model="startMenu" transition="slide-y-transition" min-width="auto" :close-on-content-click="false">
        <template v-slot:activator="{ props }">
          <Icon icon="mdi:date-range" width="30" class="my-auto"
            :class="userSession.currentTheme.value === 'light' ? 'text-gray-800' : 'text-gray-600'" />
          <v-text-field label="Start Date" placeholder="YYYY-mm-dd" hide-details="auto" variant="outlined"
            v-bind="props" v-model="startDate" :readonly="true" />
        </template>

        <v-date-picker show-adjacent-months width="350" @update:model-value="startDate = formatDate($event)"
          :color="userSession.currentTheme.value === 'light' ? '#185E5E' : '#0E3838'" />
      </v-menu>

      <Icon icon="mdi:arrow-left-right" width="30"
        :class="userSession.currentTheme.value === 'light' ? 'text-gray-800' : 'text-gray-600'" />

      <v-menu v-model="endMenu" transition="slide-y-transition" min-width="auto"  :close-on-content-click="false">
        <template v-slot:activator="{ props }">
          <v-text-field label="End Date" placeholder="YYYY-mm-dd" hide-details="auto" variant="outlined" v-bind="props"
            v-model="endDate" :readonly="true" />
          <Icon icon="mdi:date-range" width="30" class="my-auto"
            :class="userSession.currentTheme.value === 'light' ? 'text-gray-800' : 'text-gray-600'" />
        </template>

        <v-date-picker show-adjacent-months width="350" @update:model-value="endDate = formatDate($event)"
          :color="userSession.currentTheme.value === 'light' ? '#185E5E' : '#0E3838'" />
      </v-menu>
    </section>
  </main>
</template>

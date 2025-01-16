<script setup lang="ts">
import { userSession } from '@/services/sessionService';
import { Icon } from '@iconify/vue';
import { ref, watch } from 'vue';
import { debounce } from '@/services/sessionService';
import { DateTime } from 'luxon';

// Define the emit event
const emit = defineEmits<{
  (e: 'onSearchEmit', term: string, type: 'id' | 'term', start: string, end: string): void; // 'search' event with text payload
}>();

const term = ref('')
const searchType = ref<'id' | 'term'>('term')

const startDate = ref()
const endDate = ref()

const startMenu = ref(false)
const endMenu = ref(false)


const formatDate = (date: any) => {
  const d = DateTime.fromJSDate(new Date(date))
  return d.toFormat('yyyy/MM/dd')
}

const onSearchApply = debounce((text: string) => {
  emit('onSearchEmit', text === null ? '' : text, searchType.value, startDate.value, endDate.value)
}, 300)

watch([startDate, endDate], () => {
  if (startDate.value && endDate.value)
    onSearchApply(term.value)
})

watch([searchType], () => {
  if (searchType.value == 'term')
    onSearchApply(term.value)
})

</script>

<template>
  <main class="flex flex-col gap-[15px] md:justify-between w-full rounded-md md:gap-0 md:flex-row lg:flex-row">
    <!-- Search Section -->
    <section class="flex items-center gap-[10px] xs:w-full sm:w-full lg:w-[425px] md:w-[325px]">
      <div class="flex gap-[10px] xs:w-1/2 sm:w-1/2 items-center">
        <Icon icon="mdi:magnify" width="30"
          :class="userSession.currentTheme.value === 'light' ? 'text-light-contrast' : 'text-dark-contrast'" />
        <v-text-field bg-color="transparent" hide-details="auto" clearable :label="$t('home.search')"
          :placeholder="$t('home.placeholders.term')" variant="outlined" v-model="term"
          v-on:update:model-value="onSearchApply" />
      </div>

      <div class="sm:w-1/2 xs:w-1/2 md:w-[100px]">
        <v-select v-model="searchType" :label="$t('home.search_type.label')" hide-details variant="outlined" :items="[
          { text: $t('home.search_type.term'), value: 'term' },
          { text: $t('home.search_type.id'), value: 'id' }
        ]" item-value="value" item-title="text"></v-select>
      </div>
    </section>

    <!-- Dates Section -->
    <section class="flex items-center gap-[10px] w-full lg:w-[600px] md:w-[400px] sm:w-full">
      <v-menu v-model="startMenu" transition="slide-y-transition" min-width="auto" :close-on-content-click="false">
        <template v-slot:activator="{ props }">
          <Icon icon="mdi:date-range" width="30" class="my-auto"
            :class="userSession.currentTheme.value === 'light' ? 'text-gray-800' : 'text-gray-600'" />
          <v-text-field :label="$t('home.start_date')" placeholder="YYYY-mm-dd" hide-details="auto" variant="outlined"
            v-bind="props" v-model="startDate" :readonly="true" />
        </template>

        <v-date-picker :hide-header="true" show-adjacent-months width="350"
          @update:model-value="startDate = formatDate($event)"
          :color="userSession.currentTheme.value === 'light' ? '#185E5E' : '#0E3838'" />
      </v-menu>

      <Icon icon="mdi:arrow-left-right" width="30"
        :class="userSession.currentTheme.value === 'light' ? 'text-gray-800' : 'text-gray-600'" />

      <v-menu v-model="endMenu" transition="slide-y-transition" min-width="auto" :close-on-content-click="false">
        <template v-slot:activator="{ props }">
          <v-text-field :label="$t('home.end_date')" placeholder="YYYY-mm-dd" hide-details="auto" variant="outlined"
            v-bind="props" v-model="endDate" :readonly="true" />
          <Icon icon="mdi:date-range" width="30" class="my-auto"
            :class="userSession.currentTheme.value === 'light' ? 'text-gray-800' : 'text-gray-600'" />
        </template>

        <v-date-picker :hide-header="true" show-adjacent-months width="350"
          @update:model-value="endDate = formatDate($event)"
          :color="userSession.currentTheme.value === 'light' ? '#185E5E' : '#0E3838'" />
      </v-menu>
    </section>
  </main>
</template>

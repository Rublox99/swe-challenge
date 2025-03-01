<script setup lang="ts">
import { RouterView } from 'vue-router'

import { Icon } from '@iconify/vue'
import { userSession, toggleTheme } from './services/sessionService';
import { useLocale, useTheme } from 'vuetify/lib/framework.mjs';
import { ref } from 'vue';
import { useI18n } from 'vue-i18n';

const { locale } = useI18n() /* Templates Locale */
const { current } = useLocale() /* Vuetify Locale */

const theme = useTheme()

const language = ref('en')

const switchTheme = () => {
  toggleTheme()
  theme.global.name.value = userSession.currentTheme.value!
}

const onLanguageChange = () => {
  const lang = language.value

  locale.value = lang
  current.value = lang
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

      <section class="flex gap-2.5 items-center justify-center w-auto h-auto">
        <v-select class="text-gray-500" density="compact" :hide-details="true" base-color="white"
          color="white" :item-color="userSession.currentTheme.value === 'light' ? 'black' : 'white'" v-model="language"
          v-on:update:model-value="onLanguageChange" :label="$t('home.language')" :items="[
            { lang: 'en', label: 'English' },
            { lang: 'es', label: 'Español' }
          ]" :item-title="'label'" :item-value="'lang'" variant="outlined">
        </v-select>

        <v-tooltip
          :text="userSession.currentTheme.value === 'light' ? $t('tooltips.dark_theme') : $t('tooltips.light_theme')"
          location="bottom">
          <template v-slot:activator="{ props }">
            <Icon @click="switchTheme" v-bind="props"
              class="transition-all hover:opacity-80 hover:cursor-pointer text-light-inv_contrast"
              :icon="userSession.currentTheme.value === 'light' ? 'solar:moon-bold' : 'solar:sun-bold'" width="32">
            </Icon>
          </template>
        </v-tooltip>
      </section>
    </header>

    <RouterView />
  </main>
</template>
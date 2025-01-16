<script setup>
import { userSession } from '@/services/sessionService';
</script>

<template>
    <section class="w-full h-auto flex flex-col md:h-[50px] md:flex-row md:justify-between">
        <span
            class="text-center md:text-left px-4 py-2 rounded-full text-cyan-50 bg-gradient-to-r animate-gradient-bounce bg-400% from-red-800 via-blue-500 to-indigo-600">
            {{$t('home.done_by')}} <a href="https://github.com/Rublox99" target="_blank" class="font-semibold">@Rublox99</a>
        </span>
        <a href="https://rubendev.online" target="_blank" class="p-2 font-semibold text-center" :class="{
            'text-light-contrast': userSession.currentTheme.value === 'light',
            'text-dark-contrast': userSession.currentTheme.value === 'dark'
        }">
            rubendev.online
        </a>
    </section>
</template>

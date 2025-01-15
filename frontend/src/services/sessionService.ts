import { ref } from 'vue'
import type { UserSession } from '@/interfaces/UserSession'

const userSession: UserSession = {
    currentTheme: ref<'light' | 'dark'>('light'),
}

const toggleTheme = () => {
    const activeTheme = userSession.currentTheme.value === 'light' ? 'dark' : 'light'
    userSession.currentTheme.value = activeTheme
}

export {
    userSession,
    toggleTheme
}
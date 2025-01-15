import { ref } from 'vue'
import type { UserSession } from '@/interfaces/UserSession'

const userSession: UserSession = {
    currentTheme: ref<'light' | 'dark'>('light'),
}

const toggleTheme = () => {
    const activeTheme = userSession.currentTheme.value === 'light' ? 'dark' : 'light'
    userSession.currentTheme.value = activeTheme
}

const debounce = (func: Function, delay: number) => {
    let timeout: ReturnType<typeof setTimeout>

    return (...args: any[]) => {
        clearTimeout(timeout)
        timeout = setTimeout(() => func(...args), delay)
    }
}

export {
    userSession,
    toggleTheme,
    debounce
}
import { ref } from "vue"

interface UserSession {
    currentTheme: ReturnType<typeof ref<'light' | 'dark'>>
}

export type {
    UserSession
}
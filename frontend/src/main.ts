import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

// Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'

import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import '@mdi/font/css/materialdesignicons.css'

import { en, es } from 'vuetify/locale'
import { createI18n } from 'vue-i18n'
import { enI18n } from './assets/i18n/en'
import { esI18n } from './assets/i18n/es'

const app = createApp(App)
const i18n = createI18n({
    locale: 'en',
    fallbackLocale: 'en',
    messages: {
        en: enI18n,
        es: esI18n
    }
})

const vuetify = createVuetify({
    components,
    directives,
    icons: {
        defaultSet: 'mdi',
    },
    locale: {
        locale: 'en',
        fallback: 'es',
        messages: { en, es }
    }
})

app.use(router)
app.use(vuetify)
app.use(i18n)

app.mount('#app')

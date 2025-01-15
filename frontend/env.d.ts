/// <reference types="vite/client" />

// Declaration for importing .vue files
declare module '*.vue' {
    import { DefineComponent } from 'vue'
    const component: DefineComponent<{}, {}, any>
    export default component
  }
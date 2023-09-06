import { createApp } from 'vue'
import App from './App.vue'
import pinia from './store'
import { initStores } from './store'
const app = createApp(App)

app.use(pinia)
app.mount('#app')

initStores()


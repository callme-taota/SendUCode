import { createRouter , createWebHashHistory  } from "vue-router"

import home from "../pages/home/index.vue"


const routes = [
  { path : "/home" , component:home },
  { path : "/" , component:home },

]
const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router

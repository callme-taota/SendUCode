import { createRouter , createWebHashHistory  } from "vue-router"

import home from "../pages/home/index.vue"
import setting from "../pages/setting/index.vue"

const routes = [
  { path : "/home" , component:home },
  { path : "/" , component:home },
  { path : "/setting" , component:setting },

]
const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router

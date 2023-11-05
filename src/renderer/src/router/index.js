import { createRouter , createWebHashHistory  } from "vue-router"

import home from "../pages/home/index.vue"
import find from "../pages/find/index.vue"
import setting from "../pages/setting/index.vue"
import user from "../pages/user/index.vue"
import tool from "../pages/tool/index.vue"

const routes = [
  { path : "/home" , component:home },
  { path : "/" , component:home },
  { path : "/find" , component:find },
  { path : "/setting" , component:setting },
  { path : "/user" , component:user },
  { path : "/tool" , component:tool },

]
const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router

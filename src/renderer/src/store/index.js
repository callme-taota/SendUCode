import { createPinia } from 'pinia'

import { useAppStore } from './modules/app';

const pinia = createPinia()
export default pinia;


export function initStores(){
  const appStore = useAppStore()
}

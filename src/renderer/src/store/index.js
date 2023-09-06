import { createPinia } from 'pinia'

import { usePlaylistStore } from './modules/playlist';
import { useAppStore } from './modules/app';

const pinia = createPinia()
export default pinia;


export function initStores(){
  const appStore = useAppStore()
  const PlaylistStore = usePlaylistStore()
}

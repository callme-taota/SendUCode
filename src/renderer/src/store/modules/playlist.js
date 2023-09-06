import { defineStore } from "pinia";

export const usePlaylistStore = defineStore("PlaylistStore",{
  state:() => ({
    list : []
  })
})

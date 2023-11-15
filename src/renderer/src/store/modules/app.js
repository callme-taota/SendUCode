import { defineStore } from "pinia";
import { GetSystemType } from "../../utils"

const matchMedia = window.matchMedia('(prefers-color-scheme: dark)')
export const changeListener = matchMedia.addEventListener("change", function() {
  if (this.matches) {
    document.documentElement.setAttribute('theme', 'dark');
  } else {
    document.documentElement.removeAttribute('theme')
  }
})

export const useAppStore = defineStore("appStore", {
  state: () => ({
    platform: GetSystemType(),
    isDeviceDarkTheme: matchMedia.matches,
    themeSetting: "auto",
  }),
  actions: {
    SetThemeLight() {
      this.themeSetting = "light"
      document.documentElement.removeAttribute('theme')
      CleanAutoChangeTheme()
    },
    SetThemeDark() {
      this.themeSetting = "dark"
      document.documentElement.setAttribute('theme', 'dark');
      CleanAutoChangeTheme()
    },
    SetThemeAuto() {
      this.themeSetting = "auto"
      if (this.isDeviceDarkTheme == true) {
        document.documentElement.setAttribute('theme', 'dark');
      } else {
        document.documentElement.removeAttribute('theme')
      }
      this.AutoChangeTheme()
    },
    AutoChangeTheme() {
      changeListener
    },
    CleanAutoChangeTheme(){
      matchMedia.removeEventListener(changeListener)
    }
  },
  getters: {
    isDarkTheme() {
      if (this.themeSetting == "auto") return this.isDeviceDarkTheme
      return this.themeSetting == "dark" ? true : false
    }
  }
})

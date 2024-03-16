import { defineStore } from "pinia";
import { GetSystemType } from "../../utils"

const matchMedia = window.matchMedia('(prefers-color-scheme: dark)')
const changeListener = function () {
  if (matchMedia.matches) {
    document.documentElement.setAttribute('theme', 'dark');
  } else {
    document.documentElement.removeAttribute('theme');
  }
};

export const useAppStore = defineStore("appStore", {
  state: () => ({
    platform: GetSystemType(),
    isDeviceDarkTheme: matchMedia.matches,
    themeSetting: "auto"
  }),
  actions: {
    SetThemeLight() {
      this.themeSetting = "light"
      document.documentElement.removeAttribute('theme')
      this.CleanAutoChangeTheme()
    },
    SetThemeDark() {
      this.themeSetting = "dark"
      document.documentElement.setAttribute('theme', 'dark');
      this.CleanAutoChangeTheme()
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
      matchMedia.addEventListener("change", changeListener)
    },
    CleanAutoChangeTheme() {
      matchMedia.removeEventListener("change", changeListener)
    },
    SetTheme(value) {
      switch (value) {
        case 1:
          this.SetThemeLight()
          return
        case 2:
          this.SetThemeDark()
          return
        case 3:
          this.SetThemeAuto()
          return
      }
      return
    }
  },
  getters: {
    isDarkTheme() {
      if (this.themeSetting == "auto") return this.isDeviceDarkTheme
      return this.themeSetting == "dark" ? true : false
    },
    btnColor(active) {
      let isdark = false
      if (this.themeSetting == "auto") { isdark = this.isDeviceDarkTheme }
      else { isdark = this.themeSetting == "dark" ? true : false }
      if (isdark) {
        if (active) {
          return "#ffffff"
        } else {
          return "#CCCCCC"
        }
      } else {
        if (active) {
          return "#222222"
        } else {
          return "#000000"
        }
      }
    },
    qrColor_blank() {
      let isdark = false
      if (this.themeSetting == "auto") { isdark = this.isDeviceDarkTheme }
      else { isdark = this.themeSetting == "dark" ? true : false }
      if (isdark) {
        return "rgba(255,255,255,0)"
      } else {
        return "rgba(255,255,255,0)"
      }
    },
    qrColor_fill() {
      let isdark = false
      if (this.themeSetting == "auto") { isdark = this.isDeviceDarkTheme }
      else { isdark = this.themeSetting == "dark" ? true : false }
      if (isdark) {
        return "rgba(255,255,255,0.5)"
      } else {
        return "rgba(0,0,0,0.4)"
      }
    }
  }
})

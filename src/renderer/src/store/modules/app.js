import { defineStore } from "pinia";
import { GetSystemType } from "../../utils"

export const useAppStore = defineStore("appStore",{
  state:() => ({
    platform:GetSystemType(),
    isDeviceDarkTheme : window.matchMedia("(prefers-color-scheme: dark)").matches,
    themeSetting : "auto"
  }),
  actions:{
    SetThemeLight(){
      this.themeSetting = "light"
      document.documentElement.removeAttribute('theme')
    },
    SetThemeDark(){
      this.themeSetting = "dark"
      document.documentElement.setAttribute('theme', 'dark');
    },
    SetThemeAuto(){
      this.themeSetting = "auto"
      if (this.isDeviceDarkTheme == true) {
          document.documentElement.setAttribute('theme', 'dark');
      }else{
          document.documentElement.removeAttribute('theme')
      }
    }
  },
  getters:{
    isDarkTheme(){
      if(this.themeSetting == "auto") return this.isDeviceDarkTheme
      return this.themeSetting == "dark" ? true : false
    }
  }
})

import { defineStore } from "pinia";
import { CheckUsingSession } from "../../api/user";
import { getMsg } from "../../api/msg";
import { MsgWs } from "../../api";
import { CreateUser,DeleteUser } from "../../api/user";

const getUserData = () => {
  return window.electronHandler.ipcRenderer.getStoreValue("userData");
}
const setUserID = (userID = "") => {
  let userData = getUserData()
  userData.userID = userID;
  window.electronHandler.ipcRenderer.setStoreValue("userData", userData);
}
const setSession = (session = "") => {
  let userData = getUserData()
  userData.session = session;
  window.electronHandler.ipcRenderer.setStoreValue("userData", userData);
}

export const useUserStore = defineStore("userStore", {
  state: () => ({
    userID: getUserData().userID,
    session: getUserData().session,
    conn: false,
    checked: false,
    msgList: [],
    noticeList: [],
  }),
  actions: {
    async checkSession() {
      let session = this.session
      let res = await CheckUsingSession(session)
      if (res.ok == "false") {
        setUserID()
        setSession()
        this.userID = getUserData().userID
        this.session = getUserData().session
        this.checked = false
        return false
      } else {
        this.checked = true
        return true
      }
    },
    async getMsgs(_, flag) {
      if (flag == false) {
        this.conn = false
      }
      let res = await getMsg()
      if (res.ok == "false") {
        this.msgList = []
        return false
      }
      this.msgList = res
    },
    async listiner() {
      MsgWs.Start("/user/ws", {
        "session": this.session
      })
      this.conn = true
      MsgWs.Subscribe(this.getMsgs)
    },
    async setUserID(userid) {
      let res = await CreateUser(userid)
      if (res.ok != "true") {
        return false
      }
      setUserID(userid)
      setSession(res.session)
      this.session = getUserData().session
      this.userID = getUserData().userID
      this.checked = true
    },
    async delUserData() {
      await DeleteUser(this.session)
      setUserID("")
      setSession("")
      this.userID = getUserData().userID
      this.session = getUserData().session
      this.checked = false
    },
    addNotice(text) {
      let t = new Date().getTime();
      let noticeList = this.noticeList;
      let obj = { text, time: t }
      noticeList.push(obj);
      this.noticeList = noticeList
      setTimeout(() => {
        this.DelNotice(t);
      }, 2000);
    },
    DelNotice(time) {
      let noticeList = this.noticeList;
      noticeList = noticeList.filter(obj => obj.time !== time);
      this.noticeList = noticeList
    }
  },
  getters: {

  }
})

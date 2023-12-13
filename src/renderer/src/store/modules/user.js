import { defineStore } from "pinia";
import { CheckUsingSession } from "../../api/user";
import { getMsg } from "../../api/msg";
import { MsgWs } from "../../api";
import { CreateUser } from "../../api/user";

const getUserData = () => {
  return window.electronHandler.ipcRenderer.getStoreValue("userData");
}
const setUserID = (userID) => {
  let userData = getUserData()
  userData.userID = userID;
  window.electronHandler.ipcRenderer.setStoreValue("userData",userData);
}
const setSession = (session) => {
  let userData = getUserData()
  userData.session = session;
  window.electronHandler.ipcRenderer.setStoreValue("userData",userData);
}

export const useUserStore = defineStore("userStore", {
  state: () => ({
    userID: getUserData().userID,
    session: getUserData().session,
    conn: false,
    msgList: [],
    noticeList: [],
  }),
  actions: {
    async checkSession() {
      let res = await CheckUsingSession(this.session)
      if (res.ok == false) {
        setUserID()
        setSession()
        this.userID = getUserData().userID
        this.session = getUserData().session
        return false
      } else {
        return true
      }
    },
    async getMsgs(_, flag) {
      if (flag == false) {
        this.conn = false
      }
      let res = await getMsg()
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
      if (res.ok != true || res.ok != "true") {
        return false
      }
      setSession(res.session)
      this.session = getUserData().session
    },
    addNotice(text) {
      let t = new Date().getTime();
      let noticeList = this.noticeList;
      let obj = { text, time: t }
      noticeList.push(obj);
      this.noticeList = noticeList
      setTimeout(() => {
        this.DelMessage(t);
      }, 2000);
    },
    DelMessage(time) {
      let noticeList = this.noticeList;
      noticeList = noticeList.filter(obj => obj.time !== time);
      this.noticeList = noticeList
    }


  },
  getters: {

  }
})

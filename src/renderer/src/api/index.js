import axios from 'axios';
import { storeToRefs } from 'pinia';
import { useUserStore } from '../store/modules/user';

const URL = 'http://localhost:3003'

/**
 * 核心 推送异步数据 返回promise对象，需要使用then进行读取
 * @param {string} apiaddr
 * @returns {object} data
 */
export const AxiosPost = async (apiaddr, data) => {
  const userStore = useUserStore()
  const { session } = storeToRefs(userStore)
  let posturl = URL + apiaddr;
  try {
    const res = await axios.post(posturl, data, {
      headers: {
        'session': session.value,
      }
    });
    return res
  } catch (error) {
    console.error(error);
    throw error;
  }
}

/**
 * 核心 发送异步GET请求 返回promise对象，需要使用then进行读取
 * @param {string} apiaddr - API 地址
 * @param {object} params - 查询参数对象
 * @returns {Promise<object>} - Promise 对象，解析为响应数据
 */
export const AxiosGet = async (apiaddr, params = {}) => {
  const userStore = useUserStore()
  const { session } = storeToRefs(userStore)
  const getUrl = URL + apiaddr;
  try {
    const response = await axios.get(getUrl, {
      params,
      headers: {
        'session': session.value,
      }
    });
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
};

class WS {
  constructor(obj) {
    let wsURL = obj.wsURL
    if (wsURL == "") {
      return false;
    }
    this.wsURL = wsURL
    this.subscribers = []
    this.status = false
    this.context = {}
  }

  Start = (herf, obj) => {
    if(this.status) return
    let url = this.wsURL + herf + "?"
    for (let key in obj) {
      url += key + "=" + obj[key] + "&"
    }
    this.context = {herf,obj}
    const socket = new WebSocket(url)
    this.socket = socket
    this.on()
    this.listener()
    this.off()
    this.err()
  }

  Reconnect = () => {
    if(this.statue) return
    let {herf,obj} = this.context
    let url = wsURL + herf + "?"
    for (let key in obj) {
      url += key + "=" + obj[key] + "&"
    }
    const socket = new WebSocket(url)
    this.socket = socket
    this.on()
    this.listener()
    this.off()
    this.err()
  }

  on = () => {
    // 在连接打开时，设置会话标识到WebSocket的headers中
    this.socket.addEventListener('open', (event) => {
      console.log(event)
      this.status = true
    });
  }

  listener = () => {
    this.socket.addEventListener('message', (event) => {
      this.notifySubscribers(event);
    });
  }

  off = () => {
    this.socket.addEventListener('close', (event) => {
      this.status = false
      this.notifySubscribers(event);
    });
  }

  err = () => {
    this.socket.addEventListener('error', (event) => {
      this.status = false
      this.notifySubscribers(event);
      console.error('WebSocket 连接出现错误:', event);
    });
  }

  notifySubscribers = (event) => {
    this.subscribers.forEach(subscriber => {
      try {
        if (typeof subscriber.fn === 'function') {
          subscriber.fn(event,this.status);
        }
      } catch (e) {
        console.error(e);
      }
    });
  }

  Subscribe = (fn, key) => {
    const characters = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
    const randomText = characters[Math.floor(Math.random() * characters.length)] + characters[Math.floor(Math.random() * characters.length)] + characters[Math.floor(Math.random() * characters.length)] + characters[Math.floor(Math.random() * characters.length)];
    if (key === "" || key === undefined || key === null) key = "rrvS__" + new Date().getSeconds() + randomText;
    fn.key = key;
    let obj = { fn, key };
    try {
      if (typeof fn === 'function') {
        this.subscribers.push(obj);
      }
    } catch (e) {
      console.error(e);
    }
    return obj
  }
}

export const MsgWs = new WS({
  wsURL : 'ws://localhost:3003',
});


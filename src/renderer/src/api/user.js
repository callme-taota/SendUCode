import { AxiosPost, AxiosDelete } from "./index"

export const CreateUser = async (userid) =>{
  let fullURL = "/user?userid=" + userid
  // let fullURL = "/user"
  let res = await AxiosPost(fullURL,{userid})
  return res.data
}

export const CheckUsingSession = async (session) =>{
  let fullURL = "/user/check?session=" + session
  let res = await AxiosPost(fullURL,{session})
  return res.data
}

export const DeleteUser = async (session) => {
  let fullURL = "/user?session=" + session
  let res = await AxiosDelete(fullURL,{session})
  return res.data
}


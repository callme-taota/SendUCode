import { AxiosPost, AxiosDelete } from "./index"

export const CreateUser = async (userid) =>{
  let fullURL = "/user?userid=" + userid
  let res = await AxiosPost(fullURL)
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

export const TestPost = async () => {
  let fullURL = "/test?a=1&b=2"
  let obj = {a : 1, b :"xxx"}
  let res = await AxiosPost(fullURL,obj)
  return res
}

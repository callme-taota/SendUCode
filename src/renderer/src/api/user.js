import { AxiosPost } from "./index"

export const CreateUser = async (userid) =>{
  let res = await AxiosPost("/user",{userid})
  return res
}

export const CheckUsingSession = async (session) =>{
  let res = await AxiosPost("/user/check",{session})
  return res
}

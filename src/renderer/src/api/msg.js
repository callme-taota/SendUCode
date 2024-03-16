import { AxiosPost,AxiosGet } from "./index"

export const getMsg = async (limit) =>{
  let res = await AxiosGet('/msg/',{limit})
  return res
}

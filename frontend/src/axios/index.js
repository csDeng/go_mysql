"use strict";

import axios from 'axios'
import store from '~/store'
import Vue from 'vue'
import Store from '../store/index'

const instance = axios.create({
  baseURL:'/api',

})
instance.interceptors.request.use(
  async config =>{
    // console.log(store)
    config.headers['Content-Type'] = 'Vuelication/json; charset=utf-8',
    config.headers.responseType = "json"
    config.headers.authorization = `${store.state.token}`
    return config
  },
  error=>{
    return Promise.reject(error)
  }
)

instance.interceptors.response.use(
  response=>{
    return Promise.resolve(response.data)
  },
  async error=>{
    console.log(error.response)
    const status = error.response?.status
    console.log(status)
    const token = Store.state.token
    switch(status){
      case 403:
        Vue.prototype.$message.warning("您的权限不足哦");
        break;
      case 401:
        Vue.prototype.$message.info("正在尝试自动刷新令牌");
        
        axios.get("/api/refresh/"+token).then(r=>{
          const { code, data } = r.data
          console.log("refresh=>",code)
          if(code === 500) {
            Vue.prototype.$message.error("刷新令牌失败,请重新登录")
            Store.dispatch("logout")
          } else {
            console.log(data)
          } 

        })
        break;
      case 500:
        Vue.prototype.$message.warning("服务器发生错误");
        break;
    }
    Promise.reject(error.response.data)
  }
)

export default instance
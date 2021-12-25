"use strict";
import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate';
Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    isLogin: false,
    token: null,
    uid: null,
    username: null,
    level: ''
  },
  plugins: [createPersistedState()],
  mutations: {
    login(state, data) {
      state.isLogin = true
      const { user: { id: uid, username, level }, token } = data
      state.token = token
      state.uid = uid
      state.username = username
      state.level = level
      // console.log("登录", state)
    },
    logout(state) {
      state.isLogin = false
      state.token = null
      state.uid = ""
      state.username = ""
      state.level = ""
      // console.log("退出登录")
    }
  },
  actions: {
    login(context, data) {
      context.commit('login', data)
    },
    logout({ commit }) {
      commit('logout')
    }
  }
})

export default store
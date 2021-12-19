import Vue from 'vue'
import App from './App.vue'
import less from 'less'
import router from './router'
import store from '~/store/index.js'
import api from '~/api'
import '../../lazy.js'
import 'ant-design-vue/dist/antd.css'
Vue.config.productionTip = false
Vue.prototype.$api = api

Vue.use(less)

const app = new Vue({
  render: h => h(App),
  router,
  store
}).$mount('#app');
export default app;

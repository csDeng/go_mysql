import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

export default new VueRouter({
  mode: 'history',
  base:'/app',
  routes: [
    {
      path:'/',
      component: () => import('~/views/web/layout.vue'),
      redirect:'index',
      children:[
        {
          path: 'index',
          component: ()=>import('~/views/web/Index')
        },
        {
          path:'Category',
          component: ()=>import('~/views/web/Category')
        },
        {
          path:'Article',
          component: ()=>import('~/views/web/Article')
        },
        // {
        //   path:'/Tags',
        //   component: ()=>import('~/views/web/Tags')
        // },
        {
          path: 'Mine',
          component: ()=>import('~/views/web/Mine')
        },
        {
          path:'Update',
          component: ()=>import('~/views/web/Update')
        },
        {
          path:'Detail',
          component: ()=>import('@/articleDetail')
        },
      ]
    },
    {
      path: '/Login',
      component: ()=>import('~/views/web/Login')
    },
    {
      path: '/Register',
      component: ()=>import('~/views/web/Register')
    },
    {
      path: '/403',
      component: () => import('~/views/other/403.vue')
    },
    {
      path: '*',
      component: () => import('~/views/other/404.vue')
    }
  ]
})
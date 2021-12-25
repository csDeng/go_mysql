import Vue from 'vue';
import vueRouter from 'vue-router';
import layout from '@/admin/layout/layout.vue';
Vue.use(vueRouter)
const router = new vueRouter({
    mode: 'history',
    base:'/admin',
    routes: [
        {
            path: '/',
            component: layout,
            redirect:'index',
            children: [
                {
                    path: 'index',
                    component: () => import('~/views/admin/index.vue')
                },
                {
                    path: 'tag',
                    component: () => import('~/views/admin/tags/TagsAdmin.vue')
                },
                {
                    path: 'article',
                    component: () => import('~/views/admin/article/articlesAdmin.vue')
                },
                {
                    path: 'Edit',
                    component: () => import('~/views/admin/article/Edit')
                },
                {
                    path: 'user',
                    component: () => import('~/views/admin/usersAdmin.vue')
                },
                {
                    path:'about',
                    component: ()=> import('~/views/admin/about.vue')
                }
            ]
        },
        {
            path:'/403',
            component: () => import('~/views/other/403.vue')
        },
        {
            path:'*',
            component: () => import('~/views/other/404.vue')
        }
    ]
})

export default router;
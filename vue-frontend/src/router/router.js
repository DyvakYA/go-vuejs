import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../components/Home'
import Users from '../components/Users'

Vue.use(VueRouter)

const router = new VueRouter({
    mode: 'history',
    routes: [
        {
            path: '/',
            alias: '/home',
            name: 'Home',
            component: Home,
            meta: {title: 'Home'}
        },
        {
            path: '/users',
            name: 'Users',
            component: Users,
            meta: {title: 'Users'}
        }
    ]
})

export default router

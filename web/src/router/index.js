import Vue from 'vue'
import Router from 'vue-router'
import Hello from '@/components/Hello'
import Login from '@/components/Login'
import authMiddleware from '@/middleware/authMiddleware'

Vue.use(Router);

const router = new Router({
    mode: 'history',
    routes: [
        {
            path: '/',
            name: 'Hello',
            component: Hello,
            meta: { requiresAuth: true },
        },
        {
            path: '/login',
            name: 'Login Area',
            component: Login
        }
    ]
})

router.beforeEach(authMiddleware)

export default router

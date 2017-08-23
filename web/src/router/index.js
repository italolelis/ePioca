import Vue from 'vue'
import authMiddleware from '@/middleware/authMiddleware'
import Router from 'vue-router';
import Hello from '@/pages/Hello';
import Login from '@/pages/Login';
import Auction from '@/pages/auction'

Vue.use(Router);

const router = new Router({
    mode: 'history',
    routes: [
        {
            path: '/',
            name: 'Hello',
            meta: { requiresAuth: true },
            component: Hello
        },
        {
            path: '/login',
            name: 'Login Area',
            component: Login
        },
        {
            path: '/auction',
            name: 'Auction',
            component: Auction,
            meta: { requiresAuth: true }
        }
    ]
});

router.beforeEach(authMiddleware);

export default router

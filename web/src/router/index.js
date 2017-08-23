import Vue from 'vue'
import Router from 'vue-router'

import authMiddleware from '@/middleware/authMiddleware'
import roleMiddleware from '@/middleware/roleMiddleware'

import Login from '@/pages/Login'
import Dashboard from '@/pages/Dashboard'
import CreateAuction from '@/pages/CreateAuction'
import ViewAuction from '@/pages/CreateAuction'

Vue.use(Router);

const router = new Router({
    mode: 'history',
    routes: [
        {
            path: '/',
            name: 'Dashboard',
            meta: { requiresAuth: true },
            component: Dashboard
        },
        {
            path: '/login',
            name: 'Login Area',
            component: Login
        },
        {
            path: '/auctions/create',
            name: 'Create Auction',
            component: CreateAuction,
            meta: { requiresAuth: true, buyerOnly: true }
        },
        {
            path: '/auctions/:id',
            name: 'View Auction',
            component: ViewAuction,
            meta: { requiresAuth: true, supplierOnly: true }
        }
    ]
});

router.beforeEach(authMiddleware);
router.beforeEach(roleMiddleware);

export default router

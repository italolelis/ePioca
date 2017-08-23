import { getUserRole } from '@/api'

const middleware = (to, from, next) => {
    if (to.matched.some(route => route.meta.buyerOnly)) {
        if (getUserRole() !== 'buyer') {
            return next({ path: '/' })
        }
    }

    if (to.matched.some(route => route.meta.supplierOnly)) {
        if (getUserRole() === 'buyer') {
            return next({ path: '/' })
        }
    }

    return next()
}

export default middleware

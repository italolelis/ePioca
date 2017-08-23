// Stub to simulate being logged out
const getAccessToken = () => false

const middleware = (to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAuth)) {
        if (!getAccessToken()) {
            return next({
                path: '/login',
                query: { redirect: to.fullPath },
            })
        }

        return next()
    }

    return next()
}

export default middleware

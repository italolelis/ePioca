import client from '@/api'

const TAPIOCA_URI = 'https://gw-staging.hellofresh.com/scm'

export const getIngredientsForWeekAndDc = (week, dc) =>
    client.get(`${TAPIOCA_URI}/product-summaries`, {
        params: { week, dc }
    })

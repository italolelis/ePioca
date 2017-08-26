import client from '@/api'
import config from '@/config'

const TAPIOCA_URI = `${config.gateway.uri}/scm`

export const getIngredientsForWeekAndDc = (week, dc) =>
    client.get(`${TAPIOCA_URI}/product-summaries`, {
        params: { week, dc }
    })

import client from '@/api'
import { gateway } from '@/config'

const TAPIOCA_URI = `${gateway.uri}/scm`

export const getIngredientsForWeekAndDc = (week, dc) =>
    client.get(`${TAPIOCA_URI}/product-summaries`, {
        params: { week, dc }
    })

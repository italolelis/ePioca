import axios from 'axios'
import config from '@/config'
import { setAccessToken } from '@/api'

const SIGNIN_SERVICE_URI = '/login'

export const login = (username, password, country) => {
    // Remove the access token so it doesn't send with the request
    setAccessToken(null)

    return axios.post(
        `${config.gateway.uri}${SIGNIN_SERVICE_URI}?country=${country}`,
        { username, password }
    )
}

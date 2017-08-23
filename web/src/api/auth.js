import axios from 'axios'
import config from '@/config'
import { setAccessToken } from '@/api'

export const login = (username, password) => {
    // Remove the access token so it doesn't send with the request
    setAccessToken(null)

    return axios.post(
        `${config.gateway.uri}/login?country=us`,
        { username, password }
    )
}

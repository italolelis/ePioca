import axios from 'axios'
import { login } from '@/api/auth'
import config from '@/config'

const headers = {
    Accept: 'application/vnd.epioca.v1+json',
}

const client = axios.create({
    baseURL: config.gateway.uri,
    headers,
})

export const setAccessToken = (token) => {
    if (token) {
        localStorage.setItem('access_token', token)
        client.defaults.headers.common.Authorization = `Bearer ${token}`
        return
    }

    localStorage.removeItem('access_token')
    delete client.defaults.headers.common.Authorization
}

export const getAccessToken = () => localStorage.getItem('access_token')

if (getAccessToken()) {
    setAccessToken(getAccessToken())
}

// Add an interceptor to deal with forbidden requests
client.interceptors.response.use(response => response, error => {
    if (error.response) {
        if (error.response.status == 401) {
            // Probably token expired. Go to login page.
            window.location.href = '/login'
        }

        console.error(error.response.data)
    }

    if (error.request) {
        console.error(error.request)
    } else {
        console.log(error.message, error.config)
    }
})

export default client

import axios from 'axios'
import { login } from '@/api/auth'
import config from '@/config'

const headers = {
    Accept: 'application/vnd.epioca.v1+json',
}

const client = axios.create({
    baseURL: config.service.uri,
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

export const setUserRole = (rolesArray) => {
    if (rolesArray.indexOf('ROLE_SUPER_ADMIN') === -1) {
        return localStorage.setItem('user_role', 'supplier')
    }

    return localStorage.setItem('user_role', 'buyer')
}

export const getUserRole = () => localStorage.getItem('user_role')

export const setUserId = (id) => localStorage.setItem('user_id', id)

export const getUserId = () => localStorage.getItem('user_id')

export const setUserName = (name) => localStorage.setItem('user_name', name)

export const getUserName = () => localStorage.getItem('user_name')

// Set axios defaults if we have stuff in localStorage
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

    throw error
})

export default client

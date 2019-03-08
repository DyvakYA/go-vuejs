import axios from 'axios'

export const AXIOS = axios.create({
    // baseURL: 'http://localhost:8088',
    headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
    },
    withCredentials:
        true,
    credentials:
        'same-origin',
    crossdomain:
        true
})

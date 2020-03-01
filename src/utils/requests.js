import axios from 'axios'
import config from "@config"
// import { Message } from 'element-ui'
// import store from '@/store'
// import router from '@/router'
// import { getToken, getTokenExpire } from '@/utils/auth'

// create an axios instance
const requests = axios.create({
    baseURL: config.backendBaseUrl, // url = base url + request url
    withCredentials: false, // send cookies when cross-domain requests
    timeout: config.requestTimeout, // request timeout
    headers: {'Content-Type':'application/json;charset=UTF-8'},
});


export default requests

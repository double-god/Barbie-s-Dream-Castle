// src/service/apiClient.js,只负责连接后端 API
import axios from 'axios';

// src/service/apiClient.js
const api = axios.create({
    // 把它改成一个“相对路径”
    // 这会告诉 React：“请向你当前所在的域名发起 /api 请求”
    // (Nginx 会拦截这个请求)
    baseURL: '/api',

    // 允许跨域请求携带凭证
    withCredentials: true,
    headers: {
        'Content-Type': 'application/json',
    }
});

// 请求拦截器：添加 token 到 header
api.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('jwt_token');
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// 响应拦截器：处理错误
api.interceptors.response.use(
    (response) => response.data,
    (error) => {
        if (error.response?.status === 401) {
            // 清除失效的 token
            localStorage.removeItem('jwt_token');
            window.location.href = '/login';
        }
        return Promise.reject(error.response?.data || error);
    }
);

// API 方法


export const getUserProfile = () => {
    return api.get('/user/me');
};

export const updateUserProfile = (data) => {
    return api.put('/user', data);
};

export default api;
// src/service/apiClient.js只负责登陆注册api的实现
import api from './apiClient';

export const loginApi = (username, password) => {
    return api.post('/login', { username, password });
};


export const registerApi = (username, password) => {
    return api.post('/register', { username, password });
};

// 兼容旧代码：导出默认 api 实例，便于从 auth.service 直接使用 api
export default api;
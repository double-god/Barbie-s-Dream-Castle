// src/context/AuthContext.js
import React, { createContext, useState, useContext } from 'react';
import api from '../api/axios';

const AuthContext = createContext(null);

export const AuthProvider = ({ children }) => {
    const [token, setToken] = useState(localStorage.getItem('jwt_token'));

    // (注意：user 状态是可选的，这里我们只管理 token)

    const login = async (username, password) => {
        // try...catch 用于捕获网络错误或 404/500
        try {
            const response = await api.post('/login', { username, password });

            // 检查 *业务* 错误码
            if (response.data.code === 0 && response.data.data.token) {
                const newToken = response.data.data.token;
                localStorage.setItem('jwt_token', newToken);
                setToken(newToken);
                return { success: true };
            } else {
                return { success: false, error: response.data.msg || '登录失败' };
            }
        } catch (err) {
            return { success: false, error: err.response?.data?.msg || '网络或服务器错误' };
        }
    };

    const register = async (username, password, email) => {
        try {
            const response = await api.post('/register', { username, password, email });
            if (response.data.code === 0) {
                return { success: true };
            } else {
                return { success: false, error: response.data.msg || '注册失败' };
            }
        } catch (err) {
            return { success: false, error: err.response?.data?.msg || '网络或服务器错误' };
        }
    };

    const logout = () => {
        localStorage.removeItem('jwt_token');
        setToken(null);
    };

    const authContextValue = {
        token,
        login,
        register,
        logout,
    };

    return (
        <AuthContext.Provider value={authContextValue}>
            {children}
        </AuthContext.Provider>
    );
};

// 导出自定义 hook，方便其他组件使用
export const useAuth = () => useContext(AuthContext);
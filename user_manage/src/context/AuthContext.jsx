import { useState } from 'react';
import { loginApi, registerApi } from '../api/axios';
import { AuthContext } from './AuthContext.js';

export const AuthProvider = ({ children }) => {
    const [token, setToken] = useState(localStorage.getItem('jwt_token'));
    const [user, setUser] = useState(() => {
        try {
            const savedUser = localStorage.getItem('user');
            return savedUser && savedUser !== 'undefined' ? JSON.parse(savedUser) : null;
        } catch (error) {
            console.error('Error parsing user data:', error);
            return null;
        }
    });

    const login = async (username, password) => {
        try {
            const response = await loginApi(username, password);

            // 检查业务错误码
            if (response.code === 0 && response.data?.token) {
                const newToken = response.data.token;
                const userData = response.data.user || null;

                setToken(newToken);
                setUser(userData);
                localStorage.setItem('jwt_token', newToken);
                if (userData) {
                    localStorage.setItem('user', JSON.stringify(userData));
                }

                return { success: true };
            } else {
                // 【【【【【 修改点：优先使用后端返回的详细错误 】】】】】
                throw new Error(response.data?.error || response.msg || '登录失败');
            }
        } catch (error) {
            console.error('Login failed:', error);
            throw new Error(error.message || '登录失败');
        }
    };

    // 【【【【【 修改点 】】】】】
    const register = async (username, password) => {
        try {
            // 不再传递 email
            const response = await registerApi(username, password);
            // 【【【【【 修改点：增加后端错误处理 】】】】】
            if (response.code === 0) {
                return { success: true };
            } else {
                throw new Error(response.data?.error || response.msg || '注册失败');
            }

        } catch (error) {
            console.error('Register failed:', error);
            throw new Error(error.message || '注册失败');
        }
    };

    const logout = () => {
        setToken(null);
        setUser(null);
        localStorage.removeItem('jwt_token');
        localStorage.removeItem('user');
    };

    return (
        <AuthContext.Provider value={{ token, user, login, register, logout }}>
            {children}
        </AuthContext.Provider>
    );
};
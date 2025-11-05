import { useState } from 'react';
import { loginApi, registerApi } from '../../../services/auth.service';
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
                // 成功请求，但业务失败（例如 code != 0）
                throw new Error(response.data?.error || response.msg || '登录失败');
            }
        } catch (error) {
            console.error('Login failed:', error);
            // 从 error 对象 (即后端 JSON) 中提取错误信息
            throw new Error(error.data?.error || error.msg || '登录失败');
        }
    };


    const register = async (username, password) => {
        try {
            // 不再传递 email
            const response = await registerApi(username, password);

            if (response.code === 0) {
                return { success: true };
            } else {
                // 成功请求，但业务失败（例如 code != 0）
                throw new Error(response.data?.error || response.msg || '注册失败');
            }

        } catch (error) {
            console.error('Register failed:', error);
            // 从 error 对象 (即后端 JSON) 中提取错误信息
            // 这样 "该学号已被注册" 就能被正确抛出
            throw new Error(error.data?.error || error.msg || '注册失败');
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
// src/pages/Profile.jsx
import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import api from '../api/axios';
import { useAuth } from '../context/useAuth';
import './AuthPage.css';

export default function ProfilePage() {
    const navigate = useNavigate();
    const [user, setUser] = useState(null);
    const [error, setError] = useState('');
    const { logout } = useAuth(); // 获取 logout 方法

    useEffect(() => {
        // 页面加载时，获取 "我" 的信息
        // api 实例会自动带上 Token
        api.get('/user/me')
            .then(response => {
                if (response.code === 0) {
                    setUser(response.data.data);
                } else {
                    setError(response.data.msg);
                }
            })
            .catch(err => {
                if (err.response?.status === 401) {
                    setError('登录已过期，请重新登录');
                    logout(); // 自动退出登录
                } else {
                    setError(err.response?.data?.msg || '获取用户信息失败');
                }
            });
    }, [logout]);

    if (error) {
        return (
            <div className="auth-container">
                <div className="auth-card">
                    <p className="form-error">{error}</p>
                    <button className="form-button" onClick={() => navigate('/login')}>
                        返回登录
                    </button>
                </div>
            </div>
        );
    }

    if (!user) {
        return (
            <div className="auth-container">
                <div className="auth-card">
                    <p>正在加载用户信息...</p>
                </div>
            </div>
        );
    }

    return (
        <div className="auth-container">
            <div className="auth-card">
                <h2>个人资料</h2>
                <div className="profile-info">
                    <div className="info-group">
                        <label>学号</label>
                        <p>{user.username}</p>
                    </div>
                    <div className="info-group">
                        <label>ID</label>
                        <p>{user.ID}</p>
                    </div>
                    <div className="info-group">
                        <label>简介</label>
                        <p>{user.bio || '未设置'}</p>
                    </div>
                </div>
                <button className="form-button" onClick={logout}>退出登录</button>
            </div>
        </div>
    );
}
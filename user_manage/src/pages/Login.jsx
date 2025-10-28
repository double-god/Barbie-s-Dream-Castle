// src/pages/Login.jsx
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import './Login.css';

export default function LoginPage() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [loading, setLoading] = useState(false);
    const [success, setSuccess] = useState(false);
    const navigate = useNavigate();
    const { login } = useAuth(); // 从 Context 获取 login 方法

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');
        setLoading(true);
        setSuccess(false);

        try {
            const result = await login(username, password);
            setSuccess(true);
            // 显示成功消息后短暂延迟再跳转，让用户看到反馈
            setTimeout(() => {
                navigate('/profile');
            }, 1000);
        } catch (err) {
            setError(err?.message || '登录失败');
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="login-container">
            <form className="login-form" onSubmit={handleSubmit}>
                <h2>登录</h2>
                <div className="form-group">
                    <label className="sr-only">用户名</label>
                    <input className="form-input" type="text" value={username} onChange={e => setUsername(e.target.value)} placeholder="用户名" required />
                </div>
                <div className="form-group">
                    <label className="sr-only">密码</label>
                    <input
                        className="form-input"
                        type="password"
                        value={password}
                        onChange={e => setPassword(e.target.value)}
                        placeholder="密码"
                        required
                        disabled={loading}
                    />
                </div>
                <button
                    className={`form-button ${loading ? 'loading' : ''} ${success ? 'success' : ''}`}
                    type="submit"
                    disabled={loading}
                >
                    {loading ? '登录中...' : success ? '登录成功！' : '登录'}
                </button>
            </form>
            {error && <p className="form-error">{error}</p>}
            {success && <p className="form-success">登录成功！正在跳转...</p>}
            <button
                className="form-secondary-button"
                onClick={() => navigate('/register')}
                disabled={loading || success}
            >
                没有账号？去注册
            </button>
        </div>
    );
}
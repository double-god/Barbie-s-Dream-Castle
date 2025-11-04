// src/pages/AuthPage.jsx
import React, { useState } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import { useAuth } from '../context/useAuth';
import './AuthPage.css'; // 使用新的CSS

// 登录表单
const LoginForm = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [loading, setLoading] = useState(false);
    const [success, setSuccess] = useState(false);
    const navigate = useNavigate();
    const { login } = useAuth();

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');
        setLoading(true);
        setSuccess(false);

        try {
            await login(username, password);
            setSuccess(true);
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
        <form className="auth-form" onSubmit={handleSubmit}>
            <div className="form-group">
                <input
                    className="form-input"
                    type="text"
                    value={username}
                    onChange={e => setUsername(e.target.value)}
                    placeholder="请输入10位学号" // 修改提示
                    required
                />
            </div>
            <div className="form-group">
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
            {error && <p className="form-error">{error}</p>}
            {success && <p className="form-success">登录成功！正在跳转...</p>}
            <button
                className={`form-button ${loading ? 'loading' : ''} ${success ? 'success' : ''}`}
                type="submit"
                disabled={loading}
            >
                {loading ? '登录中...' : success ? '登录成功！' : '登录'}
            </button>
        </form>
    );
};

// 注册表单
const RegisterForm = ({ onRegisterSuccess }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [loading, setLoading] = useState(false);
    const { register } = useAuth();

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');
        setLoading(true);

        try {
            await register(username, password);
            onRegisterSuccess(); // 调用父组件传来的成功回调
        } catch (err) {
            setError(err?.message || '注册失败');
        } finally {
            setLoading(false);
        }
    };

    return (
        <form className="auth-form" onSubmit={handleSubmit}>
            <div className="form-group">
                <input
                    className="form-input"
                    type="text"
                    value={username}
                    onChange={e => setUsername(e.target.value)}
                    placeholder="请输入10位学号" // 修改提示
                    required
                />
            </div>
            <div className="form-group">
                <input
                    className="form-input"
                    type="password"
                    value={password}
                    onChange={e => setPassword(e.target.value)}
                    placeholder="设置密码"
                    required
                    disabled={loading}
                />
            </div>
            {error && <p className="form-error">{error}</p>}
            <button
                className={`form-button ${loading ? 'loading' : ''}`}
                type="submit"
                disabled={loading}
            >
                {loading ? '注册中...' : '注册'}
            </button>
        </form>
    );
};


// 统一的认证页面
export default function AuthPage() {
    const location = useLocation();
    // 根据路由路径决定默认激活的 tab
    const [activeTab, setActiveTab] = useState(
        location.pathname === '/register' ? 'register' : 'login'
    );
    const [registerSuccess, setRegisterSuccess] = useState(false);

    // 注册成功后的回调
    const handleRegisterSuccess = () => {
        setRegisterSuccess(true);
        // 自动切换到登录 tab
        setActiveTab('login');
    };

    return (
        <div className="auth-container">
            <div className="auth-card">
                <div className="auth-tabs">
                    <button
                        className={`auth-tab ${activeTab === 'login' ? 'active' : ''}`}
                        onClick={() => setActiveTab('login')}
                    >
                        登录
                    </button>
                    <button
                        className={`auth-tab ${activeTab === 'register' ? 'active' : ''}`}
                        onClick={() => setActiveTab('register')}
                    >
                        注册
                    </button>
                </div>

                {registerSuccess && activeTab === 'login' && (
                    <p className="form-success">注册成功！请登录。</p>
                )}

                {activeTab === 'login' ? (
                    <LoginForm />
                ) : (
                    <RegisterForm onRegisterSuccess={handleRegisterSuccess} />
                )}
            </div>
        </div>
    );
}
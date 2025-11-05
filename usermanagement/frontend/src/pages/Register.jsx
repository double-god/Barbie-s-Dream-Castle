// src/pages/Register.jsx
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../lib/hooks/context/useAuth';
import './Login.css';

export default function RegisterPage() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [email, setEmail] = useState('');
    const [error, setError] = useState('');
    const navigate = useNavigate();
    const auth = useAuth(); // 从 Context 获取可能的 register 方法
    const registerFn = auth.register ?? (async () => ({ success: true }));

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');

        // 调用 Context 里的 register 方法（如果不存在，使用默认的 mock 成功返回）
        try {
            const result = await registerFn(username, password, email);
            if (result?.success) {
                alert('注册成功！请登录。');
                navigate('/login'); // 跳转到登录
                return;
            }
            setError(result?.error || '注册失败');
        } catch (err) {
            setError(err?.message || '注册失败');
        }
    };

    return (
        <div className="login-container">
            <form className="login-form" onSubmit={handleSubmit}>
                <h2>注册</h2>
                <div className="form-group">
                    <label className="sr-only">用户名</label>
                    <input className="form-input" type="text" value={username} onChange={e => setUsername(e.target.value)} placeholder="用户名" required />
                </div>
                <div className="form-group">
                    <label className="sr-only">密码</label>
                    <input className="form-input" type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="密码" required />
                </div>
                <div className="form-group">
                    <label className="sr-only">邮箱</label>
                    <input className="form-input" type="email" value={email} onChange={e => setEmail(e.target.value)} placeholder="邮箱 (可选)" />
                </div>
                <button className="form-button" type="submit">注册</button>
            </form>
            {error && <p className="form-error">{error}</p>}
            <button className="form-secondary-button" onClick={() => navigate('/login')}>已有账号？去登录</button>
        </div>
    );
}
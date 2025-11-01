// src/App.jsx
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { useAuth } from './context/useAuth';
import LoginPage from './pages/Login.jsx';
import RegisterPage from './pages/Register.jsx';
import ProfilePage from './pages/profile.jsx';

// (关键) 保护组件：如果用户没有 Token，就重定向到登录页
const ProtectedRoute = ({ children }) => {
  const { token } = useAuth();
  if (!token) {
    return <Navigate to="/login" replace />;
  }
  return children;
};

function App() {
  const { token } = useAuth();
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />

        {/* /profile 页面必须被保护 */}
        <Route
          path="/profile"
          element={
            <ProtectedRoute>
              <ProfilePage />
            </ProtectedRoute>
          }
        />

        {/* 默认首页：如果已登录，去 /profile，否则去 /login */}
        <Route
          path="/"
          element={
            token ? <Navigate to="/profile" /> : <Navigate to="/login" />
          }
        />
      </Routes>
    </BrowserRouter>
  );
} export default App;

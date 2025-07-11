import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

const LoginPage = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const res = await axios.post("/login", { email, password });
      localStorage.setItem("token", res.data.token);
      navigate("/select-stocks");
    } catch (err) {
      setError("Login failed. Please check credentials.");
    }
  };

  return (
    <div className="centered">
      <img src="/logo.png" alt="Logo" width={120} style={{ marginBottom: 20 }} />
      <form onSubmit={handleLogin}>
        <input type="email" placeholder="Email" required value={email} onChange={(e) => setEmail(e.target.value)} />
        <input type="password" placeholder="Password" required value={password} onChange={(e) => setPassword(e.target.value)} />
        <button type="submit">Login</button>
        {error && <p style={{ color: "red" }}>{error}</p>}
        <p>Don't have an account? <a href="/register">Register</a></p>
      </form>
    </div>
  );
};

export default LoginPage;

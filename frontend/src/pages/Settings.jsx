import React, { useEffect, useState } from "react";
import axios from "axios";
import Header from "../components/Header";

const Settings = () => {
  const [form, setForm] = useState({
    first_name: "",
    last_name: "",
    email: "",
    password: "",
    base_currency: "USD"
  });

  useEffect(() => {
    const loadSettings = async () => {
      const res = await axios.get("/api/settings", {
        headers: { Authorization: `Bearer ${localStorage.getItem("token")}` },
      });
      setForm((f) => ({ ...f, ...res.data }));
    };
    loadSettings();
  }, []);

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSave = async () => {
    await axios.post("/api/settings", form, {
      headers: { Authorization: `Bearer ${localStorage.getItem("token")}` },
    });
    alert("Settings updated");
  };

  return (
    <div className="settings">
      <Header />
      <h2>Account Settings</h2>
      <input name="first_name" placeholder="First Name" value={form.first_name} onChange={handleChange} />
      <input name="last_name" placeholder="Last Name" value={form.last_name} onChange={handleChange} />
      <input name="email" placeholder="Email" value={form.email} onChange={handleChange} />
      <input name="password" type="password" placeholder="New Password" value={form.password} onChange={handleChange} />
      <input name="base_currency" placeholder="Currency (e.g. USD, EUR)" value={form.base_currency} onChange={handleChange} />
      <button onClick={handleSave}>Save</button>
    </div>
  );
};

export default Settings;

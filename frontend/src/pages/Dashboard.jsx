import React, { useEffect, useState } from "react";
import axios from "axios";
import StockGraph from "../components/StockGraph";
import StockInfoBox from "../components/StockInfoBox";
import Header from "../components/Header";

const Dashboard = () => {
  const [stocks, setStocks] = useState([]);
  const [data, setData] = useState({});
  const [lastUpdate, setLastUpdate] = useState("");

  useEffect(() => {
    const fetchData = async () => {
      const token = localStorage.getItem("token");

      try {
        const favRes = await axios.get("/api/favorites", {
          headers: { Authorization: `Bearer ${token}` },
        });

        setStocks(favRes.data.map((s) => s.symbol));

        const stockRes = await axios.get("/api/stocks", {
          headers: { Authorization: `Bearer ${token}` },
          params: { symbol: favRes.data.map((s) => s.symbol) },
        });

        setData(stockRes.data);
        setLastUpdate(new Date().toLocaleTimeString());
      } catch (err) {
        console.error("Failed to load stock data", err);
      }
    };

    fetchData();
    const interval = setInterval(fetchData, 5 * 60 * 1000); // refresh every 5 min
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="dashboard">
      <Header />
      <h2>My Stocks</h2>
      <p>Last update: {lastUpdate}</p>
      <StockGraph stockData={data} />
      <StockInfoBox stockData={data} />
    </div>
  );
};

export default Dashboard;

import React, { useState, useEffect } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const predefinedStocks = [
  { symbol: "AAPL", name: "Apple" },
  { symbol: "GOOG", name: "Google" },
  { symbol: "MSFT", name: "Microsoft" },
  { symbol: "TSLA", name: "Tesla" },
  { symbol: "AMZN", name: "Amazon" },
  { symbol: "^GSPC", name: "S&P 500 (Market Index)" },
  // ... add more or load dynamically
];

const StockSelection = () => {
  const [selected, setSelected] = useState([]);
  const [search, setSearch] = useState("");
  const navigate = useNavigate();

  const handleToggle = (symbol) => {
    if (selected.includes(symbol)) {
      setSelected(selected.filter((s) => s !== symbol));
    } else if (selected.length < 50) {
      setSelected([...selected, symbol]);
    }
  };

  const handleSave = async () => {
    const payload = selected.map((symbol, i) => {
      const match = predefinedStocks.find(s => s.symbol === symbol);
      return {
        symbol,
        display_name: match?.name || symbol,
        color: `#${Math.floor(Math.random()*16777215).toString(16)}`,
        is_market_index: symbol.startsWith("^")
      };
    });

    try {
      await axios.post("/api/favorites", payload, {
        headers: { Authorization: `Bearer ${localStorage.getItem("token")}` }
      });
      navigate("/dashboard");
    } catch (err) {
      alert("Error saving stock selections");
    }
  };

  const filteredStocks = predefinedStocks.filter(s =>
    s.name.toLowerCase().includes(search.toLowerCase()) || s.symbol.toLowerCase().includes(search.toLowerCase())
  );

  return (
    <div className="selection-container">
      <h2>Select up to 50 Stocks</h2>
      <input
        type="text"
        placeholder="Search Stocks..."
        value={search}
        onChange={(e) => setSearch(e.target.value)}
      />
      <ul className="stock-list">
        {filteredStocks.map((stock) => (
          <li key={stock.symbol}>
            <label>
              <input
                type="checkbox"
                checked={selected.includes(stock.symbol)}
                onChange={() => handleToggle(stock.symbol)}
              />
              {stock.name}: {stock.symbol}
            </label>
          </li>
        ))}
      </ul>
      <button onClick={handleSave}>Done</button>
    </div>
  );
};

export default StockSelection;

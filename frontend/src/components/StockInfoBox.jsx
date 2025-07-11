import React from "react";

const StockInfoBox = ({ stockData }) => {
  return (
    <div className="info-box">
      <h3>Stock Info</h3>
      {Object.entries(stockData).map(([symbol, info]) => (
        <div key={symbol} className="stock-block">
          <h4>{info.name} ({symbol})</h4>
          <p>Market Cap: {info.market_cap}</p>
          <p>Open: {info.open}</p>
          <p>Day Low: {info.day_low}</p>
          <p>Day High: {info.day_high}</p>
          <p>Volume: {info.volume}</p>
          <p>Currency: {info.currency}</p>
        </div>
      ))}
    </div>
  );
};

export default StockInfoBox;

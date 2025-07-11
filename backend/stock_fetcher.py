#!/usr/bin/env python3

import sys
import json
import yfinance as yf

def fetch_stock_data(symbols):
    data = {}

    for symbol in symbols:
        try:
            stock = yf.Ticker(symbol)
            info = stock.info
            hist = stock.history(period="1mo")

            data[symbol] = {
                "symbol": symbol,
                "name": info.get("shortName", "N/A"),
                "market_cap": info.get("marketCap"),
                "previous_close": info.get("previousClose"),
                "open": info.get("open"),
                "day_low": info.get("dayLow"),
                "day_high": info.get("dayHigh"),
                "volume": info.get("volume"),
                "currency": info.get("currency"),
                "history": hist.reset_index().to_dict(orient="records"),
                "last_updated": info.get("regularMarketTime")
            }
        except Exception as e:
            data[symbol] = {
                "error": str(e)
            }

    return data

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print(json.dumps({"error": "No symbol provided"}))
        sys.exit(1)

    symbols = sys.argv[1:]
    result = fetch_stock_data(symbols)
    print(json.dumps(result))

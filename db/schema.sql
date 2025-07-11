CREATE DATABASE IF NOT EXISTS stocktracker;
USE stocktracker;

-- Users table
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    ip_address VARCHAR(45),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Favorite stocks
CREATE TABLE favorite_stocks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    symbol VARCHAR(10),
    display_name VARCHAR(100),
    color VARCHAR(7),
    is_market_index BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Stock alerts
CREATE TABLE stock_alerts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    symbol VARCHAR(10),
    target_price DECIMAL(10, 2),
    direction ENUM('above', 'below'),
    alert_sent BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- User settings
CREATE TABLE user_settings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    base_currency VARCHAR(10) DEFAULT 'USD',
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Stock history (optional caching)
CREATE TABLE stock_history (
    id INT AUTO_INCREMENT PRIMARY KEY,
    symbol VARCHAR(10),
    json_data TEXT,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

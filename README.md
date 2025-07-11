# 📈 Stock Tracker

A full-stack, real-time stock tracking and alerting web application built with **React**, **Go**, **Python**, and **MySQL** — fully containerized with **Docker**. Users can register, select up to 50 stocks, view live graphing and history, set alerts, and customize their settings.

---

## 📁 Project Directory Structure

stock-tracker/
├── backend/
│   ├── main.go
│   ├── handlers/
│   ├── models/
│   ├── middleware/
│   ├── utils/
│   ├── stock_fetcher.py
│   ├── requirements.txt
│   └── go.mod
│
├── frontend/
│   ├── public/
│   ├── src/
│   │   ├── pages/
│   │   ├── components/
│   │   └── index.js
│   ├── Dockerfile
│   └── package.json
│
├── db/
│   └── schema.sql
├── .env
├── docker-compose.yml
└── README.md

---

<details>
<summary>🧠 Backend Stack</summary>

- **Language:** Go (Golang)
- **Framework:** `gin-gonic`
- **Authentication:** JWT (stored in client storage)
- **Email Alerts:** SMTP via `net/smtp`
- **Data Fetching:** Calls a Python script (`stock_fetcher.py`) using `os/exec`
- **Python Dependencies:** `yfinance`, `pandas`, `numpy`
- **Database:** MySQL 8.x via GORM ORM

</details>

<details>
<summary>🎨 Frontend Stack</summary>

- **Library:** React (with React Router v6)
- **Graphing:** Chart.js via `react-chartjs-2`
- **API:** Axios for HTTP requests
- **Styling:** CSS modules + responsive layout
- **Background:** [particles.js](https://vincentgarreau.com/particles.js/)
- **Session:** JWT stored in `localStorage`

</details>

<details>
<summary>⚙️ Technology Stack</summary>

- **Frontend:** React + Chart.js
- **Backend:** Go + Python
- **Database:** MySQL
- **Containerization:** Docker
- **Email + Alerts:** SMTP
- **Currency Conversion:** exchangerate.host
- **Reverse Proxy (optional):** Nginx

</details>

---

## 🛠️ Installation Guide

### Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop)
- [Docker Compose](https://docs.docker.com/compose/)
- (Optional) Python installed locally for testing `stock_fetcher.py`

---

### 1. Clone the Repository

```bash
git clone https://github.com/JoshBaneyCS/stock-tracker.git
cd stock-tracker
```
### 2. Setup Environment Variables
Create a .env file in the project root:

```.env
PORT=8080
DB_HOST=db
DB_PORT=3306
DB_USER=root
DB_PASS=password
DB_NAME=stocktracker

JWT_SECRET=supersecretjwtkey
EMAIL_USER=your_email@gmail.com
EMAIL_PASS=your_email_password
```
### 3. Build and Run!
```bash
docker-compose up --build
```
This will:

Build and run the Go backend + Python inside 1 container

Serve the React frontend via Nginx on port 3000

Set up a MySQL database and schema

### 4. Open in Browser
Frontend: http://localhost:3000 (Or whatever you have your port mapped to)

Backend API: http://localhost:8080 (Or whatever you have your port mapped to)

## 🔐 Default Routes

Method	Endpoint	Description
POST	/register	Register a new user
POST	/login	Login and get JWT
GET	/api/favorites	Get user’s saved stocks
POST	/api/favorites	Save user’s selected stocks
GET	/api/stocks	Fetch live stock data
GET	/api/settings	Get user profile/settings
POST	/api/settings	Update name, email, etc.
POST	/api/alerts	Create price alert
GET	/api/alerts	View all alerts
DELETE	/api/alerts/:id	Delete alert

JWT must be sent as a Bearer token in Authorization headers for protected endpoints.


## 📧 Features

- ✅ User authentication (JWT-based login/register)
- ✅ Select up to 50 favorite stocks
- ✅ Market index toggle support (e.g., S&P 500)
- ✅ Real-time stock graphing (using Chart.js)
- ✅ Custom color-coded lines per stock
- ✅ Search bar to quickly find stocks
- ✅ Historical data box (market cap, high/low, etc.)
- ✅ Auto-refresh every 5 minutes
- ✅ Email alerts when price reaches or falls below threshold
- ✅ Currency conversion via exchangerate.host
- ✅ Customizable account settings (name, email, password, currency)
- ✅ IP address is logged on login and registration
- ✅ Dockerized with flexible port management via `.env`
- ✅ Front page includes `particles.js` animated background and logo

## 🔐 Security Notes

- Passwords are hashed securely using **bcrypt**
- JWTs are used for stateless authentication
- Frontend stores JWT in `localStorage`
- CSRF is avoided by using token-only auth
- SMTP email alerts use secure environment variables

## 🧪 Development Notes

To test the Python stock data script locally:
```bash
cd backend
pip install -r requirements.txt
python3 stock_fetcher.py AAPL TSLA
```

This will return a JSON response for the given stock tickers.


## 📜 License

This project is licensed under the MIT License. You may use, modify, and distribute it freely, with proper attribution.

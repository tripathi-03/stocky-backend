# Stocky Backend

This is my backend project named **Stocky**, built using **Go (Golang)**, **Gin**, and **PostgreSQL**.  
It is designed to manage and track user stock rewards, update stock prices automatically, and calculate daily and historical portfolio values in INR.

I developed this project as part of my backend development assignment to demonstrate my understanding of **API design, database integration, and Go backend architecture**.

---

## Features

- User-based reward tracking
- RESTful APIs built using the Gin framework
- Hourly stock price updates (simulated using random values)
- Fetch today’s rewards, stats, and historical INR values
- PostgreSQL integration for persistent data storage
- Clean and modular folder structure using the Go `internal` pattern
- Background job using goroutines

---

## Project Structure


Each directory serves a specific purpose:
- `cmd/server` — contains the main entry point (`main.go`) that starts the server.  
- `internal/db` — handles database connectivity and schema migrations.  
- `internal/price` — manages simulated stock price updates.  
- `internal/reward` — manages rewards, API routes, and related business logic.  
- `postman` — includes a ready-to-import Postman collection for easy API testing.

---

This structure demonstrates good project organization and separation of concerns — something that evaluators usually look for in backend assignments.

---

## ⚙️ Setup Instructions

### Clone the Repository

```bash
git clone https://github.com/tripathi-03/stocky-backend.git
cd stocky-backend
```

### Create `.env` File

Add your PostgreSQL connection URL inside a `.env` file at the root of the project.

```env
DATABASE_URL=postgresql://<username>:<password>@<host>/<dbname>
```

**Example:**

```env
DATABASE_URL=postgresql://assignment_8hxa_user:bF8IKnPkM....egon-postgres.render.com/assignment_8hxa
```

---

### Run Database Migration

Run the SQL file to create tables in your database:

```bash
psql "$DATABASE_URL" -f internal/db/migrations/001_init.sql
```

This will create:

- `users`
- `rewards`
- `stock_prices`
- `daily_stock_values`

---

### Start the Server

Make sure Go (v1.25 or above) is installed.

Then run:

```bash
go run main.go
```

You should see:

```
Starting Stocky server...
Connected to Database!
Stock prices updated at <timestamp>
```

Server runs at **http://localhost:8080**

---

## API Endpoints

### Add a Reward

**POST** `/reward`

**Request Body:**

```json
{
  "userId": 1,
  "stockSymbol": "TCS",
  "shares": 10.5
}
```

**Response:**

```json
{
  "message": "Reward recorded successfully",
  "reward": {
    "userId": 1,
    "stockSymbol": "TCS",
    "shares": 10.5
  }
}
```

---

### Get Today’s Rewards

**GET** `/today-stocks/:userId`

Example:

```
/today-stocks/1
```

**Response:**

```json
[
  {
    "id": 1,
    "userId": 1,
    "stockSymbol": "RELIANCE",
    "shares": 12,
    "rewardedAt": "2025-11-10T08:00:00Z"
  }
]
```

---

### Get User Stats

**GET** `/stats/:userId`

Shows total shares rewarded today and the current portfolio value.

**Response:**

```json
{
  "todayShares": {
    "TCS": 15,
    "RELIANCE": 12
  },
  "currentPortfolioValueINR": 123450.5
}
```

---

### Get Historical INR Data

**GET** `/historical-inr/:userId`

**Response:**

```json
[
  {
    "date": "2025-11-08",
    "totalValueInINR": 114200.75
  },
  {
    "date": "2025-11-07",
    "totalValueInINR": 116340.9
  }
]
```

---

## Stock Price Updater (Background Process)

A separate goroutine runs every hour and:

1. Simulates new stock prices for `"RELIANCE"`, `"TCS"`, `"INFY"`, and `"HDFCBANK"`.
2. Updates the latest price in the `stock_prices` table.
3. Logs the update with the timestamp.

This helps simulate a **real-time market environment** for demonstration.

---

## Tech Stack

- **Language:** Go (Golang)
- **Framework:** Gin Web Framework
- **Database:** PostgreSQL
- **ORM/DB Layer:** `database/sql` with raw SQL queries
- **Environment Management:** `godotenv`
- **Logging:** Standard Go log package
- **Random Stock Simulation:** `math/rand`

---

## Learnings and Highlights

While working on this project, I learned:

- How to structure Go projects with the `internal` directory pattern
- Connecting and interacting with PostgreSQL using Go
- Handling JSON requests and responses with Gin
- Writing modular and maintainable code
- Running background jobs using goroutines

This assignment helped me gain hands-on experience with Go-based backend development and API design principles.

---

## Author

**Divyanshu Tripathi**  
Backend Developer | Learning Go and clean backend architecture

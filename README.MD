# 🎵 Concerts API

**Concerts** is a lightweight, high-performance REST API built with Go and MySQL for listing concerts and handling ticket bookings with precise control over booking windows and availability.

---

## 🚀 Features

- 🔍 **Search concerts** currently available
- 🎫 **Book tickets** within limited-time windows
- ⚡ **High concurrency** support (e.g. 500+ users/sec)
- 🔐 Prevents **overbooking** using SQL transactions + row locking
- ✅ Built with **Clean Architecture** for maintainability and testability

---

## 🛠️ Tech Stack

- **Golang**
- **MySQL** (transaction-safe with indexed queries)

---

## 📦 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/iandanarko/concerts.git
cd concerts
```

### 2. Install dep
```bash
make dep
```

### 3. Set env
```bash
cp env.sample .env
```

### 4. Migrate
```bash
make migrate
```

### 5. Run
```bash
make run-api
```

## API Docs
- Run the app
- Go to `http://{host}:{port}/swagger/index.html`

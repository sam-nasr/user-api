# User API (Go + net/http + mux)

A simple REST API built in Go to manage users.
This project demonstrates backend fundamentals using native Go packages and Gorilla Mux.

---

## 🚀 Features

* Create user
* Get all users
* Update user
* Delete user
* Input validation
* Logging middleware
* Standardized JSON responses

---

## 📁 Project Structure

```
user-api/
│
├── cmd/api/main.go
├── internal/
│   ├── handlers/
│   ├── services/
│   ├── models/
│   └── middleware/
```

---

## 🛠 Requirements

* Go 1.18+

Check installation:

```
go version
```

---

## ▶️ Run the Project

```bash
go run cmd/api/main.go
```

Server will start on:

```
http://localhost:8080
```

---

## 📌 API Endpoints

---

### 🔹 Get All Users

```
GET /users
```

---

### 🔹 Create User

```
POST /users
```

Body:

```json
{
  "id":1,
  "name": "Samer",
  "age": 25
}
```

---

### 🔹 Update User

```
PUT /users/{id}
```

Body:

```json
{
  "id":1,
  "name": "Updated Name",
  "age": 30
}
```

---

### 🔹 Delete User

```
DELETE /users/{id}
```

---

## 📦 Response Format

### ✅ Success

```json
{
  "data": {...},
  "error": null
}
```

### ❌ Error

```json
{
  "data": null,
  "error": "error message"
}
```

---

## 📊 Logging

Each request is logged:

```
[REQUEST] GET /users
[RESPONSE] 200 GET /users took 120µs
```

---

## ⚠️ Notes

* Data is stored in memory (no database)
* Data will reset when server restarts

---

## 📚 What I Learned

* Building REST APIs with `net/http`
* Routing using Gorilla Mux
* Structuring Go projects (handlers, services, models)
* Middleware implementation
* Request validation and error handling

---

## 🚀 Future Improvements

* Add database (MongoDB)
* Improve validation layer
* Add authentication
* Add pagination

---

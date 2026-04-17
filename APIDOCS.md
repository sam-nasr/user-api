## API Documentation

### Base URL

```
http://localhost:8080
```

---

### 🔹 Get All Users

**GET** `/users`

Response:

```json
{
  "data": [
    {
      "id": 1,
      "name": "Samer",
      "age": 25
    }
  ],
  "error": null
}
```

---

### 🔹 Create User

**POST** `/users`

Body:

```json
{
  "id":1,
  "name": "Samer",
  "age": 25
}
```

Response:

```json
{
  "data": {
    "id": 1,
    "name": "Samer",
    "age": 25
  },
  "error": null
}
```

---

### 🔹 Update User

**PUT** `/users/{id}`

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

**DELETE** `/users/{id}`

Response:

```json
{
  "data": null,
  "error": null
}
```

---

### 🔹 Error Response Format

```json
{
  "data": null,
  "error": "error message"
}
```

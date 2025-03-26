# Student Management API

This is a simple Go-based REST API for managing student records. The API supports the following operations:
- Get all students (`GET`)
- Adding a new student (`POST`)
- Updating an existing student (`PUT`)
- Deleting a student (`DELETE`)

## Prerequisites
- Go 1.16+
- A REST client (Postman, curl, etc.)

## How to Run
1. Clone the repository:
   ```sh
   git clone https://github.com/your-repo/student-api.git
   cd student-api
   ```
2. Run the server:
   ```sh
   go run main.go
   ```
   The server will start on `http://localhost:5000`.

## API Endpoints

### 1️⃣ Get All Students
**Endpoint:** `GET /get`

**Response:**
```json
{
  "status": "S",
  "errMsg": "",
  "students": [
    {
      "id": 101,
      "name": "manoj",
      "mobile": 9988112201,
      "email": "manoj@acb.com",
      "address": "Chennai",
      "dob": "01/01/2001"
    },
    ...
  ]
}
```

### 2️⃣ Add a Student
**Endpoint:** `POST /post`

**Request Body:**
```json
{
  "id": 105,
  "name": "new user",
  "mobile": 9988776609,
  "email": "user@bca.com",
  "address": "mysuru",
  "dob": "09/10/2002"
}
```

**Response:**
```json
{
  "status": "S",
  "errMsg": "",
  "students": [ ... updated list of students ... ]
}
```

### 3️⃣ Update a Student
**Endpoint:** `PUT /put`

**Request Body:**
```json
{
  "id": 102,
  "name": "Vimal Updated",
  "mobile": 9988112202,
  "email": "vimal_updated@acb.com",
  "address": "Kerala",
  "dob": "02/02/2000"
}
```

**Response:**
```json
{
  "status": "S",
  "errMsg": "",
  "students": [ ... updated list of students ... ]
}
```

### 4️⃣ Delete a Student
**Endpoint:** `DELETE /delete`

**Request Body:**
```json
{
  "id": 102
}
```

**Response:**
```json
{
  "status": "S",
  "errMsg": "",
  "students": [ ... updated list of students ... ]
}
```

## Error Responses
| HTTP Code | Message                     |
|-----------|-----------------------------|
| 400       | Bad Request                 |
| 404       | Student not found           |
| 409       | Student already exists      |
| 500       | Internal Server Error       |


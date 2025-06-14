# 🏥 Hospital Portal API

A simple Go backend for managing hospital patient data with role-based access for doctors and receptionists.

---

## Tech Stack

- Go + Chi Router  
- PostgreSQL + GORM  
- JWT Auth (via HttpOnly cookies)  

---

## Getting Started

1. **Clone & Setup**
   ```bash
   git clone https://github.com/your-username/hospital-portal.git
   cd hospital-portal
   go mod tidy

2. Set .env
    ```bash
    DB_HOST=localhost
    DB_USER=postgres
    DB_PASSWORD=yourpassword
    DB_NAME=hospitaldb
    JWT_SECRET=your_jwt_secret
3.run
    ```bash
    go run cmd/main.go

## API Routes

### Auth

| Method | Endpoint   | Access         |
|--------|------------|----------------|
| POST   | `/signup`  | Public         |
| POST   | `/login`   | Public         |
| POST   | `/logout`  | Authenticated  |

### Patients

| Method | Endpoint           | Access              |
|--------|--------------------|---------------------|
| GET    | `/patients`        | Doctor, Receptionist|
| POST   | `/patients`        | Receptionist only   |
| PUT    | `/patients/{id}`   | Receptionist only   |
| DELETE | `/patients/{id}`   | Receptionist only   |

## postman docs
[📦 Postman Collection](hospitalAPi.postman_collection.json)
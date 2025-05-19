# Notification API

A simple, containerized notification system using Go, PostgreSQL, RabbitMQ, and Docker.  
Supports Email, SMS (Twilio), and In-App notifications.

---

## Features

- REST API for sending notifications  
- Supports Email, SMS, and In-App notifications  
- PostgreSQL for persistent storage  
- RabbitMQ for message queueing  
- Worker service for async notification delivery  
- Docker Compose for easy setup  

---

## Tech Stack

- Go (Gin, GORM)  
- PostgreSQL  
- RabbitMQ  
- Docker & Docker Compose  

---

## Getting Started

### Prerequisites

- [Docker Desktop](https://docs.docker.com/desktop/)

### Setup

1. **Clone this repository:**

   ```bash
   git clone https://github.com/rahulkumar7858/notification-api.git
   cd notification-api
   ```

2. **Configure environment variables in `.env`:**

   ```
   # PostgreSQL
   DB_HOST=postgres
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=notifications
   DB_PORT=5432

   # RabbitMQ
   RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/

   # Email
   SMTP_HOST=smtp.gmail.com
   SMTP_PORT=587
   SMTP_USER=your_email@gmail.com
   SMTP_PASSWORD=your_app_password

   # Twilio (SMS)
   TWILIO_SID=your_account_sid
   TWILIO_TOKEN=your_auth_token
   TWILIO_PHONE=+1234567890
   ```

3. **Build and start all services:**

   ```bash
   docker compose up --build
   ```

---

## API Usage

### Send a Notification

**POST** `/notifications`  
**Content-Type:** `application/json`

```json
{
  "type": "email",       // "email", "sms", or "inapp"
  "to": "user@example.com", // email, phone, or user ID
  "message": "Hello from the notification system"
}
```

### Example (Email)

```bash
curl -X POST http://localhost:8080/notifications \
-H "Content-Type: application/json" \
-d '{"type":"email","to":"user@example.com","message":"Test notification"}'
```

### Example (SMS)

```bash
curl -X POST http://localhost:8080/notifications \
-H "Content-Type: application/json" \
-d '{"type":"sms","to":"+1234567890","message":"SMS test"}'
```

### Example (In-App)

```bash
curl -X POST http://localhost:8080/notifications \
-H "Content-Type: application/json" \
-d '{"type":"inapp","to":"user_id_or_username","message":"In-app test"}'
```

---

## Monitoring

- **PostgreSQL:**  
  ```bash
  docker exec -it notification-api-postgres-1 psql -U postgres -d notifications
  ```

- **RabbitMQ Management UI:**  
  [http://localhost:15672](http://localhost:15672) (user: guest / pass: guest)

- **API Logs:**  
  ```bash
  docker logs notification-api-api-1
  ```

- **Worker Logs:**  
  ```bash
  docker logs notification-api-worker-1
  ```

---

## Project Structure

```
notification-api/
├── docker-compose.yaml
├── .env
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── worker/
│       └── main.go
├── internal/
│   ├── db/
│   │   └── postgres.go
│   ├── models/
│   │   └── models.go
│   └── handlers/
│       └── notifications.go
└── migrations/
    └── 001_init.sql
```

---

## License

MIT

# CodeShip

CodeShip is a SaaS platform designed for CI/CD automation, offering seamless integration for pipelines, deployments, and CLI-based management. It allows users to host their own applications while leveraging CodeShip for CI/CD functionalities.

## Features
- **Self-Hosted CI/CD Pipelines**: Manage and execute CI/CD pipelines on your infrastructure.
- **CLI Tool**: Simplifies pipeline initialization and deployment.
- **REST API**: Provides endpoints for managing pipelines and deployments.
- **Extensibility**: Built with Go, designed for high performance and flexibility.

---

## Project Structure
```plaintext
codeship/
├── api/                    # Backend API
│   ├── handlers/           # HTTP Handlers
│   ├── middleware/         # Middleware เช่น Authentication
│   ├── models/             # Models สำหรับข้อมูล (เช่น Database schema)
│   ├── routes/             # การกำหนด API Routes
│   ├── services/           # Business Logic / Service Layer
│   ├── main.go             # Entry point ของ API
│   ├── config/             # Configuration เช่น Env variables
├── cli/                    # CLI Application
│   ├── commands/           # คำสั่ง CLI เช่น deploy, init
│   ├── utils/              # Utilities สำหรับ CLI
│   ├── main.go             # Entry point ของ CLI
├── internal/               # Utilities ภายในโปรเจกต์
│   ├── db/                 # Database Utilities
│   ├── logger/             # Logging utilities
│   ├── config/             # Configuration Loader
│   ├── queue/              # Queue Utilities เช่น RabbitMQ, Kafka
├── pkg/                    # Packages ที่สามารถนำไปใช้ซ้ำได้
│   ├── auth/               # Authentication/Authorization
│   ├── pipeline/           # Pipeline Utilities
│   ├── ssh/                # SSH Utilities
│   ├── docker/             # Docker utilities
├── scripts/                # Scripts สำหรับการพัฒนาและ Deployment
│   ├── migrate.sh          # Migration Script
│   ├── deploy.sh           # Deploy Script
├── configs/                # Configuration Files
│   ├── development.yaml    # Development Config
│   ├── production.yaml     # Production Config
├── Dockerfile              # Dockerfile สำหรับ Backend
├── go.mod                  # Go Modules
├── go.sum                  # Go Dependencies
├── README.md               # คำอธิบายโปรเจกต์
```

---

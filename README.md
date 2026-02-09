# FastAPI Digital Library → Gin Orchestrator API

## 📘 Project Overview
**Digital Library** is a RESTful backend service originally built using **FastAPI (Python)** and later **migrated to Go using the Gin framework** as part of the **Week-5 Orchestrator API assignment**.

The project demonstrates:
- Language & framework migration (**Python → Go**)
- Clean Architecture implementation
- Production-ready REST API practices
- Middleware, validation, and structured JSON responses
- Self-documenting API concepts

---

## 🧭 Project Evolution

### 🔹 Week 3 – FastAPI Implementation
- Built using **FastAPI (Python)**
- Focused on CRUD APIs, Pydantic validation, and middleware
- Used in-memory data storage
- Swagger (OpenAPI) documentation enabled by default

### 🔹 Week 5 – Gin (Go) Migration
- Migrated to **Gin framework (Go)**
- Re-architected using **Clean Architecture**
- Introduced clear separation of concerns:
  - Handler (HTTP)
  - Usecase (Business logic)
  - Repository (Data layer)
- Improved error handling and response consistency
- Implemented production-style middleware

---

## ⚙️ Server Details (Current – Go Version)

- **Framework:** Gin  
- **Language:** Go  
- **Architecture:** Clean Architecture  
- **API Type:** REST  
- **Data Storage:** In-memory (as per assignment)  
- **Documentation:** Swagger (optional / future scope)

---

## ✨ Features
- Full CRUD operations on books
- Strict input validation
- Clean Architecture layering
- Centralized error handling
- Middleware for request logging
- Processing time tracking via custom headers
- Consistent JSON API responses
- Background Task Processing via Worker Pool

---

## 📚 Book Data Model
Each book contains:
- `id` (int)
- `title` (non-empty string)
- `author` (string)
- `year` (1000–2026)
- `isbn` (10 or 13 characters)

---

## 🚀 Setup Instructions

### 1️⃣ Clone the Repository
```bash
git clone https://github.com/<your-username>/fastapi-digital-library.git
cd fastapi-digital-library

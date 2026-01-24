# FastAPI Digital Library

## 📘 Project Overview
FastAPI Digital Library is a RESTful backend server built using **FastAPI** to manage books in a digital library.  
The application demonstrates proper API design, data validation, middleware usage, and self-documenting APIs using Swagger.

---

## ⚙️ Server Details
- **Framework:** FastAPI  
- **Language:** Python  
- **API Type:** REST  
- **Data Storage:** In-memory (for demo/learning purposes)  
- **Documentation:** Swagger UI (OpenAPI)

---

## ✨ Features
- Create, read, update, and delete books
- Input validation using Pydantic
- Custom error handling for bad requests
- Global middleware for request logging
- Processing time tracking using custom response headers
- Auto-generated API documentation

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

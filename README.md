# ⭐ Dragon's Hoard Server
Backend API and service layer for the D&D virtual tabletop 'Dragon's Hoard,' providing game management, real time session tooling, and access to source material.

---

## 🌐 Live Demo
**Live Site:** [Coming Soon]()  
**API Endpoint (optional):** [Coming Soon]()  
<!--**Demo Video:** _(TBD)_-->

---

## 🚀 Tech Stack
**Backend:** Golang, [Echo](https://echo.labstack.com/) web framework  
**Database** PostgreSQL via [pgx](https://github.com/jackc/pgx) driver  
**Infrastructure:** AWS (planned deployment)  
**Planned** Redis, Elasticsearch

---

## 📖 Overview
This project is the backend for a Dungeons & Dragons virtual tabletop (VTT) application. It provides a structured API to manage games, characters, and source material, designed to support a smoother and more intuitive gameplay experience than existing options. Built with Golang and PostgreSQL, it emphasizes clean architecture and maintainable code.

<!--
---

## 🛠️ Features

-->
---

## 📂 Project Structure

dhs/  
├── cmd/  
│ └── server/  
│ └── main.go # Entry point  
├── internal/  
│ ├── config/ # Configuration  
│ ├── handlers/ # HTTP handlers for API endpoints  
│ ├── models/ # Database models / structs  
│ ├── routes/ # Route definitions  
│ ├── services/ # Business logic layer  
│ └── repository/ # Database access layer  
├── go.mod # Go modules file  
└── go.sum # Go dependencies checksum

---

## 🧠 Key Technical Concepts

- **RESTful API design** – clean and consistent endpoints for managing games and characters  
- **Database design** – normalized PostgreSQL schema accessed via 'pgx'  
- **Clean architecture** – separation of concerns between controllers, services, and repositories  
- **Robustness & maintainability** – structured error handling, validation, and scalable code organization

<!--
---

## 🧩 Challenges & Solutions
-->
---

## 🧪 Running Locally

### 1. Clone the repository
```bash
git clone https://github.com/rainhunt/dhs.git
```
### 2. Install dependencies
```bash
go mod download
```
### 3. Set up ENVs
```bash
export PORT="1323"
```
### 4. Start the backend
```bash
go run main.go
```

---

## 🗺️ Roadmap / Future Improvements

1. **User authentication & profiles** – implement JWT-based login, registration, and user profile management
2. **CRUD endpoints for reference material** – implement comprehensive CRUD endpoints for the SRD and Homebrew content
3. **Data validation & error handling** – implement input validation, structured errors, and consistent API responses
4. **Database enhancements** – normalize PostgreSQL schema, optimize queries, and add indexing
5. **Real time updates** – add WebSocket support for live game updates and notifications
6. **Caching & search** – integrate Redis for caching and Elasticsearch for fuzzy search
7. **Automated testing** – unit and integration tests for controllers, services, and repositories

<!--
---

## 🎨 Screenshots
-->
---

## 👤 Author

Rain Golombek
Fullstack Developer — React • Node.js • TypeScript
📍 Ramat Gan, Israel
🔗 LinkedIn: [Rain](https://www.linkedin.com/in/rain-golombek-fullstack/)

🔗 Portfolio: [Website](https://rainhunt.github.io/webPortfolio/)

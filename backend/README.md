Tentu, ini adalah struktur file `README.md` lengkap yang sudah saya optimasi agar terlihat sangat profesional di GitHub. File ini menonjolkan kemampuan kamu dalam **System Design**, **Automation**, dan **AI Integration**.

Salin seluruh teks di bawah ini ke dalam file `README.md` di root project kamu:

```markdown
# 🎓 Grades Management System (AI-Powered EdTech Pipeline)

[![Go Report Card](https://goreportcard.com/badge/github.com/Fadlirmn/grades_management)](https://goreportcard.com/report/github.com/Fadlirmn/grades_management)
> 🚀 An automated intelligence platform that syncs classroom data, evaluates learning objectives, and generates AI recommendations using **Go**, **Gemini AI**, and **Google Sheets API**.

---

## 💡 The Problem & Solution

### ❗ The Problem
Teachers often struggle with:
- **Data Silos**: Manually syncing data from physical spreadsheets to databases is time-consuming.
- **Surface-Level Grading**: Only recording scores without understanding which **Learning Objectives** are missed.
- **Scalability**: Providing personalized feedback for 40+ students per class is nearly impossible manually.

### ✅ The Solution
This system automates the entire pedagogical feedback loop:
1. **Extract**: Automatically pulls student grades from **Google Sheets** via a background Worker.
2. **Transform**: Processes raw scores into objective-based achievement levels (Belum, Cukup, Menguasai).
3. **AI Insight**: Leverages **Gemini-1.5-Flash** to generate concise, actionable 8-word learning recommendations.
4. **Automate**: Uses **Cron Jobs** to ensure data is always up-to-date without manual triggers.

---

## 🏗️ Technical Architecture

The project follows **Clean Architecture** principles (Repository, Service, Handler) to ensure maintainability and testability.



**Workflow:**
- **Sync Worker**: Periodically fetches data from Google Sheets API using Service Account authentication.
- **Upsert Logic**: Prevents data duplication using PostgreSQL `ON CONFLICT` constraints.
- **Analysis Worker**: Batches pending records (Min 20, Max 50) to optimize LLM token usage and costs.

---

## 🚀 Key Features

- 🔄 **Automated Google Sheets Sync**: Bi-directional data flow using Google Sheets API v4.
- 🧠 **AI Recommendation Engine**: Integrated with Google Gemini to provide automated teacher-like feedback.
- ⏱️ **Scheduled Batching**: Efficient background processing using `robfig/cron/v3`.
- 🔐 **Robust Security**: 
    - JWT Authentication with **Refresh Token** rotation.
    - **RBAC** (Role-Based Access Control) for Admin, Teacher, Parent, and Student.
- 🛠️ **Clean Code**: Implemented with Dependency Injection and Interface-based design.

---

## 🧑‍💻 Tech Stack

- **Language**: Go (Golang) 1.22+
- **Database**: PostgreSQL (Driver: `pgx/v5`, Library: `sqlx`)
- **AI Platform**: Google Gemini AI SDK
- **Cron Engine**: `robfig/cron/v3`
- **External API**: Google Sheets API v4
- **Web Framework**: Gin Gonic
- **Security**: JWT & Bcrypt

---

## 🧱 Project Structure

```text
.
├── config/             # DB, AI, and Google Sheets initialization
├── handlers/           # HTTP Request handlers & Input validation
├── middleware/         # Auth & Role-based access control (RBAC)
├── models/             # Database schemas & JSON structures
├── repository/         # PostgreSQL queries (Data Access Layer)
├── services/           # Business logic & Third-party service wrappers
├── worker/             # Background jobs (Sync & AI Analysis)
├── utils/              # JWT & String manipulation helpers
├── .env.example        # Environment variable template
└── main.go             # Application entry point & Cron setup
```

---

## ⚙️ Setup & Installation

### 1. Environment Configuration
Create a `.env` file based on `.env.example`:
```env
DB_URL=postgres://postgres:password@localhost:5432/grades_db?sslmode=disable
GEMINI_API_KEY=YOUR_GEMINI_API_KEY
SHEETS_ID=YOUR_SPREADSHEET_ID
GOOGLE_CREDENTIALS_PATH=credentials.json
JWT_SECRET=your_super_secret_key
```

### 2. Google Cloud Setup
1. Enable **Google Sheets API** in your Google Cloud Console.
2. Create a **Service Account** and download the `credentials.json` file.
3. Place `credentials.json` in the project root.
4. **Important**: Share your Google Sheet with the Service Account's email address (Editor access).

### 3. Run the Application
```bash
# Install dependencies
go mod tidy

# Run migrations (if applicable)
# go run cmd/migrate/main.go

# Start the server and cron workers
go run main.go
```

---

## 🧠 Roadmap & Future Work

- [x] **Phase 1**: Core Database & Subject Mapping.
- [x] **Phase 2**: JWT Auth & Token Rotation.
- [x] **Phase 3**: Background Sync Worker (Google Sheets).
- [x] **Phase 4**: AI Integration (Batch Recommendation).
- [ ] **Phase 5**: Real-time Notification System (WhatsApp/Telegram).
- [ ] **Phase 6**: Interactive Dashboard with Next.js & Recharts.

---

## 📌 Author
**Sumbul** - *Backend Developer & System Architect*
Passionate about automating workflows and building scalable EdTech solutions.

---

### ⭐ Support
If you find this project useful, feel free to **Star** the repository and contribute!
```

**Langkah selanjutnya:**
1. Buat file bernama `README.md` di folder utama project kamu.
2. Paste kode di atas.
3. Jangan lupa tambahkan file `.env.example` yang berisi daftar variabel lingkungan tanpa nilai rahasianya, agar orang lain tahu apa saja yang perlu diatur.

Readme ini bakal bikin profil GitHub kamu terlihat sangat "senior" karena menjelaskan *logic* di balik kode, bukan cuma cara pakainya saja! Ada bagian lain yang mau kamu tambahkan?
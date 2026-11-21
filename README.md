# 🎓 Grades Management System

AI-powered school grades management platform: Go backend syncing data from Google Sheets, running AI-based learning recommendations, with a Vue 3 frontend for Admin/Teacher/Parent/Student roles.

Monorepo layout:
```
.
├── backend/   # Go API (Gin, PostgreSQL, JWT auth, cron workers, Gemini AI)
└── frontend/  # Vue 3 + Vuetify SPA (Vite)
```

## Backend (`backend/`)

- **Stack**: Go, Gin, PostgreSQL (`pgx`/`sqlx`), JWT auth with refresh tokens, RBAC (Admin/Teacher/Parent/Student).
- **Automation**: `robfig/cron` workers sync grades from Google Sheets and batch-generate AI recommendations via Gemini.

```bash
cd backend
cp .env.example .env   # fill in DB_URL, GEMINI_API_KEY, SHEETS_ID, GOOGLE_CREDENTIALS_PATH, JWT_SECRET
go mod tidy
go run main.go
```

## Frontend (`frontend/`)

- **Stack**: Vue 3, Vuetify, Pinia, Vue Router, Vite, UnoCSS.
- Calls the backend API for auth, grades, and AI recommendations.

```bash
cd frontend
cp .env.example .env   # set API base URL to point at the backend
npm install
npm run dev     # local dev
npm run build   # production build -> dist/
```

## Deployment notes

- Frontend: static build, deployable to Vercel/Cloudflare Pages/GitHub Pages. On Pages platforms set **root directory = `frontend`**. GitHub Pages project pages need `base: '/<repo-name>/'` added to `vite.config.js`; Vercel/Cloudflare/Netlify don't.
- Backend: needs a host that runs a persistent process (cron workers run in-process), e.g. Render, Railway, Fly.io, or a VPS — not a static/serverless-only platform.

## Roles

- **Admin**: full access (users, schools, database).
- **Teacher**: input grades, manage classes, view AI recommendations per class.
- **Parent**: monitor child's grades and AI development suggestions.
- **Student**: view own grades and AI study tips.

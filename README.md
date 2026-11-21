# 🎓 Grades Management System

AI-powered school grades platform. A Go backend syncs grades from Google Sheets into PostgreSQL and generates AI-based learning recommendations (Gemini); a Vue 3 frontend gives Admin/Teacher/Parent/Student each their own dashboard.

```
.
├── backend/   # Go API
└── frontend/  # Vue 3 SPA
```

## Stack

| Layer | Tech |
|---|---|
| Backend | Go, Gin, PostgreSQL (`pgx` + `sqlx`), JWT auth + refresh tokens, `robfig/cron`, Google Sheets API, Gemini AI |
| Frontend | Vue 3, Vuetify, Pinia, Vue Router, Vite, UnoCSS, TypeScript |

## Frontend pages

| Route | Page | Roles |
|---|---|---|
| `/login`, `/register` | Auth | everyone |
| `/dashboard` | Overview stats | all |
| `/dashboard/students` | Student detail & grades | admin, teacher, parent |
| `/dashboard/subjects` | Subjects | admin, teacher |
| `/dashboard/assignments` | Assignments | admin, teacher, student |
| `/dashboard/progress` | Progress & AI recommendation | admin, teacher, parent |
| `/dashboard/reports` | Reports | admin, teacher |
| `/dashboard/admin-data` | Raw CRUD over all tables | admin |
| `/dashboard/profile` | Profile edit | all |

Role-based routing is enforced client-side in `frontend/src/router/index.ts`.

### Dashboard page

![Dashboard](docs/screenshots/dashboard.png)

- Stat cards: total students, active subjects, weekly progress %, average score (with trend deltas).
- Weekly progress trend chart (submissions vs. engagement).
- Top subjects ranked by score, with an AI-generated suggestion for the weakest one.
- Secondary stats: faculty active, system uptime, pending reviews.
- Upcoming assignments table (name, due date, status, submission count).

![Login](docs/screenshots/login.png)

## Demo (no backend required)

The frontend works standalone: every API call falls back to mock data when the backend is unreachable (`frontend/src/api/*.ts`). Deployed as a static site (Vercel/Cloudflare Pages/GitHub Pages), it's a full click-through demo with no server.

Login with any password:
- `admin@edu.com`
- `teacher@edu.com`
- `student@edu.com`
- `parent@edu.com`

## Backend API

Base path `/api` (JWT-protected, role-gated per route) plus `/auth` (public):

| Group | Endpoints |
|---|---|
| `/auth` | `POST /register`, `POST /login`, `POST /refresh` |
| `/api/assign` | CRUD, admin/teacher write, all roles read |
| `/api/objective` | CRUD, admin/teacher write, all roles read |
| `/api/progress` | CRUD, admin/teacher write, all roles read (includes AI recommendation) |
| `/api/score` | CRUD, admin/teacher write, all roles read |
| `/api/student` | CRUD, admin only write, all roles read |
| `/api/subject` | CRUD, admin only write, all roles read |
| `/api/users` | CRUD, admin only |

A cron job runs hourly: syncs grades from Google Sheets into Postgres, then batches pending records through Gemini for AI recommendations.

## Running locally

**Backend**
```bash
cd backend
cp .env.example .env   # DB_HOST/PORT/USER/PASSWORD/NAME, API_KEY (Gemini), SHEETS_ID, GOOGLE_CREDENTIALS_PATH
go mod tidy
go run main.go         # listens on :8080
```

**Frontend**
```bash
cd frontend
npm install
npm run dev             # http://localhost:5173, uses mock data if backend isn't running
npm run build            # production build -> dist/
```

## Deploying

- **Frontend only (demo/portfolio)**: static build, deploy to Vercel/Cloudflare Pages/GitHub Pages with **root directory = `frontend`**.
  - **GitHub Pages**: a workflow is already set up (`.github/workflows/deploy-pages.yml`) — it builds `frontend/` with the right base path and deploys on every push to `main`. One-time setup: in the repo's Settings → Pages, set **Source: GitHub Actions**. Site then lives at `https://<user>.github.io/grades_management/`.
  - Vercel/Cloudflare/Netlify don't need a base path — leave `vite.config.js`'s `base` as default (`/`).
- **Full stack on a VPS**: `docker-compose.yml` at repo root runs Postgres + backend + frontend (nginx) together.
  ```bash
  cp .env.example .env   # fill in DB_PASSWORD, API_KEY, SHEETS_ID, VITE_API_URL (your VPS's public address)
  # place your Google service account file at backend/credentials.json
  docker compose up -d --build
  ```
  Backend needs a host that keeps a persistent process running (the cron worker runs in-process) — a VPS, or Render/Railway/Fly.io. Not deployable to a static/serverless-only platform.

## Roles

- **Admin** — full access: users, schools, raw data tables.
- **Teacher** — input grades, manage classes, view AI recommendations for their class.
- **Parent** — monitor their child's grades and AI development suggestions.
- **Student** — view own grades and AI study tips.

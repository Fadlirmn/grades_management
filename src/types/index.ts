// ─── Role Types ───────────────────────────────────────────
export type UserRole = 'admin' | 'teacher' | 'student' | 'parent'

// ─── Auth ─────────────────────────────────────────────────
export interface LoginPayload {
  email: string
  password: string
}

export interface RegisterPayload {
  username: string
  name: string
  email: string
  password: string
  role: UserRole
}

export interface AuthTokens {
  accessToken: string
  refreshToken: string
}

export interface AuthResponse {
  user: User
  tokens: AuthTokens
}

// ─── User ─────────────────────────────────────────────────
export interface User {
  id: string
  username?: string
  name: string
  email: string
  role: UserRole
  avatar?: string
  department?: string
  phone?: string
  joinedAt?: string
  status: 'active' | 'inactive'
}

// ─── Student ──────────────────────────────────────────────
export interface Student {
  id: string
  name: string
  initials: string
  studentId: string    // e.g. STU-2024-8812
  grade: string        // e.g. "11-A"
  major: string
  gpa: number
  attendance: number   // percentage
  ranking: number
  totalStudents: number
  status: 'excellent' | 'good' | 'average' | 'at-risk'
  avatarUrl?: string
}

// ─── Subject ──────────────────────────────────────────────
export interface Subject {
  id: string
  name: string
  instructor: string
  icon: string
  color: string
  studentCount: number
  avgScore: number
  masteryRate: number  // percentage
}

// ─── Assignment ───────────────────────────────────────────
export type AssignmentStatus = 'in-progress' | 'completed' | 'pending' | 'overdue'

export interface Assignment {
  id: string
  name: string
  category: string
  subject: string
  dueDate: string
  status: AssignmentStatus
  submittedCount: number
  totalCount: number
  avgScore: number | null
  icon: string
  iconColor: string
}

// ─── Grade / Score ────────────────────────────────────────
export type MasteryStatus = 'menguasai' | 'cukup' | 'belum'

export interface Grade {
  id: string
  studentId: string
  studentName: string
  initials: string
  subject: string
  week: string
  learningObjective: string
  score: number
  status: MasteryStatus
  aiRecommendation: string
}

// ─── Progress ─────────────────────────────────────────────
export interface WeeklyProgress {
  week: number
  label: string
  title: string
  description: string
  status: 'current' | 'completed' | 'pending'
}

export interface MasteryBreakdown {
  name: string
  percentage: number
  color: string
}

// ─── Dashboard Stats ──────────────────────────────────────
export interface DashboardStat {
  label: string
  value: string | number
  icon: string
  iconBg: string
  iconBgDark: string
  iconColor: string
  trend: string
  trendColor: string
  bar?: number
}

// ─── Activity Log ─────────────────────────────────────────
export interface ActivityLog {
  id: string
  time: string
  text: string
  subtext?: string
  dotColor: string
}

// ─── AI Insight ───────────────────────────────────────────
export interface AIInsight {
  title: string
  description: string
  type: 'suggestion' | 'alert' | 'forecast'
}

// ─── Report ───────────────────────────────────────────────
export interface Report {
  id: string
  title: string
  description: string
  lastGenerated: string
  icon: string
  color: string
  bgLight: string
  bgDark: string
}

// ─── Enrolled Subject (Student Detail) ────────────────────
export interface EnrolledSubject {
  name: string
  instructor: string
  grade: string
  credits: number
  status: MasteryStatus
  icon: string
  iconBg: string
  iconBgDark: string
  iconColor: string
}

// ═══════════════════════════════════════════════════════════
// DB Table Models (Admin CRUD)
// ═══════════════════════════════════════════════════════════

export interface DBUser {
  id?: number
  name: string
  email: string
  password?: string
  role: UserRole
  created_at?: string
  updated_at?: string
}

export interface DBStudent {
  id?: number
  user_id: number
  nis: string
  class: string
  parent_name?: string
  created_at?: string
  updated_at?: string
}

export interface DBSubject {
  id?: number
  name: string
  code: string
  teacher_id: number
  description?: string
  created_at?: string
  updated_at?: string
}

export interface DBAssignment {
  id?: number
  subject_id: number
  title: string
  description?: string
  due_date: string
  created_at?: string
  updated_at?: string
}

export interface DBObjective {
  id?: number
  subject_id: number
  week: number
  description: string
  created_at?: string
  updated_at?: string
}

export interface DBProgress {
  id?: number
  student_id: number
  objective_id: number
  status: MasteryStatus
  notes?: string
  created_at?: string
  updated_at?: string
}

export interface DBScore {
  id?: number
  student_id: number
  assignment_id: number
  score: number
  feedback?: string
  created_at?: string
  updated_at?: string
}

// Generic type for table rows
export type DBTableRow = DBUser | DBStudent | DBSubject | DBAssignment | DBObjective | DBProgress | DBScore

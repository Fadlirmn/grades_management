import apiClient from './client'
import type {
  DBUser, DBStudent, DBSubject, DBAssignment,
  DBObjective, DBProgress, DBScore,
} from '@/types'

// ─── Generic CRUD helpers ────────────────────────────────
// Semua endpoint mengikuti pola: GET /table, POST /table, PUT /table/:id, DELETE /table/:id

type TableName = 'users' | 'students' | 'subjects' | 'assignments' | 'objectives' | 'progress' | 'scores'

async function getAll<T>(table: TableName): Promise<T[]> {
  try {
    const { data } = await apiClient.get<T[]>(`/${table}`)
    return data
  } catch {
    return []
  }
}

async function create<T>(table: TableName, payload: Partial<T>): Promise<T> {
  const { data } = await apiClient.post<T>(`/${table}`, payload)
  return data
}

async function update<T>(table: TableName, id: number, payload: Partial<T>): Promise<T> {
  const { data } = await apiClient.put<T>(`/${table}/${id}`, payload)
  return data
}

async function remove(table: TableName, id: number): Promise<void> {
  await apiClient.delete(`/${table}/${id}`)
}

// ─── Users ───────────────────────────────────────────────
export const usersApi = {
  getAll: () => getAll<DBUser>('users'),
  create: (p: Partial<DBUser>) => create<DBUser>('users', p),
  update: (id: number, p: Partial<DBUser>) => update<DBUser>('users', id, p),
  delete: (id: number) => remove('users', id),
}

// ─── Students ────────────────────────────────────────────
export const studentsApi = {
  getAll: () => getAll<DBStudent>('students'),
  create: (p: Partial<DBStudent>) => create<DBStudent>('students', p),
  update: (id: number, p: Partial<DBStudent>) => update<DBStudent>('students', id, p),
  delete: (id: number) => remove('students', id),
}

// ─── Subjects ────────────────────────────────────────────
export const subjectsApi = {
  getAll: () => getAll<DBSubject>('subjects'),
  create: (p: Partial<DBSubject>) => create<DBSubject>('subjects', p),
  update: (id: number, p: Partial<DBSubject>) => update<DBSubject>('subjects', id, p),
  delete: (id: number) => remove('subjects', id),
}

// ─── Assignments ─────────────────────────────────────────
export const assignmentsApi = {
  getAll: () => getAll<DBAssignment>('assignments'),
  create: (p: Partial<DBAssignment>) => create<DBAssignment>('assignments', p),
  update: (id: number, p: Partial<DBAssignment>) => update<DBAssignment>('assignments', id, p),
  delete: (id: number) => remove('assignments', id),
}

// ─── Objectives ──────────────────────────────────────────
export const objectivesApi = {
  getAll: () => getAll<DBObjective>('objectives'),
  create: (p: Partial<DBObjective>) => create<DBObjective>('objectives', p),
  update: (id: number, p: Partial<DBObjective>) => update<DBObjective>('objectives', id, p),
  delete: (id: number) => remove('objectives', id),
}

// ─── Progress ────────────────────────────────────────────
export const progressApi = {
  getAll: () => getAll<DBProgress>('progress'),
  create: (p: Partial<DBProgress>) => create<DBProgress>('progress', p),
  update: (id: number, p: Partial<DBProgress>) => update<DBProgress>('progress', id, p),
  delete: (id: number) => remove('progress', id),
}

// ─── Scores ──────────────────────────────────────────────
export const scoresApi = {
  getAll: () => getAll<DBScore>('scores'),
  create: (p: Partial<DBScore>) => create<DBScore>('scores', p),
  update: (id: number, p: Partial<DBScore>) => update<DBScore>('scores', id, p),
  delete: (id: number) => remove('scores', id),
}

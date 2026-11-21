import apiClient from './client'
import type { Grade, Student, DashboardStat } from '@/types'

/**
 * Fetch all grades / progress records
 */
export async function fetchGrades(): Promise<Grade[]> {
  try {
    const { data } = await apiClient.get<Grade[]>('/grades')
    return data
  } catch {
    return mockGrades()
  }
}

/**
 * Fetch student detail by ID
 */
export async function fetchStudentById(id: string): Promise<Student> {
  try {
    const { data } = await apiClient.get<Student>(`/students/${id}`)
    return data
  } catch {
    return mockStudent()
  }
}

/**
 * Fetch dashboard stats
 */
export async function fetchDashboardStats(): Promise<DashboardStat[]> {
  try {
    const { data } = await apiClient.get<DashboardStat[]>('/dashboard/stats')
    return data
  } catch {
    return mockStats()
  }
}

// ─── Mock Data ───────────────────────────────────────────

function mockGrades(): Grade[] {
  return [
    { id: '1', studentId: 's1', studentName: 'Adrian Miller', initials: 'AM', subject: 'Mathematics', week: 'W-12', learningObjective: 'Advanced Calculus & Int...', score: 92, status: 'menguasai', aiRecommendation: 'Unlock advanced honors physics modules' },
    { id: '2', studentId: 's2', studentName: 'Sarah Toussaint', initials: 'ST', subject: 'English Literature', week: 'W-12', learningObjective: 'Modernist Narrative Devi...', score: 74, status: 'cukup', aiRecommendation: 'Assign Joyce & Woolf supplemental reading' },
    { id: '3', studentId: 's3', studentName: 'Kevin Bryant', initials: 'KB', subject: 'World History', week: 'W-11', learningObjective: 'Cold War Geo-politics', score: 58, status: 'belum', aiRecommendation: 'Schedule 1-on-1 for timeline mapping' },
    { id: '4', studentId: 's4', studentName: 'Li Zhang', initials: 'LZ', subject: 'Computer Science', week: 'W-12', learningObjective: 'Algorithm Complexity An...', score: 98, status: 'menguasai', aiRecommendation: 'Invite to National Coding Olympiad program' },
  ]
}

function mockStudent(): Student {
  return {
    id: 's1', name: 'Alexander Chen', initials: 'AC', studentId: 'STU-2024-8812',
    grade: '11-A', major: 'Computer Science', gpa: 3.92, attendance: 98,
    ranking: 4, totalStudents: 120, status: 'excellent',
  }
}

function mockStats(): DashboardStat[] {
  return [
    { label: 'Total Students', value: '124', icon: 'mdi-account-group', iconBg: 'blue-lighten-4', iconBgDark: '#1e3a5f', iconColor: 'primary', trend: '+4%', trendColor: 'success' },
    { label: 'Active Subjects', value: '8', icon: 'mdi-bookshelf', iconBg: 'orange-lighten-4', iconBgDark: '#4a2c00', iconColor: 'tertiary', trend: 'Stable', trendColor: 'grey' },
    { label: 'Weekly Progress', value: '82%', icon: 'mdi-trending-up', iconBg: 'green-lighten-4', iconBgDark: '#14532d', iconColor: 'success', trend: '+12%', trendColor: 'success', bar: 82 },
    { label: 'Average Score', value: '85.5', icon: 'mdi-star', iconBg: 'yellow-lighten-4', iconBgDark: '#422006', iconColor: 'yellow-darken-3', trend: '-0.2%', trendColor: 'error' },
  ]
}

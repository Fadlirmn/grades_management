<template>
  <div>
    <!-- Header -->
    <div class="d-flex flex-column flex-md-row justify-space-between align-start mb-6 mb-md-8 gap-4">
      <div>
        <h1 class="text-h5 text-md-h4 font-weight-bold text-high-emphasis mb-1">
          Welcome back, {{ authStore.userName.split(' ')[0] || 'User' }}.
        </h1>
        <p class="text-body-2 text-medium-emphasis">
          <template v-if="authStore.isAdmin">System overview and management tools.</template>
          <template v-else-if="authStore.isTeacher">Your academic overview for <a href="#" class="text-primary font-weight-bold text-decoration-underline">Spring Semester 2024</a>.</template>
          <template v-else-if="authStore.isStudent">Your personal academic progress and assignments.</template>
          <template v-else>Your child's academic performance overview.</template>
        </p>
      </div>
      <div class="d-flex gap-3">
        <v-btn v-if="authStore.hasRole('admin', 'teacher')" variant="outlined" prepend-icon="mdi-download" rounded="lg" class="text-none font-weight-bold text-medium-emphasis">Export PDF</v-btn>
        <v-btn v-if="authStore.hasRole('admin', 'teacher')" color="primary" elevation="0" prepend-icon="mdi-plus" rounded="lg" class="text-none font-weight-bold">New Assignment</v-btn>
      </div>
    </div>

    <!-- Role Badge -->
    <v-chip :color="roleBadge.color" variant="tonal" size="small" class="font-weight-bold text-uppercase mb-5" :prepend-icon="roleBadge.icon">
      {{ roleBadge.label }} Dashboard
    </v-chip>

    <!-- Stat Cards -->
    <div class="grid grid-cols-4 gap-4 gap-md-5 mb-7">
      <v-card v-for="s in filteredStats" :key="s.label" color="surface" variant="flat" rounded="xl" class="pa-5 border-thin">
        <div class="d-flex justify-space-between align-center mb-3">
          <v-avatar :color="isDark ? s.iconBgDark : s.iconBg" size="40" rounded="lg"><v-icon :color="s.iconColor" size="20">{{ s.icon }}</v-icon></v-avatar>
          <v-chip :color="s.trendColor" size="x-small" variant="tonal" class="font-weight-bold">{{ s.trend }}</v-chip>
        </div>
        <div class="text-caption font-weight-medium text-medium-emphasis">{{ s.label }}</div>
        <div class="text-h4 font-weight-bold text-high-emphasis">{{ s.value }}</div>
      </v-card>
    </div>

    <!-- Charts Row (Admin + Teacher) -->
    <div v-if="authStore.hasRole('admin', 'teacher')" class="grid grid-cols-3 gap-4 gap-md-5 mb-7">
      <v-card color="surface" variant="flat" rounded="xl" class="pa-5 pa-md-6 col-span-2 border-thin">
        <div class="d-flex flex-column flex-sm-row justify-space-between align-start mb-5 gap-2">
          <div>
            <h3 class="text-subtitle-1 font-weight-bold text-high-emphasis">Weekly Progress Trend</h3>
            <p class="text-caption text-medium-emphasis">Class engagement and submission rates</p>
          </div>
          <div class="d-flex align-center gap-2 text-caption text-medium-emphasis">
            <span class="d-inline-block rounded-circle bg-primary" style="width:8px;height:8px;"></span> Submissions
            <span class="d-inline-block rounded-circle bg-blue-lighten-3 ml-2" style="width:8px;height:8px;"></span> Engagement
          </div>
        </div>
        <div style="height:220px;">
          <canvas ref="lineChartRef"></canvas>
        </div>
      </v-card>

      <v-card color="surface" variant="flat" rounded="xl" class="pa-5 pa-md-6 border-thin">
        <h3 class="text-subtitle-1 font-weight-bold text-high-emphasis mb-4">Top Subjects</h3>
        <div v-for="sub in topSubjects" :key="sub.name" class="mb-4">
          <div class="d-flex justify-space-between mb-1">
            <span class="text-body-2 font-weight-bold text-high-emphasis">{{ sub.name }}</span>
            <span class="text-body-2 font-weight-bold text-primary">{{ sub.score }}</span>
          </div>
          <v-progress-linear :model-value="sub.score" color="primary" bg-color="surface-variant" rounded height="6"></v-progress-linear>
        </div>

        <v-card color="surface-variant" variant="flat" rounded="lg" class="pa-3 mt-5">
          <div class="d-flex align-center gap-1 mb-2">
            <v-icon size="16" color="tertiary">mdi-lightbulb-on</v-icon>
            <span class="text-uppercase font-weight-bold text-medium-emphasis" style="font-size:0.6rem;letter-spacing:0.08em;">AI Suggestion</span>
          </div>
          <p class="text-body-2 font-italic ma-0 text-medium-emphasis" style="line-height:1.5;">"Physics 101 participation is lagging. Consider adding a gamified quiz to boost interest."</p>
        </v-card>
      </v-card>
    </div>

    <!-- Student Personal Progress (Student role) -->
    <div v-if="authStore.isStudent" class="grid grid-cols-2 gap-4 mb-7">
      <v-card color="surface" variant="flat" rounded="xl" class="pa-6 border-thin">
        <h3 class="text-subtitle-1 font-weight-bold text-high-emphasis mb-4">My Grades This Week</h3>
        <div v-for="g in studentGrades" :key="g.subject" class="d-flex justify-space-between align-center mb-3 pb-3 border-b-thin">
          <div>
            <div class="text-body-2 font-weight-bold text-high-emphasis">{{ g.subject }}</div>
            <div class="text-caption text-medium-emphasis">{{ g.objective }}</div>
          </div>
          <div class="d-flex align-center gap-3">
            <span class="text-body-2 font-weight-bold text-high-emphasis">{{ g.score }}</span>
            <v-chip :color="g.chipColor" size="x-small" variant="tonal" class="font-weight-bold text-uppercase">{{ g.status }}</v-chip>
          </div>
        </div>
      </v-card>

      <v-card color="surface" variant="flat" rounded="xl" class="pa-6 border-thin">
        <h3 class="text-subtitle-1 font-weight-bold text-high-emphasis mb-4">Upcoming Deadlines</h3>
        <div v-for="d in deadlines" :key="d.name" class="d-flex align-center gap-3 mb-3 pb-3 border-b-thin">
          <v-avatar :color="d.color" size="36" rounded="lg"><v-icon size="16" color="white">{{ d.icon }}</v-icon></v-avatar>
          <div class="flex-1">
            <div class="text-body-2 font-weight-bold text-high-emphasis">{{ d.name }}</div>
            <div class="text-caption text-medium-emphasis">{{ d.due }}</div>
          </div>
          <v-chip :color="d.urgency" size="x-small" variant="tonal" class="font-weight-bold">{{ d.daysLeft }}</v-chip>
        </div>
      </v-card>
    </div>

    <!-- Parent: Child Overview -->
    <div v-if="authStore.isParent" class="grid grid-cols-2 gap-4 mb-7">
      <v-card color="surface" variant="flat" rounded="xl" class="pa-6 border-thin">
        <h3 class="text-subtitle-1 font-weight-bold text-high-emphasis mb-4">
          <v-icon start size="18" color="primary">mdi-account-child-circle</v-icon>
          Alexander's Performance
        </h3>
        <div v-for="s in childSubjects" :key="s.name" class="d-flex justify-space-between align-center mb-3 pb-3 border-b-thin">
            <span class="text-body-2 font-weight-bold text-high-emphasis">{{ s.name }}</span>
            <span class="text-body-2 font-weight-bold text-primary">{{ s.grade }}</span>
        </div>
      </v-card>

      <v-card color="surface" variant="flat" rounded="xl" class="pa-6 border-thin">
        <h3 class="text-subtitle-1 font-weight-bold text-high-emphasis mb-4">
          <v-icon start size="18" color="tertiary">mdi-sparkles</v-icon>
          AI Parent Insights
        </h3>
        <v-alert type="info" variant="tonal" rounded="lg" class="mb-3">
          <div class="text-body-2 text-high-emphasis">Alexander excels in STEM subjects. Consider encouraging participation in the upcoming Science Olympiad.</div>
        </v-alert>
        <v-alert type="warning" variant="tonal" rounded="lg">
          <div class="text-body-2 text-high-emphasis">English Literature scores have dipped by 8% this week. A tutor session is recommended.</div>
        </v-alert>
      </v-card>
    </div>

    <!-- Admin: System Metrics -->
    <div v-if="authStore.isAdmin" class="grid grid-cols-3 gap-4 mb-7">
      <v-card v-for="m in adminMetrics" :key="m.label" color="surface" variant="flat" rounded="xl" class="pa-5 border-thin">
        <div class="d-flex align-center gap-3 mb-3">
          <v-avatar :color="m.bg" size="40" rounded="lg"><v-icon :color="m.color" size="20">{{ m.icon }}</v-icon></v-avatar>
          <div>
            <div class="text-caption text-medium-emphasis">{{ m.label }}</div>
            <div class="text-h5 font-weight-bold text-high-emphasis">{{ m.value }}</div>
          </div>
        </div>
        <v-progress-linear :model-value="m.pct" :color="m.color" bg-color="surface-variant" rounded height="5"></v-progress-linear>
      </v-card>
    </div>

    <!-- Upcoming Assignments (Admin + Teacher) -->
    <div v-if="authStore.hasRole('admin', 'teacher')">
      <div class="d-flex justify-space-between align-center mb-4">
        <h3 class="text-subtitle-1 font-weight-bold text-high-emphasis">Upcoming Assignments</h3>
        <a href="#" class="text-body-2 font-weight-bold text-primary text-decoration-none">View Schedule</a>
      </div>

      <v-card color="surface" variant="flat" rounded="xl" class="border-thin">
        <v-table>
          <thead>
            <tr>
              <th class="text-uppercase font-weight-bold text-medium-emphasis" style="font-size:0.65rem;letter-spacing:0.08em;">Assignment Name</th>
              <th class="text-uppercase font-weight-bold text-medium-emphasis d-none d-sm-table-cell" style="font-size:0.65rem;letter-spacing:0.08em;">Due Date</th>
              <th class="text-uppercase font-weight-bold text-medium-emphasis" style="font-size:0.65rem;letter-spacing:0.08em;">Status</th>
              <th class="text-uppercase font-weight-bold text-medium-emphasis d-none d-sm-table-cell" style="font-size:0.65rem;letter-spacing:0.08em;">Submissions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="a in assignments" :key="a.name" class="cursor-pointer">
              <td>
                <div class="d-flex align-center gap-3">
                  <v-avatar :color="a.iconColor" size="36" rounded="lg"><v-icon size="16" color="white">{{ a.icon }}</v-icon></v-avatar>
                  <div>
                    <div class="text-body-2 font-weight-bold text-high-emphasis">{{ a.name }}</div>
                    <div class="text-caption text-medium-emphasis">{{ a.category }}</div>
                  </div>
                </div>
              </td>
              <td class="text-body-2 text-medium-emphasis d-none d-sm-table-cell">{{ a.due }}</td>
              <td><v-chip :color="a.statusColor" size="x-small" variant="tonal" class="font-weight-bold text-uppercase">{{ a.status }}</v-chip></td>
              <td class="text-body-2 font-weight-bold text-high-emphasis d-none d-sm-table-cell">{{ a.submissions }}</td>
            </tr>
          </tbody>
        </v-table>
      </v-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useTheme } from 'vuetify'
import { useAuthStore } from '@/stores/auth'
import { fetchDashboardStats } from '@/api/grades'
import apiClient from '@/api/client'
import type { DashboardStat } from '@/types'
import Chart from 'chart.js/auto'

const theme = useTheme()
const authStore = useAuthStore()
const isDark = computed(() => theme.global.name.value === 'dark')
const lineChartRef = ref<HTMLCanvasElement | null>(null)
let chartInstance: Chart | null = null

// ─── Role Badge ──────────────────────────────────────────
const roleBadge = computed(() => {
  const map: Record<string, { color: string; icon: string; label: string }> = {
    admin:   { color: 'error',   icon: 'mdi-shield-crown',     label: 'Admin' },
    teacher: { color: 'primary', icon: 'mdi-human-male-board',  label: 'Teacher' },
    student: { color: 'success', icon: 'mdi-school',            label: 'Student' },
    parent:  { color: 'info',    icon: 'mdi-account-child',     label: 'Parent' },
  }
  return map[authStore.userRole || 'teacher'] || map.teacher
})

// ─── Stats (loaded from API) ─────────────────────────────
const stats = ref<DashboardStat[]>([])

const filteredStats = computed(() => {
  if (authStore.isStudent) {
    return stats.value.filter(s => ['Weekly Progress', 'Average Score'].includes(s.label))
  }
  if (authStore.isParent) {
    return stats.value.filter(s => ['Weekly Progress', 'Average Score', 'Active Subjects'].includes(s.label))
  }
  return stats.value
})

// ─── Reactive data (API-first, mock fallback) ────────────
const topSubjects = ref([
  { name: 'Calculus II', score: 94.2 },
  { name: 'World Literature', score: 88.5 },
  { name: 'Digital Arts', score: 82.1 },
  { name: 'Physics 101', score: 76.4 },
])

const assignments = ref([
  { name: 'Mid-Term Calculus Essay', category: 'Mathematics Tier A', due: 'Oct 24, 2023', status: 'In Progress', statusColor: 'primary', submissions: '88/124', icon: 'mdi-file-document', iconColor: 'tertiary' },
  { name: 'Physics Lab Report: Optics', category: 'Natural Sciences', due: 'Oct 26, 2023', status: 'MENGUASAI', statusColor: 'success', submissions: '112/124', icon: 'mdi-flask', iconColor: 'primary' },
  { name: 'Digital Renaissance Portfolio', category: 'Arts & Design', due: 'Oct 30, 2023', status: 'BELUM', statusColor: 'error', submissions: '12/124', icon: 'mdi-palette', iconColor: 'purple' },
])

const studentGrades = ref([
  { subject: 'Advanced Calculus', objective: 'Integration Techniques', score: 94, status: 'MENGUASAI', chipColor: 'primary' },
  { subject: 'Computer Science II', objective: 'Algorithm Complexity', score: 98, status: 'MENGUASAI', chipColor: 'primary' },
  { subject: 'World History', objective: 'Cold War Analysis', score: 72, status: 'CUKUP', chipColor: 'warning' },
])

const deadlines = ref([
  { name: 'Calculus Problem Set #12', due: 'Oct 24, 2023', daysLeft: '2 days', urgency: 'error', icon: 'mdi-calculator', color: 'primary' },
  { name: 'CS Project Submission', due: 'Oct 28, 2023', daysLeft: '6 days', urgency: 'warning', icon: 'mdi-code-braces', color: 'tertiary' },
  { name: 'History Essay Draft', due: 'Nov 1, 2023', daysLeft: '10 days', urgency: 'success', icon: 'mdi-book-open-variant', color: 'success' },
])

const childSubjects = ref([
  { name: 'Advanced Calculus', grade: 'A+' },
  { name: 'Computer Science II', grade: 'A' },
  { name: 'World History', grade: 'B+' },
  { name: 'English Literature', grade: 'B' },
])

const adminMetrics = ref([
  { label: 'Faculty Active', value: '42', icon: 'mdi-human-male-board', bg: 'blue-lighten-4', color: 'primary', pct: 84 },
  { label: 'System Uptime', value: '99.8%', icon: 'mdi-server', bg: 'green-lighten-4', color: 'success', pct: 99.8 },
  { label: 'Pending Reviews', value: '18', icon: 'mdi-clock-alert', bg: 'orange-lighten-4', color: 'tertiary', pct: 36 },
])

// ─── Load from API (fallback to mock already set above) ──
onMounted(async () => {
  // Stats
  stats.value = await fetchDashboardStats()

  // Top Subjects
  try {
    const { data } = await apiClient.get('/dashboard/top-subjects')
    if (data?.length) topSubjects.value = data
  } catch { /* keep mock */ }

  // Assignments
  try {
    const { data } = await apiClient.get('/dashboard/assignments')
    if (data?.length) assignments.value = data
  } catch { /* keep mock */ }

  // Student grades (student role)
  if (authStore.isStudent) {
    try {
      const { data } = await apiClient.get('/dashboard/my-grades')
      if (data?.length) studentGrades.value = data
    } catch { /* keep mock */ }

    try {
      const { data } = await apiClient.get('/dashboard/my-deadlines')
      if (data?.length) deadlines.value = data
    } catch { /* keep mock */ }
  }

  // Parent child subjects
  if (authStore.isParent) {
    try {
      const { data } = await apiClient.get('/dashboard/child-subjects')
      if (data?.length) childSubjects.value = data
    } catch { /* keep mock */ }
  }

  // Admin metrics
  if (authStore.isAdmin) {
    try {
      const { data } = await apiClient.get('/dashboard/admin-metrics')
      if (data?.length) adminMetrics.value = data
    } catch { /* keep mock */ }
  }

  createChart()
})

// ─── Chart ───────────────────────────────────────────────
const createChart = (): void => {
  if (chartInstance) chartInstance.destroy()
  if (!lineChartRef.value) return

  const gridColor = isDark.value ? '#252836' : '#f3f4f6'
  const tickColor = isDark.value ? '#6177A5' : '#9ca3af'

  chartInstance = new Chart(lineChartRef.value, {
    type: 'line',
    data: {
      labels: ['MON', 'TUE', 'WED', 'THU', 'FRI', 'SAT'],
      datasets: [
        { label: 'Submissions', data: [65, 72, 68, 80, 85, 78], borderColor: '#3B82F6', backgroundColor: 'rgba(59,130,246,0.08)', fill: true, tension: 0.4, pointRadius: 0, borderWidth: 2 },
        { label: 'Engagement', data: [55, 60, 62, 70, 75, 72], borderColor: '#93c5fd', backgroundColor: 'rgba(147,197,253,0.05)', fill: true, tension: 0.4, pointRadius: 0, borderWidth: 2 },
      ],
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: { legend: { display: false } },
      scales: {
        y: { display: false },
        x: { grid: { color: gridColor }, ticks: { color: tickColor, font: { size: 11, family: 'Inter' } } },
      },
    },
  })
}

watch(isDark, () => setTimeout(createChart, 100))
</script>


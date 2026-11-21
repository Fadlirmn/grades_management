<template>
  <div>
    <div class="d-flex flex-column flex-sm-row justify-space-between align-start mb-6 gap-3">
      <div>
        <h1 class="text-h5 text-md-h4 font-weight-bold text-high-emphasis">Assignments</h1>
        <p class="text-body-2 mt-1 text-medium-emphasis">Create, manage and track all class assignments.</p>
      </div>
      <v-btn color="primary" elevation="0" prepend-icon="mdi-plus" rounded="lg" class="text-none font-weight-bold">Create Assignment</v-btn>
    </div>

    <v-btn-toggle v-model="activeTab" mandatory density="compact" rounded="lg" color="primary" variant="outlined" class="mb-5">
      <v-btn v-for="t in tabs" :key="t" :value="t" class="text-none text-body-2">{{ t }}</v-btn>
    </v-btn-toggle>

    <v-card color="surface" variant="flat" rounded="xl" class="border-thin">
      <v-table>
        <thead>
          <tr>
            <th class="text-uppercase font-weight-bold text-medium-emphasis" style="font-size:0.65rem;letter-spacing:0.08em;">Assignment</th>
            <th class="text-uppercase font-weight-bold text-medium-emphasis d-none d-md-table-cell" style="font-size:0.65rem;letter-spacing:0.08em;">Subject</th>
            <th class="text-uppercase font-weight-bold text-medium-emphasis d-none d-sm-table-cell" style="font-size:0.65rem;letter-spacing:0.08em;">Due Date</th>
            <th class="text-uppercase font-weight-bold text-medium-emphasis" style="font-size:0.65rem;letter-spacing:0.08em;">Status</th>
            <th class="text-uppercase font-weight-bold text-medium-emphasis d-none d-sm-table-cell" style="font-size:0.65rem;letter-spacing:0.08em;">Submissions</th>
            <th class="text-uppercase font-weight-bold text-medium-emphasis d-none d-md-table-cell" style="font-size:0.65rem;letter-spacing:0.08em;">Avg Score</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="a in assignmentData" :key="a.name">
            <td>
              <div class="d-flex align-center gap-3">
                <v-avatar :color="a.iconColor" size="36" rounded="lg"><v-icon size="16" color="white">{{ a.icon }}</v-icon></v-avatar>
                <div>
                  <div class="text-body-2 font-weight-bold text-high-emphasis">{{ a.name }}</div>
                  <div class="text-caption text-medium-emphasis">{{ a.category }}</div>
                </div>
              </div>
            </td>
            <td class="text-body-2 text-medium-emphasis d-none d-md-table-cell">{{ a.subject }}</td>
            <td class="text-body-2 text-medium-emphasis d-none d-sm-table-cell">{{ a.due }}</td>
            <td><v-chip :color="a.statusColor" size="x-small" variant="tonal" class="font-weight-bold text-uppercase">{{ a.status }}</v-chip></td>
            <td class="d-none d-sm-table-cell">
              <div class="d-flex align-center gap-2">
                <span class="text-body-2 font-weight-bold text-high-emphasis">{{ a.submitted }}/{{ a.total }}</span>
                <v-progress-linear :model-value="a.submitted / a.total * 100" color="primary" bg-color="surface-variant" rounded height="5" style="width:60px;"></v-progress-linear>
              </div>
            </td>
            <td class="text-body-2 font-weight-bold text-high-emphasis d-none d-md-table-cell">{{ a.avg ?? '—' }}</td>
            <td>
              <v-menu>
                <template v-slot:activator="{ props }"><v-icon v-bind="props" class="text-medium-emphasis" size="18">mdi-dots-vertical</v-icon></template>
                <v-list density="compact"><v-list-item title="Edit"></v-list-item><v-list-item title="View Details"></v-list-item><v-list-item title="Delete"></v-list-item></v-list>
              </v-menu>
            </td>
          </tr>
        </tbody>
      </v-table>
    </v-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const activeTab = ref<string>('All')
const tabs: string[] = ['All', 'In Progress', 'Completed', 'Overdue']

interface AssignmentRow {
  name: string; category: string; subject: string; due: string
  status: string; statusColor: string; submitted: number; total: number
  avg: number | null; icon: string; iconColor: string
}

const assignmentData: AssignmentRow[] = [
  { name: 'Mid-Term Calculus Essay', category: 'Mathematics Tier A', subject: 'Advanced Calculus', due: 'Oct 24, 2023', status: 'In Progress', statusColor: 'primary', submitted: 88, total: 124, avg: 82.4, icon: 'mdi-file-document', iconColor: 'tertiary' },
  { name: 'Physics Lab Report: Optics', category: 'Natural Sciences', subject: 'Physics 101', due: 'Oct 26, 2023', status: 'Completed', statusColor: 'success', submitted: 112, total: 124, avg: 87.1, icon: 'mdi-flask', iconColor: 'primary' },
  { name: 'Digital Renaissance Portfolio', category: 'Arts & Design', subject: 'Digital Arts', due: 'Oct 30, 2023', status: 'Pending', statusColor: 'warning', submitted: 12, total: 124, avg: null, icon: 'mdi-palette', iconColor: 'purple' },
  { name: 'World War II Analysis', category: 'Historical Studies', subject: 'World History', due: 'Nov 2, 2023', status: 'In Progress', statusColor: 'primary', submitted: 65, total: 124, avg: 78.9, icon: 'mdi-book-open-variant', iconColor: 'success' },
  { name: 'Algorithm Design Challenge', category: 'Computer Science', subject: 'CS II', due: 'Nov 5, 2023', status: 'Overdue', statusColor: 'error', submitted: 98, total: 124, avg: 91.3, icon: 'mdi-code-braces', iconColor: 'teal-darken-1' },
]
</script>

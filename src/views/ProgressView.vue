<template>
  <div>
    <div class="d-flex flex-column flex-md-row justify-space-between align-start mb-7 gap-4">
      <div>
        <h1 class="text-h5 text-md-h4 font-weight-bold text-high-emphasis mb-1">Student Progress</h1>
        <p class="text-body-2 text-medium-emphasis" style="max-width:500px;">Detailed academic trajectory and AI-generated insights for your current cohort.</p>
      </div>
      <div class="d-flex gap-3">
        <v-btn variant="text" color="primary" class="text-none font-weight-bold">Export CSV</v-btn>
        <v-btn color="primary" elevation="0" rounded="lg" class="text-none font-weight-bold">Generate Full Report</v-btn>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4 mb-7">
      <v-card color="surface" variant="flat" rounded="xl" class="pa-5 border-thin">
        <div class="text-uppercase font-weight-bold text-medium-emphasis mb-2" style="font-size:0.65rem;letter-spacing:0.1em;">Average Score</div>
        <div class="text-h4 font-weight-bold text-high-emphasis">84.2 <v-chip color="success" size="x-small" variant="tonal" class="font-weight-bold ml-1">+2.4%</v-chip></div>
      </v-card>
      <v-card color="surface" variant="flat" rounded="xl" class="pa-5 border-thin">
        <div class="text-uppercase font-weight-bold text-medium-emphasis mb-2" style="font-size:0.65rem;letter-spacing:0.1em;">Mastery Rate</div>
        <div class="text-h4 font-weight-bold text-high-emphasis">68% <span class="text-body-2 text-medium-emphasis font-weight-medium">12 Students</span></div>
      </v-card>
      <v-card color="surface" variant="flat" rounded="xl" class="pa-5 border-thin">
        <div class="text-uppercase font-weight-bold text-medium-emphasis mb-2" style="font-size:0.65rem;letter-spacing:0.1em;">Active Assignments</div>
        <div class="text-h4 font-weight-bold text-high-emphasis">06 <span class="text-body-2 text-error font-weight-bold">2 Due Soon</span></div>
      </v-card>
      <v-card color="surface-variant" variant="flat" rounded="xl" class="pa-5">
        <div class="d-flex align-center gap-1 mb-2">
          <v-icon size="14" color="tertiary">mdi-sparkles</v-icon>
          <span class="text-uppercase font-weight-bold text-tertiary" style="font-size:0.65rem;letter-spacing:0.1em;">AI Alert</span>
        </div>
        <p class="text-body-2 ma-0 text-medium-emphasis" style="line-height:1.5;">Mathematics logic performance dropped 12% this week. Focus on Calculus basics.</p>
      </v-card>
    </div>

    <div class="d-flex flex-column flex-sm-row justify-space-between align-start align-sm-center mb-4 gap-3">
      <v-btn-toggle v-model="activeTab" mandatory density="compact" rounded="lg" color="primary" variant="outlined">
        <v-btn value="all" class="text-none text-body-2">All Students</v-btn>
        <v-btn value="struggling" class="text-none text-body-2">Struggling</v-btn>
        <v-btn value="excelling" class="text-none text-body-2">Excelling</v-btn>
      </v-btn-toggle>
      <div class="d-flex gap-2">
        <v-btn variant="outlined" size="small" class="text-none font-weight-medium text-medium-emphasis" rounded="lg"><v-icon start>mdi-filter-variant</v-icon>Filter</v-btn>
        <v-btn variant="outlined" size="small" class="text-none font-weight-medium text-medium-emphasis" rounded="lg"><v-icon start>mdi-sort</v-icon>Sort By</v-btn>
      </div>
    </div>

    <div class="d-flex flex-column gap-2 mb-5">
      <v-card v-for="s in gradeRows" :key="s.studentName" color="surface" variant="flat" rounded="xl" class="pa-4 cursor-pointer border-thin" hover @click="$router.push('/dashboard/students')">
        <div class="d-none d-md-grid grid-cols-7 items-center gap-3">
          <div class="d-flex align-center gap-3">
            <v-avatar color="surface-variant" size="36"><span class="text-caption font-weight-bold text-primary">{{ s.initials }}</span></v-avatar>
            <span class="text-body-2 font-weight-bold text-high-emphasis">{{ s.studentName }}</span>
          </div>
          <div class="text-body-2 text-medium-emphasis">{{ s.subject }}</div>
          <div class="text-body-2 text-medium-emphasis font-weight-medium">{{ s.week }}</div>
          <div class="text-body-2 text-medium-emphasis">{{ s.learningObjective }}</div>
          <div class="text-body-2 font-weight-bold text-high-emphasis">{{ s.score }}</div>
          <div><v-chip :color="statusColor(s.status)" size="x-small" variant="tonal" class="font-weight-bold text-uppercase">{{ s.status }}</v-chip></div>
          <div class="d-flex align-center gap-1 text-caption font-italic text-medium-emphasis">
            <v-icon size="14" color="tertiary">mdi-sparkles</v-icon>{{ s.aiRecommendation }}
          </div>
        </div>
        <div class="d-md-none">
          <div class="d-flex align-center gap-3 mb-3">
            <v-avatar color="surface-variant" size="36"><span class="text-caption font-weight-bold text-primary">{{ s.initials }}</span></v-avatar>
            <div class="flex-1">
              <div class="text-body-2 font-weight-bold text-high-emphasis">{{ s.studentName }}</div>
              <div class="text-caption text-medium-emphasis">{{ s.subject }} • {{ s.week }}</div>
            </div>
            <v-chip :color="statusColor(s.status)" size="x-small" variant="tonal" class="font-weight-bold text-uppercase">{{ s.status }}</v-chip>
          </div>
          <div class="d-flex justify-space-between align-center">
            <span class="text-caption text-medium-emphasis">{{ s.learningObjective }}</span>
            <span class="text-body-2 font-weight-bold text-high-emphasis">{{ s.score }}</span>
          </div>
        </div>
      </v-card>
    </div>

    <div class="d-flex flex-column flex-sm-row justify-space-between align-center mb-8 gap-3">
      <span class="text-body-2 text-disabled">Showing 1 to 4 of 28 students</span>
      <v-pagination v-model="page" :length="3" size="small" density="compact" rounded="lg" color="primary"></v-pagination>
    </div>

    <v-card color="surface" variant="flat" rounded="xl" class="pa-5 pa-md-6 border-thin">
      <div class="d-flex flex-column flex-sm-row gap-4 gap-sm-5">
        <v-avatar color="tertiary" size="56" rounded="xl" class="flex-shrink-0"><v-icon color="white" size="28">mdi-sparkles</v-icon></v-avatar>
        <div class="flex-1">
          <h3 class="text-subtitle-1 font-weight-bold text-high-emphasis mb-2">Automated Learning Path Adjustment</h3>
          <p class="text-body-2 mb-4 text-medium-emphasis" style="line-height:1.6;">Based on this week's progress, 14% of the class is struggling with <strong class="text-high-emphasis">Abstract Concept Visualization</strong>. The AI recommends integrating interactive geometric models into next Monday's session.</p>
          <div class="d-flex flex-column flex-sm-row gap-3">
            <v-btn color="primary" elevation="0" rounded="lg" class="text-none font-weight-bold">Apply Recommendations</v-btn>
            <v-btn variant="outlined" rounded="lg" class="text-none font-weight-bold text-medium-emphasis">Review Detailed Analytics</v-btn>
          </div>
        </div>
      </div>
    </v-card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { fetchGrades } from '@/api/grades'
import type { Grade, MasteryStatus } from '@/types'

const activeTab = ref<string>('all')
const page = ref<number>(1)
const gradeRows = ref<Grade[]>([])

onMounted(async () => {
  gradeRows.value = await fetchGrades()
})

const statusColor = (status: MasteryStatus): string => {
  const map: Record<MasteryStatus, string> = {
    menguasai: 'primary',
    cukup: 'warning',
    belum: 'error',
  }
  return map[status] || 'grey'
}
</script>

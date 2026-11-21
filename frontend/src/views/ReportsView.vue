<template>
  <div>
    <div class="d-flex flex-column flex-sm-row justify-space-between align-start mb-7 gap-3">
      <div>
        <h1 class="text-h5 text-md-h4 font-weight-bold text-high-emphasis">Reports</h1>
        <p class="text-body-2 mt-1 text-medium-emphasis">Generate and download academic performance reports.</p>
      </div>
      <v-btn color="primary" elevation="0" prepend-icon="mdi-file-chart" rounded="lg" class="text-none font-weight-bold">Generate Report</v-btn>
    </div>

    <div class="grid grid-cols-2 gap-4 gap-md-5 mb-8">
      <v-card v-for="r in reports" :key="r.title" color="surface" variant="flat" rounded="xl" class="pa-6 border-thin" hover>
        <v-avatar :color="r.bg" size="48" rounded="lg" class="mb-4"><v-icon :color="r.color" size="24">{{ r.icon }}</v-icon></v-avatar>
        <h3 class="text-body-1 font-weight-bold text-high-emphasis mb-1">{{ r.title }}</h3>
        <p class="text-body-2 mb-4 text-medium-emphasis" style="line-height:1.5;">{{ r.desc }}</p>
        <div class="d-flex flex-column flex-sm-row justify-space-between align-start align-sm-center gap-2">
          <span class="text-caption text-disabled">Last generated: {{ r.date }}</span>
          <v-btn variant="text" color="primary" size="small" class="text-none font-weight-bold" append-icon="mdi-download">Download</v-btn>
        </div>
      </v-card>
    </div>

    <v-card color="surface" variant="flat" rounded="xl" class="pa-5 pa-md-6 mb-6 border-thin">
      <div class="d-flex justify-space-between align-center mb-5">
        <h3 class="text-subtitle-1 font-weight-bold text-high-emphasis">Recent Activity</h3>
        <span class="text-uppercase font-weight-bold text-medium-emphasis cursor-pointer" style="font-size:0.65rem;letter-spacing:0.1em;">View Full System Log</span>
      </div>
      <v-timeline side="end" density="compact" truncate-line="both">
        <v-timeline-item v-for="a in activities" :key="a.time" :dot-color="a.dotColor" size="x-small">
          <div>
            <div class="text-caption text-disabled">{{ a.time }}</div>
            <div class="text-body-2 text-high-emphasis" v-html="a.text"></div>
            <div class="text-caption text-medium-emphasis" v-if="a.sub">{{ a.sub }}</div>
          </div>
        </v-timeline-item>
      </v-timeline>
    </v-card>

    <v-card rounded="xl" color="surface-variant" class="pa-5 pa-md-6">
      <div class="d-flex flex-column flex-sm-row align-center gap-4 gap-sm-5">
        <v-avatar color="surface" size="56" rounded="xl" class="flex-shrink-0"><v-icon color="primary" size="28">mdi-sparkles</v-icon></v-avatar>
        <div class="flex-1">
          <h3 class="text-subtitle-1 font-weight-bold text-high-emphasis mb-1">Predictive GPA Analysis</h3>
          <p class="text-body-2 text-medium-emphasis ma-0" style="line-height:1.5;">Our AI models suggest a 4.2% increase in graduation rates based on current mid-term grade trajectories.</p>
        </div>
        <v-btn color="error" elevation="0" size="large" rounded="lg" class="text-none font-weight-bold flex-shrink-0">Run Forecast Report</v-btn>
      </div>
    </v-card>
  </div>
</template>

<script setup lang="ts">
interface ReportCard { title: string; desc: string; date: string; icon: string; bg: string; color: string }
interface Activity { time: string; text: string; sub?: string; dotColor: string }

const reports: ReportCard[] = [
  { title: 'Semester Performance', desc: 'Overall academic performance summary for the current semester.', date: 'Oct 20, 2023', icon: 'mdi-chart-arc', bg: 'blue-lighten-4', color: 'primary' },
  { title: 'Student Progress Report', desc: 'Individual student learning trajectory across all subjects.', date: 'Oct 18, 2023', icon: 'mdi-account-check', bg: 'green-lighten-4', color: 'success' },
  { title: 'Subject Analytics', desc: 'Deep dive into per-subject scores and mastery rates.', date: 'Oct 15, 2023', icon: 'mdi-book-search', bg: 'purple-lighten-4', color: 'purple' },
  { title: 'Faculty Performance', desc: 'Teacher performance metrics and completion rates.', date: 'Oct 12, 2023', icon: 'mdi-human-male-board', bg: 'orange-lighten-4', color: 'tertiary' },
]

const activities: Activity[] = [
  { time: '2 mins ago', text: 'Grade override approved for <strong>St. ID #829</strong>', sub: 'Actioned by Admin', dotColor: 'grey' },
  { time: '45 mins ago', text: 'Faculty report generated: <em>Q3 Academic Performance</em>', dotColor: 'grey' },
  { time: '2 hours ago', text: 'New faculty account created: <strong class="text-primary">Prof. Liam Vance</strong>', dotColor: 'error' },
  { time: 'Yesterday', text: 'System backup completed successfully', sub: 'Cloud Storage: 4.2TB / 10TB', dotColor: 'grey' },
]
</script>

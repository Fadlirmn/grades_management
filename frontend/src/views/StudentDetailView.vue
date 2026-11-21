<template>
  <div>
    <div class="d-flex align-center rounded-pill px-4 py-2 mb-6 bg-surface-variant" style="max-width:380px;">
      <v-icon size="18" class="text-medium-emphasis">mdi-magnify</v-icon>
      <input type="text" placeholder="Search student records..." class="ml-2 text-body-2 text-high-emphasis" style="outline:none;border:none;background:transparent;width:100%;">
    </div>

    <div class="grid grid-cols-3 gap-4 gap-md-5 mb-8">
      <!-- Profile Card -->
      <v-card color="surface" variant="flat" rounded="xl" class="pa-5 pa-md-6 col-span-2 border-thin">
        <div class="d-flex flex-column flex-sm-row gap-5">
          <div class="position-relative flex-shrink-0">
            <v-avatar size="100" rounded="lg">
              <v-img src="https://ui-avatars.com/api/?name=Alexander+Chen&background=e5e7eb&color=374151&size=200&bold=true&font-size=0.4"></v-img>
            </v-avatar>
          </div>
          <div class="flex-1">
            <div class="d-flex align-center gap-3 mb-1 flex-wrap">
              <h1 class="text-h5 font-weight-bold text-high-emphasis">Alexander Chen</h1>
              <v-chip color="primary" size="x-small" variant="tonal" class="font-weight-bold">GRADE 11-A</v-chip>
            </div>
            <p class="text-body-2 mb-4 d-flex align-center gap-1 text-medium-emphasis">
              <v-icon size="14">mdi-card-account-details-outline</v-icon> STU-2024-8812 • Computer Science Major
            </p>
            <div class="d-flex flex-wrap gap-3">
              <v-card color="surface-variant" variant="flat" rounded="xl" class="pa-3 flex-1" style="min-width:90px;">
                <div class="text-uppercase font-weight-bold text-disabled" style="font-size:0.6rem;letter-spacing:0.1em;">GPA</div>
                <div class="text-h5 font-weight-bold text-primary">3.92</div>
              </v-card>
              <v-card color="surface-variant" variant="flat" rounded="xl" class="pa-3 flex-1" style="min-width:90px;">
                <div class="text-uppercase font-weight-bold text-disabled" style="font-size:0.6rem;letter-spacing:0.1em;">Attendance</div>
                <div class="text-h5 font-weight-bold text-high-emphasis">98%</div>
              </v-card>
              <v-card color="surface-variant" variant="flat" rounded="xl" class="pa-3 flex-1" style="min-width:90px;">
                <div class="text-uppercase font-weight-bold text-disabled" style="font-size:0.6rem;letter-spacing:0.1em;">Ranking</div>
                <div class="text-h5 font-weight-bold text-high-emphasis">#4 <span class="text-body-2 text-medium-emphasis font-weight-medium">/ 120</span></div>
              </v-card>
            </div>
          </div>
        </div>
      </v-card>

      <!-- Status Summary -->
      <v-card rounded="xl" class="pa-6 d-flex flex-column justify-space-between" color="primary" theme="dark">
        <div>
          <div class="text-uppercase font-weight-bold mb-2" style="font-size:0.6rem;letter-spacing:0.12em;opacity:0.6;">Status Summary</div>
          <h2 class="text-h5 font-weight-bold mb-3">Excellent Academic Standing</h2>
          <p class="text-body-2 mb-4" style="opacity:0.8;line-height:1.5;">Alexander is currently performing in the top 5% of his cohort with consistent growth in STEM subjects.</p>
        </div>
        <v-btn color="white" variant="flat" rounded="xl" class="text-primary text-none font-weight-bold" block>Download Semester Report</v-btn>
      </v-card>

      <!-- Weekly Progress -->
      <v-card color="surface" variant="flat" rounded="xl" class="pa-5 pa-md-6 border-thin">
        <div class="d-flex justify-space-between align-center mb-5">
          <h3 class="text-subtitle-1 font-weight-bold text-high-emphasis">Weekly Progress</h3>
          <v-icon class="text-medium-emphasis">mdi-dots-horizontal</v-icon>
        </div>
        <v-timeline side="end" density="compact" truncate-line="both">
          <v-timeline-item v-for="w in weeks" :key="w.week" :dot-color="w.dotColor" size="x-small">
            <div>
              <div class="text-uppercase font-weight-bold text-disabled" style="font-size:0.6rem;letter-spacing:0.1em;">{{ w.weekLabel }}</div>
              <div class="text-body-2 font-weight-bold text-high-emphasis">{{ w.title }}</div>
              <div class="text-caption text-medium-emphasis">{{ w.desc }}</div>
            </div>
          </v-timeline-item>
        </v-timeline>
        <v-card color="surface-variant" variant="flat" rounded="lg" class="pa-3 mt-4">
          <div class="text-uppercase font-weight-bold text-disabled" style="font-size:0.6rem;letter-spacing:0.1em;">Current Semester</div>
          <div class="text-body-1 font-weight-bold text-high-emphasis">Fall 2024</div>
        </v-card>
      </v-card>

      <!-- Mastery Breakdown -->
      <v-card color="surface" variant="flat" rounded="xl" class="pa-5 pa-md-6 col-span-2 border-thin">
        <h3 class="text-subtitle-1 font-weight-bold text-high-emphasis mb-5">Objective Mastery Breakdown</h3>
        <div class="grid grid-cols-2 gap-4 mb-6" style="column-gap:32px;">
          <div v-for="m in masteryItems" :key="m.name">
            <div class="d-flex justify-space-between mb-1">
              <span class="text-body-2 font-weight-bold text-high-emphasis">{{ m.name }}</span>
              <span class="text-body-2 font-weight-bold text-high-emphasis">{{ m.percentage }}%</span>
            </div>
            <v-progress-linear :model-value="m.percentage" :color="m.color" bg-color="surface-variant" rounded height="6"></v-progress-linear>
          </div>
        </div>

        <v-card color="surface-variant" variant="flat" rounded="xl" class="pa-5">
          <div class="d-flex align-center gap-2 mb-4">
            <v-avatar color="surface" size="28" rounded="lg"><v-icon size="16" color="tertiary">mdi-sparkles</v-icon></v-avatar>
            <h4 class="text-body-1 font-weight-bold text-high-emphasis">AI Learning Synthesis</h4>
          </div>
          <div class="grid grid-cols-2 gap-3 mb-3">
            <v-card color="surface" variant="flat" rounded="lg" class="pa-4 border-thin">
              <h5 class="text-body-2 font-weight-bold text-high-emphasis mb-2 d-flex align-center gap-1"><v-icon size="16" color="primary">mdi-trending-up</v-icon> Focus on Humanities</h5>
              <p class="text-caption ma-0 text-medium-emphasis" style="line-height:1.5;">Alexander's STEM performance is exceptional. Suggest integrating more argumentative essays in History to bolster Verbal Expression scores.</p>
            </v-card>
            <v-card color="surface" variant="flat" rounded="lg" class="pa-4 border-thin">
              <h5 class="text-body-2 font-weight-bold text-high-emphasis mb-2 d-flex align-center gap-1"><v-icon size="16" color="tertiary">mdi-lightbulb-on</v-icon> Extended Learning</h5>
              <p class="text-caption ma-0 text-medium-emphasis" style="line-height:1.5;">Consider nominating for the State Math Olympiad. Current mastery levels suggest he is ready for advanced competitive challenges.</p>
            </v-card>
          </div>
          <div class="text-right">
            <a href="#" class="text-body-2 font-weight-bold text-primary text-decoration-none">View Full Analysis Plan →</a>
          </div>
        </v-card>
      </v-card>
    </div>

    <v-card color="surface" variant="flat" rounded="xl" class="pa-5 pa-md-6 border-thin">
      <h3 class="text-subtitle-1 font-weight-bold text-high-emphasis mb-5">Enrolled Subjects & Performance</h3>
      <v-table>
        <thead>
          <tr>
            <th class="text-uppercase font-weight-bold text-disabled" style="font-size:0.6rem;letter-spacing:0.1em;">Subject</th>
            <th class="text-uppercase font-weight-bold text-disabled d-none d-sm-table-cell" style="font-size:0.6rem;letter-spacing:0.1em;">Instructor</th>
            <th class="text-uppercase font-weight-bold text-disabled" style="font-size:0.6rem;letter-spacing:0.1em;">Grade</th>
            <th class="text-uppercase font-weight-bold text-disabled d-none d-sm-table-cell" style="font-size:0.6rem;letter-spacing:0.1em;">Credits</th>
            <th class="text-uppercase font-weight-bold text-disabled" style="font-size:0.6rem;letter-spacing:0.1em;">Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="s in subjects" :key="s.name">
            <td>
              <div class="d-flex align-center gap-3">
                <v-avatar :color="s.iconBg" size="32" rounded="lg"><v-icon size="16" :color="s.iconColor">{{ s.icon }}</v-icon></v-avatar>
                <span class="text-body-2 font-weight-bold text-high-emphasis">{{ s.name }}</span>
              </div>
            </td>
            <td class="text-body-2 text-medium-emphasis d-none d-sm-table-cell">{{ s.instructor }}</td>
            <td class="text-body-2 font-weight-bold text-primary">{{ s.grade }}</td>
            <td class="text-body-2 text-medium-emphasis d-none d-sm-table-cell">{{ s.credits }}</td>
            <td><v-chip :color="s.statusColor" size="x-small" variant="tonal" class="font-weight-bold text-uppercase">{{ s.status }}</v-chip></td>
          </tr>
        </tbody>
      </v-table>
    </v-card>
  </div>
</template>

<script setup lang="ts">
import type { EnrolledSubject, MasteryBreakdown, WeeklyProgress } from '@/types'

const weeks = [
  { week: 12, weekLabel: 'WEEK 12 (CURRENT)', title: 'Advanced Calculus Midterm', desc: 'Score: 94% • High Distinction', dotColor: 'success' },
  { week: 11, weekLabel: 'WEEK 11', title: 'Data Structures Project', desc: 'Submitted 2 days early • Pending Review', dotColor: 'primary' },
  { week: 10, weekLabel: 'WEEK 10', title: 'Physics Lab Report', desc: 'Score: 88% • Distinction', dotColor: 'primary' },
  { week: 9, weekLabel: 'WEEK 09', title: 'World History Essay', desc: 'Score: 82% • Merit', dotColor: 'grey' },
]

const masteryItems: MasteryBreakdown[] = [
  { name: 'Mathematical Analysis', percentage: 96, color: 'primary' },
  { name: 'Verbal Expression', percentage: 74, color: 'tertiary' },
  { name: 'Logic & Algorithmics', percentage: 92, color: 'primary' },
  { name: 'Critical Research', percentage: 79, color: 'brown-darken-2' },
  { name: 'Scientific Inquiry', percentage: 88, color: 'primary' },
  { name: 'Creative Arts', percentage: 65, color: 'error' },
]

const subjects = [
  { name: 'Advanced Calculus', instructor: 'Dr. Sarah Miller', grade: 'A+', credits: '4.0', status: 'MENGUASAI', statusColor: 'grey', icon: 'mdi-function-variant', iconBg: 'blue-lighten-4', iconColor: 'primary' },
  { name: 'Computer Science II', instructor: 'Prof. Alan Turing', grade: 'A', credits: '4.0', status: 'MENGUASAI', statusColor: 'grey', icon: 'mdi-xml', iconBg: 'orange-lighten-4', iconColor: 'tertiary' },
  { name: 'Modern World History', instructor: 'Mrs. Elena Rossi', grade: 'B+', credits: '3.0', status: 'CUKUP', statusColor: 'warning', icon: 'mdi-earth', iconBg: 'purple-lighten-4', iconColor: 'purple' },
]
</script>

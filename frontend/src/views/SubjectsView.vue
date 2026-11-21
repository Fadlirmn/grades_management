<template>
  <div>
    <div class="d-flex flex-column flex-sm-row justify-space-between align-start mb-7 gap-3">
      <div>
        <h1 class="text-h5 text-md-h4 font-weight-bold text-high-emphasis">Subjects</h1>
        <p class="text-body-2 mt-1 text-medium-emphasis">Manage curriculum subjects and their instructors.</p>
      </div>
      <v-btn color="primary" elevation="0" prepend-icon="mdi-plus" rounded="lg" class="text-none font-weight-bold">Add Subject</v-btn>
    </div>

    <div class="grid grid-cols-3 gap-4 gap-md-5">
      <v-card v-for="s in subjects" :key="s.name" color="surface" variant="flat" rounded="xl" class="pa-6 border-thin" hover>
        <div class="d-flex justify-space-between align-center mb-4">
          <v-avatar :color="s.bg" size="44" rounded="lg"><v-icon :color="s.color" size="22">{{ s.icon }}</v-icon></v-avatar>
          <v-menu>
            <template v-slot:activator="{ props }"><v-icon v-bind="props" class="text-medium-emphasis">mdi-dots-vertical</v-icon></template>
            <v-list density="compact"><v-list-item title="Edit"></v-list-item><v-list-item title="Delete"></v-list-item></v-list>
          </v-menu>
        </div>
        <h3 class="text-body-1 font-weight-bold text-high-emphasis mb-1">{{ s.name }}</h3>
        <p class="text-body-2 mb-4 text-medium-emphasis">{{ s.instructor }}</p>
        <div class="d-flex gap-4 mb-4">
          <div><div class="text-body-1 font-weight-bold text-high-emphasis">{{ s.students }}</div><div class="text-caption text-medium-emphasis">Students</div></div>
          <div><div class="text-body-1 font-weight-bold text-high-emphasis">{{ s.avg }}</div><div class="text-caption text-medium-emphasis">Avg Score</div></div>
          <div><div class="text-body-1 font-weight-bold text-high-emphasis">{{ s.mastery }}%</div><div class="text-caption text-medium-emphasis">Mastery</div></div>
        </div>
        <v-progress-linear :model-value="s.mastery" color="primary" bg-color="surface-variant" rounded height="6"></v-progress-linear>
      </v-card>
    </div>
  </div>
</template>

<script setup lang="ts">
interface SubjectCard {
  name: string; instructor: string; students: number; avg: number; mastery: number
  icon: string; bg: string; color: string
}

const subjects: SubjectCard[] = [
  { name: 'Advanced Calculus', instructor: 'Dr. Sarah Miller', students: 32, avg: 88.4, mastery: 72, icon: 'mdi-function-variant', bg: 'blue-lighten-4', color: 'primary' },
  { name: 'Computer Science II', instructor: 'Prof. Alan Turing', students: 28, avg: 91.2, mastery: 84, icon: 'mdi-xml', bg: 'orange-lighten-4', color: 'tertiary' },
  { name: 'English Literature', instructor: 'Sarah Whitmore', students: 35, avg: 76.8, mastery: 58, icon: 'mdi-book-open-variant', bg: 'green-lighten-4', color: 'success' },
  { name: 'World History', instructor: 'Mrs. Elena Rossi', students: 30, avg: 82.1, mastery: 65, icon: 'mdi-earth', bg: 'purple-lighten-4', color: 'purple' },
  { name: 'Physics 101', instructor: 'Dr. James Newton', students: 26, avg: 79.5, mastery: 55, icon: 'mdi-atom', bg: 'red-lighten-4', color: 'error' },
  { name: 'Digital Arts', instructor: 'Ms. Clara Vinci', students: 22, avg: 93.7, mastery: 90, icon: 'mdi-palette', bg: 'pink-lighten-4', color: 'pink' },
]
</script>

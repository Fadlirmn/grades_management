<template>
  <div>
    <!-- Header -->
    <div class="d-flex flex-column flex-sm-row justify-space-between align-start mb-6 gap-3">
      <div>
        <h1 class="text-h5 text-md-h4 font-weight-bold text-high-emphasis">Data Management</h1>
        <p class="text-body-2 mt-1 text-medium-emphasis">Manage all database records. Select a table to view, create, edit or delete entries.</p>
      </div>
      <v-btn color="primary" elevation="0" prepend-icon="mdi-plus" rounded="lg" class="text-none font-weight-bold" @click="openCreateDialog">
        Add {{ currentTableLabel }}
      </v-btn>
    </div>

    <!-- Table Tabs -->
    <v-card color="surface" variant="flat" rounded="xl" class="border-thin mb-6">
      <v-tabs v-model="activeTable" color="primary" slider-color="primary" density="comfortable" show-arrows>
        <v-tab v-for="t in tables" :key="t.key" :value="t.key" class="text-none text-body-2 font-weight-bold">
          <v-icon start size="18">{{ t.icon }}</v-icon>
          {{ t.label }}
        </v-tab>
      </v-tabs>
    </v-card>

    <!-- Search & Filters -->
    <div class="d-flex flex-column flex-sm-row gap-3 mb-4">
      <div class="d-flex align-center rounded-xl px-4 py-2 bg-surface-variant flex-1" style="max-width:400px;">
        <v-icon size="18" class="text-medium-emphasis">mdi-magnify</v-icon>
        <input v-model="searchQuery" type="text" :placeholder="`Search ${currentTableLabel}...`" class="ml-2 text-body-2 text-high-emphasis" style="outline:none;border:none;background:transparent;width:100%;">
      </div>
      <v-chip variant="tonal" color="primary" class="font-weight-bold">
        {{ filteredRows.length }} records
      </v-chip>
    </div>

    <!-- Data Table -->
    <v-card color="surface" variant="flat" rounded="xl" class="border-thin">
      <v-table density="comfortable" hover>
        <thead>
          <tr>
            <th v-for="col in currentColumns" :key="col.key" class="text-uppercase font-weight-bold text-medium-emphasis" style="font-size:0.65rem;letter-spacing:0.08em;">
              {{ col.label }}
            </th>
            <th class="text-uppercase font-weight-bold text-medium-emphasis text-right" style="font-size:0.65rem;letter-spacing:0.08em;">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-if="isLoading">
            <td :colspan="currentColumns.length + 1" class="text-center py-8">
              <v-progress-circular indeterminate color="primary" size="32"></v-progress-circular>
              <div class="text-body-2 text-medium-emphasis mt-2">Loading data...</div>
            </td>
          </tr>
          <tr v-else-if="filteredRows.length === 0">
            <td :colspan="currentColumns.length + 1" class="text-center py-8">
              <v-icon size="48" class="text-disabled mb-2">mdi-database-off-outline</v-icon>
              <div class="text-body-2 text-medium-emphasis">No records found</div>
            </td>
          </tr>
          <tr v-for="row in filteredRows" :key="(row as any).id" class="cursor-pointer">
            <td v-for="col in currentColumns" :key="col.key" class="text-body-2" :class="col.key === 'id' ? 'text-medium-emphasis font-weight-bold' : 'text-high-emphasis'">
              <template v-if="col.key === 'role'">
                <v-chip :color="roleColor((row as any)[col.key])" size="x-small" variant="tonal" class="font-weight-bold text-capitalize">
                  {{ (row as any)[col.key] }}
                </v-chip>
              </template>
              <template v-else-if="col.key === 'status'">
                <v-chip :color="statusColor((row as any)[col.key])" size="x-small" variant="tonal" class="font-weight-bold text-uppercase">
                  {{ (row as any)[col.key] }}
                </v-chip>
              </template>
              <template v-else-if="col.key === 'password'">
                <span class="text-disabled">••••••••</span>
              </template>
              <template v-else>
                {{ truncate((row as any)[col.key]) }}
              </template>
            </td>
            <td class="text-right">
              <v-btn icon variant="text" size="small" color="primary" @click.stop="openEditDialog(row)">
                <v-icon size="18">mdi-pencil</v-icon>
              </v-btn>
              <v-btn icon variant="text" size="small" color="error" @click.stop="confirmDelete(row)">
                <v-icon size="18">mdi-delete</v-icon>
              </v-btn>
            </td>
          </tr>
        </tbody>
      </v-table>
    </v-card>

    <!-- Create/Edit Dialog -->
    <v-dialog v-model="formDialog" max-width="600" persistent>
      <v-card color="surface" rounded="xl" class="pa-6">
        <div class="d-flex justify-space-between align-center mb-5">
          <h3 class="text-h6 font-weight-bold text-high-emphasis">
            {{ isEditing ? 'Edit' : 'Create' }} {{ currentTableLabel }}
          </h3>
          <v-btn icon variant="text" size="small" @click="formDialog = false">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </div>

        <v-form @submit.prevent="handleSubmit">
          <div class="d-flex flex-column gap-4">
            <div v-for="col in editableColumns" :key="col.key">
              <label class="text-body-2 font-weight-bold text-high-emphasis d-block mb-2">{{ col.label }}</label>
              <v-select
                v-if="col.key === 'role'"
                v-model="formData[col.key]"
                :items="['admin', 'teacher', 'student', 'parent']"
                variant="solo"
                density="comfortable"
                bg-color="surface-variant"
                flat
                hide-details
                rounded="lg"
              ></v-select>
              <v-select
                v-else-if="col.key === 'status'"
                v-model="formData[col.key]"
                :items="['menguasai', 'cukup', 'belum']"
                variant="solo"
                density="comfortable"
                bg-color="surface-variant"
                flat
                hide-details
                rounded="lg"
              ></v-select>
              <v-text-field
                v-else-if="col.type === 'number'"
                v-model.number="formData[col.key]"
                type="number"
                variant="solo"
                density="comfortable"
                bg-color="surface-variant"
                flat
                hide-details
                rounded="lg"
              ></v-text-field>
              <v-textarea
                v-else-if="col.type === 'textarea'"
                v-model="formData[col.key]"
                variant="solo"
                density="comfortable"
                bg-color="surface-variant"
                flat
                hide-details
                rounded="lg"
                rows="3"
              ></v-textarea>
              <v-text-field
                v-else
                v-model="formData[col.key]"
                :type="col.key === 'password' ? 'password' : col.type === 'date' ? 'date' : 'text'"
                variant="solo"
                density="comfortable"
                bg-color="surface-variant"
                flat
                hide-details
                rounded="lg"
              ></v-text-field>
            </div>
          </div>

          <div class="d-flex justify-end gap-3 mt-6">
            <v-btn variant="outlined" rounded="lg" class="text-none font-weight-bold text-medium-emphasis" @click="formDialog = false">Cancel</v-btn>
            <v-btn type="submit" color="primary" elevation="0" rounded="lg" class="text-none font-weight-bold" :loading="isSaving">
              {{ isEditing ? 'Update' : 'Create' }}
            </v-btn>
          </div>
        </v-form>
      </v-card>
    </v-dialog>

    <!-- Delete Confirm Dialog -->
    <v-dialog v-model="deleteDialog" max-width="420">
      <v-card color="surface" rounded="xl" class="pa-6 text-center">
        <v-avatar color="error" size="56" class="mb-4"><v-icon color="white" size="28">mdi-delete-alert</v-icon></v-avatar>
        <h3 class="text-h6 font-weight-bold text-high-emphasis mb-2">Delete Record</h3>
        <p class="text-body-2 text-medium-emphasis mb-6">Are you sure you want to delete this record? This action cannot be undone.</p>
        <div class="d-flex justify-center gap-3">
          <v-btn variant="outlined" rounded="lg" class="text-none font-weight-bold text-medium-emphasis" @click="deleteDialog = false">Cancel</v-btn>
          <v-btn color="error" elevation="0" rounded="lg" class="text-none font-weight-bold" :loading="isDeleting" @click="handleDelete">Delete</v-btn>
        </div>
      </v-card>
    </v-dialog>

    <!-- Snackbar -->
    <v-snackbar v-model="snackbar" :color="snackbarColor" rounded="lg" timeout="3000">
      {{ snackbarText }}
    </v-snackbar>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue'
import {
  usersApi, studentsApi, subjectsApi, assignmentsApi,
  objectivesApi, progressApi, scoresApi,
} from '@/api/admin'

// ─── Table Config ────────────────────────────────────────
interface ColumnDef {
  key: string
  label: string
  type?: 'text' | 'number' | 'date' | 'textarea'
  editable?: boolean
  hidden?: boolean // hidden from table but shown in form
}

interface TableDef {
  key: string
  label: string
  icon: string
  columns: ColumnDef[]
  api: {
    getAll: () => Promise<any[]>
    create: (p: any) => Promise<any>
    update: (id: number, p: any) => Promise<any>
    delete: (id: number) => Promise<void>
  }
}

const tables: TableDef[] = [
  {
    key: 'users', label: 'Users', icon: 'mdi-account-group',
    columns: [
      { key: 'id', label: 'ID' },
      { key: 'name', label: 'Name', editable: true },
      { key: 'email', label: 'Email', editable: true },
      { key: 'password', label: 'Password', editable: true, hidden: true },
      { key: 'role', label: 'Role', editable: true },
      { key: 'created_at', label: 'Created' },
    ],
    api: usersApi,
  },
  {
    key: 'students', label: 'Students', icon: 'mdi-school',
    columns: [
      { key: 'id', label: 'ID' },
      { key: 'user_id', label: 'User ID', type: 'number', editable: true },
      { key: 'nis', label: 'NIS', editable: true },
      { key: 'class', label: 'Class', editable: true },
      { key: 'parent_name', label: 'Parent Name', editable: true },
      { key: 'created_at', label: 'Created' },
    ],
    api: studentsApi,
  },
  {
    key: 'subjects', label: 'Subjects', icon: 'mdi-book-open-variant',
    columns: [
      { key: 'id', label: 'ID' },
      { key: 'name', label: 'Name', editable: true },
      { key: 'code', label: 'Code', editable: true },
      { key: 'teacher_id', label: 'Teacher ID', type: 'number', editable: true },
      { key: 'description', label: 'Description', type: 'textarea', editable: true },
      { key: 'created_at', label: 'Created' },
    ],
    api: subjectsApi,
  },
  {
    key: 'assignments', label: 'Assignments', icon: 'mdi-clipboard-text-outline',
    columns: [
      { key: 'id', label: 'ID' },
      { key: 'subject_id', label: 'Subject ID', type: 'number', editable: true },
      { key: 'title', label: 'Title', editable: true },
      { key: 'description', label: 'Description', type: 'textarea', editable: true },
      { key: 'due_date', label: 'Due Date', type: 'date', editable: true },
      { key: 'created_at', label: 'Created' },
    ],
    api: assignmentsApi,
  },
  {
    key: 'objectives', label: 'Objectives', icon: 'mdi-target',
    columns: [
      { key: 'id', label: 'ID' },
      { key: 'subject_id', label: 'Subject ID', type: 'number', editable: true },
      { key: 'week', label: 'Week', type: 'number', editable: true },
      { key: 'description', label: 'Description', type: 'textarea', editable: true },
      { key: 'created_at', label: 'Created' },
    ],
    api: objectivesApi,
  },
  {
    key: 'progress', label: 'Progress', icon: 'mdi-chart-line-variant',
    columns: [
      { key: 'id', label: 'ID' },
      { key: 'student_id', label: 'Student ID', type: 'number', editable: true },
      { key: 'objective_id', label: 'Objective ID', type: 'number', editable: true },
      { key: 'status', label: 'Status', editable: true },
      { key: 'notes', label: 'Notes', type: 'textarea', editable: true },
      { key: 'created_at', label: 'Created' },
    ],
    api: progressApi,
  },
  {
    key: 'scores', label: 'Scores', icon: 'mdi-counter',
    columns: [
      { key: 'id', label: 'ID' },
      { key: 'student_id', label: 'Student ID', type: 'number', editable: true },
      { key: 'assignment_id', label: 'Assignment ID', type: 'number', editable: true },
      { key: 'score', label: 'Score', type: 'number', editable: true },
      { key: 'feedback', label: 'Feedback', type: 'textarea', editable: true },
      { key: 'created_at', label: 'Created' },
    ],
    api: scoresApi,
  },
]

// ─── State ───────────────────────────────────────────────
const activeTable = ref<string>('users')
const searchQuery = ref('')
const rows = ref<any[]>([])
const isLoading = ref(false)
const isSaving = ref(false)
const isDeleting = ref(false)

const formDialog = ref(false)
const deleteDialog = ref(false)
const isEditing = ref(false)
const formData = ref<Record<string, any>>({})
const deleteTarget = ref<any>(null)

const snackbar = ref(false)
const snackbarText = ref('')
const snackbarColor = ref('success')

// ─── Computed ────────────────────────────────────────────
const currentTable = computed(() => tables.find(t => t.key === activeTable.value)!)
const currentTableLabel = computed(() => currentTable.value.label)
const currentColumns = computed(() => currentTable.value.columns.filter(c => !c.hidden))
const editableColumns = computed(() => currentTable.value.columns.filter(c => c.editable))

const filteredRows = computed(() => {
  if (!searchQuery.value) return rows.value
  const q = searchQuery.value.toLowerCase()
  return rows.value.filter(row =>
    Object.values(row).some(v =>
      String(v ?? '').toLowerCase().includes(q)
    )
  )
})

// ─── Data Loading ────────────────────────────────────────
const loadData = async (): Promise<void> => {
  isLoading.value = true
  rows.value = []
  try {
    rows.value = await currentTable.value.api.getAll()
  } catch {
    showSnack('Failed to load data', 'error')
  } finally {
    isLoading.value = false
  }
}

watch(activeTable, () => {
  searchQuery.value = ''
  loadData()
})

onMounted(loadData)

// ─── CRUD Actions ────────────────────────────────────────
const openCreateDialog = (): void => {
  isEditing.value = false
  formData.value = {}
  editableColumns.value.forEach(col => {
    formData.value[col.key] = col.type === 'number' ? 0 : ''
  })
  formDialog.value = true
}

const openEditDialog = (row: any): void => {
  isEditing.value = true
  formData.value = { ...row }
  formDialog.value = true
}

const handleSubmit = async (): Promise<void> => {
  isSaving.value = true
  try {
    if (isEditing.value) {
      await currentTable.value.api.update(formData.value.id, formData.value)
      showSnack(`${currentTableLabel.value} updated successfully`, 'success')
    } else {
      await currentTable.value.api.create(formData.value)
      showSnack(`${currentTableLabel.value} created successfully`, 'success')
    }
    formDialog.value = false
    await loadData()
  } catch {
    showSnack(`Failed to save ${currentTableLabel.value}`, 'error')
  } finally {
    isSaving.value = false
  }
}

const confirmDelete = (row: any): void => {
  deleteTarget.value = row
  deleteDialog.value = true
}

const handleDelete = async (): Promise<void> => {
  isDeleting.value = true
  try {
    await currentTable.value.api.delete(deleteTarget.value.id)
    showSnack('Record deleted successfully', 'success')
    deleteDialog.value = false
    await loadData()
  } catch {
    showSnack('Failed to delete record', 'error')
  } finally {
    isDeleting.value = false
  }
}

// ─── Helpers ─────────────────────────────────────────────
const showSnack = (text: string, color: string): void => {
  snackbarText.value = text
  snackbarColor.value = color
  snackbar.value = true
}

const roleColor = (role: string): string => {
  const map: Record<string, string> = { admin: 'error', teacher: 'primary', student: 'success', parent: 'info' }
  return map[role] || 'grey'
}

const statusColor = (status: string): string => {
  const map: Record<string, string> = { menguasai: 'primary', cukup: 'warning', belum: 'error' }
  return map[status] || 'grey'
}

const truncate = (val: any, len = 40): string => {
  const s = String(val ?? '—')
  return s.length > len ? s.slice(0, len) + '…' : s
}
</script>

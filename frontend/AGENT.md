## Business Logic Rules
- **Admin:** Akses penuh (manajemen user, sekolah, & database).
- **Teacher:** Input nilai, kelola kelas, dan melihat rekomendasi AI untuk satu kelas.
- **Parent:** Monitoring nilai anak secara real-time dan melihat saran pengembangan.
- **Student:** Melihat nilai pribadi dan tips belajar dari AI.

## Technical Specifics
- Gunakan `@` sebagai alias untuk `src/`.
- Gunakan `v-slot` untuk kustomisasi tabel Vuetify.
- Gunakan UnoCSS untuk spacing dan layouting yang tidak tercover oleh Vuetify props (misal: `class="grid grid-cols-3 gap-4"`).
- Integrasi AI akan menggunakan API backend (Go) yang memproses data nilai menjadi string prompt untuk LLM.
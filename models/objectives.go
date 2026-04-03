package models

type Objective struct {
    ObjectiveId int    `db:"id" json:"objective_id"`
    SubjectId   int    `db:"subject_id" json:"subject_id"`
    Week        int    `db:"week" json:"objective_week"`
    Description string `db:"description" json:"objective_description"` // Ubah ke string
}
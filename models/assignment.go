package models

type Assignment struct{
	AssignmentId int `db:"id" json:"assignment_id"`
	ObjectiveId int `db:"objective_id" json:"objective_id"`
	AssignmentTitle string `db:"title" json:"assignment_title"`
	AssignmentType string `db:"type" json:"assignment_type"`
}
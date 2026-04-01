package models

import "time"

type Assignment struct{
	AssignmentId int `db:"id" json:"assignment_id"`
	ObjectiveId int `db:"objective_id" json:"objective_id"`
	AssignmentTitle string `db:"title" json:"assignment_title"`
	AssignmentType string `db:"type" json:"assignment_type"`
	Deadline time.Time `db:"deadline" json:"deadline"`
	UpdateAt time.Time `db:"update_at" json:"update_at"`
}
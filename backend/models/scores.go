package models

type Scores struct{
	ScoreID int `db:"id" json:"score_id"`
	StudentId int `db:"student_id" json:"student_id"`
	AssignmentId int `db:"assignment_id" json:"assignment_id"`
	Score int `db:"score" json:"score"`
}
package models

type Progress struct{
	ProgressId int `db:"id" json:"progress_id"`
	StudentId int `db:"student_id" json:"student_id"`
	ObjectiveId int `db:"objective_id" json:"objective_id"`
	Week int `db:"week" json:"week"`
	FinalScore int `db:"final_score" json:"final_score"`
	Status string `db:"status" json:"status"`
	Recommendatiaon string `db:"recommendation" json:"recommendation"`
}
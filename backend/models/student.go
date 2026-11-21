package models

type Student struct{
	StudentId int ` db:"id" json:"student_id"`
	StudentName string ` db:"name" json:"student_name"`
	StudentClass string ` db:"class" json:"student_class"`
}
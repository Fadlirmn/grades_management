package models

type Student struct{
	StudentId int ` db:"id" json:"student_id"`
	StudentName int ` db:"name" json:"student_name"`
	StudentClass int ` db:"class" json:"student_class"`
}
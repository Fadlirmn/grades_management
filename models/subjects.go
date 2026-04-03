package models

type Subject struct{
	SubjectID int `db:"id" json:"id"`
	SubjectName string `db:"name" json:"name"`
}
package models

type Subject struct{
	SubjectID int `db:"id" json:"id_subject"`
	SubjectName string `db:"name" json:"name"`
}
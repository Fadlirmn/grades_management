package repository

import (
	"grades-management/models"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type SubjectRepository interface {
	FindAllSubject() []models.Subject
	SaveSubject(Subject models.Subject)
	UpdateSubject(id int, Subject models.Subject) error
	DeleteSubject(id int) error
}

type SubjectRepo struct {
	db *sqlx.DB
}

func NewSubjectRepository(db *sqlx.DB) SubjectRepository {
	return &SubjectRepo{db: db}
}

func (r *SubjectRepo) FindAllSubject() []models.Subject {
	var subject []models.Subject
	 err := r.db.Select(&subject,"SELECT * FROM subject")

	if err != nil {
		log.Println("error query", err)
		return nil
	}
	return subject
}

func (r *SubjectRepo) SaveSubject(subject models.Subject) {
	_, err := r.db.NamedExec(
		`INSERT INTO subject(name) VALUES (:name)`,
		subject,
	)
	if err != nil {
		log.Println("fail Add Subject", err)
	}
}

func (r *SubjectRepo) UpdateSubject(subjectId int, subject models.Subject) error {
	_, err := r.db.Exec("UPDATE subject SET name=$1 WHERE id=$3", subject.SubjectName, subjectId)
	return err
}

func (r *SubjectRepo) DeleteSubject(subjectId int) error {
	_, err := r.db.Exec("DELETE FROM subject WHERE id=$1", subjectId)
	return err
}

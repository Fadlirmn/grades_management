package repository

import (
	"grades-management/models"
	"github.com/jmoiron/sqlx"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type StudentRepository interface {
	FindAllStudent() []models.Student
	SaveStudent(student models.Student)
	UpdateStudent(id int, student models.Student) error
	DeleteStudent(id int) error
}

type studentRepo struct {
	db *sqlx.DB
}

func NewStudentRepository(db *sqlx.DB) StudentRepository {
	return &studentRepo{db: db}
}

func (r *studentRepo) FindAllStudent() []models.Student {
	var students []models.Student
	 err := r.db.Select(&students,"SELECT * FROM students")

	if err != nil {
		log.Println("error query", err)
		return nil
	}
	return students
}

func (r *studentRepo) SaveStudent(student models.Student) {
	_, err := r.db.NamedExec(
		`INSERT INTO students(name, class) VALUES(:name, :class, )`, 
	student,
)
	if err != nil {
		log.Println("fail add Student", err)
	}
}

func (r *studentRepo) UpdateStudent(idStudent int, student models.Student) error {
	_, err := r.db.Exec("UPDATE students SET name=$1, class=$2 WHERE id=$3", student.StudentName, student.StudentClass, idStudent)
	return err
}

func (r *studentRepo) DeleteStudent(idStudent int) error {
	_, err := r.db.Exec("DELETE FROM students WHERE id=$1", idStudent)
	return err
}

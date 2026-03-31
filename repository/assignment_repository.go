package repository

import (
	"grades-management/models"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type AssignmentRepository interface {
	FindAllAssignment() []models.Assignment
	SaveAssignment(Assignment models.Assignment)
	UpdateAssignment(id int, Assignment models.Assignment) error
	DeleteAssignment(id int) error
}

type AssignmentRepo struct {
	db *sqlx.DB
}

func NewAssignmentRepository(db *sqlx.DB) AssignmentRepository {
	return &AssignmentRepo{db: db}
}

func (r *AssignmentRepo) FindAllAssignment() []models.Assignment {
	var assignments []models.Assignment
	 err := r.db.Select(&assignments,"SELECT * FROM assignments")

	if err != nil {
		log.Println("error query", err)
		return nil
	}
	return assignments
}

func (r *AssignmentRepo) SaveAssignment(assignment models.Assignment) {
	_, err := r.db.NamedExec(
		`INSERT INTO assignments(name_assignment,assignment_code) VALUES (:name_assignment,:assignment_code)`,
		assignment,
	)
	if err != nil {
		log.Println("fail Add Assignment", err)
	}
}

func (r *AssignmentRepo) UpdateAssignment(assignmentId int, assignment models.Assignment) error {
	_, err := r.db.Exec("UPDATE assignments SET objective_id=$1, title=$2, type=$3 WHERE id_assignment=$3", assignment.AssignmentTitle, assignment.AssignmentType, assignmentId)
	return err
}

func (r *AssignmentRepo) DeleteAssignment(assignmentId int) error {
	_, err := r.db.Exec("DELETE FROM assignments WHERE id_assignment=$1", assignmentId)
	return err
}

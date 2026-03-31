package repository

import (
	"grades-management/models"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type ObjectiveRepository interface {
	FindAllObjective() []models.Objective
	SaveObjective(Objective models.Objective)
	UpdateObjective(id int, Objective models.Objective) error
	DeleteObjective(id int) error
}

type ObjectiveRepo struct {
	db *sqlx.DB
}

func NewObjectiveRepository(db *sqlx.DB) ObjectiveRepository {
	return &ObjectiveRepo{db: db}
}

func (r *ObjectiveRepo) FindAllObjective() []models.Objective {
	var objectives []models.Objective
	 err := r.db.Select(&objectives,"SELECT * FROM objectives")

	if err != nil {
		log.Println("error query", err)
		return nil
	}
	return objectives
}

func (r *ObjectiveRepo) SaveObjective(objectives models.Objective) {
	_, err := r.db.NamedExec(
		`INSERT INTO objectives(subject_id,week,description) VALUES (:subject_id,:week,:description)`,
		objectives,
	)
	if err != nil {
		log.Println("fail Add Objective", err)
	}
}

func (r *ObjectiveRepo) UpdateObjective(objectivesId int, objectives models.Objective) error {
	_, err := r.db.Exec("UPDATE objectives SET subject_id=$1, week=$2, description=$3 WHERE id=$3", objectives.SubjectId, objectives.Week, objectives.Description, objectivesId)
	return err
}

func (r *ObjectiveRepo) DeleteObjective(objectivesId int) error {
	_, err := r.db.Exec("DELETE FROM objectives WHERE id=$1", objectivesId)
	return err
}

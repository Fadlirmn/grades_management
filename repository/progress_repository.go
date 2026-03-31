package repository

import (
	"grades-management/models"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type ProgressRepository interface {
	FindAllProgress() []models.Progress
	SaveProgress(Progress models.Progress)
	UpdateProgress(id int, Progress models.Progress) error
	DeleteProgress(id int) error
	FindAnalysisByStudentId(studentId int)([]models.Progress,error)
}

type ProgressRepo struct {
	db *sqlx.DB
}

func NewProgressRepository(db *sqlx.DB) ProgressRepository {
	return &ProgressRepo{db: db}
}

func (r *ProgressRepo) FindAllProgress() []models.Progress {
	var progress []models.Progress
	 err := r.db.Select(&progress,"SELECT * FROM progress")

	if err != nil {
		log.Println("error query", err)
		return nil
	}
	return progress
}

func (r *ProgressRepo) SaveProgress(progress models.Progress) {
	_, err := r.db.NamedExec(
		`INSERT INTO progress(student_id,objective_id,week,final_score,status,recommendation) VALUES (:student_id,:objective_id,:week,:final_score,:status,:recommendation)`,
		progress,
	)
	if err != nil {
		log.Println("fail Add Progress", err)
	}
}

func (r *ProgressRepo) UpdateProgress(progressId int, progress models.Progress) error {
	_, err := r.db.Exec("UPDATE progresss SET student_id=$1, objective_id=$2, week=$3,final_score=$4,status=$5,recommendation=$6 WHERE id=$7", progress.StudentId, progress.ObjectiveId, progress.Week,progress.FinalScore,progress.Status,progress.Recommendatiaon, progressId)
	return err
}

func (r *ProgressRepo) DeleteProgress(progressId int) error {
	_, err := r.db.Exec("DELETE FROM progresss WHERE id=$1", progressId)
	return err
}


func (r *ProgressRepo)FindAnalysisByStudentId(studentId int)([]models.Progress,error)  {
	var progress []models.Progress
	err := r.db.Select(&progress,"SELECT * FROM grades WHERE student_id=$1",studentId)
	if err != nil {
		log.Printf("error query",err)
	}
	return progress,nil
}
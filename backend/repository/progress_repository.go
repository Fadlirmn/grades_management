package repository

import (
	"grades-management/models"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type ProgressRepository interface {
	FindAllProgress() []models.Progress
	SaveProgress(Progress models.Progress)
	UpdateProgress(id int, Progress models.Progress) error
	DeleteProgress(id int) error
	FindAnalysisByStudentId(studentId int) ([]models.Progress, error)
	CountPending()(int, error)
	FindPendingAnalysis(studentId int) ([]models.Progress, error)
	UpdateRecommendation(progressId int, progress string) error
	UpdateBatchRecommendation(result []models.AIRecommendation) error 
	UpsertFromSheets(studentId, objId,week,score int,status string)error
}

type ProgressRepo struct {
	db *sqlx.DB
}

func NewProgressRepository(db *sqlx.DB) ProgressRepository {
	return &ProgressRepo{db: db}
}

func (r *ProgressRepo) FindAllProgress() []models.Progress {
	var progress []models.Progress
	err := r.db.Select(&progress, "SELECT * FROM progress")

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
	_, err := r.db.Exec("UPDATE progresss SET student_id=$1, objective_id=$2, week=$3,final_score=$4,status=$5,recommendation=$6 WHERE id=$7", progress.StudentId, progress.ObjectiveId, progress.Week, progress.FinalScore, progress.Status, progress.Recommendatiaon, progressId)
	return err
}

func (r *ProgressRepo) DeleteProgress(progressId int) error {
	_, err := r.db.Exec("DELETE FROM progresss WHERE id=$1", progressId)
	return err
}

func (r *ProgressRepo) FindAnalysisByStudentId(studentId int) ([]models.Progress, error) {
	var progress []models.Progress
	err := r.db.Select(&progress, "SELECT * FROM progress WHERE student_id=$1", studentId)
	if err != nil {
		log.Printf("error query")
	}
	return progress, nil
}

func (r * ProgressRepo)CountPending()(int, error)  {
	var CountPending int
	err:= r.db.Get(&CountPending,"SELECT COUNT(*) FROM progress WHERE COALESCE(recommendation, '') = ''")
	return CountPending,err
}

//join progress C obejctives
func (r *ProgressRepo) FindPendingAnalysis(studentId int) ([]models.Progress, error) {
	var progress []models.Progress

	query:=`
        SELECT 
            p.id, 
            p.student_id, 
            p.objective_id, 
            p.final_score, 
            p.week, 
            p.status,
            COALESCE(p.recommendation, '') as recommendation, 
            s.name as subject_name, 
            o.description as objective_desc 
        FROM progress p 
        JOIN objectives o ON p.objective_id = o.id 
        JOIN subjects s ON o.subject_id = s.id 
        WHERE COALESCE(p.recommendation, '') = ''`

	if studentId != 0 {
		query += " AND p.student_id = $1 LIMIT 50"
		err := r.db.Select(&progress,query, studentId)
		if err != nil {
			log.Printf("error query student id")
		}
	} else {
		query += " LIMIT 50"
		err := r.db.Select(&progress, query)
		if err != nil {
			log.Printf("error query FindPendingAnalysis: %v", err)
		}
	}
	if len(progress) > 0 {
        log.Printf("DEBUG: Baris pertama yang ditarik memiliki ProgressId: %d", progress[0].ProgressId)
    }
	return progress, nil
}

func (r *ProgressRepo) UpdateRecommendation(progressId int, progress string) error {
	_, err := r.db.Exec("UPDATE progress SET recommendation=$1 WHERE id=$2",  progress, progressId)
	return err
}
func (r *ProgressRepo) UpdateBatchRecommendation(result []models.AIRecommendation) error {
	if len(result)==0 {
		return nil
	}
	tx,err:= r.db.Begin()
	if err != nil {
		return  err
	}
	query:=`UPDATE progress SET recommendation=$1 WHERE id=$2`

	stmt,err:=tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for _,res:= range result{
		_, err:= stmt.Exec(res.Rec,res.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (r *ProgressRepo) UpsertFromSheets(studentId, objId,week,score int,status string)error {
	query:= `INSERT INTO progress(student_id,objective_id,week,final_score,status)
			VALUES ($1,$2,$3,$4,$5)
			ON CONFLICT (student_id,objective_id,week)
			DO UPDATE SET final_score = EXCLUDED.final_score, status= EXCLUDED.status, recommendation = NULL`

	_,err:= r.db.Exec(query,studentId,objId,week,score,status)
	return err
}
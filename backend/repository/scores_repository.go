package repository

import (
	"grades-management/models"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type ScoreRepository interface {
	FindAllScore() []models.Scores
	SaveScore(Score models.Scores)
	UpdateScore(id int, Score models.Scores) error
	DeleteScore(id int) error
}

type ScoreRepo struct {
	db *sqlx.DB
}

func NewScoreRepository(db *sqlx.DB) ScoreRepository {
	return &ScoreRepo{db: db}
}

func (r *ScoreRepo) FindAllScore() []models.Scores {
	var scores []models.Scores
	 err := r.db.Select(&scores,"SELECT * FROM scores")

	if err != nil {
		log.Println("error query", err)
		return nil
	}
	return scores
}

func (r *ScoreRepo) SaveScore(scores models.Scores) {
	_, err := r.db.NamedExec(
		`INSERT INTO scores(student_id,assignment_id,scores) VALUES (:student_id,:assignment_id,:scores)`,
		scores,
	)
	if err != nil {
		log.Println("fail Add Score", err)
	}
}

func (r *ScoreRepo) UpdateScore(scoresId int, scores models.Scores) error {
	_, err := r.db.Exec("UPDATE scores SET student_id=$1, assignment_id=$2, scores=$3 WHERE id=$3", scores.StudentId, scores.AssignmentId, scores.Score, scoresId)
	return err
}

func (r *ScoreRepo) DeleteScore(scoresId int) error {
	_, err := r.db.Exec("DELETE FROM scores WHERE id=$1", scoresId)
	return err
}

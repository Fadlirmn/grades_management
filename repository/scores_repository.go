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
	var score []models.Scores
	 err := r.db.Select(&score,"SELECT * FROM score")

	if err != nil {
		log.Println("error query", err)
		return nil
	}
	return score
}

func (r *ScoreRepo) SaveScore(score models.Scores) {
	_, err := r.db.NamedExec(
		`INSERT INTO score(student_id,assignment_id,score) VALUES (:student_id,:assignment_id,:score)`,
		score,
	)
	if err != nil {
		log.Println("fail Add Score", err)
	}
}

func (r *ScoreRepo) UpdateScore(scoreId int, score models.Scores) error {
	_, err := r.db.Exec("UPDATE score SET student_id=$1, assignment_id=$2, score=$3 WHERE id=$3", score.StudentId, score.AssignmentId, score.Score, scoreId)
	return err
}

func (r *ScoreRepo) DeleteScore(scoreId int) error {
	_, err := r.db.Exec("DELETE FROM score WHERE id=$1", scoreId)
	return err
}

package services

import(
	"grades-management/models"
	"grades-management/repository"
)

type ScoresService struct{
	scoreRepo repository.ScoreRepository
}

func NewScoresService(score repository.ScoreRepository) *ScoresService  {
	return & ScoresService{
		scoreRepo: score,
	}
}

func (s *ScoresService) GetScores()([]models.Scores)  {
	return s.scoreRepo.FindAllScore()
}

func (s *ScoresService) CreateScores(Scores models.Scores)  {
		s.scoreRepo.SaveScore(Scores)
}

func (s *ScoresService) UpdateScores(id int,Scores models.Scores)  error{
	return s.scoreRepo.UpdateScore(id, Scores)
}
func (s *ScoresService) DeleteScores(id int) error {
	return s.scoreRepo.DeleteScore(id)
}
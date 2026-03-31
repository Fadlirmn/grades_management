package services

import(
	"grades-management/models"
	"grades-management/repository"
)

type ProgressService struct{
	progressRepo repository.ProgressRepository
}

func NewProgressService(progress repository.ProgressRepository) *ProgressService  {
	return & ProgressService{
		progressRepo: progress,
	}
}

func (s *ProgressService) GetProgress()([]models.Progress)  {
	return s.progressRepo.FindAllProgress()
}

func (s *ProgressService) CreateProgress(Progress models.Progress)  {
		s.progressRepo.SaveProgress(Progress)
}

func (s *ProgressService) UpdateProgress(id int,Progress models.Progress)  error{
	return s.progressRepo.UpdateProgress(id, Progress)
}
func (s *ProgressService) DeleteProgress(id int) error {
	return s.progressRepo.DeleteProgress(id)
}
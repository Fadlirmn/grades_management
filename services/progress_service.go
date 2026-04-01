package services

import (
	"grades-management/models"
	"grades-management/repository"
	"log"
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

func (s *ProgressService)FindAnalysisByStudentId(studentId int) ([] models.Progress,error) {
	result, err:= s.progressRepo.FindAnalysisByStudentId(studentId)
	if err != nil {
		log.Printf("Failed fetch data for id %d: %v", studentId, err)
		return nil, err
	}
	
	if len(result)==0{
		log.Printf("recommendation not found, for ID %d",studentId)
		return nil,nil
	}
	return result,nil
}

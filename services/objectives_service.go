package services

import(
	"grades-management/models"
	"grades-management/repository"
)

type ObjectiveService struct{
	objectiveRepo repository.ObjectiveRepository
}

func NewObjectiveService(objective repository.ObjectiveRepository) *ObjectiveService  {
	return & ObjectiveService{
		objectiveRepo: objective,
	}
}

func (s *ObjectiveService) GetObjective()([]models.Objective)  {
	return s.objectiveRepo.FindAllObjective()
}

func (s *ObjectiveService) CreateObjective(Objective models.Objective)  {
		s.objectiveRepo.SaveObjective(Objective)
}

func (s *ObjectiveService) UpdateObjective(id int,Objective models.Objective)  error{
	return s.objectiveRepo.UpdateObjective(id, Objective)
}
func (s *ObjectiveService) DeleteObjective(id int) error {
	return s.objectiveRepo.DeleteObjective(id)
}
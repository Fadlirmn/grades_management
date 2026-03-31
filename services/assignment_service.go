package services

import(
	"grades-management/models"
	"grades-management/repository"
)

type AssignService struct{
	assignRepo repository.AssignmentRepository
}

func NewAssignService(assign repository.AssignmentRepository) *AssignService  {
	return & AssignService{
		assignRepo: assign,
	}
}

func (s *AssignService) GetAssignments()([]models.Assignment)  {
	return s.assignRepo.FindAllAssignment()
}

func (s *AssignService) CreateAssignment(Assignment models.Assignment)  {
		s.assignRepo.SaveAssignment(Assignment)
}

func (s *AssignService) UpdateAssignment(id int,Assignment models.Assignment)  error{
	return s.assignRepo.UpdateAssignment(id, Assignment)
}
func (s *AssignService) DeleteAssignment(id int) error {
	return s.assignRepo.DeleteAssignment(id)
}
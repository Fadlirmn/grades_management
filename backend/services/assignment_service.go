package services

import (
	"fmt"
	"grades-management/models"
	"grades-management/repository"
	"time"
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

func (s *AssignService) CreateAssignment(assignment models.Assignment)  {
	assignment.UpdateAt = time.Now().UTC()

	err := s.assignRepo.SaveAssignment(assignment)
	if err != nil {
		fmt.Printf("Failed Save Assignment")
		return
	}
	fmt.Println("Successs Save Assignment")
}

func (s *AssignService) UpdateAssignment(id int,assignment models.Assignment)  error{
	assignment.UpdateAt = time.Now().UTC()

	err := s.assignRepo.UpdateAssignment(id,assignment)
	if err != nil {
		fmt.Printf("Failed update Assignment")
		
	}
	fmt.Println("Successs update Assignment")
	return err
}
func (s *AssignService) DeleteAssignment(id int) error {
	return s.assignRepo.DeleteAssignment(id)
}
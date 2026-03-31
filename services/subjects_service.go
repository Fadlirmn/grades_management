package services

import(
	"grades-management/models"
	"grades-management/repository"
)

type SubjectService struct{
	subjectRepo repository.SubjectRepository
}

func NewSubjectService(subject repository.SubjectRepository) *SubjectService  {
	return & SubjectService{
		subjectRepo: subject,
	}
}

func (s *SubjectService) GetSubject()([]models.Subject)  {
	return s.subjectRepo.FindAllSubject()
}

func (s *SubjectService) CreateSubject(Subject models.Subject)  {
		s.subjectRepo.SaveSubject(Subject)
}

func (s *SubjectService) UpdateSubject(id int,Subject models.Subject)  error{
	return s.subjectRepo.UpdateSubject(id, Subject)
}
func (s *SubjectService) DeleteSubject(id int) error {
	return s.subjectRepo.DeleteSubject(id)
}
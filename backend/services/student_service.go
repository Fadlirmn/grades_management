package services

import(
	"grades-management/models"
	"grades-management/repository"
)

type StudentService struct{
	studentRepo repository.StudentRepository
}

func NewStudentService(student repository.StudentRepository) *StudentService  {
	return & StudentService{
		studentRepo: student,
	}
}

func (s *StudentService) GetStudent()([]models.Student)  {
	return s.studentRepo.FindAllStudent()
}

func (s *StudentService) CreateStudent(Student models.Student)  {
		s.studentRepo.SaveStudent(Student)
}

func (s *StudentService) UpdateStudent(id int,Student models.Student)  error{
	return s.studentRepo.UpdateStudent(id, Student)
}
func (s *StudentService) DeleteStudent(id int) error {
	return s.studentRepo.DeleteStudent(id)
}
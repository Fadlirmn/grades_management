package services

import(
	"grades-management/models"
	"grades-management/repository"
)


type UserService struct{
	repo repository.UserRepository
}
func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUsers()[]models.User  {
	return  s.repo.FindAll()
}

func (s *UserService) FindByUsername(username string) (*models.User, error) {
	return s.repo.FindByUsername(username)
}

func (s *UserService) CreateUser(user models.User)  {
	s.repo.Save(user)
}

func (s *UserService) UpdateUser(id string,user models.User)  error{
	return s.repo.Update(id, user)
}
func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}
package user

import (
	"log"

	"backend/models"

	"github.com/mashingan/smapping"
)

//UserService is a contract.....
type UserService interface {
	Update(user models.USER) models.USER
	Profile(userID string) models.USER
}

type userService struct {
	userRepository UsersRepository
}

//NewUserService creates a new instance of UserService
func NewUserService(userRepo UsersRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user models.USER) models.USER {
	userToUpdate := models.USER{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) models.USER {
	return service.userRepository.ProfileUser(userID)
}

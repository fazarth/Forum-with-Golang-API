package user

import (
	"log"

	"backend/models"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

//AuthService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential(username string, password string) interface{}
	CreateUser(user models.USER) models.USER
	FindByUserName(username string) models.USER
	IsDuplicateUserName(username string) bool
	IsDuplicateMSISDN(msisdn int) bool
}

type authService struct {
	userRepository UsersRepository
}

//NewAuthService creates a new instance of AuthService
func NewAuthService(userRep UsersRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(USERNAME string, PASSWORD string) interface{} {
	res := service.userRepository.VerifyCredential(USERNAME, PASSWORD)
	if v, ok := res.(models.USER); ok {
		comparedPassword := comparePassword(v.PASSWORD, []byte(PASSWORD))
		if v.USERNAME == USERNAME && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user models.USER) models.USER {
	userToCreate := models.USER{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *authService) FindByUserName(username string) models.USER {
	return service.userRepository.FindByUserName(username)
}

func (service *authService) IsDuplicateUserName(username string) bool {
	res := service.userRepository.IsDuplicateUserName(username)
	return !(res.Error == nil)
}

func (service *authService) IsDuplicateMSISDN(msisdn int) bool {
	res := service.userRepository.IsDuplicateMSISDN(msisdn)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

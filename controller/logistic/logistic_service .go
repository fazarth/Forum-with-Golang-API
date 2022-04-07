package logistic

import (
	"fmt"
	"log"

	"backend/models"

	"github.com/mashingan/smapping"
)

//LogisticService is a ....
type LogisticService interface {
	InsertLogistic(b models.LOGISTIC) models.LOGISTIC
	UpdateLogistic(b models.LOGISTIC) models.LOGISTIC
	DeleteLogistic(b models.LOGISTIC)
	GetAllLogistic() []models.LOGISTIC
	FindLogisticByOriginName(OriginName string, DestinationName string) models.LOGISTIC
	FindLogisticById(logisticId uint64) models.LOGISTIC
	IsAllowedToEdit(userID string, logisticId uint64) bool
}

type logisticService struct {
	logisticRepository LogisticRepository
}

//NewLogisticService .....
func NewLogisticService(logisticRepo LogisticRepository) LogisticService {
	return &logisticService{
		logisticRepository: logisticRepo,
	}
}

func (service *logisticService) InsertLogistic(b models.LOGISTIC) models.LOGISTIC {
	logistic := models.LOGISTIC{}
	err := smapping.FillStruct(&logistic, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.logisticRepository.InsertLogistic(logistic)
	return res
}

func (service *logisticService) UpdateLogistic(b models.LOGISTIC) models.LOGISTIC {
	res := service.logisticRepository.UpdateLogistic(b)
	return res
}

func (service *logisticService) GetAllLogistic() []models.LOGISTIC {
	return service.logisticRepository.GetAllLogistic()
}

func (service *logisticService) FindLogisticByOriginName(OriginName string, DestinationName string) models.LOGISTIC {
	return service.logisticRepository.FindLogisticByOriginName(OriginName, DestinationName)
}

func (service *logisticService) FindLogisticById(logisticId uint64) models.LOGISTIC {
	return service.logisticRepository.FindLogisticById(logisticId)
}

func (service *logisticService) DeleteLogistic(b models.LOGISTIC) {
	service.logisticRepository.DeleteLogistic(b)
}

func (service *logisticService) IsAllowedToEdit(userID string, logisticId uint64) bool {
	b := service.logisticRepository.FindLogisticById(logisticId)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return userID == id
}

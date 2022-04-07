package logistic

import (
	"backend/models"

	"gorm.io/gorm"
)

//ModuleRepository is contract what userRepository can do to db
type LogisticRepository interface {
	InsertLogistic(b models.LOGISTIC) models.LOGISTIC
	GetAllLogistic() []models.LOGISTIC
	FindLogisticByOriginName(OriginName string, DestinationName string) models.LOGISTIC
	FindLogisticById(logisticId uint64) models.LOGISTIC
	UpdateLogistic(b models.LOGISTIC) models.LOGISTIC
	DeleteLogistic(b models.LOGISTIC)
}

type logisticConnection struct {
	connection *gorm.DB
}

//NewLogisticRepository creates an instance LogisticRepository
func NewLogisticRepository(dbConn *gorm.DB) LogisticRepository {
	return &logisticConnection{
		connection: dbConn,
	}
}

func (db *logisticConnection) InsertLogistic(b models.LOGISTIC) models.LOGISTIC {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *logisticConnection) GetAllLogistic() []models.LOGISTIC {
	var getlogistic []models.LOGISTIC
	db.connection.Find(&getlogistic)
	return getlogistic
}

func (db *logisticConnection) FindLogisticByOriginName(OriginName string, DestinationName string) models.LOGISTIC {
	var findlogisticID models.LOGISTIC
	db.connection.Find(&findlogisticID).Where("ORIGIN_NAME = ? AND DESTINATION_NAME = ?", OriginName, DestinationName)
	return findlogisticID
}

func (db *logisticConnection) FindLogisticById(logisticId uint64) models.LOGISTIC {
	var findlogisticID models.LOGISTIC
	db.connection.Find(&findlogisticID).Where("LOGISTIC_ID = ?", logisticId)
	return findlogisticID
}

func (db *logisticConnection) UpdateLogistic(b models.LOGISTIC) models.LOGISTIC {
	var updatelogistic models.LOGISTIC
	db.connection.Find(&updatelogistic).Where("LOGISTIC_ID = ?", b.LOGISTIC_ID)
	updatelogistic.COMMENT = b.COMMENT
	db.connection.Updates(&b)
	return b
}

func (db *logisticConnection) DeleteLogistic(b models.LOGISTIC) {
	db.connection.Delete(&b)
}

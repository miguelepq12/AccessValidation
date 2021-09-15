package application

import (
	"AccessValidation/app/domain/entity"
	"AccessValidation/app/domain/repository"
	"AccessValidation/app/infrastructure/kafka/access/producer"
	"encoding/json"
	"errors"
	"fmt"
)

type AccessUseCase struct {
	accessRepository repository.AccessRepository
	accessRecordProducer *producer.AccessRecordKafkaProducer
}

func NewAccessUseCase(accessRepository repository.AccessRepository, producer *producer.AccessRecordKafkaProducer) *AccessUseCase {
	return &AccessUseCase{accessRepository: accessRepository,accessRecordProducer: producer}
}

func (c *AccessUseCase) ValidateAccess(access *entity.AccessInformation) error  {
	hasAccess, err := c.accessRepository.HasAccess(access.Id,access.SensorId)
	if err!=nil{
		return err
	}

	if hasAccess {
		accessRecord := entity.NewAccessRecord(access)
		accessRecordByte, err := json.Marshal(accessRecord)
		if err != nil{
			fmt.Println("... Error to send kafka message to marshall")
		}

		err = c.accessRecordProducer.PushMessage(accessRecordByte)
		if err != nil {
			fmt.Println("... Error to send kafka message")
		}
	}else{
		return errors.New("access denied")
	}

	return nil
}

package entity

import (
	"fmt"
	"time"
)

type AccessRecord struct {
	Id string `json:"id"`
	Name  string `json:"name"`
	Rol string `json:"rol"`
	Place string `json:"place"`
	AccessTime time.Time `json:"access_time"`
}

func NewAccessRecord(information *AccessInformation) *AccessRecord {
	return &AccessRecord{
		Id:         information.Id,
		Name:       fmt.Sprintf("%s %s",information.FirstName,information.LastName),
		Rol:        information.Rol,
		Place:      information.SensorPlace,
		AccessTime: information.AccessTime,
	}
}

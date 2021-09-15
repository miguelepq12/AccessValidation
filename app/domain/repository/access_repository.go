package repository

type AccessRepository interface {
	HasAccess(idPersonalInformation string, idSensor string) (bool, error)
}

package oracle

import (
	"fmt"
)

type AccessMysqlRepository struct {
	mysqlDb *MysqlConnection
}

func NewAccessMysqlRepository(mysqlDb *MysqlConnection) *AccessMysqlRepository {
	return &AccessMysqlRepository{mysqlDb: mysqlDb}
}
func (c *AccessMysqlRepository) HasAccess(idPersonalInformation string, idSensor string) (bool, error) {
	var (
		hasAccess bool
	)
	err := c.mysqlDb.db.QueryRow("SELECT id_personal,id_sensor,has_access "+
		"FROM access_personal "+
		"WHERE id_personal = ? AND id_sensor = ?",
		idPersonalInformation, idSensor).
		Scan(&idPersonalInformation, &idSensor, &hasAccess)

	if err != nil {
		fmt.Println("...error get row")
	}
	return hasAccess, nil
}

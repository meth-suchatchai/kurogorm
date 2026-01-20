package kurogorm

import (
	dbmodels "kurogorm/daos"
	"log"

	"github.com/meth-suchatchai/kurostatemachine"
)

func (c *defaultClient) CreateScene(data *dbmodels.Scene) (*dbmodels.Scene, error) {
	result := c.orm.Create(&data)
	if result.Error != nil {
		log.Printf("create scene failed: %v", result.Error)
		return nil, result.Error
	}

	return data, nil
}

func (c *defaultClient) UpdateScene(id uint, params map[string]interface{}) error {
	return c.orm.Model(&dbmodels.Scene{}).Where("id = ?", id).Updates(params).Error
}

func (c *defaultClient) UpdateStatusScene(id uint, status kurostatemachine.State) error {
	return c.orm.Model(&dbmodels.Scene{}).Where("id = ?", id).Update("status", status).Error
}

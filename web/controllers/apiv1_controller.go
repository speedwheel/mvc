package controllers

import (
	"github.com/kataras/iris"
	"github.com/speedwheel/mvc/datamodels"
	"github.com/speedwheel/mvc/services"
)

type ApiV1Controller struct {
	Ctx     iris.Context
	Service services.PropertyService
}

func (c *ApiV1Controller) GetSetProperty() interface{} {
	err := c.Service.Create(datamodels.Property{
		Name:  "2 Bedroom Apartment",
		Price: 100000,
		Rooms: 2,
	})
	if err != nil {
		return map[string]interface{}{"error": err}
	}
	return map[string]interface{}{"success": true}
}

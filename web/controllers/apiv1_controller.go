package controllers

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/speedwheel/mvc/datamodels"
	"github.com/speedwheel/mvc/services"
)

type ApiV1Controller struct {
	Ctx     iris.Context
	Service services.PropertyService
}

func (c *ApiV1Controller) GetSetProperty() datamodels.Property {
	property := datamodels.Property{
		Name:  "2 Bedroom Apartment",
		Price: 100000,
		Rooms: 2,
	}
	err := c.Service.Create(property)
	fmt.Println(err)
	return property
}

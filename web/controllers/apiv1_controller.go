package controllers

import (
	"github.com/kataras/iris"
	"github.com/speedwheel/mvc/datasource/db/aerospike"
	"github.com/speedwheel/mvc/services"
)

type ApiV1Controller struct {
	Ctx     iris.Context
	Service services.PropertyService
}

func (c *ApiV1Controller) GetSetProperty() interface{} {
	err := c.Service.Create(aerospike.RowObj{
		"name":  "2 Bedroom Apartment",
		"price": 100000,
		"rooms": 2,
	})
	if err != nil {
		return map[string]interface{}{"error": err}
	}
	return map[string]interface{}{"success": true}
}

package controllers

import "github.com/kataras/iris"

type ApiV1Controller struct {
	Ctx iris.Context
}

func (c *ApiV1Controller) GetProperty() {
	c.Ctx.JSON(iris.Map{"id": "1"})
}

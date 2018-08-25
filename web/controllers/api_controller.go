package controllers

import (
	"net/http"

	"github.com/kataras/iris"
)

type ApiController struct {
	Ctx iris.Context
}

func (c *ApiController) Get() {
	c.Ctx.StatusCode(http.StatusForbidden)
}

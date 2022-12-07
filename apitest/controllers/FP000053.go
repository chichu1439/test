package controllers

import (
	"github.com/astaxie/beego"
)

// api query
type FP000053Controller struct {
	beego.Controller
}

// @Title
// @Description api query
// @Param body  body  models.FP000053I "Input parameter"
// @Success 200  {object} models.FP000053O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000053Controller) Post()() {
}

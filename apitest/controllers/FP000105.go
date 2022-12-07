package controllers

import (
	"github.com/astaxie/beego"
)

// balance adjustment apply
type FP000105Controller struct {
	beego.Controller
}

// @Title
// @Description balance adjustment apply
// @Param body  body  models.FP000105I "Input parameter"
// @Success 200  {object} models.FP000105O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000105Controller) Post()() {
}

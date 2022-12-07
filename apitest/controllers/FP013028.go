package controllers

import (
	"github.com/astaxie/beego"
)

// adjust mockinfo
type FP013028Controller struct {
	beego.Controller
}

// @Title
// @Description adjust mockinfo
// @Param body  body  models.FP013028I "Input parameter"
// @Success 200  {object} models.FP013028O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP013028Controller) Post()() {
}

package controllers

import (
	"github.com/astaxie/beego"
)

// report file query
type FP000030Controller struct {
	beego.Controller
}

// @Title
// @Description report file query
// @Param body  body  models.FP000030I "Input parameter"
// @Success 200  {object} models.FP000030O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000030Controller) Post()() {
}

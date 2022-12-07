package controllers

import (
	"github.com/astaxie/beego"
)

// participant unsuspension
type FP000005Controller struct {
	beego.Controller
}

// @Title
// @Description participant unsuspension
// @Param body  body  models.FP000005I "Input parameter"
// @Success 200  {object} models.FP000005O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000005Controller) Post()() {
}

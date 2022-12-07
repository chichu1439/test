package controllers

import (
	"github.com/astaxie/beego"
)

// participant update
type FP000047Controller struct {
	beego.Controller
}

// @Title
// @Description participant update
// @Param body  body  models.FP000047I "Input parameter"
// @Success 200  {object} models.FP000047O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000047Controller) Post()() {
}

package controllers

import (
	"github.com/astaxie/beego"
)

// participant suspension
type FP000004Controller struct {
	beego.Controller
}

// @Title
// @Description participant suspension
// @Param body  body  models.FP000004I "Input parameter"
// @Success 200  {object} models.FP000004O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000004Controller) Post()() {
}

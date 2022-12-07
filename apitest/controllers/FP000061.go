package controllers

import (
	"github.com/astaxie/beego"
)

// values update
type FP000061Controller struct {
	beego.Controller
}

// @Title
// @Description values update
// @Param body  body  models.FP000061I "Input parameter"
// @Success 200  {object} models.FP000061O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000061Controller) Post()() {
}

package controllers

import (
	"github.com/astaxie/beego"
)

// participant agency relation deletion
type FP000017Controller struct {
	beego.Controller
}

// @Title
// @Description participant agency relation deletion
// @Param body  body  models.FP000017I "Input parameter"
// @Success 200  {object} models.FP000017O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000017Controller) Post()() {
}

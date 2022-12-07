package controllers

import (
	"github.com/astaxie/beego"
)

// addressing register
type FP020001Controller struct {
	beego.Controller
}

// @Title
// @Description addressing register
// @Param body  body  models.FP020001I "Input parameter"
// @Success 200  {object} models.FP020001O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP020001Controller) Post()() {
}

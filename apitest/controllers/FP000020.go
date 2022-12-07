package controllers

import (
	"github.com/astaxie/beego"
)

// alert-balance setting
type FP000020Controller struct {
	beego.Controller
}

// @Title
// @Description alert-balance setting
// @Param body  body  models.FP000020I "Input parameter"
// @Success 200  {object} models.FP000020O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000020Controller) Post()() {
}

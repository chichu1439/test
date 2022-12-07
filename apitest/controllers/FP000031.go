package controllers

import (
	"github.com/astaxie/beego"
)

// monthly report
type FP000031Controller struct {
	beego.Controller
}

// @Title
// @Description monthly report
// @Param body  body  models.FP000031I "Input parameter"
// @Success 200  {object} models.FP000031O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000031Controller) Post()() {
}

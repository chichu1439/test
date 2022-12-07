package controllers

import (
	"github.com/astaxie/beego"
)

// participant add
type FP000046Controller struct {
	beego.Controller
}

// @Title
// @Description participant add
// @Param body  body  models.FP000046I "Input parameter"
// @Success 200  {object} models.FP000046O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000046Controller) Post()() {
}

package controllers

import (
	"github.com/astaxie/beego"
)

// ransaction limit update
type FP000024Controller struct {
	beego.Controller
}

// @Title
// @Description ransaction limit update
// @Param body  body  models.FP000024I "Input parameter"
// @Success 200  {object} models.FP000024O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000024Controller) Post()() {
}

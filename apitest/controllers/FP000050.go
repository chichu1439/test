package controllers

import (
	"github.com/astaxie/beego"
)

// scenario list
type FP000050Controller struct {
	beego.Controller
}

// @Title
// @Description scenario list
// @Param body  body  models.FP000050I "Input parameter"
// @Success 200  {object} models.FP000050O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000050Controller) Post()() {
}

package controllers

import (
	"github.com/astaxie/beego"
)

// scenario update
type FP000055Controller struct {
	beego.Controller
}

// @Title
// @Description scenario update
// @Param body  body  models.FP000055I "Input parameter"
// @Success 200  {object} models.FP000055O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000055Controller) Post()() {
}

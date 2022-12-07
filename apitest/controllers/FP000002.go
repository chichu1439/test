package controllers

import (
	"github.com/astaxie/beego"
)

// participant detail inquiry
type FP000002Controller struct {
	beego.Controller
}

// @Title
// @Description participant detail inquiry
// @Param body  body  models.FP000002I "Input parameter"
// @Success 200  {object} models.FP000002O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000002Controller) Post()() {
}

package controllers

import (
	"github.com/astaxie/beego"
)

// other participant detail inquiry
type FP000010Controller struct {
	beego.Controller
}

// @Title
// @Description other participant detail inquiry
// @Param body  body  models.FP000010I "Input parameter"
// @Success 200  {object} models.FP000010O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000010Controller) Post()() {
}

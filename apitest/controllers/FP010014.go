package controllers

import (
	"github.com/astaxie/beego"
)

// payment detail enquiry
type FP010014Controller struct {
	beego.Controller
}

// @Title
// @Description payment detail enquiry
// @Param body  body  models.FP010014I "Input parameter"
// @Success 200  {object} models.FP010014O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010014Controller) Post()() {
}

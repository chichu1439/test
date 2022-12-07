package controllers

import (
	"github.com/astaxie/beego"
)

// sweeping enquiry
type FP010013Controller struct {
	beego.Controller
}

// @Title
// @Description sweeping enquiry
// @Param body  body  models.FP010013I "Input parameter"
// @Success 200  {object} models.FP010013O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010013Controller) Post()() {
}

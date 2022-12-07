package controllers

import (
	"github.com/astaxie/beego"
)

// participant agency relation detail inquiry
type FP000008Controller struct {
	beego.Controller
}

// @Title
// @Description participant agency relation detail inquiry
// @Param body  body  models.FP000008I "Input parameter"
// @Success 200  {object} models.FP000008O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000008Controller) Post()() {
}

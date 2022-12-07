package controllers

import (
	"github.com/astaxie/beego"
)

// desc
type SMP0000008Controller struct {
	beego.Controller
}

// @Title
// @Description desc
// @Param body  body  models.SMP0000008I "Input parameter"
// @Success 200  {object} models.SMP0000008O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *SMP0000008Controller) Post()() {
}

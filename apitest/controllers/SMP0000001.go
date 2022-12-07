package controllers

import (
	"github.com/astaxie/beego"
)

// desc
type SMP0000001Controller struct {
	beego.Controller
}

// @Title
// @Description desc
// @Param body  body  models.SMP0000001I "Input parameter"
// @Success 200  {object} models.SMP0000001O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *SMP0000001Controller) Post()() {
}

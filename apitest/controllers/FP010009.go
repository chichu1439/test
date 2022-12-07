package controllers

import (
	"github.com/astaxie/beego"
)

// message redeliver
type FP010009Controller struct {
	beego.Controller
}

// @Title
// @Description message redeliver
// @Param body  body  models.FP010009I "Input parameter"
// @Success 200  {object} models.FP010009O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010009Controller) Post()() {
}

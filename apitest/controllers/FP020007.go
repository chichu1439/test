package controllers

import (
	"github.com/astaxie/beego"
)

// execl add
type FP020007Controller struct {
	beego.Controller
}

// @Title
// @Description execl add
// @Param body  body  models.FP020007I "Input parameter"
// @Success 200  {object} models.FP020007O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP020007Controller) Post()() {
}

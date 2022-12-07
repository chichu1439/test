package controllers

import (
	"github.com/astaxie/beego"
)

// values delete
type FP000062Controller struct {
	beego.Controller
}

// @Title
// @Description values delete
// @Param body  body  models.FP000062I "Input parameter"
// @Success 200  {object} models.FP000062O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000062Controller) Post()() {
}

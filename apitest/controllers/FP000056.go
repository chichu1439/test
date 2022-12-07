package controllers

import (
	"github.com/astaxie/beego"
)

// scenario delete
type FP000056Controller struct {
	beego.Controller
}

// @Title
// @Description scenario delete
// @Param body  body  models.FP000056I "Input parameter"
// @Success 200  {object} models.FP000056O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000056Controller) Post()() {
}

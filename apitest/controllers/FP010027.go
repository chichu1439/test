package controllers

import (
	"github.com/astaxie/beego"
)

// 8583bill-payment
type FP010027Controller struct {
	beego.Controller
}

// @Title
// @Description 8583bill-payment
// @Param body  body  models.FP0100023I "Input parameter"
// @Success 200  {object} models.FP0100023O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010027Controller) Post() {
}

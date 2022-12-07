package controllers

import (
	"github.com/astaxie/beego"
)

// 8583request to pay
type FP010019Controller struct {
	beego.Controller
}

// @Title
// @Description 8583request to pay
// @Param body  body  models.FP010019I "Input parameter"
// @Success 200  {object} models.FP010019O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010019Controller) Post() {
}

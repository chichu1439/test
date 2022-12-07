package controllers

import (
	"github.com/astaxie/beego"
)

// 8583sign off
type FP010025Controller struct {
	beego.Controller
}

// @Title
// @Description 8583sign off
// @Param body  body  models.Spec0200 "Input parameter"
// @Success 200  {object} models.Spec0210  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010025Controller) Post() {
}

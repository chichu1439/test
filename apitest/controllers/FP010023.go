package controllers

import (
	"github.com/astaxie/beego"
)

// 8583sign on
type FP010023Controller struct {
	beego.Controller
}

// @Title
// @Description 8583sign on
// @Param body  body  models.FP0100023I "Input parameter"
// @Success 200  {object} models.FP0100023O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010023Controller) Post() {
}

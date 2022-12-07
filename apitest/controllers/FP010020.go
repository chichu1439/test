package controllers

import (
	"github.com/astaxie/beego"
)

// 8583echo check
type FP010020Controller struct {
	beego.Controller
}

// @Title
// @Description 8583echo check
// @Param body  body  models.FP010023I "Input parameter"
// @Success 200  {object} models.FP010023O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010020Controller) Post() {
}

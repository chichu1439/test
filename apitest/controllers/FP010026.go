package controllers

import (
	"github.com/astaxie/beego"
)

// 8583credit notification
type FP010026Controller struct {
	beego.Controller
}

// @Title
// @Description 8583credit notification
// @Param body  body  models.SpecCredit0200 "Input parameter"
// @Success 200  {object} models.SpecCredit0200  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010026Controller) Post() {
}

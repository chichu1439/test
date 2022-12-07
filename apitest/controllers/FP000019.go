package controllers

import (
	"github.com/astaxie/beego"
)

// debit credit
type FP000019Controller struct {
	beego.Controller
}

// @Title
// @Description debit credit
// @Param body  body  models.FP000019I "Input parameter"
// @Success 200  {object} models.FP000019O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000019Controller) Post()() {
}

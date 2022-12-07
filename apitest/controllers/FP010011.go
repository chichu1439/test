package controllers

import (
	"github.com/astaxie/beego"
)

// payment-return detail
type FP010011Controller struct {
	beego.Controller
}

// @Title
// @Description payment-return detail
// @Param body  body  models.FP010011I "Input parameter"
// @Success 200  {object} models.FP010011O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010011Controller) Post()() {
}

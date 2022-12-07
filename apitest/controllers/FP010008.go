package controllers

import (
	"github.com/astaxie/beego"
)

// message-detail inquiry
type FP010008Controller struct {
	beego.Controller
}

// @Title
// @Description message-detail inquiry
// @Param body  body  models.FP010008I "Input parameter"
// @Success 200  {object} models.FP010008O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010008Controller) Post()() {
}

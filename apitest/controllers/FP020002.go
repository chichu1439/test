package controllers

import (
	"github.com/astaxie/beego"
)

// addressing cancel
type FP020002Controller struct {
	beego.Controller
}

// @Title
// @Description addressing cancel
// @Param body  body  models.FP020002I "Input parameter"
// @Success 200  {object} models.FP020002O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP020002Controller) Post()() {
}

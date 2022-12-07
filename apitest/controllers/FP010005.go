package controllers

import (
	"github.com/astaxie/beego"
)

// payment summary
type FP010005Controller struct {
	beego.Controller
}

// @Title
// @Description payment summary
// @Param body  body  models.FP910002I "Input parameter"
// @Success 200  {object} models.FP910002O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010005Controller) Post() {
}

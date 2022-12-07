package controllers

import (
	"github.com/astaxie/beego"
)

// participant relation list inquiry
type FP000007Controller struct {
	beego.Controller
}

// @Title
// @Description participant relation list inquiry
// @Param body  body  models.FP000007I "Input parameter"
// @Success 200  {object} models.FP000007O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000007Controller) Post()() {
}

package controllers

import (
	"github.com/astaxie/beego"
)

// participant inquiry
type FP000001Controller struct {
	beego.Controller
}

// @Title
// @Description participant inquiry
// @Param body  body  models.FP000001I "Input parameter"
// @Success 200  {object} models.FP000001O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000001Controller) Post()() {
}

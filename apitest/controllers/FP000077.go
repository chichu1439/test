package controllers

import (
	"github.com/astaxie/beego"
)

// alert inquiry
type FP000077Controller struct {
	beego.Controller
}

// @Title
// @Description alert inquiry
// @Param body  body  models.FP000077I "Input parameter"
// @Success 200  {object} models.FP000077O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000077Controller) Post()() {
}

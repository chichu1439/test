package controllers

import (
	"github.com/astaxie/beego"
)

// participant amendment
type FP000003Controller struct {
	beego.Controller
}

// @Title
// @Description participant amendment
// @Param body  body  models.FP000003I "Input parameter"
// @Success 200  {object} models.FP000003O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000003Controller) Post()() {
}

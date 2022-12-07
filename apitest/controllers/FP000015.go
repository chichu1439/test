package controllers

import (
	"github.com/astaxie/beego"
)

// settlement account inquiry
type FP000015Controller struct {
	beego.Controller
}

// @Title
// @Description settlement account inquiry
// @Param body  body  models.FP000015I "Input parameter"
// @Success 200  {object} models.FP000015O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000015Controller) Post()() {
}

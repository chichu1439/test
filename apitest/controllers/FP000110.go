package controllers

import (
	"github.com/astaxie/beego"
)

// checker enquiry
type FP000110Controller struct {
	beego.Controller
}

// @Title
// @Description checker enquiry
// @Param body  body  models.FP000110I "Input parameter"
// @Success 200  {object} models.FP000110O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000110Controller) Post()() {
}

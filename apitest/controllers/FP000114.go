package controllers

import (
	"github.com/astaxie/beego"
)

// history enquiry
type FP000114Controller struct {
	beego.Controller
}

// @Title
// @Description history enquiry
// @Param body  body  models.FP000114I "Input parameter"
// @Success 200  {object} models.FP000114O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000114Controller) Post()() {
}

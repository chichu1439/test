package controllers

import (
	"github.com/astaxie/beego"
)

// adjustment enquiry
type FP000108Controller struct {
	beego.Controller
}

// @Title
// @Description adjustment enquiry
// @Param body  body  models.FP000108I "Input parameter"
// @Success 200  {object} models.FP000108O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000108Controller) Post()() {
}

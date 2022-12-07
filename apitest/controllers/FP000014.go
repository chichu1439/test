package controllers

import (
	"github.com/astaxie/beego"
)

// verification
type FP000014Controller struct {
	beego.Controller
}

// @Title
// @Description verification
// @Param body  body  models.FP000014I "Input parameter"
// @Success 200  {object} models.FP000014O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000014Controller) Post()() {
}

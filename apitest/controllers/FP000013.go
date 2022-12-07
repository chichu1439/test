package controllers

import (
	"github.com/astaxie/beego"
)

// random code creation
type FP000013Controller struct {
	beego.Controller
}

// @Title
// @Description random code creation
// @Param body  body  models.FP000013I "Input parameter"
// @Success 200  {object} models.FP000013O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000013Controller) Post()() {
}

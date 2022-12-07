package controllers

import (
	"github.com/astaxie/beego"
)

// settlement account detail inquiry
type FP000018Controller struct {
	beego.Controller
}

// @Title
// @Description settlement account detail inquiry
// @Param body  body  models.FP000018I "Input parameter"
// @Success 200  {object} models.FP000018O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000018Controller) Post()() {
}

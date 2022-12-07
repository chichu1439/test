package controllers

import (
	"github.com/astaxie/beego"
)

// host start
type FP000066Controller struct {
	beego.Controller
}

// @Title
// @Description host start
// @Param body  body  models.FP000066I "Input parameter"
// @Success 200  {object} models.FP000066O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000066Controller) Post()() {
}

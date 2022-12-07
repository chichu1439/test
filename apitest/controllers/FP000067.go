package controllers

import (
	"github.com/astaxie/beego"
)

// host query
type FP000067Controller struct {
	beego.Controller
}

// @Title
// @Description host query
// @Param body  body  models.FP000067I "Input parameter"
// @Success 200  {object} models.FP000067O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000067Controller) Post()() {
}

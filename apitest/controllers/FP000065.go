package controllers

import (
	"github.com/astaxie/beego"
)

// host list
type FP000065Controller struct {
	beego.Controller
}

// @Title
// @Description host list
// @Param body  body  models.FP000065I "Input parameter"
// @Success 200  {object} models.FP000065O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000065Controller) Post()() {
}

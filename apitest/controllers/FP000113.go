package controllers

import (
	"github.com/astaxie/beego"
)

// payment creat list
type FP000113Controller struct {
	beego.Controller
}

// @Title
// @Description payment creat list
// @Param body  body  models.FP000113I "Input parameter"
// @Success 200  {object} models.FP000113O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000113Controller) Post()() {
}

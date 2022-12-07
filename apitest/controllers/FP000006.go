package controllers

import (
	"github.com/astaxie/beego"
)

// participant creation
type FP000006Controller struct {
	beego.Controller
}

// @Title
// @Description participant creation
// @Param body  body  models.FP000006I "Input parameter"
// @Success 200  {object} models.FP000006O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000006Controller) Post()() {
}

package controllers

import (
	"github.com/astaxie/beego"
)

// participant list
type FP000044Controller struct {
	beego.Controller
}

// @Title
// @Description participant list
// @Param body  body  models.FP000044I "Input parameter"
// @Success 200  {object} models.FP000044O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000044Controller) Post()() {
}

package controllers

import (
	"github.com/astaxie/beego"
)

// prodfee update
type FP000040Controller struct {
	beego.Controller
}

// @Title
// @Description prodfee update
// @Param body  body  models.FP000040I "Input parameter"
// @Success 200  {object} models.FP000040O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000040Controller) Post()() {
}

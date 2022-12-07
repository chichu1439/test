package controllers

import (
	"github.com/astaxie/beego"
)

// aog wait
type FP000100Controller struct {
	beego.Controller
}

// @Title
// @Description aog wait
// @Param body  body  models.FP000100I "Input parameter"
// @Success 200  {object} models.FP000100O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000100Controller) Post()() {
}

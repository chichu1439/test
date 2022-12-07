package controllers

import (
	"github.com/astaxie/beego"
)

// sweeping update
type FP000029Controller struct {
	beego.Controller
}

// @Title
// @Description sweeping update
// @Param body  body  models.FP000029I "Input parameter"
// @Success 200  {object} models.FP000029O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000029Controller) Post()() {
}

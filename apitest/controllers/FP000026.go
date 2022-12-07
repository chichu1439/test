package controllers

import (
	"github.com/astaxie/beego"
)

// query history
type FP000026Controller struct {
	beego.Controller
}

// @Title
// @Description query history
// @Param body  body  models.FP000026I "Input parameter"
// @Success 200  {object} models.FP000026O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000026Controller) Post()() {
}

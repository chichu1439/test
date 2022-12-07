package controllers

import (
	"github.com/astaxie/beego"
)

// payment enum list
type FP000112Controller struct {
	beego.Controller
}

// @Title
// @Description payment enum list
// @Param body  body  models.FP000112I "Input parameter"
// @Success 200  {object} models.FP000112O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000112Controller) Post()() {
}

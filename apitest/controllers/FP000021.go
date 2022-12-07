package controllers

import (
	"github.com/astaxie/beego"
)

// transaction limit add
type FP000021Controller struct {
	beego.Controller
}

// @Title
// @Description transaction limit add
// @Param body  body  models.FP000021I "Input parameter"
// @Success 200  {object} models.FP000021O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000021Controller) Post()() {
}

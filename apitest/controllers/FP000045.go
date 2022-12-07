package controllers

import (
	"github.com/astaxie/beego"
)

// transaction query
type FP000045Controller struct {
	beego.Controller
}

// @Title
// @Description transaction query
// @Param body  body  models.FP000045I "Input parameter"
// @Success 200  {object} models.FP000045O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000045Controller) Post()() {
}

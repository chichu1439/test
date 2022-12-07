package controllers

import (
	"github.com/astaxie/beego"
)

// adjustment confirm
type FP000109Controller struct {
	beego.Controller
}

// @Title
// @Description adjustment confirm
// @Param body  body  models.FP000109I "Input parameter"
// @Success 200  {object} models.FP000109O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000109Controller) Post()() {
}

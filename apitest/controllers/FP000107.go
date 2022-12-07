package controllers

import (
	"github.com/astaxie/beego"
)

// adjustment delete
type FP000107Controller struct {
	beego.Controller
}

// @Title
// @Description adjustment delete
// @Param body  body  models.FP000107I "Input parameter"
// @Success 200  {object} models.FP000107O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000107Controller) Post()() {
}

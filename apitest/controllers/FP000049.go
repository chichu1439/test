package controllers

import (
	"github.com/astaxie/beego"
)

// transaction trigger
type FP000049Controller struct {
	beego.Controller
}

// @Title
// @Description transaction trigger
// @Param body  body  models.FP000049I "Input parameter"
// @Success 200  {object} models.FP000049O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000049Controller) Post()() {
}

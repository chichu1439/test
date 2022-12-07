package controllers

import (
	"github.com/astaxie/beego"
)

// netting param update
type FP000104Controller struct {
	beego.Controller
}

// @Title
// @Description netting param update
// @Param body  body  models.FP000104I "Input parameter"
// @Success 200  {object} models.FP000104O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000104Controller) Post()() {
}

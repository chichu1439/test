package controllers

import (
	"github.com/astaxie/beego"
)

// addressing amendment
type FP020003Controller struct {
	beego.Controller
}

// @Title
// @Description addressing amendment
// @Param body  body  models.FP020003I "Input parameter"
// @Success 200  {object} models.FP020003O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP020003Controller) Post()() {
}

package controllers

import (
	"github.com/astaxie/beego"
)

// addressing inquiry
type FP020004Controller struct {
	beego.Controller
}

// @Title
// @Description addressing inquiry
// @Param body  body  models.FP020004I "Input parameter"
// @Success 200  {object} models.FP020004O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP020004Controller) Post()() {
}

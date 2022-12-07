package controllers

import (
	"github.com/astaxie/beego"
)

// limit query list
type FP000025Controller struct {
	beego.Controller
}

// @Title
// @Description limit query list
// @Param body  body  models.FP000025I "Input parameter"
// @Success 200  {object} models.FP000025O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000025Controller) Post()() {
}

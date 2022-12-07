package controllers

import (
	"github.com/astaxie/beego"
)

// apis chema
type FP000063Controller struct {
	beego.Controller
}

// @Title
// @Description apis chema
// @Param body  body  models.FP000063I "Input parameter"
// @Success 200  {object} models.FP000063O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000063Controller) Post()() {
}

package controllers

import (
	"github.com/astaxie/beego"
)

// sample query
type FP000102Controller struct {
	beego.Controller
}

// @Title
// @Description sample query
// @Param body  body  models.FP000102I "Input parameter"
// @Success 200  {object} models.FP000102O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000102Controller) Post()() {
}

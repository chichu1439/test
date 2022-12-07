package controllers

import (
	"github.com/astaxie/beego"
)

// settlement account creation
type FP000012Controller struct {
	beego.Controller
}

// @Title
// @Description settlement account creation
// @Param body  body  models.FP000012I "Input parameter"
// @Success 200  {object} models.FP000012O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000012Controller) Post()() {
}

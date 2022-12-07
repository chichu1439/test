package controllers

import (
	"github.com/astaxie/beego"
)

// participant agency relation creation
type FP000009Controller struct {
	beego.Controller
}

// @Title
// @Description participant agency relation creation
// @Param body  body  models.FP000009I "Input parameter"
// @Success 200  {object} models.FP000009O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000009Controller) Post()() {
}

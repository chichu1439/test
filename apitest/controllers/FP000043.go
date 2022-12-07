package controllers

import (
	"github.com/astaxie/beego"
)

// report log query
type FP000043Controller struct {
	beego.Controller
}

// @Title
// @Description report log query
// @Param body  body  models.FP000043I "Input parameter"
// @Success 200  {object} models.FP000043O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000043Controller) Post()() {
}

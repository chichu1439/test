package controllers

import (
	"github.com/astaxie/beego"
)

// values creation
type FP000060Controller struct {
	beego.Controller
}

// @Title
// @Description values creation
// @Param body  body  models.FP000060I "Input parameter"
// @Success 200  {object} models.FP000060O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000060Controller) Post()() {
}

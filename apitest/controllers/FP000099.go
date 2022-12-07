package controllers

import (
	"github.com/astaxie/beego"
)

// get messages info
type FP000099Controller struct {
	beego.Controller
}

// @Title
// @Description get messages info
// @Param body  body  models.FP000099I "Input parameter"
// @Success 200  {object} models.FP000099O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000099Controller) Post()() {
}

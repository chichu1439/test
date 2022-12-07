package controllers

import (
	"github.com/astaxie/beego"
)

// request to pay
type FP010018Controller struct {
	beego.Controller
}

// @Title
// @tags xxxx
// @Description request to pay
// @Param body  body  models.FP010018I "Input parameter"
// @Success 200  {object} models.FP010018O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010018Controller) Post() {
}

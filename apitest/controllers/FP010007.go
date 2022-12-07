package controllers

import (
	"github.com/astaxie/beego"
)

// message listin quiry
type FP010007Controller struct {
	beego.Controller
}

// @Title
// @Description message listin quiry
// @Param body  body  models.FP010007I "Input parameter"
// @Success 200  {object} models.FP010007O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010007Controller) Post()() {
}

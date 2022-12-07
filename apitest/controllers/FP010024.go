package controllers

import (
	"github.com/astaxie/beego"
)

// 20022sign off
type FP010024Controller struct {
	beego.Controller
}

// @Title
// @Description 20022sign off
// @Param body  body  models.FPSSignOff003 "Input parameter"
// @Success 200  {object} models.FPSSignOff004  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010024Controller) Post() {
}

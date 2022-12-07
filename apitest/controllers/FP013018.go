package controllers

import (
	"github.com/astaxie/beego"
)

// v3 request to pay
type FP013018Controller struct {
	beego.Controller
}

// @Title
// @Description v3 request to pay
// @Param body  body  models.FpsPain013 "Input parameter"
// @Success 200  {object} models.FpsPain014  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP013018Controller) Post() {
}

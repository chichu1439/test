package controllers

import (
	"github.com/astaxie/beego"
)

// echo check
type FP010021Controller struct {
	beego.Controller
}

// @Title
// @Description echo check
// @Param body  body  models.FPSADMNo005 "Input parameter"
// @Success 200  {object} models.FPSADMNo006  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010021Controller) Post() {
}

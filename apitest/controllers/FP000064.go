package controllers

import (
	"github.com/astaxie/beego"
)

// aig serviceID
type FP000064Controller struct {
	beego.Controller
}

// @Title
// @Description aig serviceID
// @Param body  body  models.FP000064I "Input parameter"
// @Success 200  {object} models.FP000064O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000064Controller) Post()() {
}

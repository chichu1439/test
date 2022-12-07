package controllers

import (
	"github.com/astaxie/beego"
)

// settlement account balancead justment
type FP000016Controller struct {
	beego.Controller
}

// @Title
// @Description settlement account balancead justment
// @Param body  body  models.FP000016I "Input parameter"
// @Success 200  {object} models.FP000016O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000016Controller) Post()() {
}

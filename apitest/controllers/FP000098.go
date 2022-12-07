package controllers

import (
	"github.com/astaxie/beego"
)

// query redis message
type FP000098Controller struct {
	beego.Controller
}

// @Title
// @Description query redis message
// @Param body  body  models.FP000098I "Input parameter"
// @Success 200  {object} models.FP000098O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000098Controller) Post()() {
}

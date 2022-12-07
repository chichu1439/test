package controllers

import (
	"github.com/astaxie/beego"
)

// execl download
type FP020010Controller struct {
	beego.Controller
}

// @Title
// @Description execl download
// @Param body  body  models.FP020010I "Input parameter"
// @Success 200  {object} models.FP020010O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP020010Controller) Post()() {
}

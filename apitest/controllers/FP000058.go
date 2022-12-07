package controllers

import (
	"github.com/astaxie/beego"
)

// step update
type FP000058Controller struct {
	beego.Controller
}

// @Title
// @Description step update
// @Param body  body  models.FP000058I "Input parameter"
// @Success 200  {object} models.FP000058O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000058Controller) Post()() {
}

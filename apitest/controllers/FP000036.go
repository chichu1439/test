package controllers

import (
	"github.com/astaxie/beego"
)

// balance alter noti
type FP000036Controller struct {
	beego.Controller
}

// @Title
// @Description balance alter noti
// @Param body  body  models.FP000036I "Input parameter"
// @Success 200  {object} models.FP000036O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000036Controller) Post()() {
}

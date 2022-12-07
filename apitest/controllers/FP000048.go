package controllers

import (
	"github.com/astaxie/beego"
)

// participant delete
type FP000048Controller struct {
	beego.Controller
}

// @Title
// @Description participant delete
// @Param body  body  models.FP000048I "Input parameter"
// @Success 200  {object} models.FP000048O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000048Controller) Post()() {
}

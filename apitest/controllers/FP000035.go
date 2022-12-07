package controllers

import (
	"github.com/astaxie/beego"
)

// download file
type FP000035Controller struct {
	beego.Controller
}

// @Title
// @Description download file
// @Param body  body  models.FP000035I "Input parameter"
// @Success 200  {object} models.FP000035O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000035Controller) Post()() {
}

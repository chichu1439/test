package controllers

import (
	"github.com/astaxie/beego"
)

// execl list
type FP020009Controller struct {
	beego.Controller
}

// @Title
// @Description execl list
// @Param body  body  models.RequestSfp0200009 "Input parameter"
// @Success 200  {object} models.ResponseSfp0200009  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP020009Controller) Post() {
}

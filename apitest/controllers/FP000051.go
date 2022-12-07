package controllers

import (
	"github.com/astaxie/beego"
)

// scenario detail
type FP000051Controller struct {
	beego.Controller
}

// @Title
// @Description scenario detail
// @Param body  body  models.FP000051I "Input parameter"
// @Success 200  {object} models.FP000051O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000051Controller) Post()() {
}

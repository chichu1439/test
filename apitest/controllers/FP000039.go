package controllers

import (
	"github.com/astaxie/beego"
)

// history query list
type FP000039Controller struct {
	beego.Controller
}

// @Title
// @Description history query list
// @Param body  body  models.FP000039I "Input parameter"
// @Success 200  {object} models.FP000039O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000039Controller) Post()() {
}

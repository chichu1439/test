package controllers

import (
	"github.com/astaxie/beego"
)

// scenario creation
type FP0000572Controller struct {
	beego.Controller
}

// @Title
// @Description scenario creation
// @Param body  body  models.FP000057I "Input parameter"
// @Success 200  {object} models.FP000057O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP0000572Controller) Post() {
}

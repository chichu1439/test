package controllers

import (
	"github.com/astaxie/beego"
)

// step creation
type FP000057Controller struct {
	beego.Controller
}

// @Title
// @Description step creation
// @Param body  body  models.FP000057I "Input parameter"
// @Success 200  {object} models.FP000057O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000057Controller) Post() {
}

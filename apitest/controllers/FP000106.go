package controllers

import (
	"github.com/astaxie/beego"
)

// adjustment update
type FP000106Controller struct {
	beego.Controller
}

// @Title
// @Description adjustment update
// @Param body  body  models.FP000106I "Input parameter"
// @Success 200  {object} models.FP000106O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000106Controller) Post()() {
}

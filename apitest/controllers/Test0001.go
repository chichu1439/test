package controllers

import (
	"github.com/astaxie/beego"
)

// desc
type Test0001Controller struct {
	beego.Controller
}

// @Title
// @Description desc
// @Param body  body  models.Test0001I "Input parameter"
// @Success 200  {object} models.Test0001O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *Test0001Controller) Post()() {
}

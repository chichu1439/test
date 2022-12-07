package controllers

import (
	"github.com/astaxie/beego"
)

// prod fee query
type FP000041Controller struct {
	beego.Controller
}

// @Title
// @Description prod fee query
// @Param body  body  models.FP000041I "Input parameter"
// @Success 200  {object} models.FP000041O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000041Controller) Post()() {
}

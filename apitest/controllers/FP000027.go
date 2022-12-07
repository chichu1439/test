package controllers

import (
	"github.com/astaxie/beego"
)

// prod fee insert
type FP000027Controller struct {
	beego.Controller
}

// @Title
// @Description prod fee insert
// @Param body  body  models.FP000027I "Input parameter"
// @Success 200  {object} models.FP000027O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000027Controller) Post()() {
}

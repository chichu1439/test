package controllers

import (
	"github.com/astaxie/beego"
)

// partclearing query
type FP000070Controller struct {
	beego.Controller
}

// @Title
// @Description partclearing query
// @Param body  body  models.FP000070I "Input parameter"
// @Success 200  {object} models.FP000070O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000070Controller) Post()() {
}

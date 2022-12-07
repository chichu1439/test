package controllers

import (
	"github.com/astaxie/beego"
)

// alert balance set
type FP000078Controller struct {
	beego.Controller
}

// @Title
// @Description alert balance set
// @Param body  body  models.FP000078I "Input parameter"
// @Success 200  {object} models.FP000078O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000078Controller) Post()() {
}

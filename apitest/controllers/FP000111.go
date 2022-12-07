package controllers

import (
	"github.com/astaxie/beego"
)

// result enquiry
type FP000111Controller struct {
	beego.Controller
}

// @Title
// @Description result enquiry
// @Param body  body  models.FP000111I "Input parameter"
// @Success 200  {object} models.FP000111O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000111Controller) Post()() {
}

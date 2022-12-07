package controllers

import (
	"github.com/astaxie/beego"
)

// param enquiry
type FP000103Controller struct {
	beego.Controller
}

// @Title
// @Description param enquiry
// @Param body  body  models.FP000103I "Input parameter"
// @Success 200  {object} models.FP000103O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000103Controller) Post()() {
}

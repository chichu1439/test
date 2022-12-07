package controllers

import (
	"github.com/astaxie/beego"
)

// rtgs-participant inquiry
type FP000011Controller struct {
	beego.Controller
}

// @Title
// @Description rtgs-participant inquiry
// @Param body  body  models.FP000011I "Input parameter"
// @Success 200  {object} models.FP000011O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000011Controller) Post()() {
}

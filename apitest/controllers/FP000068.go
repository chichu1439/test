package controllers

import (
	"github.com/astaxie/beego"
)

// scenario enum query
type FP000068Controller struct {
	beego.Controller
}

// @Title
// @Description scenario enum query
// @Param body  body  models.FP000068I "Input parameter"
// @Success 200  {object} models.FP000068O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000068Controller) Post()() {
}

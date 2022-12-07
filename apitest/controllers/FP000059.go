package controllers

import (
	"github.com/astaxie/beego"
)

// step delete
type FP000059Controller struct {
	beego.Controller
}

// @Title
// @Description step delete
// @Param body  body  models.FP000059I "Input parameter"
// @Success 200  {object} models.FP000059O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000059Controller) Post()() {
}

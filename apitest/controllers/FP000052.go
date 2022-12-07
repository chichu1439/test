package controllers

import (
	"github.com/astaxie/beego"
)

// step detail
type FP000052Controller struct {
	beego.Controller
}

// @Title
// @Description step detail
// @Param body  body  models.FP000052I "Input parameter"
// @Success 200  {object} models.FP000052O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000052Controller) Post()() {
}

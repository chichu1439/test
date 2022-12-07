package controllers

import (
	"github.com/astaxie/beego"
)

// direct debit
type FP010012Controller struct {
	beego.Controller
}

// @Title
// @Description direct debit
// @Param body  body  models.FPSPacs003 "Input parameter"
// @Success 200  {object} models.FPSPacs002  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010012Controller) Post() {
}

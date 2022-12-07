package controllers

import (
	"github.com/astaxie/beego"
)

// direct-debit detail
type FP010010Controller struct {
	beego.Controller
}

// @Title
// @Description direct-debit detail
// @Param body  body  models.FP910017I "Input parameter"
// @Success 200  {object} models.FP910017O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010010Controller) Post() {
}

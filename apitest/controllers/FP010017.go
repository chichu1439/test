package controllers

import (
	"github.com/astaxie/beego"
)

// 8583 credit transfer
type FP010017Controller struct {
	beego.Controller
}

// @Title
// @Description 8583 credit transfer
// @Param body  body  models.FP010017I "Input parameter"
// @Success 200  {object} models.FP010017O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010017Controller) Post() {
}

package controllers

import (
	"github.com/astaxie/beego"
)

// 20022 credit transfer
type FP010003Controller struct {
	beego.Controller
}

// @Title
// @Description 20022 credit transfer
// @Param body  body  models.FPSPacs008 "Input parameter"
// @Success 200  {object} models.FPSPacs002  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010003Controller) Post() {
}

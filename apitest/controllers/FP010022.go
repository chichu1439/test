package controllers

import (
	"github.com/astaxie/beego"
)

// 20022 sign on
type FP010022Controller struct {
	beego.Controller
}

// @Title
// @Description 20022 sign on
// @Param body  body  models.FPSSignOn001 "Input parameter"
// @Success 200  {object} models.FPSSignOn002  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010022Controller) Post() {
}

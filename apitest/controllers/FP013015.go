package controllers

import (
	"github.com/astaxie/beego"
)

// 20022 account verify v3
type FP013015Controller struct {
	beego.Controller
}

// @Title
// @Description 20022 account verify v3
// @Param body  body  models.FpsAcmt023 "Input parameter"
// @Success 200  {object} models.FpsAcmt024  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP013015Controller) Post() {
}

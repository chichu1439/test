package controllers

import (
	"github.com/astaxie/beego"
)

// repayment
type IC000070Controller struct {
	beego.Controller
}

// @Title
// @Description  repayment
// @Param	body		body 	models.IC000070Request	"Input parameter"
// @Success 200  {object} models.IC000070Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *IC000070Controller) Post() ()  {
}

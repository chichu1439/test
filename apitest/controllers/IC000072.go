package controllers

import (
	"github.com/astaxie/beego"
)

// loan list
type IC000072Controller struct {
	beego.Controller
}

// @Title
// @Description  loan list
// @Param	body		body 	models.IC000072Request	"Input parameter"
// @Success 200  {object} models.IC000072Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *IC000072Controller) Post() ()  {
}

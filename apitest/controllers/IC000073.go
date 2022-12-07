package controllers

import (
	"github.com/astaxie/beego"
)

// loan detail
type IC000073Controller struct {
	beego.Controller
}

// @Title
// @Description  loan detail
// @Param	body		body 	models.IC000073Request	"Input parameter"
// @Success 200  {object} models.IC000073Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *IC000073Controller) Post() ()  {
}

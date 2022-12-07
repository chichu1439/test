package controllers

import (
	"github.com/astaxie/beego"
)

// Create Collection Contact Info
type IC000042Controller struct {
	beego.Controller
}

// @Title
// @Description Create Collection Contact Info
// @Param	body		body 	models.IC000042I	"Input parameter"
// @Success 200  {object} models.IC000042O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000042Controller) Post() ()  {
}

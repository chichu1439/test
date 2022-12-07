package controllers

import (
	"github.com/astaxie/beego"
)

// Update Collection Contact Info
type IC000039Controller struct {
	beego.Controller
}

// @Title
// @Description Update Collection Contact Info
// @Param	body		body 	models.IC000039I	"Input parameter"
// @Success 200  {object} models.IC000039O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000039Controller) Post() ()  {
}

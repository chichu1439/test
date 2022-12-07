package controllers

import (
	"github.com/astaxie/beego"
)

// customer detail inquiry
type IC000008Controller struct {
	beego.Controller
}

// @Title
// @Description customer detail inquiry
// @Param	body		body 	models.IC000008I	"Input parameter"
// @Success 200  {object} models.IC000008O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000008Controller) Post() ()  {
}

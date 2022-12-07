package controllers

import (
	"github.com/astaxie/beego"
)

// Customer history inquiry
type IC000031Controller struct {
	beego.Controller
}

// @Title
// @Description Customer history inquiry
// @Param	body		body 	models.IC000031I	"Input parameter"
// @Success 200  {object} models.IC000031O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000031Controller) Post() ()  {
}

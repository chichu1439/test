package controllers

import (
	"github.com/astaxie/beego"
)

// Customer quota inquiry
type IC000059Controller struct {
	beego.Controller
}

// @Title
// @Description Customer quota inquiry
// @Param	body		body 	models.IC000059I	"Input parameter"
// @Success 200  {object} models.IC000059O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000059Controller) Post() ()  {
}

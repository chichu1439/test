package controllers

import (
	"github.com/astaxie/beego"
)

// Modify customer application
type IC000017Controller struct {
	beego.Controller
}

// @Title
// @Description Modify customer application
// @Param	body		body 	models.IC000017I	"Input parameter"
// @Success 200  {object} models.IC000017O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000017Controller) Post() ()  {
}

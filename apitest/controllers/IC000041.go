package controllers

import (
	"github.com/astaxie/beego"
)

// Collection Summary by status
type IC000041Controller struct {
	beego.Controller
}

// @Title
// @Description Collection Summary by status
// @Param	body		body 	models.IC000041I	"Input parameter"
// @Success 200  {object} models.IC000041O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000041Controller) Post() ()  {
}

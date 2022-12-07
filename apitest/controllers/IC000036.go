package controllers

import (
	"github.com/astaxie/beego"
)

// Create Collection Task
type IC000036Controller struct {
	beego.Controller
}

// @Title
// @Description Create Collection Task
// @Param	body		body 	models.IC000036I	"Input parameter"
// @Success 200  {object} models.IC000036O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000036Controller) Post() ()  {
}

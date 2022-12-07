package controllers

import (
	"github.com/astaxie/beego"
)

// Query Collection Task Detail
type IC000038Controller struct {
	beego.Controller
}

// @Title
// @Description Query Collection Task Detail
// @Param	body		body 	models.IC000038I	"Input parameter"
// @Success 200  {object} models.IC000038O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000038Controller) Post() ()  {
}

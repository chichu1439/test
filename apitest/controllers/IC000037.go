package controllers

import (
	"github.com/astaxie/beego"
)

// Query Collection Task List
type IC000037Controller struct {
	beego.Controller
}

// @Title
// @Description Query Collection Task List
// @Param	body		body 	models.IC000037I	"Input parameter"
// @Success 200  {object} models.IC000037O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000037Controller) Post() ()  {
}

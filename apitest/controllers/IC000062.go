package controllers

import (
	"github.com/astaxie/beego"
)

// Customer Credit Line Application
type IC000062Controller struct {
	beego.Controller
}

// @Title
// @Description Customer Credit Line Application
// @Param	body		body 	models.IC000062I	"Input parameter"
// @Success 200  {object} models.IC000062O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000062Controller) Post() ()  {
}

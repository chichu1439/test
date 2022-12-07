package controllers

import (
	"github.com/astaxie/beego"
)

// new customer approve
type IC000010Controller struct {
	beego.Controller
}

// @Title
// @Description new customer approve
// @Param	body		body 	models.IC000010I	"Input parameter"
// @Success 200  {object} models.IC000010O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000010Controller) Post() ()  {
}

package controllers

import (
	"github.com/astaxie/beego"
)
// Create product item
type PD000019Controller struct {
	beego.Controller
}

// @Title
// @Description Create product item
// @Param	body		body 	models.SPD0000019I	"Input parameter"
// @Success 200  {object} models.SPD0000019O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000019Controller) Post() {
}

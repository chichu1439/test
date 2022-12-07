package controllers

import (
	"github.com/astaxie/beego"
)
// Query Product Service Pool
type PD000069Controller struct {
	beego.Controller
}

// @Title
// @Description  Query Product Service Pool
// @Param	body		body 	models.PD000069I	"Input parameter"
// @Success 200  {object} models.PD000069O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000069Controller) Post() {
}

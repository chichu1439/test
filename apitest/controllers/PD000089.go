package controllers

import (
	"github.com/astaxie/beego"
)
// Product Limit Control
type PD000089Controller struct {
	beego.Controller
}

// @Title
// @Description  Product Limit Control
// @Param	body		body 	models.PD000089I	"Input parameter"
// @Success 200  {object} models.PD000089O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000089Controller) Post() {
}

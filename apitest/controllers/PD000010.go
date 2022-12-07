package controllers

import (
	"github.com/astaxie/beego"
)
// Delete current deposit product
type PD000010Controller struct {
	beego.Controller
}

// @Title
// @Description Delete current deposit product
// @Param	body		body 	models.SPD0000010I	"Input parameter"
// @Success 200  {object} models.SPD0000010O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000010Controller) Post() {
}

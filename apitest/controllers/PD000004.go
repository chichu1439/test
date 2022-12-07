package controllers

import (
	"github.com/astaxie/beego"
)
// Delete category
type PD000004Controller struct {
	beego.Controller
}

// @Title
// @Description Delete category
// @Param	body		body 	models.PD000004I	"Input parameter"
// @Success 200  {object} models.PD000004O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000004Controller) Post() {
}

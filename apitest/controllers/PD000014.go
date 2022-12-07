package controllers

import (
	"github.com/astaxie/beego"
)
// Delete micro loan product
type PD000014Controller struct {
	beego.Controller
}

// @Title
// @Description Delete micro loan product
// @Param	body		body 	models.SPD0000014I	"Input parameter"
// @Success 200  {object} models.SPD0000014O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000014Controller) Post() {
}

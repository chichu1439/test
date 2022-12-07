package controllers

import (
	"github.com/astaxie/beego"
)
// Delete product item
type PD000021Controller struct {
	beego.Controller
}

// @Title
// @Description Delete product item
// @Param	body		body 	models.SPD0000021I	"Input parameter"
// @Success 200  {object} models.SPD0000021O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000021Controller) Post() {
}

package controllers

import (
	"github.com/astaxie/beego"
)
// Modify product item
type PD000020Controller struct {
	beego.Controller
}

// @Title
// @Description Modify product item
// @Param	body		body 	models.SPD0000020I	"Input parameter"
// @Success 200  {object} models.SPD0000020O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000020Controller) Post() {
}

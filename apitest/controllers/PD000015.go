package controllers

import (
	"github.com/astaxie/beego"
)
// Query Saving interest strategy
type PD000015Controller struct {
	beego.Controller
}

// @Title
// @Description Query Saving interest strategy
// @Param	body		body 	models.SPD0000015I	"Input parameter"
// @Success 200  {object} models.SPD0000015O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000015Controller) Post() {
}

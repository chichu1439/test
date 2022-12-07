package controllers

import (
	"github.com/astaxie/beego"
)
// Modify Saving fee strategy
type PD000018Controller struct {
	beego.Controller
}

// @Title
// @Description Modify Saving fee strategy
// @Param	body		body 	models.SPD0000018I	"Input parameter"
// @Success 200  {object} models.SPD0000018O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000018Controller) Post() {
}

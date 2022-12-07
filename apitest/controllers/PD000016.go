package controllers

import (
	"github.com/astaxie/beego"
)
// Modify Saving interest strategy
type PD000016Controller struct {
	beego.Controller
}

// @Title
// @Description Modify Saving interest strategy
// @Param	body		body 	models.SPD0000016I	"Input parameter"
// @Success 200  {object} models.SPD0000016O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000016Controller) Post() {
}

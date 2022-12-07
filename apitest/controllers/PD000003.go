package controllers

import (
	"github.com/astaxie/beego"
)
// Modify category
type PD000003Controller struct {
	beego.Controller
}

// @Title
// @Description Modify category
// @Param	body		body 	models.PD000003I	"Input parameter"
// @Success 200  {object} models.PD000003O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000003Controller) Post() {
}

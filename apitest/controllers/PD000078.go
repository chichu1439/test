package controllers

import (
	"github.com/astaxie/beego"
)
// Query Product Fee History
type PD000078Controller struct {
	beego.Controller
}

// @Title
// @Description  Query Product Fee History
// @Param	body		body 	models.PD000078I	"Input parameter"
// @Success 200  {object} models.PD000078O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000078Controller) Post() {
}

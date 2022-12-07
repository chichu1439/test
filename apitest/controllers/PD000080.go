package controllers

import (
	"github.com/astaxie/beego"
)
// Query Product Fee History
type PD000080Controller struct {
	beego.Controller
}

// @Title
// @Description  Query Product Fee History
// @Param	body		body 	models.PD000080I	"Input parameter"
// @Success 200  {object} models.PD000080O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000080Controller) Post() {
}

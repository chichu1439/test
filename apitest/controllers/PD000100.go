package controllers

import (
	"github.com/astaxie/beego"
)
// Query the last product modification history
type PD000100Controller struct {
	beego.Controller
}

// @Title
// @Description  Query the last product modification history
// @Param	body		body 	models.PD000100I	"Input parameter"
// @Success 200  {object} models.PD000100O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000100Controller) Post() {
}

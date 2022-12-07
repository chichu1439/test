package controllers

import (
	"github.com/astaxie/beego"
)
// Product examine
type PD000096Controller struct {
	beego.Controller
}

// @Title
// @Description  Product examine
// @Param	body		body 	models.PD000096I	"Input parameter"
// @Success 200  {object} models.PD000096O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000096Controller) Post() {
}

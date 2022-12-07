package controllers

import (
	"github.com/astaxie/beego"
)
// Query Product Limit History
type PD000082Controller struct {
	beego.Controller
}

// @Title
// @Description  Query Product Limit History
// @Param	body		body 	models.PD000082I	"Input parameter"
// @Success 200  {object} models.PD000082O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000082Controller) Post() {
}

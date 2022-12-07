package controllers

import (
	"github.com/astaxie/beego"
)
// Query product list
type PD000006Controller struct {
	beego.Controller
}

// @Title
// @Description Query product list
// @Param	body		body 	models.SPD0000006I	"Input parameter"
// @Success 200  {object} models.SPD0000006O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000006Controller) Post() {
}

package controllers

import (
	"github.com/astaxie/beego"
)
// Query category tree
type PD000002Controller struct {
	beego.Controller
}

// @Title
// @Description Query category tree
// @Param	body		body 	models.PD000002I	"Input parameter"
// @Success 200  {object} models.PD000002O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000002Controller) Post() {
}

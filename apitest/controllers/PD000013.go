package controllers

import (
	"github.com/astaxie/beego"
)
// Modify micro loan product
type PD000013Controller struct {
	beego.Controller
}

// @Title
// @Description Modify micro loan product
// @Param	body		body 	models.SPD0000013I	"Input parameter"
// @Success 200  {object} models.SPD0000013O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000013Controller) Post() {
}

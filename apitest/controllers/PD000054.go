package controllers

import (
	"github.com/astaxie/beego"
)
// Modify Category Product
type PD000054Controller struct {
	beego.Controller
}

// @Title
// @Description  Modify Category Product
// @Param	body		body 	models.SPD0000054I	"Input parameter"
// @Success 200  {object} models.SPD0000054O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000054Controller) Post() {
}

package controllers

import (
	"github.com/astaxie/beego"
)
// Query Saving fee strategy
type PD000017Controller struct {
	beego.Controller
}

// @Title
// @Description Query Saving fee strategy
// @Param	body		body 	models.SPD0000017I	"Input parameter"
// @Success 200  {object} models.SPD0000017O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000017Controller) Post() {
}

package controllers

import (
	"github.com/astaxie/beego"
)
// Modify current deposit product
type PD000009Controller struct {
	beego.Controller
}

// @Title
// @Description Modify current deposit product
// @Param	body		body 	models.SPD0000009I	"Input parameter"
// @Success 200  {object} models.SPD0000009O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000009Controller) Post() {
}
